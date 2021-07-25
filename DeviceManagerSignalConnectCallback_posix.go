// +build !windows
// +build cgo

package frida_go

/*
extern void* deviceManager_onAdded(void*, void*,void*);
extern void* deviceManager_onChanged(void*, void*);
extern void* deviceManager_onRemoved(void*, void*, void*);
*/
import "C"
import (
	"unsafe"
)


//export deviceManager_onAdded
func deviceManager_onAdded(o unsafe.Pointer, rawDevice unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(deviceManager_onAddedCallBack(uintptr(o), uintptr(rawDevice), uintptr(userdata)))
}
//export deviceManager_onChanged
func deviceManager_onChanged(o unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(deviceManager_onChangedCallBack(uintptr(o), uintptr(userdata)))
}
//export deviceManager_onRemoved
func deviceManager_onRemoved(o unsafe.Pointer, rawDevice unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(deviceManager_onRemovedCallBack(uintptr(o), uintptr(rawDevice), uintptr(userdata)))
}


var(
	deviceManager_onAddedPtr=uintptr(C.deviceManager_onAdded)
	deviceManager_onChangedPtr=uintptr(C.deviceManager_onChanged)
	deviceManager_onRemovedPtr=uintptr(C.deviceManager_onRemoved)
)