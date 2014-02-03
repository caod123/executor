package executor

import (
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

type TaskRegistry struct {
	ExecutorMemoryMB int
	ExecutorDiskMB   int
	runOnces         map[string]models.RunOnce
}

func NewTaskRegistry(memoryMB int, diskMB int) *TaskRegistry {
	return &TaskRegistry{
		ExecutorMemoryMB: memoryMB,
		ExecutorDiskMB:   diskMB,
		runOnces:         make(map[string]models.RunOnce),
	}
}

func (registry *TaskRegistry) HasCapacityForRunOnce(runOnce models.RunOnce) bool {
	if runOnce.MemoryMB > registry.availableMemoryMB() {
		return false
	}

	if runOnce.DiskMB > registry.availableDiskMB() {
		return false
	}

	return true
}

func (registry *TaskRegistry) AddRunOnce(runOnce models.RunOnce) {
	registry.runOnces[runOnce.Guid] = runOnce
	return
}

func (registry *TaskRegistry) availableMemoryMB() int {
	usedMemory := 0
	for _, r := range registry.runOnces {
		usedMemory = usedMemory + r.MemoryMB
	}
	return registry.ExecutorMemoryMB - usedMemory
}

func (registry *TaskRegistry) availableDiskMB() int {
	usedDisk := 0
	for _, r := range registry.runOnces {
		usedDisk = usedDisk + r.DiskMB
	}
	return registry.ExecutorDiskMB - usedDisk
}

func (registry *TaskRegistry) RunOnces() map[string]models.RunOnce {
	return registry.runOnces
}
