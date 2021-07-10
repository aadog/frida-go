package frida_go

import (
	"fmt"
	"frida-go/cfrida"
	"unsafe"
)

const (
	G_CONNECT_AFTER	= 1 << 0
	G_CONNECT_SWAPPED = 1 << 1
)

type DeviceManager struct {
	CObj
	*DeviceManagerSignalConnect
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
	fmt.Println("DeviceManager gc")
	d.free()
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
func (d *DeviceManager) GetDeviceById(id string,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_get_device_by_id_sync(d.instance,id,timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(cfrida.G_object_ref(dev)),nil
}
func (d *DeviceManager) GetDeviceByType(tp DeviceType,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_get_device_by_type_sync(d.instance, int(tp),timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(cfrida.G_object_ref(dev)),nil
}
func (d *DeviceManager) FindDeviceById(id string,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_find_device_by_id_sync(d.instance,id,timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	return DeviceFromInst(cfrida.G_object_ref(dev)),nil
}
func (d *DeviceManager) FindDeviceByType(tp DeviceType,timeout_ms int) (*Device,error) {
	var err GError
	dev:=cfrida.Frida_device_manager_find_device_by_type_sync(d.instance, int(tp),timeout_ms,0,err.ErrInput())
	if err.IsError(){
		return nil,err.ToError()
	}
	if dev==0{
		return nil,nil
	}
	return DeviceFromInst(cfrida.G_object_ref(dev)),nil
}
func DeviceManager_Create() *DeviceManager {
	i := new(DeviceManager)
	i.instance = cfrida.Frida_device_manager_new()
	i.ptr = unsafe.Pointer(i.instance)
	i.DeviceManagerSignalConnect=NewDeviceManagerSignalConnect(i.instance)
	setFinalizer(i, (*DeviceManager).Free)
	return i
}