// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/executor"
	"github.com/cloudfoundry-incubator/executor/depot/metrics"
	"github.com/cloudfoundry-incubator/executor/depot/registry"
)

type FakeExecutorSource struct {
	CurrentCapacityStub        func() registry.Capacity
	currentCapacityMutex       sync.RWMutex
	currentCapacityArgsForCall []struct{}
	currentCapacityReturns struct {
		result1 registry.Capacity
	}
	TotalCapacityStub        func() registry.Capacity
	totalCapacityMutex       sync.RWMutex
	totalCapacityArgsForCall []struct{}
	totalCapacityReturns struct {
		result1 registry.Capacity
	}
	GetAllContainersStub        func() []executor.Container
	getAllContainersMutex       sync.RWMutex
	getAllContainersArgsForCall []struct{}
	getAllContainersReturns struct {
		result1 []executor.Container
	}
}

func (fake *FakeExecutorSource) CurrentCapacity() registry.Capacity {
	fake.currentCapacityMutex.Lock()
	fake.currentCapacityArgsForCall = append(fake.currentCapacityArgsForCall, struct{}{})
	fake.currentCapacityMutex.Unlock()
	if fake.CurrentCapacityStub != nil {
		return fake.CurrentCapacityStub()
	} else {
		return fake.currentCapacityReturns.result1
	}
}

func (fake *FakeExecutorSource) CurrentCapacityCallCount() int {
	fake.currentCapacityMutex.RLock()
	defer fake.currentCapacityMutex.RUnlock()
	return len(fake.currentCapacityArgsForCall)
}

func (fake *FakeExecutorSource) CurrentCapacityReturns(result1 registry.Capacity) {
	fake.CurrentCapacityStub = nil
	fake.currentCapacityReturns = struct {
		result1 registry.Capacity
	}{result1}
}

func (fake *FakeExecutorSource) TotalCapacity() registry.Capacity {
	fake.totalCapacityMutex.Lock()
	fake.totalCapacityArgsForCall = append(fake.totalCapacityArgsForCall, struct{}{})
	fake.totalCapacityMutex.Unlock()
	if fake.TotalCapacityStub != nil {
		return fake.TotalCapacityStub()
	} else {
		return fake.totalCapacityReturns.result1
	}
}

func (fake *FakeExecutorSource) TotalCapacityCallCount() int {
	fake.totalCapacityMutex.RLock()
	defer fake.totalCapacityMutex.RUnlock()
	return len(fake.totalCapacityArgsForCall)
}

func (fake *FakeExecutorSource) TotalCapacityReturns(result1 registry.Capacity) {
	fake.TotalCapacityStub = nil
	fake.totalCapacityReturns = struct {
		result1 registry.Capacity
	}{result1}
}

func (fake *FakeExecutorSource) GetAllContainers() []executor.Container {
	fake.getAllContainersMutex.Lock()
	fake.getAllContainersArgsForCall = append(fake.getAllContainersArgsForCall, struct{}{})
	fake.getAllContainersMutex.Unlock()
	if fake.GetAllContainersStub != nil {
		return fake.GetAllContainersStub()
	} else {
		return fake.getAllContainersReturns.result1
	}
}

func (fake *FakeExecutorSource) GetAllContainersCallCount() int {
	fake.getAllContainersMutex.RLock()
	defer fake.getAllContainersMutex.RUnlock()
	return len(fake.getAllContainersArgsForCall)
}

func (fake *FakeExecutorSource) GetAllContainersReturns(result1 []executor.Container) {
	fake.GetAllContainersStub = nil
	fake.getAllContainersReturns = struct {
		result1 []executor.Container
	}{result1}
}

var _ metrics.ExecutorSource = new(FakeExecutorSource)