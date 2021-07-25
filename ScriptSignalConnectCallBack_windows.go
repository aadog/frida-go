package frida_go

import (
	"syscall"
)

var script_onDestroyedPtr = syscall.NewCallbackCDecl(func(sc uintptr, userdata uintptr) uintptr {
	return script_onDestroyedCallBack(sc,userdata)
})
var script_onMessagePtr = syscall.NewCallbackCDecl(func(sc uintptr, rawjson uintptr, rawdata uintptr, userdata uintptr) uintptr {
	return script_onMessageCallBack(sc,rawjson,rawdata,userdata)
})

