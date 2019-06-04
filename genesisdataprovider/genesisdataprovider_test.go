//
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
//

package genesisdataprovider

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/insolar/insolar/component"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/testutils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func mockContractRequesterWithError(t *testing.T) *testutils.ContractRequesterMock {
	contractRequesterMock := testutils.NewContractRequesterMock(t)
	contractRequesterMock.SendRequestFunc = func(p context.Context, p1 *insolar.Reference, p2 string, p3 []interface{}) (r insolar.Reply, r1 error) {
		return nil, errors.New("test reasons")
	}
	return contractRequesterMock
}

func mockContractRequester(t *testing.T, res insolar.Reply) *testutils.ContractRequesterMock {
	contractRequesterMock := testutils.NewContractRequesterMock(t)
	contractRequesterMock.SendRequestFunc = func(p context.Context, p1 *insolar.Reference, p2 string, p3 []interface{}) (r insolar.Reply, r1 error) {
		return res, nil
	}
	return contractRequesterMock
}

func mockCertificateManager(t *testing.T, rootDomainRef *insolar.Reference) *testutils.CertificateManagerMock {
	certificateMock := testutils.NewCertificateMock(t)
	certificateMock.GetRootDomainReferenceFunc = func() (r *insolar.Reference) {
		return rootDomainRef
	}

	certificateManagerMock := testutils.NewCertificateManagerMock(t)
	certificateManagerMock.GetCertificateFunc = func() (r insolar.Certificate) {
		return certificateMock
	}
	return certificateManagerMock
}

func mockInfoResult(rootMemberRef insolar.Reference, oracleMemberRefs map[string]insolar.Reference, mdAdminMemberRef insolar.Reference, nodeDomainRef insolar.Reference) insolar.Reply {
	oracleMemberRefStrs := map[string]string{}
	for name, ref := range oracleMemberRefs {
		oracleMemberRefStrs[name] = ref.String()
	}
	result := map[string]interface{}{
		"root_member":     rootMemberRef.String(),
		"oracle_members":  oracleMemberRefStrs,
		"md_admin_member": mdAdminMemberRef.String(),
		"node_domain":     nodeDomainRef.String(),
	}
	resJSON, _ := json.Marshal(result)
	resSer, _ := insolar.MarshalArgs(resJSON, nil)
	return &reply.CallMethod{Result: resSer}
}

func TestNew(t *testing.T) {
	contractRequester := mockContractRequester(t, nil)
	certificateManager := mockCertificateManager(t, nil)

	result, err := New()

	cm := &component.Manager{}
	cm.Inject(contractRequester, certificateManager, result)

	require.NoError(t, err)
	require.Equal(t, result.CertificateManager, certificateManager)
	require.Equal(t, result.ContractRequester, contractRequester)
}

func TestGenesisDataProvider_setInfo(t *testing.T) {
	ctx := inslogger.TestContext(t)
	rootMemberRef := testutils.RandomRef()
	oracleMemberRefs := map[string]insolar.Reference{}
	for i := 0; i < 10; i++ {
		oracleMemberRefs["oracle"+strconv.Itoa(i)] = testutils.RandomRef()
	}
	mdAdminMemberRef := testutils.RandomRef()
	nodeDomainRef := testutils.RandomRef()

	infoRes := mockInfoResult(rootMemberRef, oracleMemberRefs, mdAdminMemberRef, nodeDomainRef)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequester(t, infoRes),
	}

	err := gdp.setInfo(ctx)

	require.NoError(t, err)
	require.Equal(t, &rootMemberRef, gdp.rootMemberRef)
	require.Equal(t, &nodeDomainRef, gdp.nodeDomainRef)
}

func TestGenesisDataProvider_setInfo_ErrorSendRequest(t *testing.T) {
	ctx := inslogger.TestContext(t)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequesterWithError(t),
	}

	err := gdp.setInfo(ctx)

	require.Contains(t, err.Error(), "test reasons")
}

func TestGenesisDataProvider_GetRootDomain(t *testing.T) {
	ctx := inslogger.TestContext(t)
	rootDomainRef := testutils.RandomRef()

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, &rootDomainRef),
	}

	res := gdp.GetRootDomain(ctx)

	require.Equal(t, &rootDomainRef, res)
}

