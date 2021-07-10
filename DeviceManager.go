package frida_go

import (
	"fmt"
	"frida-go/cfrida"
	"sync"
	"syscall"
	"unsafe"
)

type DeviceManagerAddedEventFunc func(manager *DeviceManager,device *Device)
type DeviceManagerChangeEventFunc func(manager *DeviceManager)
type DeviceManagerRemovedEventFunc func(manager *DeviceManager,device *Device)
const (
	G_CONNECT_AFTER	= 1 << 0
	G_CONNECT_SWAPPED = 1 << 1
)
const (
	DeviceManage_Event_added = "added"
	DeviceManage_Event_removed = "removed"
	DeviceManage_Event_changed = "changed"
)

type DeviceManager struct {
	CObj
	added_handles sync.Map
	added_cb uintptr
	changed_handles sync.Map
	changed_cb uintptr
	removed_handles sync.Map
	removed_cb uintptr
	allrawhandles sync.Map
}
func (d *DeviceManager) Close()error{
	var err GError
	cfrida.Frida_device_manager_close_sync(d.instance,0,err.ErrInput())
	if err.IsError(){
	    return err.ToError()
	}
	return nil
}
func (d *DeviceManager) Free() {
	d.allrawhandles.Range(func(key, value interface{}) bool {
		cfrida.G_signal_handler_disconnect(d.instance,key.(int64))
		return true
	})
	cfrida.G_object_unref(d.instance)
}
func (d *DeviceManager) EnumerateDevices()([]*Device,error){
	var err GError
	rawDevices:=cfrida.Frida_device_manager_enumerate_devices_sync(d.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	defer cfrida.G_object_unref(rawDevices)
	devices:=make([]*Device,0)
	n:=cfrida.Frida_device_list_size(rawDevices)
	for i := 0; i <n; i++ {
		devices=append(devices,DeviceFromInst(cfrida.Frida_device_list_get(rawDevices,i)))
	}
	return devices,nil
}
func (d *DeviceManager) AddRemoteDevice(address string,ops *RemoteDeviceOptions)(*Device,error){
	if ops!=nil{
		if ops.Certificate!=""{
			panic("no imp")
		}
		if ops.Origin!=""{
			cfrida.Frida_remote_device_options_set_origin(d.instance,ops.Origin)
		}
		if ops.Token!=""{
			cfrida.Frida_remote_device_options_set_token(d.instance,ops.Token)
		}
		if ops.KeepaliveInterval!=0{
			cfrida.Frida_remote_device_options_set_keepalive_interval(d.instance,ops.KeepaliveInterval)
		}
	}
	var err GError
	r:=cfrida.Frida_device_manager_add_remote_device_sync(d.instance,address, ops.instance,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(r),nil
}
func (d *DeviceManager) RemoveRemoteDevice(address string)error{
	var err GError
	cfrida.Frida_device_manager_remove_remote_device_sync(d.instance,address,0,err.ErrInput())
	if err.IsError(){
		return err.ToError()
	}
	return nil
}
func (d *DeviceManager) added_callback(_ uintptr,dev uintptr,userdata uintptr)uintptr{
	ld,ok:=d.added_handles.Load(userdata)
	if ok !=true{
		fmt.Println("找不到这个事件")
		return 0
	}
	fn:=ld.(DeviceManagerAddedEventFunc)
	cfrida.G_object_unref(dev)
	fn(d,DeviceFromInst(dev))
	return 0
}
func (d *DeviceManager) changed_callback(dev uintptr,userdata uintptr)uintptr{
	ld,ok:=d.changed_handles.Load(userdata)
	if ok !=true{
		fmt.Println("找不到这个事件")
		return 0
	}
	fn:=ld.(DeviceManagerChangeEventFunc)
	fn(d)
	return 0
}
func (d *DeviceManager) removed_callback(mgr uintptr,dev uintptr,userdata uintptr)uintptr{
	ld,ok:=d.removed_handles.Load(userdata)
	if ok !=true{
		fmt.Println("找不到这个事件")
		return 0
	}
	fn:=ld.(DeviceManagerRemovedEventFunc)
	cfrida.G_object_unref(dev)
	fn(d,DeviceFromInst(dev))
	return 0
}

func (d *DeviceManager) OnAdded(eventFunc DeviceManagerAddedEventFunc)int64{
	h:=d.SignalConnectData(DeviceManage_Event_added,eventFunc)
	d.allrawhandles.Store(h,h)
	return h
}
func (d *DeviceManager) OnChanged(eventFunc DeviceManagerChangeEventFunc)int64{
	h:=d.SignalConnectData(DeviceManage_Event_changed,eventFunc)
	d.allrawhandles.Store(h,h)
	return h
}
func (d *DeviceManager) OnRemoved(eventFunc DeviceManagerRemovedEventFunc)int64{
	h:=d.SignalConnectData(DeviceManage_Event_removed,eventFunc)
	d.allrawhandles.Store(h,h)
	return h
}

func (d *DeviceManager) SignalConnectData(event string,handle interface{})int64{
	switch event{
	case DeviceManage_Event_added:
		userdata:=uintptr(unsafe.Pointer(&handle))
		d.added_handles.Store(userdata,handle)
		return cfrida.G_signal_connect_data(d.instance,event,d.added_cb,userdata,0,G_CONNECT_AFTER)
	case DeviceManage_Event_changed:
		userdata:=uintptr(unsafe.Pointer(&handle))
		d.changed_handles.Store(userdata,handle)
		return cfrida.G_signal_connect_data(d.instance,event,d.changed_cb,userdata,0,G_CONNECT_AFTER)
	case DeviceManage_Event_removed:
		userdata:=uintptr(unsafe.Pointer(&handle))
		d.removed_handles.Store(userdata,handle)
		return cfrida.G_signal_connect_data(d.instance,event,d.removed_cb,userdata,0,G_CONNECT_AFTER)
	default:
		panic("不支持这个事件")
	}
}
func (d *DeviceManager) GetDeviceById(id string,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_get_device_by_id_sync(d.instance,id,timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(dev),nil
}
func (d *DeviceManager) GetDeviceByType(tp DeviceType,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_get_device_by_type_sync(d.instance, int(tp),timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(dev),nil
}
func (d *DeviceManager) FindDeviceById(id string,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_find_device_by_id_sync(d.instance,id,timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(dev),nil
}
func (d *DeviceManager) FindDeviceByType(tp DeviceType,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_find_device_by_type_sync(d.instance, int(tp),timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(dev),nil
}
func DeviceManager_Create() *DeviceManager {
	i := new(DeviceManager)
	i.instance = cfrida.Frida_device_manager_new()
	i.ptr = unsafe.Pointer(i.instance)
	i.added_handles=sync.Map{}
	i.added_cb=syscall.NewCallback(i.added_callback)
	i.changed_handles=sync.Map{}
	i.changed_cb=syscall.NewCallback(i.changed_callback)
	i.removed_handles=sync.Map{}
	i.removed_cb=syscall.NewCallback(i.removed_callback)
	i.allrawhandles=sync.Map{}
	setFinalizer(i, (*DeviceManager).Free)
	return i
}