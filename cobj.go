package frida_go

import "unsafe"

type CObj struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// Instance 返回对象实例指针。
//
// Return object instance pointer.
func (i *CObj) Instance() uintptr {
	return i.instance
}

// UnsafeAddr 获取一个不安全的地址。
//
// Get an unsafe address.
func (i *CObj) UnsafeAddr() unsafe.Pointer {
	return i.ptr
}

// IsValid 检测地址是否为空。
//
// Check if the address is empty.
func (i *CObj) IsValid() bool {
	return i.instance != 0
}
func (i *CObj) ClearPointer()  {
	i.instance=0
	i.ptr=unsafe.Pointer(i.instance)
}
