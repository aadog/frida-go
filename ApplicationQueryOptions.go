package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type ApplicationQueryOptions struct {
	CObj
	Identifiers []string
	Scope FridaScope
}

func (a *ApplicationQueryOptions) Free() {
	cfrida.G_object_unref(a.instance)
}

func NewApplicationQueryOptions()*ApplicationQueryOptions{
	r := new(ApplicationQueryOptions)
	r.instance = cfrida.Frida_application_query_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*ApplicationQueryOptions).Free)
	return r
}
