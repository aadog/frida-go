package frida_go

import (
"fmt"
"frida-go/cfrida"
"unsafe"
)

const (
	ScriptEvent_message = "message"
	ScriptEvent_destroyed = "destroyed"
)

type Script struct {
	CObj
}

func (p *Script) Pid() int {
	return cfrida.Frida_session_get_pid(p.instance)
}

func (p *Script) Load() error {
	var err GError
	cfrida.Frida_script_load_sync(p.instance,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}
func (p *Script) UnLoad() error {
	var err GError
	cfrida.Frida_script_unload_sync(p.instance,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}
func (p *Script) Eternalize() error {
	var err GError
	cfrida.Frida_script_eternalize_sync(p.instance,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}

func (p *Script) Description() string {
	if p.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.Script()")
}

func (p *Script) Free() {
	cfrida.G_object_unref(p.instance)
}


// ScriptFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func ScriptFromInst(inst uintptr) *Script {
	dl:=new(Script)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	cfrida.G_signal_connect_data(dl.instance,"message",0,0,0,G_CONNECT_AFTER)
	cfrida.G_signal_connect_data(dl.instance,"destroyed",0,0,0,G_CONNECT_AFTER)
	setFinalizer(dl, (*Script).Free)
	return dl
}

