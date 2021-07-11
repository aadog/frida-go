package frida_go

import (
	"fmt"
	"frida-go/cfrida"
	"unsafe"
)

type CrashDetails struct {
	CObj
}



func (c *CrashDetails) Pid() uint {
	return cfrida.Frida_crash_get_pid(c.instance)
}

func (c *CrashDetails) ProcessName() string {
	return cfrida.Frida_crash_get_process_name(c.instance)
}
func (c *CrashDetails) Summary() string {
	return cfrida.Frida_crash_get_summary(c.instance)
}
func (c *CrashDetails) Report() string {
	return cfrida.Frida_crash_get_report(c.instance)
}

func (c *CrashDetails) Parameters() map[string]interface{} {
	return cfrida.Frida_crash_get_parameters(c.instance)
}

func (c *CrashDetails) Description() string {
	if c.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.CrashDetails(pid: %d, processName: %s, summary: %s",c.Pid(),c.ProcessName(),c.Summary())
}


func (c *CrashDetails) Free() {
	cfrida.G_object_unref(c.instance)
}

// CrashDetailsFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func CrashDetailsFromInst(inst uintptr) *CrashDetails {
	dl:=new(CrashDetails)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(dl.instance)
	setFinalizer(dl, (*CrashDetails).Free)
	return dl
}
