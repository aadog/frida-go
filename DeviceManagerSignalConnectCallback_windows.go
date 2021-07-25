// +build windows
package frida_go

import (
	"syscall"
)


var deviceManager_onAddedPtr = syscall.NewCallbackCDecl(func(o uintptr, rawDevice uintptr,userdata uintptr) uintptr {
	return deviceManager_onAddedCallBack(uintptr,rawDevice,userdata)
})
var deviceManager_onChangedPtr = syscall.NewCallbackCDecl(func(o uintptr,userdata uintptr) uintptr {
	return deviceManager_onChangedCallBack(o,userdata)
})
var deviceManager_onRemovedPtr = syscall.NewCallbackCDecl(func(o uintptr, rawDevice uintptr,userdata uintptr) uintptr {
	return deviceManager_onRemovedCallBack(o,rawDevice,userdata)
})
