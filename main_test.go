package main

import (
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("returns exactly size elements", func(t *testing.T) {
		const size = 100
		if got := len(generateRandomElements(size)); got != size {
			t.Fatalf("len = %d, want %d", got, size)
		}
	})

	t.Run("size 1", func(t *testing.T) {
		if got := len(generateRandomElements(1)); got != 1 {
			t.Fatalf("len = %d, want 1", got)
		}
	})

	t.Run("non-negative values", func(t *testing.T) {
		for _, v := range generateRandomElements(1000) {
			if v < 0 {
				t.Fatalf("got negative value %d", v)
			}
		}
	})

	t.Run("panics on non-positive size", func(t *testing.T) {
		for _, size := range []int{0, -1, -100} {
			func() {
				defer func() {
					if recover() == nil {
						t.Errorf("size %d: expected panic, got none", size)
					}
				}()
				generateRandomElements(size)
			}()
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"single zero", []int{0}, 0},
		{"max first", []int{9, 1, 2, 3}, 9},
		{"max last", []int{1, 2, 3, 9}, 9},
		{"max middle", []int{1, 9, 2}, 9},
		{"duplicates", []int{5, 5, 5}, 5},
		{"all negative", []int{-9, -2, -5}, -2},
		{"mixed", []int{-9, 4, -2, 0}, 4},
		{"maxint", []int{0, 1<<62 - 1}, 1<<62 - 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}

	t.Run("matches brute force on random data", func(t *testing.T) {
		for range 20 {
			data := make([]int, rand.Intn(500)+1)
			for i := range data {
				data[i] = rand.Intn(2001) - 1000
			}
			want := data[0]
			for _, v := range data {
				if v > want {
					want = v
				}
			}
			if got := maximum(data); got != want {
				t.Fatalf("maximum = %d, want %d", got, want)
			}
		}
	})
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"fewer elements than chunks", []int{3, 8, 1}, 8},
		{"exactly one chunk", []int{3, 8, 1, 4, 2, 9, 7, 5}, 9},
		{"non-divisible length", []int{3, 8, 1, 4, 2, 9, 7, 5, 11}, 11},
		{"max in last element", []int{1, 2, 3, 4, 5, 6, 7, 8, 42}, 42},
		{"all negative", []int{-1, -2, -3, -4, -5, -6, -7, -8, -9}, -1},
		{"duplicates", []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}

	t.Run("agrees with maximum on random data", func(t *testing.T) {
		for range 20 {
			data := make([]int, rand.Intn(2000)+1)
			for i := range data {
				data[i] = rand.Intn(2001) - 1000
			}
			if got, want := maxChunks(data), maximum(data); got != want {
				t.Fatalf("len %d: maxChunks = %d, maximum = %d", len(data), got, want)
			}
		}
	})
}
