package main

import (
	"fmt"
	"time"
)

// function untuk Goroutine pertama
func process1(data interface{}) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Process 1: %v (Iteration %d)\n", data, i+1)
		time.Sleep(100 * time.Millisecond) // simulasi waktu proses
	}
}

// function untuk Goroutine kedua
func process2(data interface{}) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Process 2: %v (Iteration %d)\n", data, i+1)
		time.Sleep(100 * time.Millisecond) // simulasi waktu proses
	}
}

func main() {
	// data tipe interface{}
	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	// mulai Goroutine
	go process1(data1)
	go process2(data2)

	// waktu untuk kedua Goroutine menyelesaikan proses
	time.Sleep(1 * time.Second)
}