func TestGenesisDataProvider_GetRootMember(t *testing.T) {
	ctx := inslogger.TestContext(t)
	rootMemberRef := testutils.RandomRef()
	oracleMemberRefs := map[string]insolar.Reference{}
	for i := 0; i < 10; i++ {
		oracleMemberRefs["oracle"+strconv.Itoa(i)] = testutils.RandomRef()
	}
	mdAdminMemberRef := testutils.RandomRef()
	nodeDomainRef := testutils.RandomRef()

	infoRes := mockInfoResult(rootMemberRef, oracleMemberRefs, mdAdminMemberRef, nodeDomainRef)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequester(t, infoRes),
	}

	res, err := gdp.GetRootMember(ctx)

	require.NoError(t, err)
	require.Equal(t, &rootMemberRef, res)
}

func TestGenesisDataProvider_GetRootMember_AlreadySet(t *testing.T) {
	ctx := inslogger.TestContext(t)
	rootMemberRef := testutils.RandomRef()
	oracleMemberRefs := map[string]insolar.Reference{}
	for i := 0; i < 10; i++ {
		oracleMemberRefs["oracle"+strconv.Itoa(i)] = testutils.RandomRef()
	}
	mdAdminMemberRef := testutils.RandomRef()
	nodeDomainRef := testutils.RandomRef()

	newRootMemberRef := testutils.RandomRef()
	infoRes := mockInfoResult(newRootMemberRef, oracleMemberRefs, mdAdminMemberRef, nodeDomainRef)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequester(t, infoRes),
	}
	gdp.rootMemberRef = &rootMemberRef

	res, err := gdp.GetRootMember(ctx)

	require.NoError(t, err)
	require.Equal(t, &rootMemberRef, res)
	require.NotEqual(t, &newRootMemberRef, res)
}

func TestGenesisDataProvider_GetRootMember_Error(t *testing.T) {
	ctx := inslogger.TestContext(t)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequesterWithError(t),
	}

	res, err := gdp.GetRootMember(ctx)

	require.Contains(t, err.Error(), "test reasons")
	require.Nil(t, res)
}

func TestGenesisDataProvider_GetNodeDomain(t *testing.T) {
	ctx := inslogger.TestContext(t)
	rootMemberRef := testutils.RandomRef()
	oracleMemberRefs := map[string]insolar.Reference{}
	for i := 0; i < 10; i++ {
		oracleMemberRefs["oracle"+strconv.Itoa(i)] = testutils.RandomRef()
	}
	mdAdminMemberRef := testutils.RandomRef()
	nodeDomainRef := testutils.RandomRef()

	infoRes := mockInfoResult(rootMemberRef, oracleMemberRefs, mdAdminMemberRef, nodeDomainRef)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequester(t, infoRes),
	}

	res, err := gdp.GetNodeDomain(ctx)

	require.NoError(t, err)
	require.Equal(t, &nodeDomainRef, res)
}

func TestGenesisDataProvider_GetNodeDomain_AlreadySet(t *testing.T) {
	ctx := inslogger.TestContext(t)
	rootMemberRef := testutils.RandomRef()
	oracleMemberRefs := map[string]insolar.Reference{}
	for i := 0; i < 10; i++ {
		oracleMemberRefs["oracle"+strconv.Itoa(i)] = testutils.RandomRef()
	}
	mdAdminMemberRef := testutils.RandomRef()
	nodeDomainRef := testutils.RandomRef()

	newNodeDomainRef := testutils.RandomRef()
	infoRes := mockInfoResult(rootMemberRef, oracleMemberRefs, mdAdminMemberRef, newNodeDomainRef)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequester(t, infoRes),
	}
	gdp.nodeDomainRef = &nodeDomainRef

	res, err := gdp.GetNodeDomain(ctx)

	require.NoError(t, err)
	require.Equal(t, &nodeDomainRef, res)
	require.NotEqual(t, &newNodeDomainRef, res)
}

func TestGenesisDataProvider_GetNodeDomain_Error(t *testing.T) {
	ctx := inslogger.TestContext(t)

	gdp := &GenesisDataProvider{
		CertificateManager: mockCertificateManager(t, nil),
		ContractRequester:  mockContractRequesterWithError(t),
	}

	res, err := gdp.GetNodeDomain(ctx)

	require.Contains(t, err.Error(), "test reasons")
	require.Nil(t, res)
}
