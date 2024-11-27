Usage:
```func main() {
	size := uintptr(4096) // 4 KB
	addr, err := Malloc(size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Allocated memory at address: %x\n", addr)

	// Write to the allocated memory
	data := (*[4096]byte)(unsafe.Pointer(addr))
	data[0] = 42
	fmt.Println("Stored value:", data[0])
	for i := 0; i < 4096; i++ {
		data[i] = byte(i * 42)
		fmt.Println(data[i])
	}
	fmt.Println(*data)
	// Free the memory
	err = Free(addr)
	if err != nil {
		fmt.Println("Failed to free memory:", err)
	} else {
		fmt.Println("Memory freed successfully")
	}
}
```