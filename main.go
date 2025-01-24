package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	type T struct {
		v int
		p unsafe.Pointer
	}

	for i := 0; i < 100; i++ {
		v := &new(T).v
		*v = 97531

		runtime.AddCleanup(v, func(v int) {
			fmt.Println(v)
		}, 123)

		v = nil
		runtime.GC()
	}
}
