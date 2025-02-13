// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/weaveworks/eksctl/pkg/actions/podidentityassociation"
	"github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/cfn/manager"
)

type FakeStackUpdater struct {
	DescribeStackStub        func(context.Context, *manager.Stack) (*manager.Stack, error)
	describeStackMutex       sync.RWMutex
	describeStackArgsForCall []struct {
		arg1 context.Context
		arg2 *manager.Stack
	}
	describeStackReturns struct {
		result1 *manager.Stack
		result2 error
	}
	describeStackReturnsOnCall map[int]struct {
		result1 *manager.Stack
		result2 error
	}
	GetIAMServiceAccountsStub        func(context.Context) ([]*v1alpha5.ClusterIAMServiceAccount, error)
	getIAMServiceAccountsMutex       sync.RWMutex
	getIAMServiceAccountsArgsForCall []struct {
		arg1 context.Context
	}
	getIAMServiceAccountsReturns struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}
	getIAMServiceAccountsReturnsOnCall map[int]struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}
	GetStackTemplateStub        func(context.Context, string) (string, error)
	getStackTemplateMutex       sync.RWMutex
	getStackTemplateArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStackTemplateReturns struct {
		result1 string
		result2 error
	}
	getStackTemplateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListPodIdentityStackNamesStub        func(context.Context) ([]string, error)
	listPodIdentityStackNamesMutex       sync.RWMutex
	listPodIdentityStackNamesArgsForCall []struct {
		arg1 context.Context
	}
	listPodIdentityStackNamesReturns struct {
		result1 []string
		result2 error
	}
	listPodIdentityStackNamesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	MustUpdateStackStub        func(context.Context, manager.UpdateStackOptions) error
	mustUpdateStackMutex       sync.RWMutex
	mustUpdateStackArgsForCall []struct {
		arg1 context.Context
		arg2 manager.UpdateStackOptions
	}
	mustUpdateStackReturns struct {
		result1 error
	}
	mustUpdateStackReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStackUpdater) DescribeStack(arg1 context.Context, arg2 *manager.Stack) (*manager.Stack, error) {
	fake.describeStackMutex.Lock()
	ret, specificReturn := fake.describeStackReturnsOnCall[len(fake.describeStackArgsForCall)]
	fake.describeStackArgsForCall = append(fake.describeStackArgsForCall, struct {
		arg1 context.Context
		arg2 *manager.Stack
	}{arg1, arg2})
	stub := fake.DescribeStackStub
	fakeReturns := fake.describeStackReturns
	fake.recordInvocation("DescribeStack", []interface{}{arg1, arg2})
	fake.describeStackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackUpdater) DescribeStackCallCount() int {
	fake.describeStackMutex.RLock()
	defer fake.describeStackMutex.RUnlock()
	return len(fake.describeStackArgsForCall)
}

func (fake *FakeStackUpdater) DescribeStackCalls(stub func(context.Context, *manager.Stack) (*manager.Stack, error)) {
	fake.describeStackMutex.Lock()
	defer fake.describeStackMutex.Unlock()
	fake.DescribeStackStub = stub
}

