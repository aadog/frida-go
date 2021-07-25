package frida_go

/*
extern void* device_onSpawnAdded(void*, void*,void*);
extern void* device_onSpawnRemoved(void*, void*,void*);
extern void* device_onChildAdded(void*, void*, void*);
extern void* device_onChildRemoved(void*, void*, void*);
extern void* device_onProcessCrashed(void*, void*, void*);
extern void* device_onOutput(void*, void*, void*,void*, void*, void*);
extern void* device_onUninjected(void*, void*, void*);
extern void* device_onLost(void*, void*);
*/
import "C"
import (
	"unsafe"
)

//export device_onSpawnAdded
func device_onSpawnAdded(o unsafe.Pointer, rawSpawn unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onSpawnAddedCallBack(uintptr(o), uintptr(rawSpawn), uintptr(userdata)))
}
//export device_onSpawnRemoved
func device_onSpawnRemoved(o unsafe.Pointer, rawSpawn unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onSpawnRemovedCallBack(uintptr(o), uintptr(rawSpawn), uintptr(userdata)))
}
//export device_onChildAdded
func device_onChildAdded(o unsafe.Pointer, rawChild unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onChildAddedCallBack(uintptr(o), uintptr(rawChild), uintptr(userdata)))
}
//export device_onChildRemoved
func device_onChildRemoved(o unsafe.Pointer, rawChild unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onChildRemovedCallBack(uintptr(o), uintptr(rawChild), uintptr(userdata)))
}
//export device_onProcessCrashed
func device_onProcessCrashed(o unsafe.Pointer, rawCrash unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onProcessCrashedCallBack(uintptr(o), uintptr(rawCrash), uintptr(userdata)))
}
//export device_onOutput
func device_onOutput(o unsafe.Pointer,pid unsafe.Pointer,fd unsafe.Pointer,rawData unsafe.Pointer,rawDataSize unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onOutputCallBack(uintptr(o), uintptr(pid), uintptr(fd), uintptr(rawData), uintptr(rawDataSize), uintptr(userdata)))
}
//export device_onUninjected
func device_onUninjected(o unsafe.Pointer, id unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onUninjectedCallBack(uintptr(o), uintptr(id), uintptr(userdata)))
}
//export device_onLost
func device_onLost(o unsafe.Pointer,userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(device_onLostCallBack(uintptr(o), uintptr(userdata)))
}


var(
	device_onSpawnAddedPtr=uintptr(C.device_onSpawnAdded)
	device_onSpawnRemovedPtr=uintptr(C.device_onSpawnRemoved)
	device_onChildAddedPtr=uintptr(C.device_onChildAdded)
	device_onChildRemovedPtr=uintptr(C.device_onChildRemoved)
	device_onProcessCrashedPtr=uintptr(C.device_onProcessCrashed)
	device_onOutputPtr=uintptr(C.device_onOutput)
	device_onUninjectedPtr=uintptr(C.device_onUninjected)
	device_onLostPtr=uintptr(C.device_onLost)
)