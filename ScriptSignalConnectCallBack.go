package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	jsoniter "github.com/json-iterator/go"
)

func script_onDestroyedCallBack(sc uintptr, userdata uintptr) uintptr {
	v,ok:=script_onDestroyedCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(ScriptOnDestroyedEventFunc)
	h()
	return 0
}
func script_onMessageCallBack(sc uintptr, rawjson uintptr, rawdata uintptr, userdata uintptr) uintptr {
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
}

