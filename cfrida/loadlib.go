package cfrida

import (
	"syscall"
)

var(
	libfrida=loadUILib()
)

func loadUILib() *syscall.LazyDLL  {
	lib:=syscall.NewLazyDLL("frida_shared.dll")
	err:=lib.Load()
	if err!=nil{
	    panic(err)
	}
	return lib
}
func init(){
	Frida_init()
}
