// Code generated by counterfeiter. DO NOT EDIT.
package usersfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/users"
)

type FakeUserLookupper struct {
	LookupStub        func(string, string) (*users.ExecUser, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		arg1 string
		arg2 string
	}
	lookupReturns struct {
		result1 *users.ExecUser
		result2 error
	}
	lookupReturnsOnCall map[int]struct {
		result1 *users.ExecUser
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserLookupper) Lookup(arg1 string, arg2 string) (*users.ExecUser, error) {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.LookupStub
	fakeReturns := fake.lookupReturns
	fake.recordInvocation("Lookup", []interface{}{arg1, arg2})
	fake.lookupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserLookupper) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeUserLookupper) LookupCalls(stub func(string, string) (*users.ExecUser, error)) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = stub
}

func (fake *FakeUserLookupper) LookupArgsForCall(i int) (string, string) {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	argsForCall := fake.lookupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserLookupper) LookupReturns(result1 *users.ExecUser, result2 error) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 *users.ExecUser
		result2 error
	}{result1, result2}
}

func (fake *FakeUserLookupper) LookupReturnsOnCall(i int, result1 *users.ExecUser, result2 error) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 *users.ExecUser
			result2 error
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 *users.ExecUser
		result2 error
	}{result1, result2}
}

func (fake *FakeUserLookupper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserLookupper) recordInvocation(key string, args []interface{}) {
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

var _ users.UserLookupper = new(FakeUserLookupper)
