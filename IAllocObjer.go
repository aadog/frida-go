package frida_go

import "unsafe"

type IAllocObjer interface {
	Instance() uintptr
	UnsafeAddr() unsafe.Pointer
	IsValid() bool
	Free()
}