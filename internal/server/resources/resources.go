package resources

import (
	"fmt"
	"sync"
	"time"

	"github.com/lxc/incus/v6/shared/api"
)

// Cache to speed up concurrent runs.
var (
	muResources   sync.Mutex
	lastResources *api.Resources
	lastRun       time.Time
)

// GetResources returns a filled api.Resources struct ready for use by Incus.
func GetResources() (*api.Resources, error) {
	// Ensure only one concurrent run.
	muResources.Lock()
	defer muResources.Unlock()

	// Check if we ran less than 10s ago.
	if lastResources != nil && lastRun.Add(10*time.Second).Before(time.Now()) {
		return lastResources, nil
	}

	// Get CPU information
	cpu, err := GetCPU()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve CPU information: %w", err)
	}

	// Get memory information
	memory, err := GetMemory()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve memory information: %w", err)
	}

	// Get GPU information
	gpu, err := GetGPU()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve GPU information: %w", err)
	}

	// Get network card information
	network, err := GetNetwork()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve network information: %w", err)
	}

	// Get storage information
	storage, err := GetStorage()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve storage information: %w", err)
	}

	// Get USB information
	usb, err := GetUSB()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve USB information: %w", err)
	}

	// Get PCI information
	pci, err := GetPCI()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve PCI information: %w", err)
	}

	// Get system information
	system, err := GetSystem()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve system information: %w", err)
	}

	load, err := GetLoad()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve load information: %w", err)
	}

	// Build the final struct
	resources := api.Resources{
		CPU:     *cpu,
		Memory:  *memory,
		GPU:     *gpu,
		Network: *network,
		Storage: *storage,
		USB:     *usb,
		PCI:     *pci,
		System:  *system,
		Load:    *load,
	}

	// Update the cache.
	lastResources = &resources
	lastRun = time.Now()

	return &resources, nil
}
