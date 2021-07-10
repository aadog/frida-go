package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type SessionOptions struct {
	CObj
	Realm FridaRealm
	PersistTimeout int
}

func (a *SessionOptions) Free() {
	cfrida.G_object_unref(a.instance)
}

func NewSessionOptions()*SessionOptions{
	r := new(SessionOptions)
	r.instance = cfrida.Frida_session_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*SessionOptions).Free)
	return r
}
