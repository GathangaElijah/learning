package main

import "math"


// Write a function that builds a map[int]float64 where the keys are the
// numbers from 0 (inclusive) to 100,000 (exclusive) and the values are
// the square roots of those numbers (use the math.Sqrt function to
// calculate square roots). Use sync.OnceValue to generate a function
// that caches the map returned by this function and use the cached value to
// look up square roots for every 1,000th number from 0 to 100,000.

func MapProcessor() map[int]float64 {
	mp := make( map[int]float64)

		for i := 0; i < 100000; i++ {
			sq := math.Sqrt(float64(i))
			 mp[i] = sq
		}
	return mp
}
