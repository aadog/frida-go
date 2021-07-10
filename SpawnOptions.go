package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

const (
	FRIDA_STDIO_INHERIT = iota
	FRIDA_STDIO_PIPE
)

type FridaStdio int32
type SpawnOptions struct {
	CObj
	Argv []string
	Envp []string
	Env []string
	Cwd string
	Stdio FridaStdio
}

func (a *SpawnOptions) Free() {
	cfrida.G_object_unref(a.instance)
}

func NewSpawnOptions()*SpawnOptions{
	r := new(SpawnOptions)
	r.instance = cfrida.Frida_spawn_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*SpawnOptions).Free)
	return r
}
