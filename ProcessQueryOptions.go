package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type ProcessQueryOptions struct {
	CObj
	SelectPids []int
	Scope FridaScope
}

func (a *ProcessQueryOptions) Free() {
	cfrida.G_object_unref(a.instance)
}

func NewProcessQueryOptions()*ProcessQueryOptions{
	r := new(ProcessQueryOptions)
	r.instance = cfrida.Frida_process_query_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*ProcessQueryOptions).Free)
	return r
}
