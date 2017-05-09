// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/pivotal-cf/on-demand-service-broker/services"
	"github.com/pivotal-cf/on-demand-service-broker/upgrader"
)

type FakeListener struct {
	StartingStub                  func()
	startingMutex                 sync.RWMutex
	startingArgsForCall           []struct{}
	InstancesToUpgradeStub        func(instances []string)
	instancesToUpgradeMutex       sync.RWMutex
	instancesToUpgradeArgsForCall []struct {
		instances []string
	}
	InstanceUpgradeStartingStub        func(instance string, index, totalInstances int)
	instanceUpgradeStartingMutex       sync.RWMutex
	instanceUpgradeStartingArgsForCall []struct {
		instance       string
		index          int
		totalInstances int
	}
	InstanceUpgradeStartResultStub        func(status services.UpgradeOperationType)
	instanceUpgradeStartResultMutex       sync.RWMutex
	instanceUpgradeStartResultArgsForCall []struct {
		status services.UpgradeOperationType
	}
	InstanceUpgradedStub        func(instance string, result string)
	instanceUpgradedMutex       sync.RWMutex
	instanceUpgradedArgsForCall []struct {
		instance string
		result   string
	}
	WaitingForStub        func(instance string, boshTaskId int)
	waitingForMutex       sync.RWMutex
	waitingForArgsForCall []struct {
		instance   string
		boshTaskId int
	}
	ProgressStub        func(pollingInterval time.Duration, orphanCount, upgradedCount, upgradesLeftCount, deletedCount int)
	progressMutex       sync.RWMutex
	progressArgsForCall []struct {
		pollingInterval   time.Duration
		orphanCount       int
		upgradedCount     int
		upgradesLeftCount int
		deletedCount      int
	}
	FinishedStub        func(orphanCount, upgradedCount, deletedCount int)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		orphanCount   int
		upgradedCount int
		deletedCount  int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeListener) Starting() {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct{}{})
	fake.recordInvocation("Starting", []interface{}{})
	fake.startingMutex.Unlock()
	if fake.StartingStub != nil {
		fake.StartingStub()
	}
}

func (fake *FakeListener) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakeListener) InstancesToUpgrade(instances []string) {
	var instancesCopy []string
	if instances != nil {
		instancesCopy = make([]string, len(instances))
		copy(instancesCopy, instances)
	}
	fake.instancesToUpgradeMutex.Lock()
	fake.instancesToUpgradeArgsForCall = append(fake.instancesToUpgradeArgsForCall, struct {
		instances []string
	}{instancesCopy})
	fake.recordInvocation("InstancesToUpgrade", []interface{}{instancesCopy})
	fake.instancesToUpgradeMutex.Unlock()
	if fake.InstancesToUpgradeStub != nil {
		fake.InstancesToUpgradeStub(instances)
	}
}

func (fake *FakeListener) InstancesToUpgradeCallCount() int {
	fake.instancesToUpgradeMutex.RLock()
	defer fake.instancesToUpgradeMutex.RUnlock()
	return len(fake.instancesToUpgradeArgsForCall)
}

func (fake *FakeListener) InstancesToUpgradeArgsForCall(i int) []string {
	fake.instancesToUpgradeMutex.RLock()
	defer fake.instancesToUpgradeMutex.RUnlock()
	return fake.instancesToUpgradeArgsForCall[i].instances
}

func (fake *FakeListener) InstanceUpgradeStarting(instance string, index int, totalInstances int) {
	fake.instanceUpgradeStartingMutex.Lock()
	fake.instanceUpgradeStartingArgsForCall = append(fake.instanceUpgradeStartingArgsForCall, struct {
		instance       string
		index          int
		totalInstances int
	}{instance, index, totalInstances})
	fake.recordInvocation("InstanceUpgradeStarting", []interface{}{instance, index, totalInstances})
	fake.instanceUpgradeStartingMutex.Unlock()
	if fake.InstanceUpgradeStartingStub != nil {
		fake.InstanceUpgradeStartingStub(instance, index, totalInstances)
	}
}

func (fake *FakeListener) InstanceUpgradeStartingCallCount() int {
	fake.instanceUpgradeStartingMutex.RLock()
	defer fake.instanceUpgradeStartingMutex.RUnlock()
	return len(fake.instanceUpgradeStartingArgsForCall)
}

func (fake *FakeListener) InstanceUpgradeStartingArgsForCall(i int) (string, int, int) {
	fake.instanceUpgradeStartingMutex.RLock()
	defer fake.instanceUpgradeStartingMutex.RUnlock()
	return fake.instanceUpgradeStartingArgsForCall[i].instance, fake.instanceUpgradeStartingArgsForCall[i].index, fake.instanceUpgradeStartingArgsForCall[i].totalInstances
}

func (fake *FakeListener) InstanceUpgradeStartResult(status services.UpgradeOperationType) {
	fake.instanceUpgradeStartResultMutex.Lock()
	fake.instanceUpgradeStartResultArgsForCall = append(fake.instanceUpgradeStartResultArgsForCall, struct {
		status services.UpgradeOperationType
	}{status})
	fake.recordInvocation("InstanceUpgradeStartResult", []interface{}{status})
	fake.instanceUpgradeStartResultMutex.Unlock()
	if fake.InstanceUpgradeStartResultStub != nil {
		fake.InstanceUpgradeStartResultStub(status)
	}
}

