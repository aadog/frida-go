// +build !windows
// +build cgo

package frida_go

/*
extern void* bus_onDetach(void*, void* );
extern void* bus_onMessage(void*, void*, void*, void*);
 */
import "C"
import (
	"unsafe"
)

//export bus_onDetach
func bus_onDetach(sc unsafe.Pointer, userdata unsafe.Pointer)unsafe.Pointer{
	return unsafe.Pointer(bus_onDetachCallBack(uintptr(sc), uintptr(userdata)))
}

//export bus_onMessage
func bus_onMessage(sc unsafe.Pointer, rawjson unsafe.Pointer, rawdata unsafe.Pointer, userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(bus_onMessageCallBack(uintptr(sc), uintptr(rawjson), uintptr(rawdata), uintptr(userdata)))
}

var(
	bus_onDetachPtr=uintptr(C.bus_onDetach)
	bus_onMessagePtr=uintptr(C.bus_onMessage)
)