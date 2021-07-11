package frida_go

import (
	"fmt"
	"github.com/a97077088/frida-go/cfrida"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

type Bus struct {
	CObj
	*BusSignalConnect
}

func (b *Bus) Post(message interface{},data []byte)  {
	cfrida.Frida_bus_post(b.instance,jsoniter.Wrap(message).ToString(),data)
}

func (b *Bus) Attach()(bool,error){
	isatt,err:=cfrida.Frida_bus_attach_sync(b.instance,0)
	return isatt,err
}

func (b *Bus) Description() string {
	if b.instance==0{
		return ""
	}
	return fmt.Sprintf(`Frida.Bus()`)
}
func (b *Bus) IsClosed() bool {
	return cfrida.Frida_bus_is_detached(b.instance)
}

func (b *Bus) Free() {
	cfrida.G_object_unref(b.instance)
}

// BusFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func BusFromInst(inst uintptr) *Bus {
	dl:=new(Bus)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(dl.instance)
	dl.BusSignalConnect=NewBusSignalConnect(dl.instance)
	setFinalizer(dl, (*Bus).Free)
	return dl
}