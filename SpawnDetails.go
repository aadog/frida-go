package frida_go

import (
	"fmt"
	"github.com/a97077088/frida-go/cfrida"
	"unsafe"
)

type SpawnDetails struct {
	CObj
}


func (s *SpawnDetails) Description() string {
	if s.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.SpawnDetails(pid: %d, identifier: %s)",s.Pid(),s.Identifier())
}

func (p *SpawnDetails) Pid() uint {
	return cfrida.Frida_spawn_get_pid(p.instance)
}

func (p *SpawnDetails) Identifier() string {
	return cfrida.Frida_spawn_get_identifier(p.instance)
}


func (d *SpawnDetails) Free() {
	fmt.Println("spawn gc")
	cfrida.G_object_unref(d.instance)
}


// SpawnDetailsFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func SpawnDetailsFromInst(inst uintptr) *SpawnDetails {
	dl:=new(SpawnDetails)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(dl.instance)
	setFinalizer(dl, (*SpawnDetails).Free)
	return dl
}
