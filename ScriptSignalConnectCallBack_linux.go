// +build !windows
// +build linux
// +build darwin

package frida_go
/*
extern void* script_onDestroyed(void*, void*);
extern void* script_onMessage(void*, void*,void*, void*);
 */
import "C"
import (
	"unsafe"
)

//export script_onDestroyed
func script_onDestroyed(sc unsafe.Pointer, userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(script_onDestroyedCallBack(uintptr(sc), uintptr(userdata)))
}
//export script_onMessage
func script_onMessage(sc unsafe.Pointer, rawjson unsafe.Pointer, rawdata unsafe.Pointer, userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(script_onMessageCallBack(uintptr(sc), uintptr(rawjson), uintptr(rawdata), uintptr(userdata)))
}

var(
	script_onDestroyedPtr=uintptr(C.script_onDestroyed)
	script_onMessagePtr=uintptr(C.script_onMessage)
)