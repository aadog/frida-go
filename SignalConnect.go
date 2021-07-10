package frida_go

import (
	"fmt"
	"frida-go/cfrida"
	"reflect"
	"runtime"
	"sync"
	"syscall"
)

type SignalConnect struct {
	rawPtr uintptr
	callbackPtr uintptr
	handleMap sync.Map
}

func (s *SignalConnect) Free() {
	s.handleMap.Range(func(key, value interface{}) bool {
		cfrida.G_signal_handler_disconnect(s.rawPtr,key.(int64))
		return true
	})
}

func (s *SignalConnect) Call(vals ...interface{})uintptr{
	reflect.ValueOf()
	return 0
}

func NewSignalConnect(rawPtr uintptr)*SignalConnect{
	sig:=new(SignalConnect)
	sig.rawPtr=rawPtr
	sig.handleMap=sync.Map{}
	sig.callbackPtr=syscall.NewCallbackCDecl(sig.Call)
	runtime.SetFinalizer(sig,sig.Free)
	return sig
}