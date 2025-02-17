//go:build linux || darwin
// +build linux darwin

package malloc

/*
 This code has been tested and works like the windows version

*/
import (
	"fmt"
	"sync"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Memory registry to track allocations
var (
	memoryRegistry = map[uintptr]int{} // Maps address to allocated size
	registryLock   sync.Mutex          // Ensures thread-safe access to the registry
)

func PlatformMalloc(size uintptr) (uintptr, error) {
	// Using unix.Mmap to allocate memory
	data, err := unix.Mmap(
		-1,                             // No file descriptor
		0,                              // Offset
		int(size),                      // Size of memory to allocate
		unix.PROT_READ|unix.PROT_WRITE, // Memory protection
		unix.MAP_ANON|unix.MAP_PRIVATE, // Flags
	)
	if err != nil {
		return 0, fmt.Errorf("Mmap failed: %v", err)
	}
	addr := uintptr(unsafe.Pointer(&data[0]))
	// Register the allocation in the memory registry
	registryLock.Lock()
	memoryRegistry[addr] = int(size)
	registryLock.Unlock()
	return addr, nil
}

func PlatformFree(addr uintptr) error {

	// Look up the size in the memory registry
	registryLock.Lock()
	size, exists := memoryRegistry[addr]
	if exists {
		delete(memoryRegistry, addr) // Remove from registry
	}
	registryLock.Unlock()

	if !exists {
		return fmt.Errorf("invalid free: address %x not found in registry", addr)
	}
	// Reconstruct the slice and unmap the memory
	data := unsafe.Slice((*byte)(unsafe.Pointer(addr)), int(size))
	return unix.Munmap(data)
}
