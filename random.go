package main

import (
	"fmt"
	"time"
)

var (
	s [2]uint64
)

func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

func next() uint64 {
	s0, s1 := s[0], s[1]
	result := s0 + s1
	s1 ^= s0
	s[0] = rotl(s0, 55) ^ s1 ^ (s1 << 14)
	s[1] = rotl(s1, 36)
	return result
}

func main() {
	s[0], s[1] = uint64(time.Now().UnixNano()^0x3bfa8764f685bd1c), uint64(time.Now().UnixNano()^0x5a2fdc2bf68cedb3) // silly seed
	for i := 0; i < 10; i++ {
		fmt.Println(next())
	}
}