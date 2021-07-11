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
	err:=cfrida.Frida_device_manager_close_sync(d.instance,0)
	if err!=nil{
	    return err
	}
	return nil
}
func (d *DeviceManager) Free() {
	d.free()
	fmt.Println("DeviceManager gc")
	cfrida.G_object_unref(d.instance)
}
func (d *DeviceManager) EnumerateDevices()([]*Device,error){

	rawDevices,err:=cfrida.Frida_device_manager_enumerate_devices_sync(d.instance,0)
	if err!=nil{
	    return nil,err
	}
	defer cfrida.G_object_unref(rawDevices)
	devices:=make([]*Device,0)
	n:=cfrida.Frida_device_list_size(rawDevices)
	for i := 0; i <n; i++ {
		devices=append(devices,DeviceFromInst(cfrida.Frida_device_list_get(rawDevices,i)))
	}

	return devices,nil
}
func (d *DeviceManager) AddRemoteDevice(address string,ops RemoteDeviceOptions)(*Device,error){
	rawops:=cfrida.Frida_remote_device_options_new()
	defer cfrida.G_object_unref(rawops)
	if ops.Certificate!=""{
		panic("no imp")
	}
	if ops.Origin!=""{
		cfrida.Frida_remote_device_options_set_origin(rawops,ops.Origin)
	}
	if ops.Token!=""{
		cfrida.Frida_remote_device_options_set_token(rawops,ops.Token)
	}
	if ops.KeepaliveInterval!=0{
		cfrida.Frida_remote_device_options_set_keepalive_interval(rawops,ops.KeepaliveInterval)
	}

	r,err:=cfrida.Frida_device_manager_add_remote_device_sync(d.instance,address,rawops,0)
	if err!=nil{
	    return nil,err
	}
	return DeviceFromInst(r),nil
}
func (d *DeviceManager) RemoveRemoteDevice(address string)error{
	err:=cfrida.Frida_device_manager_remove_remote_device_sync(d.instance,address,0)
	if err!=nil{
	    return err
	}
	return nil
}
func (d *DeviceManager) GetDeviceById(id string,timeout_ms int) (*Device,error) {
	dev,err:=cfrida.Frida_device_manager_get_device_by_id_sync(d.instance,id,timeout_ms,0)
	if err!=nil{
	    return nil,err
	}
	return DeviceFromInst(dev),nil
}
func (d *DeviceManager) GetDeviceByType(tp DeviceType,timeout_ms int) (*Device,error) {
	dev,err:=cfrida.Frida_device_manager_get_device_by_type_sync(d.instance, int(tp),timeout_ms,0)
	if err!=nil{
	    return nil,err
	}
	return DeviceFromInst(dev),nil
}
func (d *DeviceManager) FindDeviceById(id string,timeout_ms int) (*Device,error) {
	dev,err:=cfrida.Frida_device_manager_find_device_by_id_sync(d.instance,id,timeout_ms,0)
	if err!=nil{
	    return nil,err
	}
	return DeviceFromInst(dev),nil
}
func (d *DeviceManager) FindDeviceByType(tp DeviceType,timeout_ms int) (*Device,error) {
	dev,err:=cfrida.Frida_device_manager_find_device_by_type_sync(d.instance, int(tp),timeout_ms,0)
	if err!=nil{
	    return nil,err
	}
	if dev==0{
		return nil,nil
	}
	return DeviceFromInst(dev),nil
}
func DeviceManager_Create() *DeviceManager {
	i := new(DeviceManager)
	i.instance = cfrida.Frida_device_manager_new()
	i.ptr = unsafe.Pointer(i.instance)
	i.DeviceManagerSignalConnect=NewDeviceManagerSignalConnect(i.instance)
	setFinalizer(i, (*DeviceManager).Free)
	return i
}