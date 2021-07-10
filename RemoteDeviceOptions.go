package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)

type RemoteDeviceOptions struct {
	CObj
	Certificate string
	Origin string
	Token string
	KeepaliveInterval int
}

func (r *RemoteDeviceOptions) Free() {
	cfrida.G_object_unref(r.instance)
}

func NewRemoteDeviceOptions()*RemoteDeviceOptions{
	r := new(RemoteDeviceOptions)
	r.instance = cfrida.Frida_remote_device_options_new()
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*RemoteDeviceOptions).Free)
	return r
}