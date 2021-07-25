package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	"reflect"
	"sync"
)

type DeviceManagerAddedEventFunc func(device *Device)
type DeviceManagerChangedEventFunc func()
type DeviceManagerRemovedEventFunc func(device *Device)
type DeviceManagerSignalConnect struct {
	onAddedSigs sync.Map
	onChangedSigs sync.Map
	onRemovedSigs sync.Map
	rawDeviceManagerPtr uintptr
}

var deviceManager_onAddedCallbackTable=sync.Map{}
var deviceManager_onChangedCallbackTable=sync.Map{}
var deviceManager_onRemovedCallbackTable=sync.Map{}




func (s *DeviceManagerSignalConnect) OnAdded(on DeviceManagerAddedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	deviceManager_onAddedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDeviceManagerPtr, "added", deviceManager_onAddedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onAddedSigs.Store(sigid,userdata)
	return sigid
}

func (s *DeviceManagerSignalConnect) OnChanged(on DeviceManagerChangedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	deviceManager_onChangedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDeviceManagerPtr, "changed", deviceManager_onChangedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onChangedSigs.Store(sigid,userdata)
	return sigid
}

func (s *DeviceManagerSignalConnect) OnRemoved(on DeviceManagerRemovedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	deviceManager_onRemovedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDeviceManagerPtr, "removed", deviceManager_onRemovedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onRemovedSigs.Store(sigid,userdata)
	return sigid
}

func (s *DeviceManagerSignalConnect) free() {
	if s.rawDeviceManagerPtr!=0{
		//fmt.Println("DeviceManagerSignalConnect gc")
		s.onAddedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDeviceManagerPtr, key.(int64))
			deviceManager_onAddedCallbackTable.Delete(key.(int64))
			return true
		})
		s.onChangedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDeviceManagerPtr, key.(int64))
			deviceManager_onChangedCallbackTable.Delete(key.(int64))
			return true
		})
		s.onRemovedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDeviceManagerPtr, key.(int64))
			deviceManager_onRemovedCallbackTable.Delete(key.(int64))
			return true
		})
		s.rawDeviceManagerPtr=0
	}
}


func NewDeviceManagerSignalConnect(rawPtr uintptr) *DeviceManagerSignalConnect {
	sig := new(DeviceManagerSignalConnect)
	sig.onAddedSigs = sync.Map{}
	sig.onChangedSigs = sync.Map{}
	sig.onRemovedSigs = sync.Map{}
	sig.rawDeviceManagerPtr = rawPtr
	return sig
}