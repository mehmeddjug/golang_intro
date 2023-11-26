package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Wait group is used to wait for goroutines to finish.
var wg sync.WaitGroup

// Wait group for shared resources is used to wait for goroutines to finish.
var wg_lsr sync.WaitGroup

func countPlus30(base int32) {
	defer wg.Done()
	for i := 0; i < 30; i++ {
		base += 1
		fmt.Printf("%d ", base)
	}
}

func add30toLSR(base *int32) {
	// Notify wg_lsr when the function is executed.
	defer wg_lsr.Done()
	for i := 0; i < 30; i++ {
		atomic.AddInt32(base, 1)
		runtime.Gosched()
	}
}

func ConcurrencyExample() {
	runtime.GOMAXPROCS(1) // Allocates 1 logical processor
	wg.Add(3)             // For each goroutine we add a group.

	fmt.Println("Start Goroutines")

	go countPlus30(10)
	go countPlus30(50)
	go countPlus30(100)

	fmt.Println("Waiting goroutines to finish.")
	wg.Wait()
	fmt.Println("\nExit Program")
}

func LockSharedResourcesExample() {
	runtime.GOMAXPROCS(1) // Allocates 1 logical processor
	wg_lsr.Add(3)         // For each goroutine we add a group.

	count := int32(0)
	fmt.Println("Start Goroutines LSR")

	go add30toLSR(&count)
	go add30toLSR(&count)
	go add30toLSR(&count)

	fmt.Println("Waiting goroutines to finish.")
	wg_lsr.Wait()
	fmt.Printf("Count: %d\n", count)
	fmt.Println("Exit Program")
}
