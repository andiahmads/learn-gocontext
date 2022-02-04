package context_with_timeout

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
context with timeout
selain menambahkan value ke context dan juga sinyal cancel, context juga bisa menambahkan sinyal cancel secara otomatis menggunakan timeout.
dengan menggunakan timeout kita tidak perlu  melakukan eksekusi secara manual, cancel akan otomatis dieksekusi jika waktu timeout terlewati.
penggunaan context dengan timeout sangat cocok ketika melakukan query kedatabase/http api, namun ingin menentukan batas maksimal timeoutnya.
untuk membuat contxt cancel dengan timeout kita bisa menggunakan function context.WithTimeOut(parent,duration) */

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination

}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	destination := CreateCounter(ctx)
	for num := range destination {
		fmt.Println("Counter", num)
	}
	cancel()
	fmt.Println(runtime.NumGoroutine())
}
