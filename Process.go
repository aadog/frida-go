package frida_go

import (
	"fmt"
	"frida-go/cfrida"
	"unsafe"
)

type Process struct {
	CObj
}

func (p *Process) Pid() int {
	return cfrida.Frida_process_get_pid(p.instance)
}

func (p *Process) Name() string {
	return cfrida.Frida_process_get_name(p.instance)
}

func (p *Process) Description() string {
	if p.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.ProcessDetails(pid: %d, name: %s, parameters: %s)",p.Pid(),p.Name(),"")
}

func (p *Process) Free() {
	cfrida.G_object_unref(p.instance)
}


// ProcessFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func ProcessFromInst(inst uintptr) *Process {
	dl:=new(Process)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	setFinalizer(dl, (*Process).Free)
	return dl
}
