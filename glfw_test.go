package glfw

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func TestInitAndTerminate(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetVersion(t *testing.T) {
	fmt.Println(GetVersion())
}

func TestGetVersionString(t *testing.T) {
	fmt.Println(GetVersionString())
}

func TestGetMonitors(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	monitors := GetMonitors()
	for _, m := range monitors {
		fmt.Println(m)
	}

	fmt.Println(GetPrimaryMonitor())

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetMonitorPos(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m := GetPrimaryMonitor()
	fmt.Println(GetMonitorPos(m))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetMonitorWorkArea(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m := GetPrimaryMonitor()
	fmt.Println(GetMonitorWorkArea(m))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetMonitorPhysicalSize(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m := GetPrimaryMonitor()
	fmt.Println(GetMonitorPhysicalSize(m))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetMonitorContentScale(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m := GetPrimaryMonitor()
	fmt.Println(GetMonitorContentScale(m))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetMonitorName(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m := GetPrimaryMonitor()
	fmt.Println(GetMonitorName(m))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestMonitorUserPointer(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m := GetPrimaryMonitor()
	str := "awa"
	SetMonitorUserPointer(m, unsafe.Pointer(&str))
	strGet := GetMonitorUserPointer(m)
	fmt.Println(*(*string)(strGet))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}
