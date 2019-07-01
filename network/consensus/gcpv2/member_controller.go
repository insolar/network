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

package gcpv2

import (
	"context"
	common2 "github.com/insolar/insolar/network/consensus/gcpv2/common"
	"sync"

	"github.com/insolar/insolar/network/consensus/common"
	"github.com/insolar/insolar/network/consensus/gcpv2/census"
	"github.com/insolar/insolar/network/consensus/gcpv2/core"
	"github.com/insolar/insolar/network/consensus/gcpv2/errors"
	"github.com/insolar/insolar/network/consensus/gcpv2/packets"
)

func NewConsensusMemberController(chronicle census.ConsensusChronicles, upstream core.UpstreamPulseController,
	roundFactory core.RoundControllerFactory) core.ConsensusController {

	return &ConsensusMemberController{
		upstreamPulseController: upstream,
		chronicle:               chronicle,
		roundFactory:            roundFactory,
		hasRequestedPower:       true,
		requestedPower:          100,
	}
}

type upstreamPulseController core.UpstreamPulseController

type ConsensusMemberController struct {
	/* No mutex needed. Set on construction */
	upstreamPulseController
	chronicle    census.ConsensusChronicles
	roundFactory core.RoundControllerFactory

	mutex sync.RWMutex
	/* mutex needed */
	currentRound core.RoundController

	hasRequestedPower bool
	requestedPower    common2.MemberPower
}

func (h *ConsensusMemberController) Abort() {
	panic("implement me")
}

func (h *ConsensusMemberController) RequestLeave() {
	panic("implement me")
}

func (h *ConsensusMemberController) GetActivePowerLimit() (common2.MemberPower, common.PulseNumber) {
	census := h.chronicle.GetActiveCensus()
	return census.GetOnlinePopulation().GetLocalProfile().GetPower(), census.GetPulseNumber()
}

func (h *ConsensusMemberController) _getRequestedPowerLimit() common2.MemberPower {
	if h.hasRequestedPower {
		return h.requestedPower
	}

	census := h.chronicle.GetLatestCensus()
	return census.GetOnlinePopulation().GetLocalProfile().GetDeclaredPower()
}

func (h *ConsensusMemberController) GetRequestedPowerLimit() common2.MemberPower {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h._getRequestedPowerLimit()
}

func (h *ConsensusMemberController) AdjustPowerLimit(pwl core.MemberPowerLevel) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	//TODO pick allowed levels from the mandate
	switch pwl {
	case core.PowerLevelZero:
		h.requestedPower = 0
	case core.PowerLevelMinimal:
		h.requestedPower = 25
	case core.PowerLevelReduced:
		h.requestedPower = 75
	case core.PowerLevelFull:
		h.requestedPower = 100
	default:
		panic("illegal value")
	}

	h.hasRequestedPower = true
}

func (h *ConsensusMemberController) getCurrentRound() core.RoundController {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h.currentRound
}

func (h *ConsensusMemberController) ensureRound() (core.RoundController, bool) {
	r := h.getCurrentRound()
	if r != nil {
		return r, false
	}
	return h._getOrCreateRound()
}

func (h *ConsensusMemberController) _getOrCreateRound() (core.RoundController, bool) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if h.currentRound != nil {
		return h.currentRound, false
	}

	pw := h._getRequestedPowerLimit()
	h.hasRequestedPower = false
	h.currentRound = h.roundFactory.CreateConsensusRound(h.chronicle, pw)
	h.currentRound.StartConsensusRound(h)
	return h.currentRound, true
}

func (h *ConsensusMemberController) _discardRound(toBeDiscarded core.RoundController) core.RoundController {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	round := h.currentRound
	if toBeDiscarded != nil && toBeDiscarded != round {
		//This round was already discarded
		return nil
	}
	h.currentRound = nil

	return round
}

func (h *ConsensusMemberController) discardRound(toBeDiscarded core.RoundController) {
	round := h._discardRound(toBeDiscarded)
	if round != nil {
		go round.StopConsensusRound()
	}
}

func (h *ConsensusMemberController) _processPacket(ctx context.Context, payload packets.PacketParser, from common.HostIdentityHolder, repeated bool) (core.RoundController, bool, error) {
	round, created := h.ensureRound()
	err := round.HandlePacket(ctx, payload, from)

	if ok, _ := errors.IsNextPulseArrivedError(err); ok {
		if repeated || created {
			return round, false, err
		}
		return round, false, nil
	}
	return round, err == nil, err
}

func (h *ConsensusMemberController) ProcessPacket(ctx context.Context, payload packets.PacketParser, from common.HostIdentityHolder) error {

	round, ok, err := h._processPacket(ctx, payload, from, false)
	if ok || err != nil {
		return err
	}
	h.discardRound(round)

	_, _, err = h._processPacket(ctx, payload, from, true)
	return err
}

func (h *ConsensusMemberController) MembershipConfirmed(report core.MembershipUpstreamReport, expectedCensus census.OperationalCensus) {
	// h.discardRound()
	h.upstreamPulseController.MembershipConfirmed(report, expectedCensus)
}

func (h *ConsensusMemberController) MembershipLost(graceful bool) {
	h.discardRound(nil)
	h.upstreamPulseController.MembershipLost(graceful)
}
