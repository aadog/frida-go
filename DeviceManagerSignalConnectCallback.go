package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
)


func deviceManager_onAddedCallBack(_, rawDevice uintptr,userdata uintptr) uintptr {
	v,ok:=deviceManager_onAddedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceManagerAddedEventFunc)
	h(DeviceFromInst(cfrida.G_object_ref(rawDevice)))
	return 0
}
func deviceManager_onChangedCallBack(_,userdata uintptr) uintptr {
	v,ok:=deviceManager_onChangedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceManagerChangedEventFunc)
	h()
	return 0
}
func deviceManager_onRemovedCallBack(_, rawDevice uintptr,userdata uintptr) uintptr {
	v,ok:=deviceManager_onRemovedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(DeviceManagerRemovedEventFunc)
	h(DeviceFromInst(cfrida.G_object_ref(rawDevice)))
	return 0
}
