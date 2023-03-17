// Code generated by counterfeiter. DO NOT EDIT.
package execrunnerfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/execrunner"
	lager "code.cloudfoundry.org/lager/v3"
)

type FakeProcessDepot struct {
	CreateProcessDirStub        func(lager.Logger, string, string) (string, error)
	createProcessDirMutex       sync.RWMutex
	createProcessDirArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	createProcessDirReturns struct {
		result1 string
		result2 error
	}
	createProcessDirReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	LookupProcessDirStub        func(lager.Logger, string, string) (string, error)
	lookupProcessDirMutex       sync.RWMutex
	lookupProcessDirArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	lookupProcessDirReturns struct {
		result1 string
		result2 error
	}
	lookupProcessDirReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessDepot) CreateProcessDir(arg1 lager.Logger, arg2 string, arg3 string) (string, error) {
	fake.createProcessDirMutex.Lock()
	ret, specificReturn := fake.createProcessDirReturnsOnCall[len(fake.createProcessDirArgsForCall)]
	fake.createProcessDirArgsForCall = append(fake.createProcessDirArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateProcessDirStub
	fakeReturns := fake.createProcessDirReturns
	fake.recordInvocation("CreateProcessDir", []interface{}{arg1, arg2, arg3})
	fake.createProcessDirMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcessDepot) CreateProcessDirCallCount() int {
	fake.createProcessDirMutex.RLock()
	defer fake.createProcessDirMutex.RUnlock()
	return len(fake.createProcessDirArgsForCall)
}

func (fake *FakeProcessDepot) CreateProcessDirCalls(stub func(lager.Logger, string, string) (string, error)) {
	fake.createProcessDirMutex.Lock()
	defer fake.createProcessDirMutex.Unlock()
	fake.CreateProcessDirStub = stub
}

func (fake *FakeProcessDepot) CreateProcessDirArgsForCall(i int) (lager.Logger, string, string) {
	fake.createProcessDirMutex.RLock()
	defer fake.createProcessDirMutex.RUnlock()
	argsForCall := fake.createProcessDirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeProcessDepot) CreateProcessDirReturns(result1 string, result2 error) {
	fake.createProcessDirMutex.Lock()
	defer fake.createProcessDirMutex.Unlock()
	fake.CreateProcessDirStub = nil
	fake.createProcessDirReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) CreateProcessDirReturnsOnCall(i int, result1 string, result2 error) {
	fake.createProcessDirMutex.Lock()
	defer fake.createProcessDirMutex.Unlock()
	fake.CreateProcessDirStub = nil
	if fake.createProcessDirReturnsOnCall == nil {
		fake.createProcessDirReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createProcessDirReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) LookupProcessDir(arg1 lager.Logger, arg2 string, arg3 string) (string, error) {
	fake.lookupProcessDirMutex.Lock()
	ret, specificReturn := fake.lookupProcessDirReturnsOnCall[len(fake.lookupProcessDirArgsForCall)]
	fake.lookupProcessDirArgsForCall = append(fake.lookupProcessDirArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.LookupProcessDirStub
	fakeReturns := fake.lookupProcessDirReturns
	fake.recordInvocation("LookupProcessDir", []interface{}{arg1, arg2, arg3})
	fake.lookupProcessDirMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcessDepot) LookupProcessDirCallCount() int {
	fake.lookupProcessDirMutex.RLock()
	defer fake.lookupProcessDirMutex.RUnlock()
	return len(fake.lookupProcessDirArgsForCall)
}

func (fake *FakeProcessDepot) LookupProcessDirCalls(stub func(lager.Logger, string, string) (string, error)) {
	fake.lookupProcessDirMutex.Lock()
	defer fake.lookupProcessDirMutex.Unlock()
	fake.LookupProcessDirStub = stub
}

func (fake *FakeProcessDepot) LookupProcessDirArgsForCall(i int) (lager.Logger, string, string) {
	fake.lookupProcessDirMutex.RLock()
	defer fake.lookupProcessDirMutex.RUnlock()
	argsForCall := fake.lookupProcessDirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeProcessDepot) LookupProcessDirReturns(result1 string, result2 error) {
	fake.lookupProcessDirMutex.Lock()
	defer fake.lookupProcessDirMutex.Unlock()
	fake.LookupProcessDirStub = nil
	fake.lookupProcessDirReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) LookupProcessDirReturnsOnCall(i int, result1 string, result2 error) {
	fake.lookupProcessDirMutex.Lock()
	defer fake.lookupProcessDirMutex.Unlock()
	fake.LookupProcessDirStub = nil
	if fake.lookupProcessDirReturnsOnCall == nil {
		fake.lookupProcessDirReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lookupProcessDirReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessDepot) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createProcessDirMutex.RLock()
	defer fake.createProcessDirMutex.RUnlock()
	fake.lookupProcessDirMutex.RLock()
	defer fake.lookupProcessDirMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcessDepot) recordInvocation(key string, args []interface{}) {
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

var _ execrunner.ProcessDepot = new(FakeProcessDepot)
