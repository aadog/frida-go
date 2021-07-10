package frida_go

import (
	"fmt"
	"frida-go/cfrida"
	"unsafe"
)

type Application struct {
	CObj
}

func (a *Application) Identifier() string {
	return cfrida.Frida_application_get_identifier(a.instance)
}

func (a *Application) Name() string {
	return cfrida.Frida_application_get_name(a.instance)
}

func (a *Application) Pid() int {
	return cfrida.Frida_application_get_pid(a.instance)
}
func (a *Application) Params() map[string]interface{} {
	//tb:=cfrida.Frida_application_get_parameters(a.instance)
	//fmt.Println(tb)
	return map[string]interface{}{}
}

func (a *Application) Description() string {
	if a.instance==0{
		return ""
	}
	if a.Pid()!=0{
		return fmt.Sprintf("Frida.ApplicationDetails(identifier: %s, name: %s, pid: %d, parameters: %s)",a.Identifier(),a.Name(),a.Pid(),"")
	}else{
		return fmt.Sprintf("Frida.ApplicationDetails(identifier: %s, name: %s, parameters: %s)",a.Identifier(),a.Name(),"")
	}
}

func (a *Application) Free() {
	cfrida.G_object_unref(a.instance)
}


// ApplicationFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func ApplicationFromInst(inst uintptr) *Application {
	dl:=new(Application)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(inst)
	setFinalizer(dl, (*Application).Free)
	return dl
}