func (fake *FakeListener) InstanceUpgradeStartResultCallCount() int {
	fake.instanceUpgradeStartResultMutex.RLock()
	defer fake.instanceUpgradeStartResultMutex.RUnlock()
	return len(fake.instanceUpgradeStartResultArgsForCall)
}

func (fake *FakeListener) InstanceUpgradeStartResultArgsForCall(i int) services.UpgradeOperationType {
	fake.instanceUpgradeStartResultMutex.RLock()
	defer fake.instanceUpgradeStartResultMutex.RUnlock()
	return fake.instanceUpgradeStartResultArgsForCall[i].status
}

func (fake *FakeListener) InstanceUpgraded(instance string, result string) {
	fake.instanceUpgradedMutex.Lock()
	fake.instanceUpgradedArgsForCall = append(fake.instanceUpgradedArgsForCall, struct {
		instance string
		result   string
	}{instance, result})
	fake.recordInvocation("InstanceUpgraded", []interface{}{instance, result})
	fake.instanceUpgradedMutex.Unlock()
	if fake.InstanceUpgradedStub != nil {
		fake.InstanceUpgradedStub(instance, result)
	}
}

func (fake *FakeListener) InstanceUpgradedCallCount() int {
	fake.instanceUpgradedMutex.RLock()
	defer fake.instanceUpgradedMutex.RUnlock()
	return len(fake.instanceUpgradedArgsForCall)
}

func (fake *FakeListener) InstanceUpgradedArgsForCall(i int) (string, string) {
	fake.instanceUpgradedMutex.RLock()
	defer fake.instanceUpgradedMutex.RUnlock()
	return fake.instanceUpgradedArgsForCall[i].instance, fake.instanceUpgradedArgsForCall[i].result
}

func (fake *FakeListener) WaitingFor(instance string, boshTaskId int) {
	fake.waitingForMutex.Lock()
	fake.waitingForArgsForCall = append(fake.waitingForArgsForCall, struct {
		instance   string
		boshTaskId int
	}{instance, boshTaskId})
	fake.recordInvocation("WaitingFor", []interface{}{instance, boshTaskId})
	fake.waitingForMutex.Unlock()
	if fake.WaitingForStub != nil {
		fake.WaitingForStub(instance, boshTaskId)
	}
}

func (fake *FakeListener) WaitingForCallCount() int {
	fake.waitingForMutex.RLock()
	defer fake.waitingForMutex.RUnlock()
	return len(fake.waitingForArgsForCall)
}

func (fake *FakeListener) WaitingForArgsForCall(i int) (string, int) {
	fake.waitingForMutex.RLock()
	defer fake.waitingForMutex.RUnlock()
	return fake.waitingForArgsForCall[i].instance, fake.waitingForArgsForCall[i].boshTaskId
}

func (fake *FakeListener) Progress(pollingInterval time.Duration, orphanCount int, upgradedCount int, upgradesLeftCount int, deletedCount int) {
	fake.progressMutex.Lock()
	fake.progressArgsForCall = append(fake.progressArgsForCall, struct {
		pollingInterval   time.Duration
		orphanCount       int
		upgradedCount     int
		upgradesLeftCount int
		deletedCount      int
	}{pollingInterval, orphanCount, upgradedCount, upgradesLeftCount, deletedCount})
	fake.recordInvocation("Progress", []interface{}{pollingInterval, orphanCount, upgradedCount, upgradesLeftCount, deletedCount})
	fake.progressMutex.Unlock()
	if fake.ProgressStub != nil {
		fake.ProgressStub(pollingInterval, orphanCount, upgradedCount, upgradesLeftCount, deletedCount)
	}
}

func (fake *FakeListener) ProgressCallCount() int {
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	return len(fake.progressArgsForCall)
}

func (fake *FakeListener) ProgressArgsForCall(i int) (time.Duration, int, int, int, int) {
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	return fake.progressArgsForCall[i].pollingInterval, fake.progressArgsForCall[i].orphanCount, fake.progressArgsForCall[i].upgradedCount, fake.progressArgsForCall[i].upgradesLeftCount, fake.progressArgsForCall[i].deletedCount
}

func (fake *FakeListener) Finished(orphanCount int, upgradedCount int, deletedCount int) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		orphanCount   int
		upgradedCount int
		deletedCount  int
	}{orphanCount, upgradedCount, deletedCount})
	fake.recordInvocation("Finished", []interface{}{orphanCount, upgradedCount, deletedCount})
	fake.finishedMutex.Unlock()
	if fake.FinishedStub != nil {
		fake.FinishedStub(orphanCount, upgradedCount, deletedCount)
	}
}

func (fake *FakeListener) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakeListener) FinishedArgsForCall(i int) (int, int, int) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return fake.finishedArgsForCall[i].orphanCount, fake.finishedArgsForCall[i].upgradedCount, fake.finishedArgsForCall[i].deletedCount
}

func (fake *FakeListener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.instancesToUpgradeMutex.RLock()
	defer fake.instancesToUpgradeMutex.RUnlock()
	fake.instanceUpgradeStartingMutex.RLock()
	defer fake.instanceUpgradeStartingMutex.RUnlock()
	fake.instanceUpgradeStartResultMutex.RLock()
	defer fake.instanceUpgradeStartResultMutex.RUnlock()
	fake.instanceUpgradedMutex.RLock()
	defer fake.instanceUpgradedMutex.RUnlock()
	fake.waitingForMutex.RLock()
	defer fake.waitingForMutex.RUnlock()
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeListener) recordInvocation(key string, args []interface{}) {
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

var _ upgrader.Listener = new(FakeListener)
