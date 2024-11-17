package main

import (
	"fmt"
	"runtime"
)

func main() {

	// time looping compare
	// start := time.Now()

	// for i := 1; i <= 10000000; i++ {
	// 	fmt.Printf("Looping ke %d\n", i)
	// }

	// duration := time.Since(start)
	// fmt.Printf("Waktu proses: %v\n", duration)

	// allocated memory

	var m runtime.MemStats

	// Sebelum alokasi
	runtime.ReadMemStats(&m)
	fmt.Printf("Sebelum Alokasi: Alloc = %v KB\n", m.Alloc/1024)

	// Alokasi memori
	for i := 0; i < 10; i++ {
		allocateMemory()
	}

	// Setelah alokasi
	runtime.ReadMemStats(&m)
	fmt.Printf("Setelah Alokasi: Alloc = %v KB\n", m.Alloc/1024)

	// Memaksa GC berjalan
	runtime.GC()

	// Setelah GC
	runtime.ReadMemStats(&m)
	fmt.Printf("Setelah GC: Alloc = %v KB\n", m.Alloc/1024)

}

func allocateMemory() {
	_ = make([]byte, 1<<20) // Alokasikan 1 MB
}
