package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
)

func device_onSpawnAddedCallBack(o uintptr, rawSpawn uintptr,userdata uintptr) uintptr {
	v,ok:=device_onSpawnAddedbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnSpawnAddedEventFunc)
	h(SpawnDetailsFromInst(cfrida.G_object_ref(rawSpawn)))
	return 0
}
func device_onSpawnRemovedCallBack(o uintptr, rawSpawn uintptr,userdata uintptr) uintptr {
	v,ok:=device_onSpawnRemovedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnSpawnRemovedEventFunc)
	h(SpawnDetailsFromInst(cfrida.G_object_ref(rawSpawn)))
	return 0
}
func device_onChildAddedCallBack(o uintptr, rawChild uintptr,userdata uintptr) uintptr {
	v,ok:=device_onChildAddedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnChildAddedEventFunc)
	h(ChildDetailsFromInst(cfrida.G_object_ref(rawChild)))
	return 0
}
func device_onChildRemovedCallBack(o uintptr, rawChild uintptr,userdata uintptr) uintptr {
	v,ok:=deviceManager_onAddedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnChildRemovedEventFunc)
	h(ChildDetailsFromInst(cfrida.G_object_ref(rawChild)))
	return 0
}
func device_onProcessCrashedCallBack(o uintptr, rawCrash uintptr,userdata uintptr) uintptr {
	v,ok:=device_onProcessCrashedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnProcessCrashedEventFunc)
	h(CrashDetailsFromInst(cfrida.G_object_ref(rawCrash)))
	return 0
}
func device_onOutputCallBack(o uintptr,pid uintptr,fd uintptr,rawData uintptr,rawDataSize uintptr,userdata uintptr) uintptr {
	v,ok:=device_onOutputCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnOutputEventFunc)
	data:=cfrida.CBytesToGoBytes(rawData, int(rawDataSize))
	h(data, int(fd), int(pid))
	return 0
}
func device_onUninjectedCallBack(o uintptr, id uintptr,userdata uintptr) uintptr {
	v,ok:=deviceManager_onAddedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnUninjectedEventFunc)
	h(uint(id))
	return 0
}
func device_onLostCallBack(o uintptr,userdata uintptr) uintptr {
	v,ok:=device_onLostCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceOnLostEventFunc)
	h()
	return 0
}