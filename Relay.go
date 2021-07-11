package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	"unsafe"
)
const (
	RELAY_KIND_TURN_UDP = iota
	RELAY_KIND_TURN_TCP
	RELAY_KIND_TURN_TLS
)
type RelayKind int

type Relay struct {
	CObj
}

func (r *Relay) Free() {
	cfrida.G_object_unref(r.instance)
}

// NewRelay
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func NewRelay(address string,username string,password string,kind RelayKind) *Relay {
	r := new(Relay)
	r.instance = cfrida.Frida_relay_new(address,username,password, int(kind))
	r.ptr = unsafe.Pointer(r.instance)
	setFinalizer(r, (*Relay).Free)
	return r
}

