package frida_go

import (
	"context"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"runtime"
	"testing"
	"time"
)

func TestGC(t *testing.T) {
	go func() {
		dm := DeviceManager_Create()
		d, err := dm.FindDeviceByType(DeviceType_USB, 1000)
		if err != nil {
			t.Fatal(err)
		}
		if d==nil{
			t.Fatal(errors.New("device not found"))
		}
		fmt.Println(d.Name())
	}()
	for{
		runtime.GC()
		time.Sleep(time.Second)
	}
}
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

	p, err := d.GetProcessByName("映客直播", ProcessMatchOptions{})
	if err != nil {
		t.Fatal(err)
	}
	system_session, err := d.Attach(p.Pid(), SessionOptions{})
	if err != nil {
		t.Fatal(err)
	}

	sc, err := system_session.CreateScript(`
	console.log("ok111111111")
	rpc.exports={
		add:async function(a,b,c){
			return a+b+c
		}
	}
	`, ScriptOptions{})
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
	p, err := d.GetProcessByName("映客直播", ProcessMatchOptions{})
	if err != nil {
		t.Fatal(err)
	}
	system_session, err := d.Attach(p.Pid(), SessionOptions{})
	if err != nil {
		t.Fatal(err)
	}
	//defer system_session.Detach()
	sc, err := system_session.CreateScript(`
	console.log("ok111111111")
	rpc.exports={
		listThreads: function () {
    		return ProcessDetails.enumerateThreadsSync();
  		}
	}
	`, ScriptOptions{})
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
	system_session, err := d.Attach(0, SessionOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer system_session.Detach()

	bt, err := system_session.CompileScript(`
	console.log('hello')`, ScriptOptions{})
	if err != nil {
		t.Fatal(err)
	}
	sc, err := system_session.CreateScriptFormBytes(bt, ScriptOptions{})
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

	d, err := dm.AddRemoteDevice("127.0.0.1", RemoteDeviceOptions{})
	if err != nil {
		panic(err)
	}

	p, err := d.GetProcessByName("notepad.exe", ProcessMatchOptions{})
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
	ops:=SpawnOptions{}
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
	a, err := d.FrontmostApplication(FrontmostQueryOptions{})
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
	apps, err := d.EnumerateApplications(ApplicationQueryOptions{
		Scope:       0,
	})
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
	dm.AddRemoteDevice("127.0.0.1",RemoteDeviceOptions{})
	fmt.Println("wait devices change...")
	for {
		time.Sleep(time.Second)
	}

}

func TestDevice_EnumerateProcesses(t *testing.T) {
	mgr := DeviceManager_Create()
	d, _ := mgr.GetDeviceByType(DeviceType_USB, 1000)
	ops := ProcessQueryOptions{}
	ops.SelectPids = []uint{1, 9125}
	pss, _ := d.EnumerateProcesses(ops)
	for _, process := range pss {
		fmt.Println(process.Description())
	}
}

func TestApplication_Params(t *testing.T) {
	mgr := DeviceManager_Create()
	d, _ := mgr.GetDeviceByType(DeviceType_USB, 1000)
	dict,err:=d.QuerySystemParameters()
	if err!=nil{
	    t.Fatal(err)
	}
	fmt.Println(dict)
}

func TestNewDeviceSignalConnect(t *testing.T) {
	go func() {
		mgr := DeviceManager_Create()
		d, _ := mgr.GetDeviceByType(DeviceType_USB, 1000)
		d.OnSpawnAdded(func(spawn *SpawnDetails) {
			fmt.Println(spawn)
		})
		d.OnChildAdded(func(child *ChildDetails) {
			fmt.Println("aa")
		})
		d.OnOutput(func(data []byte, fd int, pid int) {
			fmt.Println("aa")
		})

	}()
	for{
		runtime.GC()
		time.Sleep(time.Second*1)
	}
}

func TestProcessDetails_Parameters(t *testing.T) {
	mgr := DeviceManager_Create()
	d, _ := mgr.GetDeviceByType(DeviceType_USB, 1000)
	ps,err:=d.EnumerateProcesses(ProcessQueryOptions{})
	if err!=nil{
	    t.Fatal(err)
	}
	for _, p := range ps {
		fmt.Println(p.Description())
	}
	p,err:=d.FindProcessById(10675,ProcessMatchOptions{})
	if err!=nil{
	    t.Fatal(err)
	}
	fmt.Println(p.Description())
}
func TestBus(t *testing.T) {
	mgr := DeviceManager_Create()
	d, _ := mgr.GetDeviceByType(DeviceType_USB, 1000)
	bus:=d.Bus()
	bus.OnMessage(func(sjson jsoniter.Any, data []byte) {
		fmt.Println(sjson.ToString())
	})
	isatt,err:=bus.Attach()
	if err!=nil{
	    t.Fatal(err)
	}
	fmt.Println(isatt)
	for{
		time.Sleep(time.Second*1)
	}
}