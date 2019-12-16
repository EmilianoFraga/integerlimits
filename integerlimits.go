package main

import (
	"fmt"
)

func main() {
	var minI8 int8 = -0x80
	fmt.Printf("%8s: %+x\n", "minI8", minI8)

	var maxI8 int8 = 0x7f
	fmt.Printf("%8s: %+x\n", "maxI8", maxI8)

	var minI16 int16 = -0x8000
	fmt.Printf("%8s: %+x\n", "minI16", minI16)

	var maxI16 int16 = 0x7fff
	fmt.Printf("%8s: %+x\n", "maxI16", maxI16)

	var minI32 int32 = -0x80000000
	fmt.Printf("%8s: %+x\n", "minI32", minI32)

	var maxI32 int32 = 0x7fffffff
	fmt.Printf("%8s: %+x\n", "maxI32", maxI32)

	var minI64 int64 = -0x8000000000000000
	fmt.Printf("%8s: %+x\n", "minI64", minI64)

	var maxI64 int64 = 0x7fffffffffffffff
	fmt.Printf("%8s: %+x\n", "maxI64", maxI64)

	var minUI8 uint8 = 0x0
	fmt.Printf("%8s: %+x\n", "minUI8", minUI8)
	
	var maxUI8 uint8 = 0xff
	fmt.Printf("%8s: %+x\n", "maxUI8", maxUI8)

	var minUI16 uint16 = 0x0
	fmt.Printf("%8s: %+x\n", "minUI16", minUI16)

	var maxUI16 uint16 = 0xffff
	fmt.Printf("%8s: %+x\n", "maxUI16", maxUI16)

	var minUI32 uint32 = 0x0
	fmt.Printf("%8s: %+x\n", "minUI32", minUI32)

	var maxUI32 uint32 = 0xffffffff
	fmt.Printf("%8s: %+x\n", "maxUI32", maxUI32)

	var minUI64 uint64 = 0x0
	fmt.Printf("%8s: %+x\n", "minUI64", minUI64)

	var maxUI64 uint64 = 0xffffffffffffffff
	fmt.Printf("%8s: %+x\n", "maxUI64", maxUI64)
}