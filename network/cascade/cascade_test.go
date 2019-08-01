//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package cascade

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/platformpolicy"
	"github.com/insolar/insolar/testutils"
)

const (
	domainStr = ".1tJD1hMFxYYt9rHcYuvCMLdCn4AZdPfy4HPaavNWn8"
	id1Str    = "4K2V1kpVycZ6qSFsNdz2FtpNxnJs17eBNzf9rdCMcKoe"
	id2Str    = "4NwnA4HWZurKyXWNowJwYmb9CwX4gBKzwQKov1ExMf8M"
	id3Str    = "4Ss5JMkXAD9Z7cktFEdrqeMuT6jGMF1pVozTyPHZ6zT4"
	id4Str    = "4WnNSfDXkWSnFi1PgXxn8X8fhFwU2Jhe4Df82mL9rKmm"
	id5Str    = "4ahfaxgYLok1PoFu7qHhRPuRwR9fhNPTcdKn69Nkbf6U"
	id6Str    = "4ecxjG9Yw73EXtWQZ8cciGgCBaMsNS5HB2zS9XRMLzRB"
	id7Str    = "4iYFsZcZXQLTfykuzRwY19SxRja53Vm6jSf6CuTx6Kjt"
	id8Str    = "4nTZ1s5a7hdgp51RRjGTJ2DiftnGiZSvHrKkGHWYqf4b"
	id9Str    = "4rNrAAYahzvuxAFvs2bNatzUv3zUPd8jrFzQKfZ9azPJ"
	id10Str   = "4K1b7kbvUPB935DdMuLqpfmG23zMhxKcHQ9gbdmydPVZ"
	id11Str   = "4K2UQtex1jnjN2Vx8yCMcsmf1HNuMJ4NeA7TgNeVs7kk"
	id12Str   = "4K3Mi2hyZ6QKgynGv33sR5n3zWmSzdo8zv5Em7X26r1w"
)

func TestCalculateNextNodes(t *testing.T) {
	//	t.Skip()
	nodeIds := make([]insolar.Reference, 0)

	ref, err := insolar.NewReferenceFromBase58(id1Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id2Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id3Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id4Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id5Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id6Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id7Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id8Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id9Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id10Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id11Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)
	ref, err = insolar.NewReferenceFromBase58(id12Str + domainStr)
	require.NoError(t, err)
	nodeIds = append(nodeIds, *ref)

	c := insolar.Cascade{
		NodeIds:           nodeIds,
		Entropy:           insolar.Entropy{0},
		ReplicationFactor: 2,
	}
	pcs := platformpolicy.NewPlatformCryptographyScheme()

	r, _ := CalculateNextNodes(pcs, c, nil)
	require.Equal(t, []insolar.Reference{nodeIds[8], nodeIds[7]}, r)
	r, _ = CalculateNextNodes(pcs, c, &nodeIds[8])
	require.Equal(t, []insolar.Reference{nodeIds[9], nodeIds[5]}, r)
	r, _ = CalculateNextNodes(pcs, c, &nodeIds[9])
	require.Equal(t, []insolar.Reference{nodeIds[11], nodeIds[2]}, r)
}

func Test_geometricProgressionSum(t *testing.T) {
	require.Equal(t, 1022, geometricProgressionSum(2, 2, 9))
	require.Equal(t, 39, geometricProgressionSum(3, 3, 3))
}

func Test_calcHash(t *testing.T) {
	pcs := platformpolicy.NewPlatformCryptographyScheme()

	ref, err := insolar.NewReferenceFromBase58("4SxZ6BSx6qBP41nqQgtsFW5EF3JLDxYscZeVQnviPUGZ.1tJD1hMFxYYt9rHcYuvCMLdCn4AZdPfy4HPaavNWn8")
	require.NoError(t, err)
	c := []byte{0x5c, 0x57, 0x7f, 0xd, 0x80, 0xe1, 0xb8, 0xfb, 0xa2, 0x9a, 0x69, 0x17, 0x3, 0xa7, 0x66, 0x86, 0xd6,
		0xce, 0xba, 0xb6, 0xd5, 0x48, 0xcc, 0x6d, 0x18, 0xb2, 0x5d, 0xac, 0x65, 0xd1, 0x91, 0xb0, 0xa8, 0xf6, 0xa7,
		0xa7, 0x40, 0x80, 0xcb, 0xbb, 0x16, 0x82, 0xf8, 0xc9, 0x8d, 0xe1, 0x9e, 0xda, 0x2a, 0xdc, 0x94, 0x8a, 0xee,
		0x6e, 0x68, 0x9b, 0xf0, 0x33, 0xb8, 0x7a, 0x6a, 0x3d, 0x34, 0x40}
	require.Equal(t, c, calcHash(pcs, *ref, insolar.Entropy{0}))
}

func Test_getNextCascadeLayerIndexes(t *testing.T) {
	// nodeIds := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	nodeIds := make([]insolar.Reference, 0, 12)
	for i := 0; i < 11; i++ {
		nodeIds = append(nodeIds, testutils.RandomRef())
	}
	startIndex, endIndex := getNextCascadeLayerIndexes(nodeIds, nodeIds[4], 2)
	require.Equal(t, 10, startIndex)
	require.Equal(t, 12, endIndex)
	startIndex, endIndex = getNextCascadeLayerIndexes(nodeIds, nodeIds[1], 2)
	require.Equal(t, 4, startIndex)
	require.Equal(t, 6, endIndex)
	startIndex, endIndex = getNextCascadeLayerIndexes(nodeIds, nodeIds[2], 3)
	require.Equal(t, 9, startIndex)
	require.Equal(t, 12, endIndex)
	startIndex, endIndex = getNextCascadeLayerIndexes(nodeIds, testutils.RandomRef(), 2)
	require.Equal(t, len(nodeIds), startIndex)
	require.Equal(t, len(nodeIds), endIndex)
}
