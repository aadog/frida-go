package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"sync"
	"syscall"
)

type ScriptOnMessageEventFunc func(sjson jsoniter.Any, data []byte)
type ScriptOnDestroyedEventFunc func()
type ScriptSignalConnect struct {
	onMessageSigs sync.Map
	onDestroyedSigs sync.Map
	rawScriptPtr uintptr
}

func (s *ScriptSignalConnect) free() {
	if s.rawScriptPtr!=0{
		//fmt.Println("ScriptSignalConnect gc")
		s.onMessageSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawScriptPtr, key.(int64))
			script_onMessageCallbackTable.Delete(key.(int64))
			return true
		})
		s.onDestroyedSigs.Range(func(key, value interface{}) bool {
			cfrida.G_signal_handler_disconnect(s.rawScriptPtr, key.(int64))
			script_onDestroyedCallbackTable.Delete(key.(int64))
			return true
		})
		s.rawScriptPtr=0
	}
}
func (s *ScriptSignalConnect) OnDestroyed(on ScriptOnDestroyedEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	script_onDestroyedCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawScriptPtr, "destroyed", script_onDestroyedPtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onDestroyedSigs.Store(sigid,userdata)
	return sigid
}
func (s *ScriptSignalConnect) OnMessage(on ScriptOnMessageEventFunc) int64 {
	userdata:=int64(reflect.ValueOf(on).Pointer())
	script_onMessageCallbackTable.Store(userdata,on)
	sigid := cfrida.G_signal_connect_data(s.rawScriptPtr, "message", script_onMessagePtr, uintptr(userdata), 0, G_CONNECT_AFTER)
	s.onMessageSigs.Store(sigid,userdata)
	return sigid
}
var script_onDestroyedPtr = syscall.NewCallbackCDecl(func(sc uintptr, userdata uintptr) uintptr {
	v,ok:=script_onDestroyedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(ScriptOnDestroyedEventFunc)
	h()
	return 0
})
var script_onDestroyedCallbackTable=sync.Map{}
var script_onMessageCallbackTable=sync.Map{}
var script_onMessagePtr = syscall.NewCallbackCDecl(func(sc uintptr, rawjson uintptr, rawdata uintptr, userdata uintptr) uintptr {
	sjson := cfrida.CStrToGoStr(rawjson)
	jjson := jsoniter.Get([]byte(sjson))
	data := cfrida.G_bytes_to_bytes_and_unref(rawdata)
	tp := jjson.Get("type").ToString()
	if tp != "send" {
		v,ok:=script_onMessageCallbackTable.Load(int64(userdata))
		if !ok{
			return 0
		}
		h:=v.(ScriptOnMessageEventFunc)
		h(jjson,data)
	} else {
		if jjson.Get("payload").Size() < 4 {
			v,ok:=script_onMessageCallbackTable.Load(int64(userdata))
			if !ok{
				return 0
			}
			h:=v.(ScriptOnMessageEventFunc)
			h(jjson,data)
		} else {
			rpcCallbackFunc(jjson.Get("payload"))
		}
	}
	return 0
})

func NewScriptSignalConnect(rawPtr uintptr) *ScriptSignalConnect {
	sig := new(ScriptSignalConnect)
	sig.onMessageSigs = sync.Map{}
	sig.onDestroyedSigs = sync.Map{}
	sig.rawScriptPtr = rawPtr
	return sig
}