func (fake *FakeStackUpdater) DescribeStackArgsForCall(i int) (context.Context, *manager.Stack) {
	fake.describeStackMutex.RLock()
	defer fake.describeStackMutex.RUnlock()
	argsForCall := fake.describeStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackUpdater) DescribeStackReturns(result1 *manager.Stack, result2 error) {
	fake.describeStackMutex.Lock()
	defer fake.describeStackMutex.Unlock()
	fake.DescribeStackStub = nil
	fake.describeStackReturns = struct {
		result1 *manager.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) DescribeStackReturnsOnCall(i int, result1 *manager.Stack, result2 error) {
	fake.describeStackMutex.Lock()
	defer fake.describeStackMutex.Unlock()
	fake.DescribeStackStub = nil
	if fake.describeStackReturnsOnCall == nil {
		fake.describeStackReturnsOnCall = make(map[int]struct {
			result1 *manager.Stack
			result2 error
		})
	}
	fake.describeStackReturnsOnCall[i] = struct {
		result1 *manager.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) GetIAMServiceAccounts(arg1 context.Context) ([]*v1alpha5.ClusterIAMServiceAccount, error) {
	fake.getIAMServiceAccountsMutex.Lock()
	ret, specificReturn := fake.getIAMServiceAccountsReturnsOnCall[len(fake.getIAMServiceAccountsArgsForCall)]
	fake.getIAMServiceAccountsArgsForCall = append(fake.getIAMServiceAccountsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetIAMServiceAccountsStub
	fakeReturns := fake.getIAMServiceAccountsReturns
	fake.recordInvocation("GetIAMServiceAccounts", []interface{}{arg1})
	fake.getIAMServiceAccountsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackUpdater) GetIAMServiceAccountsCallCount() int {
	fake.getIAMServiceAccountsMutex.RLock()
	defer fake.getIAMServiceAccountsMutex.RUnlock()
	return len(fake.getIAMServiceAccountsArgsForCall)
}

func (fake *FakeStackUpdater) GetIAMServiceAccountsCalls(stub func(context.Context) ([]*v1alpha5.ClusterIAMServiceAccount, error)) {
	fake.getIAMServiceAccountsMutex.Lock()
	defer fake.getIAMServiceAccountsMutex.Unlock()
	fake.GetIAMServiceAccountsStub = stub
}

func (fake *FakeStackUpdater) GetIAMServiceAccountsArgsForCall(i int) context.Context {
	fake.getIAMServiceAccountsMutex.RLock()
	defer fake.getIAMServiceAccountsMutex.RUnlock()
	argsForCall := fake.getIAMServiceAccountsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStackUpdater) GetIAMServiceAccountsReturns(result1 []*v1alpha5.ClusterIAMServiceAccount, result2 error) {
	fake.getIAMServiceAccountsMutex.Lock()
	defer fake.getIAMServiceAccountsMutex.Unlock()
	fake.GetIAMServiceAccountsStub = nil
	fake.getIAMServiceAccountsReturns = struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) GetIAMServiceAccountsReturnsOnCall(i int, result1 []*v1alpha5.ClusterIAMServiceAccount, result2 error) {
	fake.getIAMServiceAccountsMutex.Lock()
	defer fake.getIAMServiceAccountsMutex.Unlock()
	fake.GetIAMServiceAccountsStub = nil
	if fake.getIAMServiceAccountsReturnsOnCall == nil {
		fake.getIAMServiceAccountsReturnsOnCall = make(map[int]struct {
			result1 []*v1alpha5.ClusterIAMServiceAccount
			result2 error
		})
	}
	fake.getIAMServiceAccountsReturnsOnCall[i] = struct {
		result1 []*v1alpha5.ClusterIAMServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) GetStackTemplate(arg1 context.Context, arg2 string) (string, error) {
	fake.getStackTemplateMutex.Lock()
	ret, specificReturn := fake.getStackTemplateReturnsOnCall[len(fake.getStackTemplateArgsForCall)]
	fake.getStackTemplateArgsForCall = append(fake.getStackTemplateArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStackTemplateStub
	fakeReturns := fake.getStackTemplateReturns
	fake.recordInvocation("GetStackTemplate", []interface{}{arg1, arg2})
	fake.getStackTemplateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackUpdater) GetStackTemplateCallCount() int {
	fake.getStackTemplateMutex.RLock()
	defer fake.getStackTemplateMutex.RUnlock()
	return len(fake.getStackTemplateArgsForCall)
}

func (fake *FakeStackUpdater) GetStackTemplateCalls(stub func(context.Context, string) (string, error)) {
	fake.getStackTemplateMutex.Lock()
	defer fake.getStackTemplateMutex.Unlock()
	fake.GetStackTemplateStub = stub
}

func (fake *FakeStackUpdater) GetStackTemplateArgsForCall(i int) (context.Context, string) {
	fake.getStackTemplateMutex.RLock()
	defer fake.getStackTemplateMutex.RUnlock()
	argsForCall := fake.getStackTemplateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackUpdater) GetStackTemplateReturns(result1 string, result2 error) {
	fake.getStackTemplateMutex.Lock()
	defer fake.getStackTemplateMutex.Unlock()
	fake.GetStackTemplateStub = nil
	fake.getStackTemplateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) GetStackTemplateReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStackTemplateMutex.Lock()
	defer fake.getStackTemplateMutex.Unlock()
	fake.GetStackTemplateStub = nil
	if fake.getStackTemplateReturnsOnCall == nil {
		fake.getStackTemplateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStackTemplateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) ListPodIdentityStackNames(arg1 context.Context) ([]string, error) {
	fake.listPodIdentityStackNamesMutex.Lock()
	ret, specificReturn := fake.listPodIdentityStackNamesReturnsOnCall[len(fake.listPodIdentityStackNamesArgsForCall)]
	fake.listPodIdentityStackNamesArgsForCall = append(fake.listPodIdentityStackNamesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListPodIdentityStackNamesStub
	fakeReturns := fake.listPodIdentityStackNamesReturns
	fake.recordInvocation("ListPodIdentityStackNames", []interface{}{arg1})
	fake.listPodIdentityStackNamesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackUpdater) ListPodIdentityStackNamesCallCount() int {
	fake.listPodIdentityStackNamesMutex.RLock()
	defer fake.listPodIdentityStackNamesMutex.RUnlock()
	return len(fake.listPodIdentityStackNamesArgsForCall)
}

func (fake *FakeStackUpdater) ListPodIdentityStackNamesCalls(stub func(context.Context) ([]string, error)) {
	fake.listPodIdentityStackNamesMutex.Lock()
	defer fake.listPodIdentityStackNamesMutex.Unlock()
	fake.ListPodIdentityStackNamesStub = stub
}

func (fake *FakeStackUpdater) ListPodIdentityStackNamesArgsForCall(i int) context.Context {
	fake.listPodIdentityStackNamesMutex.RLock()
	defer fake.listPodIdentityStackNamesMutex.RUnlock()
	argsForCall := fake.listPodIdentityStackNamesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStackUpdater) ListPodIdentityStackNamesReturns(result1 []string, result2 error) {
	fake.listPodIdentityStackNamesMutex.Lock()
	defer fake.listPodIdentityStackNamesMutex.Unlock()
	fake.ListPodIdentityStackNamesStub = nil
	fake.listPodIdentityStackNamesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) ListPodIdentityStackNamesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listPodIdentityStackNamesMutex.Lock()
	defer fake.listPodIdentityStackNamesMutex.Unlock()
	fake.ListPodIdentityStackNamesStub = nil
	if fake.listPodIdentityStackNamesReturnsOnCall == nil {
		fake.listPodIdentityStackNamesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listPodIdentityStackNamesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackUpdater) MustUpdateStack(arg1 context.Context, arg2 manager.UpdateStackOptions) error {
	fake.mustUpdateStackMutex.Lock()
	ret, specificReturn := fake.mustUpdateStackReturnsOnCall[len(fake.mustUpdateStackArgsForCall)]
	fake.mustUpdateStackArgsForCall = append(fake.mustUpdateStackArgsForCall, struct {
		arg1 context.Context
		arg2 manager.UpdateStackOptions
	}{arg1, arg2})
	stub := fake.MustUpdateStackStub
	fakeReturns := fake.mustUpdateStackReturns
	fake.recordInvocation("MustUpdateStack", []interface{}{arg1, arg2})
	fake.mustUpdateStackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStackUpdater) MustUpdateStackCallCount() int {
	fake.mustUpdateStackMutex.RLock()
	defer fake.mustUpdateStackMutex.RUnlock()
	return len(fake.mustUpdateStackArgsForCall)
}

func (fake *FakeStackUpdater) MustUpdateStackCalls(stub func(context.Context, manager.UpdateStackOptions) error) {
	fake.mustUpdateStackMutex.Lock()
	defer fake.mustUpdateStackMutex.Unlock()
	fake.MustUpdateStackStub = stub
}

func (fake *FakeStackUpdater) MustUpdateStackArgsForCall(i int) (context.Context, manager.UpdateStackOptions) {
	fake.mustUpdateStackMutex.RLock()
	defer fake.mustUpdateStackMutex.RUnlock()
	argsForCall := fake.mustUpdateStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackUpdater) MustUpdateStackReturns(result1 error) {
	fake.mustUpdateStackMutex.Lock()
	defer fake.mustUpdateStackMutex.Unlock()
	fake.MustUpdateStackStub = nil
	fake.mustUpdateStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackUpdater) MustUpdateStackReturnsOnCall(i int, result1 error) {
	fake.mustUpdateStackMutex.Lock()
	defer fake.mustUpdateStackMutex.Unlock()
	fake.MustUpdateStackStub = nil
	if fake.mustUpdateStackReturnsOnCall == nil {
		fake.mustUpdateStackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mustUpdateStackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackUpdater) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.describeStackMutex.RLock()
	defer fake.describeStackMutex.RUnlock()
	fake.getIAMServiceAccountsMutex.RLock()
	defer fake.getIAMServiceAccountsMutex.RUnlock()
	fake.getStackTemplateMutex.RLock()
	defer fake.getStackTemplateMutex.RUnlock()
	fake.listPodIdentityStackNamesMutex.RLock()
	defer fake.listPodIdentityStackNamesMutex.RUnlock()
	fake.mustUpdateStackMutex.RLock()
	defer fake.mustUpdateStackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStackUpdater) recordInvocation(key string, args []interface{}) {
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

var _ podidentityassociation.StackUpdater = new(FakeStackUpdater)
