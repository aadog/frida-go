package frida_go

import (
	"github.com/a97077088/frida-go/cfrida"
	jsoniter "github.com/json-iterator/go"
)

func bus_onDetachCallBack(sc uintptr, userdata uintptr) uintptr {
	v,ok:=bus_onDetachCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(BusOnDetachEventFunc)
	h()
	return 0
}
func bus_onMessageCallBack(sc uintptr, rawjson uintptr, rawdata uintptr, userdata uintptr) uintptr {
	sjson := cfrida.CStrToGoStr(rawjson)
	jjson := jsoniter.Get([]byte(sjson))
	data := cfrida.G_bytes_to_bytes_and_unref(rawdata)
	v,ok:=bus_onMessageCallbackTable.Load(int64(userdata))
	if !ok{
		return 0
	}
	h:=v.(BusOnMessageEventFunc)
	h(jjson,data)
	return 0
}