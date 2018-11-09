package messagebus

import (
	"bytes"
	"context"
	"encoding/gob"
	"io"

	"github.com/satori/go.uuid"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/core/reply"
)

type tape interface {
	Write(ctx context.Context, writer io.Writer) error
	GetReply(ctx context.Context, msgHash []byte) (core.Reply, error)
	SetReply(ctx context.Context, msgHash []byte, rep core.Reply) error
}

// storagetape saves and fetches message replies to/from local storage.
//
// It uses <storagetape id> + <message hash> for Value keys.
type storagetape struct {
	ls    core.LocalStorage
	pulse core.PulseNumber
	id    uuid.UUID
}

type couple struct {
	Key   []byte
	Value []byte
}

// NewTape creates new storagetape with random id.
func NewTape(ls core.LocalStorage, pulse core.PulseNumber) (*storagetape, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &storagetape{ls: ls, pulse: pulse, id: id}, nil
}

// NewTapeFromReader creates and fills a new storagetape from a stream.
//
// This is a very long operation, as it saves replies in storage until the stream is exhausted.
func NewTapeFromReader(ctx context.Context, ls core.LocalStorage, reader io.Reader) (*storagetape, error) {
	var err error
	tape := storagetape{ls: ls}

	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(&tape.pulse)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(&tape.id)
	if err != nil {
		return nil, err
	}
	for {
		var rep couple
		err = decoder.Decode(&rep)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		err = tape.setReplyBinary(ctx, rep.Key, rep.Value)
		if err != nil {
			return nil, err
		}
	}

	return &tape, nil
}

func (t *storagetape) Write(ctx context.Context, writer io.Writer) error {
	var err error

	encoder := gob.NewEncoder(writer)
	err = encoder.Encode(t.pulse)
	if err != nil {
		return err
	}
	err = encoder.Encode(t.id)
	if err != nil {
		return err
	}

	err = t.ls.Iterate(ctx, t.pulse, t.id[:], func(k, v []byte) error {
		err = encoder.Encode(&couple{
			Key:   k[len(t.id):],
			Value: v,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *storagetape) GetReply(ctx context.Context, msgHash []byte) (core.Reply, error) {
	key := bytes.Join([][]byte{t.id[:], msgHash}, nil)
	buff, err := t.ls.Get(ctx, t.pulse, key)
	if err != nil {
		return nil, err
	}

	return reply.Deserialize(bytes.NewBuffer(buff))
}

func (t *storagetape) SetReply(ctx context.Context, msgHash []byte, rep core.Reply) error {
	reader, err := reply.Serialize(rep)
	if err != nil {
		return err
	}
	buff := new(bytes.Buffer)
	_, err = buff.ReadFrom(reader)
	if err != nil {
		return err
	}
	return t.setReplyBinary(ctx, msgHash, buff.Bytes())
}

func (t *storagetape) setReplyBinary(ctx context.Context, msgHash []byte, rep []byte) error {
	key := bytes.Join([][]byte{t.id[:], msgHash}, nil)
	return t.ls.Set(ctx, t.pulse, key, rep)
}
