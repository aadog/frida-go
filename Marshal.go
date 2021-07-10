package frida_go

import (
	"frida-go/cfrida"
	"unsafe"
)
func DataFromBytes(bts uintptr)[]byte{
	size:=0
	data:=cfrida.G_bytes_get_data(bts,&size)
	return data
}

func StrvFromArray(array  []string)(uintptr,int){
	o:=make([]uintptr,0)
	for _, v := range array {
		o=append(o,cfrida.GoStrToCStr(v))
	}
	return uintptr(unsafe.Pointer(&o[0])),len(array)
}
