package frida_go

import (
	"context"
	"errors"
	"fmt"
	"frida-go/cfrida"
	"github.com/json-iterator/go"
	"log"
	"math"
	"sync"
	"unsafe"
)

const (
	RpcOperation_call = "call"
)

const (
	RpcKind_default = "frida:rpc"
)

var reqlk sync.Mutex
var rpcRequestId int64
func nextRpcRequestId() int64 {
	reqlk.Lock()
	defer reqlk.Unlock()
	currentId := rpcRequestId
	if currentId >= math.MaxInt64{
		rpcRequestId = 0
	} else {
		rpcRequestId++
	}
	return currentId
}
var rpcCallbackMap sync.Map=sync.Map{}
func rpcCallbackFunc(jjsonCallback jsoniter.Any) {
	id := jjsonCallback.Get(1).ToInt64()
	h, ok := rpcCallbackMap.LoadAndDelete(id)
	if !ok {
		return
	}
	h.(chan jsoniter.Any) <- jjsonCallback
}


type Script struct {
	*ScriptSignalConnect
	CObj
}


func (s *Script) DefaultOnMessage(jjson jsoniter.Any,data []byte){
	tp:=jjson.Get("type").ToString()
	if tp=="log"{
		log.Println(jjson.Get("payload").ToString())
	}else if tp=="error"{

		log.Println(jjson.Get("stack").ToString())
		log.Println(jjson.Get("fileName").ToString())
	}else{
		log.Println(jjson.ToString())
	}
}
func (s *Script) IsDestroyed() bool {
	return cfrida.Frida_script_is_destroyed(s.instance)
}

func (s *Script) Load() error {
	err:=cfrida.Frida_script_load_sync(s.instance,0,)
	if err!=nil{
	    return err
	}
	return nil
}
func (s *Script) UnLoad() error {
	err:=cfrida.Frida_script_unload_sync(s.instance,0)
	if err!=nil{
	    return err
	}
	return nil
}
func (s *Script) Eternalize() error {
	err:=cfrida.Frida_script_eternalize_sync(s.instance,0,)
	if err!=nil{
	    return err
	}
	return nil
}


func (s *Script) RpcCall(ctx context.Context, functionName string, args ...interface{}) (jsoniter.Any, error){
	reqid := nextRpcRequestId()
	message := []interface{}{
		RpcKind_default,
		reqid,
		RpcOperation_call,
		functionName,
		args,
	}
	ch := make(chan jsoniter.Any)
	rpcCallbackMap.Store(reqid, ch)
	s.Post(message, nil)
	select {
	case <-ctx.Done():
		_=s.ptr // ref s.ptr active gc
		return nil, ctx.Err()
	case jsresult := <-ch:
		ret := jsresult.Get(2).ToString()
		if ret != "ok" {
			return nil, errors.New(jsresult.Get(3).ToString())
		} else {
			return jsresult.Get(3), nil
		}
	}

	return nil, errors.New("未知错误")
}
func (s *Script) Post(obj interface{}, data []byte) {
	cfrida.Frida_script_post(s.rawScriptPtr, jsoniter.Wrap(obj).ToString(), data)
}

func (s *Script) Description() string {
	if s.instance==0{
		return ""
	}
	return fmt.Sprintf("Frida.Script()")
}



func (s *Script) Free() {
	if s.instance!=0{
		s.free()
		fmt.Println("script gc")
		cfrida.G_object_unref(s.instance)
		s.instance,s.ptr=0,nil
	}
}


// ScriptFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func ScriptFromInst(inst uintptr) *Script {
	dl:=new(Script)
	dl.instance=inst
	dl.ptr= unsafe.Pointer(dl.instance)
	dl.ScriptSignalConnect=NewScriptSignalConnect(dl.instance)
	setFinalizer(dl, (*Script).Free)
	return dl
}

