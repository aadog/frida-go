package frida_go

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"runtime"
	"testing"
	"time"
)

func TestDevice_InjectLibraryFileFile(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.FindDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		t.Fatal(err)
	}
	_,err=d.InjectLibraryFileFile(1,"./1.dylib","main",[]byte("aaa"))
	if err!=nil{
	    t.Fatal(err)
	}
	fmt.Println("inject ok")
}

func v(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.GetDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(d)

	p, err := d.GetProcessByName("映客直播", NewProcessMatchOptions())
	if err != nil {
		t.Fatal(err)
	}
	system_session, err := d.Attach(p.Pid(), NewSessionOptions())
	if err != nil {
		t.Fatal(err)
	}
	ops := NewScriptOptions()
	ops.Name = "test"
	sc, err := system_session.CreateScript(`
	console.log("ok111111111")
	rpc.exports={
		add:async function(a,b,c){
			return a+b+c
		}
	}
	`, NewScriptOptions())
	if err != nil {
		t.Fatal(err)
	}
	sc.OnMessage(func(sjson jsoniter.Any, data []byte) {
		fmt.Println(sjson.ToString())
	})
	err = sc.Load()
	if err != nil {
		t.Fatal(err)
	}
	defer sc.UnLoad()
}
func TestScript_Free(t *testing.T) {
	go func() {
		v(t)
	}()
	time.Sleep(time.Second * 1)
	for {
		runtime.GC()
		time.Sleep(time.Second)
	}
}
func TestScript_Post(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.FindDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		t.Fatal(err)
	}
	p, err := d.GetProcessByName("映客直播", NewProcessMatchOptions())
	if err != nil {
		t.Fatal(err)
	}
	system_session, err := d.Attach(p.Pid(), NewSessionOptions())
	if err != nil {
		t.Fatal(err)
	}
	//defer system_session.Detach()
	ops := NewScriptOptions()
	ops.Name = "test"
	sc, err := system_session.CreateScript(`
	console.log("ok111111111")
	rpc.exports={
		listThreads: function () {
    		return Process.enumerateThreadsSync();
  		}
	}
	`, NewScriptOptions())
	if err != nil {
		t.Fatal(err)
	}
	sc.OnDestroyed(func() {
		fmt.Println("脚本退出了")
	})

	sc.OnMessage(sc.DefaultOnMessage)
	err = sc.Load()
	if err != nil {
		t.Fatal(err)
	}
	sc.UnLoad()
	fmt.Println("aaa")


	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
	r, err := sc.RpcCall(ctx, "listThreads", 1, 2, 4)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r.ToString())
}
func TestSession_bytecode(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.FindDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		t.Fatal(err)
	}
	system_session, err := d.Attach(0, NewSessionOptions())
	if err != nil {
		t.Fatal(err)
	}
	defer system_session.Detach()
	ops := NewScriptOptions()
	bt, err := system_session.CompileScript(`
	console.log('hello')`, ops)
	if err != nil {
		t.Fatal(err)
	}
	sc, err := system_session.CreateScriptFormBytes(bt, NewScriptOptions())
	if err != nil {
		t.Fatal(err)
	}
	sc.OnMessage(sc.DefaultOnMessage)
	err = sc.Load()
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second*2)
}

func TestFileMonitor_Create(t *testing.T) {
	f := FileMonitor_Create("/")
	err := f.Enable()
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 10)
	defer f.Disable()
}
func TestNewSignalConnect(t *testing.T) {

}

func TestDevice(t *testing.T) {
	dm := DeviceManager_Create()
	defer dm.Close()

}

func TestDevice_OpenChannel(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.GetDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Description())
	stream, err := d.OpenChannel("tcp:22")
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()
	bt, _, err := stream.ReadAll(512)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bt))

}

func TestAttach(t *testing.T) {
	dm := DeviceManager_Create()

	d, err := dm.AddRemoteDevice("127.0.0.1", NewRemoteDeviceOptions())
	if err != nil {
		panic(err)
	}

	p, err := d.GetProcessByName("notepad.exe", NewProcessMatchOptions())
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Description())

	//ops:=NewSessionOptions()
	//ops.PersistTimeout=2000
	//session,err:=d.Attach(a.Pid(),ops)
	//if err!=nil{
	//	panic(err)
	//}
	//fmt.Println(session)
	//defer session.Detach()
	//fmt.Println(session.PersistTimeout())
	//fmt.Println(session.Description())
	//scops:=NewScriptOptions()
	//scops.Runtime=FRIDA_SCRIPT_RUNTIME_V8
	//sc,err:=session.CreateScript(`
	//console.log("aaa")
	//`,scops)
	//if err!=nil{
	//   panic(err)
	//}
	//err=sc.Load()
	//if err!=nil{
	//    panic(err)
	//}
	//defer sc.UnLoad()
	//
	//fmt.Println(sc)
	//fmt.Println(session.IsDetached())

}

func TestDevice_Spawn(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.GetDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		panic(err)
	}
	ops:=NewSpawnOptions()
	ops.Argv=[]string{"-c","cat /etc/hosts"}
	ops.Stdio=FRIDA_STDIO_PIPE
	p,err:=d.Spawn("/bin/sh",ops)
	if err!=nil{
	    t.Fatal(err)
	}
	d.Resume(p)
	fmt.Println(p)
}

func TestDevice_FrontmostApplication(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.GetDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		panic(err)
	}
	a, err := d.FrontmostApplication(NewFrontmostQueryOptions())
	if err != nil {
		panic(err)
	}
	fmt.Println(a.Description())
}

func TestApplication(t *testing.T) {
	dm := DeviceManager_Create()
	d, err := dm.GetDeviceByType(DeviceType_USB, 1000)
	if err != nil {
		panic(err)
	}
	ops := NewApplicationQueryOptions()
	apps, err := d.EnumerateApplications(ops)
	if err != nil {
		panic(err)
	}
	for _, app := range apps {
		fmt.Println(app.Description())
	}
}

func TestDeviceManager_GetDeviceById(t *testing.T) {
	dm := DeviceManager_Create()
	ls, err := dm.EnumerateDevices()
	if err != nil {
		panic(err)
	}
	for _, l := range ls {
		fmt.Println(l.Description())
	}
	d, err := dm.GetDeviceByType(DeviceType_REMOTE, 100)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(d.Description())
}

func TestManager(t *testing.T) {
	dm := DeviceManager_Create()
	defer dm.Close()
	ls, err := dm.EnumerateDevices()
	if err != nil {
		panic(err)
	}
	for _, l := range ls {
		fmt.Println(l.Description())
		fmt.Println(l.Description())
	}
	dm.OnAdded(func(device *Device) {
		fmt.Println("色号被添加1:",device.Name())
	})
	dm.OnAdded(func(device *Device) {
		fmt.Println("添加新设备:",device.Name())
	})
	dm.AddRemoteDevice("127.0.0.1",NewRemoteDeviceOptions())
	fmt.Println("wait devices change...")
	for {
		time.Sleep(time.Second)
	}

}

func TestDevice_EnumerateProcesses(t *testing.T) {
	mgr := DeviceManager_Create()
	d, _ := mgr.GetDeviceByType(DeviceType_USB, 1000)
	ops := NewProcessQueryOptions()
	ops.SelectPids = []int{1, 9125}
	pss, _ := d.EnumerateProcesses(ops)
	for _, process := range pss {
		fmt.Println(process.Description())
	}
}
