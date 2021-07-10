package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type FridaPeerOptions struct {
	CObj
	StunServer string
	Relays []*Relay
}

func (a *FridaPeerOptions) Free() {
	cfrida.G_object_unref(a.instance)
}

func NewFridaPeerOptions()*FridaPeerOptions{
	r := new(FridaPeerOptions)
	r.instance = cfrida.Frida_script_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*FridaPeerOptions).Free)
	return r
}
