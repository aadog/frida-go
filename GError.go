package frida_go

import (
	"errors"
	"frida-go/cfrida"
	"unsafe"
)

type GError struct {
	CObj
}
func (g *GError) ToError()error{
	if !g.IsError(){
		return nil
	}
	return errors.New(g.Message())
}
func (g *GError) ErrInput()uintptr{
	return uintptr(unsafe.Pointer(&g.instance))
}
func (g *GError) Message() string {
	if !g.IsError(){
		return ""
	}
	return cfrida.G_error_get_message(g.instance)
}
func (g *GError) Code() int {
	if !g.IsError(){
		return 0
	}
	return cfrida.G_error_get_code(g.instance)
}
func (g *GError) IsError()bool{
	return g.instance!=0
}
func (g *GError) Free() {
	if !g.IsError(){
		return
	}
	cfrida.G_error_free(g.instance)
}

// GErrorFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func GErrorFromInst(inst uintptr) *GError {
	dl:=new(GError)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	setFinalizer(dl, (*GError).Free)
	return dl
}

