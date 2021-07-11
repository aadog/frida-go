package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type FileMonitor struct {
	CObj
}

func (f *FileMonitor) Free() {
	cfrida.G_object_unref(f.instance)
}

func (f *FileMonitor) Enable()error{
	err:=cfrida.Frida_file_monitor_enable_sync(f.instance,0,)
	if err!=nil{
	    return err
	}
	return nil
}
func (f *FileMonitor) Disable()error{
	err:=cfrida.Frida_file_monitor_disable_sync(f.instance,0,)
	if err!=nil{
	    return err
	}
	return nil
}


// FileMonitorCreate
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func FileMonitor_Create(path string) *FileMonitor {
	dl:=new(FileMonitor)
	dl.instance=cfrida.Frida_file_monitor_new(path)
	dl.ptr= unsafe.Pointer(dl.instance)
	setFinalizer(dl, (*FileMonitor).Free)
	return dl
}
