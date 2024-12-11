// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"goformation/v4/cloudformation/types"
	"sync"

	"github.com/weaveworks/eksctl/pkg/vpc"
)

type FakeImporter struct {
	ClusterSecurityGroupStub        func() *types.Value
	clusterSecurityGroupMutex       sync.RWMutex
	clusterSecurityGroupArgsForCall []struct {
	}
	clusterSecurityGroupReturns struct {
		result1 *types.Value
	}
	clusterSecurityGroupReturnsOnCall map[int]struct {
		result1 *types.Value
	}
	ControlPlaneSecurityGroupStub        func() *types.Value
	controlPlaneSecurityGroupMutex       sync.RWMutex
	controlPlaneSecurityGroupArgsForCall []struct {
	}
	controlPlaneSecurityGroupReturns struct {
		result1 *types.Value
	}
	controlPlaneSecurityGroupReturnsOnCall map[int]struct {
		result1 *types.Value
	}
	SecurityGroupsStub        func() types.Slice
	securityGroupsMutex       sync.RWMutex
	securityGroupsArgsForCall []struct {
	}
	securityGroupsReturns struct {
		result1 types.Slice
	}
	securityGroupsReturnsOnCall map[int]struct {
		result1 types.Slice
	}
	SharedNodeSecurityGroupStub        func() *types.Value
	sharedNodeSecurityGroupMutex       sync.RWMutex
	sharedNodeSecurityGroupArgsForCall []struct {
	}
	sharedNodeSecurityGroupReturns struct {
		result1 *types.Value
	}
	sharedNodeSecurityGroupReturnsOnCall map[int]struct {
		result1 *types.Value
	}
	SubnetsPrivateStub        func() *types.Value
	subnetsPrivateMutex       sync.RWMutex
	subnetsPrivateArgsForCall []struct {
	}
	subnetsPrivateReturns struct {
		result1 *types.Value
	}
	subnetsPrivateReturnsOnCall map[int]struct {
		result1 *types.Value
	}
	SubnetsPublicStub        func() *types.Value
	subnetsPublicMutex       sync.RWMutex
	subnetsPublicArgsForCall []struct {
	}
	subnetsPublicReturns struct {
		result1 *types.Value
	}
	subnetsPublicReturnsOnCall map[int]struct {
		result1 *types.Value
	}
	VPCStub        func() *types.Value
	vPCMutex       sync.RWMutex
	vPCArgsForCall []struct {
	}
	vPCReturns struct {
		result1 *types.Value
	}
	vPCReturnsOnCall map[int]struct {
		result1 *types.Value
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImporter) ClusterSecurityGroup() *types.Value {
	fake.clusterSecurityGroupMutex.Lock()
	ret, specificReturn := fake.clusterSecurityGroupReturnsOnCall[len(fake.clusterSecurityGroupArgsForCall)]
	fake.clusterSecurityGroupArgsForCall = append(fake.clusterSecurityGroupArgsForCall, struct {
	}{})
	stub := fake.ClusterSecurityGroupStub
	fakeReturns := fake.clusterSecurityGroupReturns
	fake.recordInvocation("ClusterSecurityGroup", []interface{}{})
	fake.clusterSecurityGroupMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImporter) ClusterSecurityGroupCallCount() int {
	fake.clusterSecurityGroupMutex.RLock()
	defer fake.clusterSecurityGroupMutex.RUnlock()
	return len(fake.clusterSecurityGroupArgsForCall)
}

func (fake *FakeImporter) ClusterSecurityGroupCalls(stub func() *types.Value) {
	fake.clusterSecurityGroupMutex.Lock()
	defer fake.clusterSecurityGroupMutex.Unlock()
	fake.ClusterSecurityGroupStub = stub
}

func (fake *FakeImporter) ClusterSecurityGroupReturns(result1 *types.Value) {
	fake.clusterSecurityGroupMutex.Lock()
	defer fake.clusterSecurityGroupMutex.Unlock()
	fake.ClusterSecurityGroupStub = nil
	fake.clusterSecurityGroupReturns = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) ClusterSecurityGroupReturnsOnCall(i int, result1 *types.Value) {
	fake.clusterSecurityGroupMutex.Lock()
	defer fake.clusterSecurityGroupMutex.Unlock()
	fake.ClusterSecurityGroupStub = nil
	if fake.clusterSecurityGroupReturnsOnCall == nil {
		fake.clusterSecurityGroupReturnsOnCall = make(map[int]struct {
			result1 *types.Value
		})
	}
	fake.clusterSecurityGroupReturnsOnCall[i] = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) ControlPlaneSecurityGroup() *types.Value {
	fake.controlPlaneSecurityGroupMutex.Lock()
	ret, specificReturn := fake.controlPlaneSecurityGroupReturnsOnCall[len(fake.controlPlaneSecurityGroupArgsForCall)]
	fake.controlPlaneSecurityGroupArgsForCall = append(fake.controlPlaneSecurityGroupArgsForCall, struct {
	}{})
	stub := fake.ControlPlaneSecurityGroupStub
	fakeReturns := fake.controlPlaneSecurityGroupReturns
	fake.recordInvocation("ControlPlaneSecurityGroup", []interface{}{})
	fake.controlPlaneSecurityGroupMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImporter) ControlPlaneSecurityGroupCallCount() int {
	fake.controlPlaneSecurityGroupMutex.RLock()
	defer fake.controlPlaneSecurityGroupMutex.RUnlock()
	return len(fake.controlPlaneSecurityGroupArgsForCall)
}

func (fake *FakeImporter) ControlPlaneSecurityGroupCalls(stub func() *types.Value) {
	fake.controlPlaneSecurityGroupMutex.Lock()
	defer fake.controlPlaneSecurityGroupMutex.Unlock()
	fake.ControlPlaneSecurityGroupStub = stub
}

func (fake *FakeImporter) ControlPlaneSecurityGroupReturns(result1 *types.Value) {
	fake.controlPlaneSecurityGroupMutex.Lock()
	defer fake.controlPlaneSecurityGroupMutex.Unlock()
	fake.ControlPlaneSecurityGroupStub = nil
	fake.controlPlaneSecurityGroupReturns = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) ControlPlaneSecurityGroupReturnsOnCall(i int, result1 *types.Value) {
	fake.controlPlaneSecurityGroupMutex.Lock()
	defer fake.controlPlaneSecurityGroupMutex.Unlock()
	fake.ControlPlaneSecurityGroupStub = nil
	if fake.controlPlaneSecurityGroupReturnsOnCall == nil {
		fake.controlPlaneSecurityGroupReturnsOnCall = make(map[int]struct {
			result1 *types.Value
		})
	}
	fake.controlPlaneSecurityGroupReturnsOnCall[i] = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) SecurityGroups() types.Slice {
	fake.securityGroupsMutex.Lock()
	ret, specificReturn := fake.securityGroupsReturnsOnCall[len(fake.securityGroupsArgsForCall)]
	fake.securityGroupsArgsForCall = append(fake.securityGroupsArgsForCall, struct {
	}{})
	stub := fake.SecurityGroupsStub
	fakeReturns := fake.securityGroupsReturns
	fake.recordInvocation("SecurityGroups", []interface{}{})
	fake.securityGroupsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImporter) SecurityGroupsCallCount() int {
	fake.securityGroupsMutex.RLock()
	defer fake.securityGroupsMutex.RUnlock()
	return len(fake.securityGroupsArgsForCall)
}

func (fake *FakeImporter) SecurityGroupsCalls(stub func() types.Slice) {
	fake.securityGroupsMutex.Lock()
	defer fake.securityGroupsMutex.Unlock()
	fake.SecurityGroupsStub = stub
}

func (fake *FakeImporter) SecurityGroupsReturns(result1 types.Slice) {
	fake.securityGroupsMutex.Lock()
	defer fake.securityGroupsMutex.Unlock()
	fake.SecurityGroupsStub = nil
	fake.securityGroupsReturns = struct {
		result1 types.Slice
	}{result1}
}

func (fake *FakeImporter) SecurityGroupsReturnsOnCall(i int, result1 types.Slice) {
	fake.securityGroupsMutex.Lock()
	defer fake.securityGroupsMutex.Unlock()
	fake.SecurityGroupsStub = nil
	if fake.securityGroupsReturnsOnCall == nil {
		fake.securityGroupsReturnsOnCall = make(map[int]struct {
			result1 types.Slice
		})
	}
	fake.securityGroupsReturnsOnCall[i] = struct {
		result1 types.Slice
	}{result1}
}

func (fake *FakeImporter) SharedNodeSecurityGroup() *types.Value {
	fake.sharedNodeSecurityGroupMutex.Lock()
	ret, specificReturn := fake.sharedNodeSecurityGroupReturnsOnCall[len(fake.sharedNodeSecurityGroupArgsForCall)]
	fake.sharedNodeSecurityGroupArgsForCall = append(fake.sharedNodeSecurityGroupArgsForCall, struct {
	}{})
	stub := fake.SharedNodeSecurityGroupStub
	fakeReturns := fake.sharedNodeSecurityGroupReturns
	fake.recordInvocation("SharedNodeSecurityGroup", []interface{}{})
	fake.sharedNodeSecurityGroupMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImporter) SharedNodeSecurityGroupCallCount() int {
	fake.sharedNodeSecurityGroupMutex.RLock()
	defer fake.sharedNodeSecurityGroupMutex.RUnlock()
	return len(fake.sharedNodeSecurityGroupArgsForCall)
}

func (fake *FakeImporter) SharedNodeSecurityGroupCalls(stub func() *types.Value) {
	fake.sharedNodeSecurityGroupMutex.Lock()
	defer fake.sharedNodeSecurityGroupMutex.Unlock()
	fake.SharedNodeSecurityGroupStub = stub
}

func (fake *FakeImporter) SharedNodeSecurityGroupReturns(result1 *types.Value) {
	fake.sharedNodeSecurityGroupMutex.Lock()
	defer fake.sharedNodeSecurityGroupMutex.Unlock()
	fake.SharedNodeSecurityGroupStub = nil
	fake.sharedNodeSecurityGroupReturns = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) SharedNodeSecurityGroupReturnsOnCall(i int, result1 *types.Value) {
	fake.sharedNodeSecurityGroupMutex.Lock()
	defer fake.sharedNodeSecurityGroupMutex.Unlock()
	fake.SharedNodeSecurityGroupStub = nil
	if fake.sharedNodeSecurityGroupReturnsOnCall == nil {
		fake.sharedNodeSecurityGroupReturnsOnCall = make(map[int]struct {
			result1 *types.Value
		})
	}
	fake.sharedNodeSecurityGroupReturnsOnCall[i] = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) SubnetsPrivate() *types.Value {
	fake.subnetsPrivateMutex.Lock()
	ret, specificReturn := fake.subnetsPrivateReturnsOnCall[len(fake.subnetsPrivateArgsForCall)]
	fake.subnetsPrivateArgsForCall = append(fake.subnetsPrivateArgsForCall, struct {
	}{})
	stub := fake.SubnetsPrivateStub
	fakeReturns := fake.subnetsPrivateReturns
	fake.recordInvocation("SubnetsPrivate", []interface{}{})
	fake.subnetsPrivateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImporter) SubnetsPrivateCallCount() int {
	fake.subnetsPrivateMutex.RLock()
	defer fake.subnetsPrivateMutex.RUnlock()
	return len(fake.subnetsPrivateArgsForCall)
}

func (fake *FakeImporter) SubnetsPrivateCalls(stub func() *types.Value) {
	fake.subnetsPrivateMutex.Lock()
	defer fake.subnetsPrivateMutex.Unlock()
	fake.SubnetsPrivateStub = stub
}

func (fake *FakeImporter) SubnetsPrivateReturns(result1 *types.Value) {
	fake.subnetsPrivateMutex.Lock()
	defer fake.subnetsPrivateMutex.Unlock()
	fake.SubnetsPrivateStub = nil
	fake.subnetsPrivateReturns = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) SubnetsPrivateReturnsOnCall(i int, result1 *types.Value) {
	fake.subnetsPrivateMutex.Lock()
	defer fake.subnetsPrivateMutex.Unlock()
	fake.SubnetsPrivateStub = nil
	if fake.subnetsPrivateReturnsOnCall == nil {
		fake.subnetsPrivateReturnsOnCall = make(map[int]struct {
			result1 *types.Value
		})
	}
	fake.subnetsPrivateReturnsOnCall[i] = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) SubnetsPublic() *types.Value {
	fake.subnetsPublicMutex.Lock()
	ret, specificReturn := fake.subnetsPublicReturnsOnCall[len(fake.subnetsPublicArgsForCall)]
	fake.subnetsPublicArgsForCall = append(fake.subnetsPublicArgsForCall, struct {
	}{})
	stub := fake.SubnetsPublicStub
	fakeReturns := fake.subnetsPublicReturns
	fake.recordInvocation("SubnetsPublic", []interface{}{})
	fake.subnetsPublicMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImporter) SubnetsPublicCallCount() int {
	fake.subnetsPublicMutex.RLock()
	defer fake.subnetsPublicMutex.RUnlock()
	return len(fake.subnetsPublicArgsForCall)
}

func (fake *FakeImporter) SubnetsPublicCalls(stub func() *types.Value) {
	fake.subnetsPublicMutex.Lock()
	defer fake.subnetsPublicMutex.Unlock()
	fake.SubnetsPublicStub = stub
}

func (fake *FakeImporter) SubnetsPublicReturns(result1 *types.Value) {
	fake.subnetsPublicMutex.Lock()
	defer fake.subnetsPublicMutex.Unlock()
	fake.SubnetsPublicStub = nil
	fake.subnetsPublicReturns = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) SubnetsPublicReturnsOnCall(i int, result1 *types.Value) {
	fake.subnetsPublicMutex.Lock()
	defer fake.subnetsPublicMutex.Unlock()
	fake.SubnetsPublicStub = nil
	if fake.subnetsPublicReturnsOnCall == nil {
		fake.subnetsPublicReturnsOnCall = make(map[int]struct {
			result1 *types.Value
		})
	}
	fake.subnetsPublicReturnsOnCall[i] = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) VPC() *types.Value {
	fake.vPCMutex.Lock()
	ret, specificReturn := fake.vPCReturnsOnCall[len(fake.vPCArgsForCall)]
	fake.vPCArgsForCall = append(fake.vPCArgsForCall, struct {
	}{})
	stub := fake.VPCStub
	fakeReturns := fake.vPCReturns
	fake.recordInvocation("VPC", []interface{}{})
	fake.vPCMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImporter) VPCCallCount() int {
	fake.vPCMutex.RLock()
	defer fake.vPCMutex.RUnlock()
	return len(fake.vPCArgsForCall)
}

func (fake *FakeImporter) VPCCalls(stub func() *types.Value) {
	fake.vPCMutex.Lock()
	defer fake.vPCMutex.Unlock()
	fake.VPCStub = stub
}

func (fake *FakeImporter) VPCReturns(result1 *types.Value) {
	fake.vPCMutex.Lock()
	defer fake.vPCMutex.Unlock()
	fake.VPCStub = nil
	fake.vPCReturns = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) VPCReturnsOnCall(i int, result1 *types.Value) {
	fake.vPCMutex.Lock()
	defer fake.vPCMutex.Unlock()
	fake.VPCStub = nil
	if fake.vPCReturnsOnCall == nil {
		fake.vPCReturnsOnCall = make(map[int]struct {
			result1 *types.Value
		})
	}
	fake.vPCReturnsOnCall[i] = struct {
		result1 *types.Value
	}{result1}
}

func (fake *FakeImporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clusterSecurityGroupMutex.RLock()
	defer fake.clusterSecurityGroupMutex.RUnlock()
	fake.controlPlaneSecurityGroupMutex.RLock()
	defer fake.controlPlaneSecurityGroupMutex.RUnlock()
	fake.securityGroupsMutex.RLock()
	defer fake.securityGroupsMutex.RUnlock()
	fake.sharedNodeSecurityGroupMutex.RLock()
	defer fake.sharedNodeSecurityGroupMutex.RUnlock()
	fake.subnetsPrivateMutex.RLock()
	defer fake.subnetsPrivateMutex.RUnlock()
	fake.subnetsPublicMutex.RLock()
	defer fake.subnetsPublicMutex.RUnlock()
	fake.vPCMutex.RLock()
	defer fake.vPCMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImporter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ vpc.Importer = new(FakeImporter)
