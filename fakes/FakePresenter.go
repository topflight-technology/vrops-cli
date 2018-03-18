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
	PresentResourcesStub        func(models.Resources)
	presentResourcesMutex       sync.RWMutex
	presentResourcesArgsForCall []struct {
		arg1 models.Resources
	}
	PresentStatsStub        func(models.ListStatsResponseValuesStatListStats)
	presentStatsMutex       sync.RWMutex
	presentStatsArgsForCall []struct {
		arg1 models.ListStatsResponseValuesStatListStats
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

func (fake *FakePresenter) PresentResources(arg1 models.Resources) {
	fake.presentResourcesMutex.Lock()
	fake.presentResourcesArgsForCall = append(fake.presentResourcesArgsForCall, struct {
		arg1 models.Resources
	}{arg1})
	fake.recordInvocation("PresentResources", []interface{}{arg1})
	fake.presentResourcesMutex.Unlock()
	if fake.PresentResourcesStub != nil {
		fake.PresentResourcesStub(arg1)
	}
}

func (fake *FakePresenter) PresentResourcesCallCount() int {
	fake.presentResourcesMutex.RLock()
	defer fake.presentResourcesMutex.RUnlock()
	return len(fake.presentResourcesArgsForCall)
}

func (fake *FakePresenter) PresentResourcesArgsForCall(i int) models.Resources {
	fake.presentResourcesMutex.RLock()
	defer fake.presentResourcesMutex.RUnlock()
	return fake.presentResourcesArgsForCall[i].arg1
}

func (fake *FakePresenter) PresentStats(arg1 models.ListStatsResponseValuesStatListStats) {
	fake.presentStatsMutex.Lock()
	fake.presentStatsArgsForCall = append(fake.presentStatsArgsForCall, struct {
		arg1 models.ListStatsResponseValuesStatListStats
	}{arg1})
	fake.recordInvocation("PresentStats", []interface{}{arg1})
	fake.presentStatsMutex.Unlock()
	if fake.PresentStatsStub != nil {
		fake.PresentStatsStub(arg1)
	}
}

func (fake *FakePresenter) PresentStatsCallCount() int {
	fake.presentStatsMutex.RLock()
	defer fake.presentStatsMutex.RUnlock()
	return len(fake.presentStatsArgsForCall)
}

func (fake *FakePresenter) PresentStatsArgsForCall(i int) models.ListStatsResponseValuesStatListStats {
	fake.presentStatsMutex.RLock()
	defer fake.presentStatsMutex.RUnlock()
	return fake.presentStatsArgsForCall[i].arg1
}

func (fake *FakePresenter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.presentAdapterKindsMutex.RLock()
	defer fake.presentAdapterKindsMutex.RUnlock()
	fake.presentResourceKindsMutex.RLock()
	defer fake.presentResourceKindsMutex.RUnlock()
	fake.presentResourcesMutex.RLock()
	defer fake.presentResourcesMutex.RUnlock()
	fake.presentStatsMutex.RLock()
	defer fake.presentStatsMutex.RUnlock()
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
