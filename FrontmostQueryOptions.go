package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type FrontmostQueryOptions struct {
	CObj
	Scope FridaScope
}

func (a *FrontmostQueryOptions) Free() {
	cfrida.G_object_unref(a.instance)
}

func NewFrontmostQueryOptions()*FrontmostQueryOptions{
	r := new(FrontmostQueryOptions)
	r.instance = cfrida.Frida_frontmost_query_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*FrontmostQueryOptions).Free)
	return r
}
