package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type ChildDetails struct {
	CObj
}

func (c *ChildDetails) Pid() uint {
	return cfrida.Frida_child_get_pid(c.instance)
}

func (c *ChildDetails) ParentPid() uint {
	return cfrida.Frida_child_get_parent_pid(c.instance)
}

func (c *ChildDetails) Origin() int {
	return cfrida.Frida_child_get_origin(c.instance)
}

func (c *ChildDetails) Identifier() string {
	return cfrida.Frida_child_get_identifier(c.instance)
}

func (c *ChildDetails) Path() string {
	return cfrida.Frida_child_get_path(c.instance)
}

func (c *ChildDetails) Argv() []string {
	return cfrida.Frida_child_get_argv(c.instance)
}
func (c *ChildDetails) Envp() map[string]interface{} {
	return cfrida.Frida_child_get_envp(c.instance)
}


func (c *ChildDetails) Free() {
	cfrida.G_object_unref(c.instance)
}

// ChildDetailsFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func ChildDetailsFromInst(inst uintptr) *ChildDetails {
	dl:=new(ChildDetails)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(dl.instance)
	setFinalizer(dl, (*ChildDetails).Free)
	return dl
}
