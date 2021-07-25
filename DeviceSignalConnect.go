package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	"reflect"
	"sync"
)

type DeviceOnSpawnAddedEventFunc func(spawn *SpawnDetails)
type DeviceOnSpawnRemovedEventFunc func(spawn *SpawnDetails)
type DeviceOnChildAddedEventFunc func(child *ChildDetails)
type DeviceOnChildRemovedEventFunc func(child *ChildDetails)
type DeviceOnProcessCrashedEventFunc func(crash *CrashDetails)
type DeviceOnOutputEventFunc func(data []byte,fd int,pid int)
type DeviceOnUninjectedEventFunc func(id uint)
type DeviceOnLostEventFunc func()
var device_onSpawnAddedbackTable=sync.Map{}
var device_onSpawnRemovedCallbackTable=sync.Map{}
var device_onChildAddedCallbackTable=sync.Map{}
var device_onChildRemovedCallbackTable=sync.Map{}
var device_onProcessCrashedCallbackTable=sync.Map{}
var device_onOutputCallbackTable=sync.Map{}
var device_onUninjectedCallbackTable=sync.Map{}
var device_onLostCallbackTable=sync.Map{}

type DeviceSignalConnect struct {
	onSpawnAddedSigs sync.Map
	onSpawnRemovedSigs sync.Map
	onChildAddedSigs sync.Map
	onChildRemovedSigs sync.Map
	onProcessCrashedSigs sync.Map
	onOutputSigs sync.Map
	onUninjectedSigs sync.Map
	onLostSigs sync.Map
	rawDevicePtr uintptr
}


func (s *DeviceSignalConnect) OnSpawnAdded(on DeviceOnSpawnAddedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onSpawnAddedbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "spawn-added", device_onSpawnAddedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onSpawnAddedSigs.Store(sigid,userdata)
	return sigid
}
func (s *DeviceSignalConnect) OnSpawnRemoved(on DeviceOnSpawnRemovedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onSpawnRemovedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "spawn-removed", device_onSpawnRemovedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onSpawnRemovedSigs.Store(sigid,userdata)
	return sigid
}
func (s *DeviceSignalConnect) OnChildAdded(on DeviceOnChildAddedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onChildAddedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "child-added", device_onChildAddedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onChildAddedSigs.Store(sigid,userdata)
	return sigid
}
func (s *DeviceSignalConnect) OnChildRemoved(on DeviceOnChildRemovedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onChildRemovedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "child-removed", device_onChildRemovedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onChildRemovedSigs.Store(sigid,userdata)
	return sigid
}
func (s *DeviceSignalConnect) OnProcessCrashed(on DeviceOnProcessCrashedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onProcessCrashedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "process-crashed", device_onProcessCrashedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onProcessCrashedSigs.Store(sigid,userdata)
	return sigid
}
func (s *DeviceSignalConnect) OnOutput(on DeviceOnOutputEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onOutputCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "output", device_onOutputPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onOutputSigs.Store(sigid,userdata)
	return sigid
}
func (s *DeviceSignalConnect) OnUninjected(on DeviceOnUninjectedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onUninjectedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "uninjected", device_onUninjectedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onUninjectedSigs.Store(sigid,userdata)
	return sigid
}
func (s *DeviceSignalConnect) OnLost(on DeviceOnLostEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	device_onLostCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawDevicePtr, "lost", device_onLostPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onLostSigs.Store(sigid,userdata)
	return sigid
}

func (s *DeviceSignalConnect) free() {
	if s.rawDevicePtr!=0{
		//fmt.Println("DeviceSignalConnect gc")
		s.onSpawnAddedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onSpawnAddedbackTable.Delete(key.(int64))
			return true
		})
		s.onSpawnRemovedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onSpawnRemovedCallbackTable.Delete(key.(int64))
			return true
		})
		s.onChildAddedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onChildAddedCallbackTable.Delete(key.(int64))
			return true
		})
		s.onChildRemovedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onChildRemovedCallbackTable.Delete(key.(int64))
			return true
		})
		s.onProcessCrashedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onProcessCrashedCallbackTable.Delete(key.(int64))
			return true
		})
		s.onOutputSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onOutputCallbackTable.Delete(key.(int64))
			return true
		})
		s.onUninjectedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onUninjectedCallbackTable.Delete(key.(int64))
			return true
		})
		s.onLostSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawDevicePtr, key.(int64))
			device_onLostCallbackTable.Delete(key.(int64))
			return true
		})
		s.rawDevicePtr=0
	}
}

func NewDeviceSignalConnect(rawPtr uintptr) *DeviceSignalConnect {
	sig := new(DeviceSignalConnect)
	sig.onSpawnAddedSigs = sync.Map{}
	sig.onSpawnRemovedSigs = sync.Map{}
	sig.onChildAddedSigs = sync.Map{}
	sig.onChildRemovedSigs = sync.Map{}
	sig.onProcessCrashedSigs = sync.Map{}
	sig.onOutputSigs = sync.Map{}
	sig.onUninjectedSigs = sync.Map{}
	sig.onLostSigs = sync.Map{}
	sig.rawDevicePtr = rawPtr
	return sig
}
