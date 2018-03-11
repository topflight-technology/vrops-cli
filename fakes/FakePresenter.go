// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/rcw5/vrops-cli/models"
	"github.com/rcw5/vrops-cli/presenters"
)

type FakePresenter struct {
	PresentAdapterKindsStub        func(models.AdapterKinds)
	presentAdapterKindsMutex       sync.RWMutex
	presentAdapterKindsArgsForCall []struct {
		arg1 models.AdapterKinds
	}
	PresentResourceKindsStub        func([]string)
	presentResourceKindsMutex       sync.RWMutex
	presentResourceKindsArgsForCall []struct {
		arg1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePresenter) PresentAdapterKinds(arg1 models.AdapterKinds) {
	fake.presentAdapterKindsMutex.Lock()
	fake.presentAdapterKindsArgsForCall = append(fake.presentAdapterKindsArgsForCall, struct {
		arg1 models.AdapterKinds
	}{arg1})
	fake.recordInvocation("PresentAdapterKinds", []interface{}{arg1})
	fake.presentAdapterKindsMutex.Unlock()
	if fake.PresentAdapterKindsStub != nil {
		fake.PresentAdapterKindsStub(arg1)
	}
}

func (fake *FakePresenter) PresentAdapterKindsCallCount() int {
	fake.presentAdapterKindsMutex.RLock()
	defer fake.presentAdapterKindsMutex.RUnlock()
	return len(fake.presentAdapterKindsArgsForCall)
}

func (fake *FakePresenter) PresentAdapterKindsArgsForCall(i int) models.AdapterKinds {
	fake.presentAdapterKindsMutex.RLock()
	defer fake.presentAdapterKindsMutex.RUnlock()
	return fake.presentAdapterKindsArgsForCall[i].arg1
}

func (fake *FakePresenter) PresentResourceKinds(arg1 []string) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentResourceKindsMutex.Lock()
	fake.presentResourceKindsArgsForCall = append(fake.presentResourceKindsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("PresentResourceKinds", []interface{}{arg1Copy})
	fake.presentResourceKindsMutex.Unlock()
	if fake.PresentResourceKindsStub != nil {
		fake.PresentResourceKindsStub(arg1)
	}
}

func (fake *FakePresenter) PresentResourceKindsCallCount() int {
	fake.presentResourceKindsMutex.RLock()
	defer fake.presentResourceKindsMutex.RUnlock()
	return len(fake.presentResourceKindsArgsForCall)
}

func (fake *FakePresenter) PresentResourceKindsArgsForCall(i int) []string {
	fake.presentResourceKindsMutex.RLock()
	defer fake.presentResourceKindsMutex.RUnlock()
	return fake.presentResourceKindsArgsForCall[i].arg1
}

func (fake *FakePresenter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.presentAdapterKindsMutex.RLock()
	defer fake.presentAdapterKindsMutex.RUnlock()
	fake.presentResourceKindsMutex.RLock()
	defer fake.presentResourceKindsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePresenter) recordInvocation(key string, args []interface{}) {
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

var _ presenters.Presenter = new(FakePresenter)