package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	"unsafe"
)

func StrvFromArray(array  []string)(uintptr,int){
	o:=make([]uintptr,0)
	for _, v := range array {
		o=append(o,cfrida.GoStrToCStr(v))
	}
	return uintptr(unsafe.Pointer(&o[0])),len(array)
}
