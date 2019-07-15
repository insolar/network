///
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
///

package executor

import (
	"context"
	"testing"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

func TestFinalizationKeeper_OldCurrentPulse(t *testing.T) {
	testPulse := insolar.GenesisPulse.PulseNumber
	jkMock := NewJetKeeperMock(t)
	jkMock.TopSyncPulseMock.ExpectOnce().Return(testPulse + 1)
	fk := NewFinalizationKeeper(jkMock, nil, 100)
	err := fk.OnPulse(context.Background(), testPulse)
	require.EqualError(t, err, "Current pulse ( 65537 ) is less than last confirmed ( 65538 )")
}

func TestFinalizationKeeper_LimitExceeded(t *testing.T) {
	testPulse := insolar.GenesisPulse.PulseNumber
	limit := 10
	jkMock := NewJetKeeperMock(t)
	jkMock.TopSyncPulseMock.ExpectOnce().Return(testPulse)

	networkMock := testutils.NewTerminationHandlerMock(t)
	networkMock.LeaveMock.Return()

	fk := NewFinalizationKeeper(jkMock, networkMock, limit)
	err := fk.OnPulse(context.Background(), testPulse+insolar.PulseNumber(limit*10))
	require.Contains(t, err.Error(), "Last finalized pulse falls behind too much")
}

func TestFinalizationKeeper_HappyPath(t *testing.T) {
	testPulse := insolar.GenesisPulse.PulseNumber
	limit := 10
	jkMock := NewJetKeeperMock(t)
	jkMock.TopSyncPulseMock.ExpectOnce().Return(testPulse)

	networkMock := testutils.NewTerminationHandlerMock(t)
	networkMock.LeaveMock.Return()

	fk := NewFinalizationKeeper(jkMock, networkMock, limit)
	err := fk.OnPulse(context.Background(), testPulse+insolar.PulseNumber(limit))
	require.NoError(t, err)
}
