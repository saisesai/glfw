package glfw

import (
	"fmt"
	"runtime"
	"testing"
	"time"
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

	monitors, err := GetMonitors()
	if err != nil {
		panic(err)
	}
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

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}
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

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}
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

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}
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

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}
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

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}
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

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}
	str := "awa"
	err = SetMonitorUserPointer(m, unsafe.Pointer(&str))
	if err != nil {
		panic(err)
	}
	strGet, err := GetMonitorUserPointer(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(*(*string)(strGet))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetVideoModes(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	go func() {
		m, err := GetPrimaryMonitor()
		if err != nil {
			panic(err)
		}
		fmt.Println(GetVideoModes(m))
	}()
	time.Sleep(time.Second)

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGetVideoMode(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}
	fmt.Println(GetVideoMode(m))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestSetGamma(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}

	err = SetGamma(m, 2.0)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)

	err = SetGamma(m, 1.0)
	if err != nil {
		panic(err)
	}

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestGammaRamp(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	m, err := GetPrimaryMonitor()
	if err != nil {
		panic(err)
	}

	fmt.Println(GetGammaRamp(m))
	fmt.Println(SetGammaRamp(m, nil))

	err = Terminate()
	if err != nil {
		panic(err)
	}
}

func TestWindow(t *testing.T) {
	var err error
	err = Init()
	if err != nil {
		panic(err)
	}

	err = DefaultWindowHints()
	if err != nil {
		panic(err)
	}

	w, err := CreateWindow(1280, 720, "test", nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(GetWindowPos(w))
	fmt.Println(SetWindowPos(w, 300, 300))
	fmt.Println(GetWindowSize(w))
	fmt.Println(SetWindowSizeLimits(w, 300, 300, 1280, 720))
	fmt.Println(SetWindowAspectRatio(w, 1280, 720))
	fmt.Println(SetWindowSize(w, 888, 500))

	err = SetWindowTitle(w, "awa")
	if err != nil {
		panic(err)
	}

	err = MakeContextCurrent(w)
	if err != nil {
		panic(err)
	}

	err = SwapInterval(1)
	if err != nil {
		panic(err)
	}

	for !WindowShouldClose(w) {

		err = SwapBuffers(w)
		if err != nil {
			panic(err)
		}
		err = PollEvents()
		if err != nil {
			panic(err)
		}
	}

	err = DestroyWindow(w)
	if err != nil {
		panic(err)
	}

	err = Terminate()
	if err != nil {
		panic(err)
	}
}
