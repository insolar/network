///
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
///

package critlog

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"io"
	"sync/atomic"
)

var _ zerolog.LevelWriter = &fatalFlusher{}

type Flusher interface {
	Flush() error
}

func FatalFlusher(w io.Writer, lockPostFatal bool) zerolog.LevelWriter {
	return &fatalFlusher{w: AsLevelWriter(w), lockPostFatal: true}
}

func FatalFlusherExt(w io.Writer, lockPostFatal bool) zerolog.LevelWriter {
	return &fatalFlusher{w: AsLevelWriter(w), lockPostFatal: lockPostFatal}
}

type fatalFlusher struct {
	w             zerolog.LevelWriter
	lockPostFatal bool
	state         uint32 // atomic
}

func (w *fatalFlusher) hasFatal() bool {
	return atomic.LoadUint32(&w.state) != 0
}

func (w *fatalFlusher) setFatal() bool {
	return atomic.CompareAndSwapUint32(&w.state, 0, 1)
}

func (w *fatalFlusher) Write(p []byte) (n int, err error) {
	if !w.hasFatal() {
		return w.w.Write(p)
	}
	if w.lockPostFatal {
		lockDown()
	}
	return len(p), nil
}

func (w *fatalFlusher) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	switch {
	case level == zerolog.FatalLevel:
		if w.setFatal() {
			n, err = w.w.Write(p)
			w.flushOrClose()
			return
		}
		fallthrough
	case w.hasFatal():
		if w.lockPostFatal {
			lockDown()
		}
		return len(p), nil
	case level == zerolog.PanicLevel:
		n, err = w.w.Write(p)
		w.flush()
		return
	default:
		return w.w.Write(p)
	}
}

func (w *fatalFlusher) Flush() error {
	if f, ok := w.w.(Flusher); ok {
		return f.Flush()
	}
	return errors.New("unsupported")
}

func (w *fatalFlusher) Close() error {
	if f, ok := w.w.(io.Closer); ok {
		return f.Close()
	}
	return errors.New("unsupported")
}

func (w *fatalFlusher) flush() bool {
	return w.Flush() == nil
}

func (w *fatalFlusher) flushOrClose() {
	if w.flush() {
		return
	}
	if c, ok := w.w.(io.Closer); ok {
		_ = c.Close()
		return
	}
	// TODO any other approaches? e.g. write a long (4kB) padding?
}

func lockDown() {
	select {
	case <-context.Background().Done():
	}
}
