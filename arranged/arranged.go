package main

import (
	"fmt"
	"sync"
	"time"
)

// function untuk Goroutine pertama
func process1(data interface{}, wg *sync.WaitGroup, mutex1, mutex2 *sync.Mutex) {
	defer wg.Done() // selesai function berakhir

	for i := 0; i < 4; i++ {
		mutex1.Lock() // kunci mutex1 untuk goroutine pertama
		fmt.Printf("Process 1: %v (Iteration %d)\n", data, i+1)
		mutex2.Unlock() // buka kunci mutex2 untuk goroutine kedua
		time.Sleep(100 * time.Millisecond) // simulasi waktu proses
	}
}

// function untuk Goroutine kedua
func process2(data interface{}, wg *sync.WaitGroup, mutex1, mutex2 *sync.Mutex) {
	defer wg.Done() // selesai function berakhir

	for i := 0; i < 4; i++ {
		mutex2.Lock() // kunci mutex2 untuk goroutine kedua
		fmt.Printf("Process 2: %v (Iteration %d)\n", data, i+1)
		mutex1.Unlock() // buka kunci mutex1 untuk goroutine pertama
		time.Sleep(100 * time.Millisecond) // simulasi waktu proses
	}
}

func main() {
	var wg sync.WaitGroup
	var mutex1 sync.Mutex
	var mutex2 sync.Mutex

	// kunci mutex2 sehingga goroutine pertama yang akan dijalankan pertama kali
	mutex2.Lock()

	// data tipe interface{}
	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	// mulai Goroutine
	wg.Add(2)
	go process1(data1, &wg, &mutex1, &mutex2)
	go process2(data2, &wg, &mutex1, &mutex2)

	// tunggu kedua Goroutine selesai
	wg.Wait()
}
