package frida_go

import (
	"syscall"
)

var device_onSpawnAddedPtr = syscall.NewCallbackCDecl(func(o uintptr, rawSpawn uintptr,userdata uintptr) uintptr {
	return device_onSpawnAddedCallBack(o,rawSpawn,userdata)
})
var device_onSpawnRemovedPtr = syscall.NewCallbackCDecl(func(o uintptr, rawSpawn uintptr,userdata uintptr) uintptr {
	return device_onSpawnRemovedCallBack(o,rawSpawn,userdata)
})
var device_onChildAddedPtr = syscall.NewCallbackCDecl(func(o uintptr, rawChild uintptr,userdata uintptr) uintptr {
	return device_onChildAddedCallBack(o,rawChild,userdata)
})
var device_onChildRemovedPtr = syscall.NewCallbackCDecl(func(o uintptr, rawChild uintptr,userdata uintptr) uintptr {
	return device_onChildRemovedCallBack(o,rawChild,userdata)
})
var device_onProcessCrashedPtr = syscall.NewCallbackCDecl(func(o uintptr, rawCrash uintptr,userdata uintptr) uintptr {
	return device_onProcessCrashedCallBack(o,rawCrash,userdata)
})
var device_onOutputPtr = syscall.NewCallbackCDecl(func(o uintptr,pid uintptr,fd uintptr,rawData uintptr,rawDataSize uintptr,userdata uintptr) uintptr {
	return device_onOutputCallBack(o,pid,fd,rawData,rawDataSize,userdata)
})
var device_onUninjectedPtr = syscall.NewCallbackCDecl(func(o uintptr, id uintptr,userdata uintptr) uintptr {
	return device_onUninjectedCallBack(o,id,userdata)
})
var device_onLostPtr = syscall.NewCallbackCDecl(func(o uintptr,userdata uintptr) uintptr {
	return device_onLostCallBack(o,userdata)
})