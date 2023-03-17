// Code generated by counterfeiter. DO NOT EDIT.
package rundmcfakes

import (
	"sync"

	spec "code.cloudfoundry.org/guardian/gardener/container-spec"
	"code.cloudfoundry.org/guardian/rundmc"
	"code.cloudfoundry.org/guardian/rundmc/goci"
)

type FakeBundlerRule struct {
	ApplyStub        func(goci.Bndl, spec.DesiredContainerSpec) (goci.Bndl, error)
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 goci.Bndl
		arg2 spec.DesiredContainerSpec
	}
	applyReturns struct {
		result1 goci.Bndl
		result2 error
	}
	applyReturnsOnCall map[int]struct {
		result1 goci.Bndl
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBundlerRule) Apply(arg1 goci.Bndl, arg2 spec.DesiredContainerSpec) (goci.Bndl, error) {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 goci.Bndl
		arg2 spec.DesiredContainerSpec
	}{arg1, arg2})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1, arg2})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBundlerRule) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeBundlerRule) ApplyCalls(stub func(goci.Bndl, spec.DesiredContainerSpec) (goci.Bndl, error)) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeBundlerRule) ApplyArgsForCall(i int) (goci.Bndl, spec.DesiredContainerSpec) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBundlerRule) ApplyReturns(result1 goci.Bndl, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 goci.Bndl
		result2 error
	}{result1, result2}
}

func (fake *FakeBundlerRule) ApplyReturnsOnCall(i int, result1 goci.Bndl, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 goci.Bndl
			result2 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 goci.Bndl
		result2 error
	}{result1, result2}
}

func (fake *FakeBundlerRule) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBundlerRule) recordInvocation(key string, args []interface{}) {
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

var _ rundmc.BundlerRule = new(FakeBundlerRule)
