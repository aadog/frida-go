package frida_go

import (
	"fmt"
	"github.com/a97077088/frida-go/cfrida"
	"unsafe"
)

type ApplicationDetails struct {
	CObj
}

func (a *ApplicationDetails) Identifier() string {
	return cfrida.Frida_application_get_identifier(a.instance)
}

func (a *ApplicationDetails) Name() string {
	return cfrida.Frida_application_get_name(a.instance)
}

func (a *ApplicationDetails) Pid() int {
	return cfrida.Frida_application_get_pid(a.instance)
}
func (a *ApplicationDetails) Params() map[string]interface{} {
	return cfrida.Frida_application_get_parameters(a.instance)
}

func (a *ApplicationDetails) Description() string {
	if a.instance==0{
		return ""
	}
	if a.Pid()!=0{
		return fmt.Sprintf("Frida.ApplicationDetails(identifier: %s, name: %s, pid: %d, parameters: %s)",a.Identifier(),a.Name(),a.Pid(),"")
	}else{
		return fmt.Sprintf("Frida.ApplicationDetails(identifier: %s, name: %s, parameters: %s)",a.Identifier(),a.Name(),"")
	}
}

func (a *ApplicationDetails) Free() {
	cfrida.G_object_unref(a.instance)
}


// ApplicationFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func ApplicationFromInst(inst uintptr) *ApplicationDetails {
	dl:=new(ApplicationDetails)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	setFinalizer(dl, (*ApplicationDetails).Free)
	return dl
}