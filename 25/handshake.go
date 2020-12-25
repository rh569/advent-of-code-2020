package main

import "fmt"

const (
	PubKey1 = 10441485
	PubKey2 = 1004920
)

const (
	DIVISOR   = 20201227
	START_SUB = 7
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Enc key : %v\n", part1(PubKey1, PubKey2))
}

func part1(key1, key2 int) int {
	c1, c2 := make(chan int, 1), make(chan int, 1)

	go calculateCycles(key1, c1)
	go calculateCycles(key2, c2)

	select {
	case cycles1 := <-c1:
		fmt.Println("cycles found via key1")
		return calculateEncKey(key2, cycles1)
	case cycles2 := <-c2:
		fmt.Println("cycles found via key2")
		return calculateEncKey(key1, cycles2)
	}
}

func calculateCycles(key int, cycles chan int) {
	test := 1
	doneCycles := 0

	for test != key {
		test = doCycle(test, START_SUB)
		doneCycles++
	}

	cycles <- doneCycles
}

func calculateEncKey(pub, cycles int) int {
	enc := 1

	for i := 0; i < cycles; i++ {
		enc = doCycle(enc, pub)
	}

	return enc
}

func doCycle(in, sub int) int {
	out := in * sub
	out = out % DIVISOR
	return out
}
