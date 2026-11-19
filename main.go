package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		panic("can't be less than zero")
	}
	res := make([]int, size)
	for i := range size {
		res[i] = rand.Int()
	}
	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, i := range data {
		if i > max {
			max = i
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var chunkMx = make([]int, CHUNKS)
	chunkLen := len(data) / CHUNKS
	group := sync.WaitGroup{}
	for i := 1; i <= CHUNKS; i++ {
		end := chunkLen * i
		if i == CHUNKS {
			end = len(data)
		}
		group.Add(1)
		go func(chunk []int, idx int) {
			defer group.Done()
			chunkMx[idx] = maximum(chunk)
		}(data[chunkLen*(i-1):end], i-1)
	}
	group.Wait()
	return maximum(chunkMx)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	s := time.Now()
	mx := maximum(data)
	elapsed := time.Since(s).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", mx, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	s = time.Now()
	mx = maxChunks(data)
	elapsed = time.Since(s).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", mx, elapsed)
}
