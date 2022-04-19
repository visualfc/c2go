package main

import (
	"fmt"
	"log"
	"unsafe"
)

func a_cas(p *int32, t, s int32) int32 {
	log.Panicln("a_cas: notimpl")
	return 0
}

func C(s string) *int8 {
	n := len(s)
	ret := make([]byte, n+1)
	copy(ret, s)
	ret[n] = '\x00'
	return (*int8)(unsafe.Pointer(&ret[0]))
}

func gostring(s *int8) string {
	n, arr := 0, (*[1 << 20]byte)(unsafe.Pointer(s))
	for arr[n] != 0 {
		n++
	}
	return string(arr[:n])
}

func printf(format *int8, args ...interface{}) int32 {
	goformat := gostring(format)
	for i, arg := range args {
		if v, ok := arg.(*int8); ok {
			args[i] = gostring(v)
		}
	}
	fmt.Printf(goformat, args...)
	return 0
}

func sliceOf(v unsafe.Pointer, bytes uint) []byte {
	return (*[1 << 20]byte)(v)[:bytes]
}

func __builtin___memcpy_chk(dst unsafe.Pointer, src unsafe.Pointer, n uint, elem uint) unsafe.Pointer {
	copy(sliceOf(dst, n), sliceOf(src, n))
	return dst
}

func __builtin_object_size(unsafe.Pointer, int32) uint {
	return 1
}

func __builtin_bswap32(v uint32) uint32 {
	log.Panicln("__builtin_bswap32: notimpl")
	return v
}

func __builtin_bswap64(v uint64) uint64 {
	log.Panicln("__builtin_bswap32: notimpl")
	return v
}

type struct___locale_data struct{}
