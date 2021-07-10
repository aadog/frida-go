package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type ScriptOptions struct {
	CObj
	Name string
	Runtime FridaScriptRuntime
}

func (a *ScriptOptions) Free() {
	cfrida.G_object_unref(a.instance)
}

func NewScriptOptions()*ScriptOptions{
	r := new(ScriptOptions)
	r.instance = cfrida.Frida_script_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*ScriptOptions).Free)
	return r
}
