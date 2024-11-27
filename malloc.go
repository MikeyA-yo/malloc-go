package malloc

func Malloc(size uintptr) (uintptr, error) {
	return PlatformMalloc(size)
}

func Free(addr uintptr) error {
	return PlatformFree(addr)
}
