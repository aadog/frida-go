package frida_go

import (
	"fmt"
	"github.com/a97077088/frida-go/cfrida"
	"unsafe"
)

type ProcessDetails struct {
	CObj
}

func (p *ProcessDetails) Pid() uint {
	return cfrida.Frida_process_get_pid(p.instance)
}

func (p *ProcessDetails) Name() string {
	return cfrida.Frida_process_get_name(p.instance)
}
func (p *ProcessDetails) Parameters() map[string]interface{} {
	return cfrida.Frida_process_get_parameters(p.instance)
}

func (p *ProcessDetails) Description() string {
	if p.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.ProcessDetails(pid: %d, name: %s, parameters: %s)",p.Pid(),p.Name(),p.Parameters())
}

func (p *ProcessDetails) Free() {
	cfrida.G_object_unref(p.instance)
}


// ProcessFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func ProcessFromInst(inst uintptr) *ProcessDetails {
	dl:=new(ProcessDetails)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	setFinalizer(dl, (*ProcessDetails).Free)
	return dl
}
