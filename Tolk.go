package tolk
import (
	"fmt"
	"syscall"
	"unsafe"
	"strings"
)

var lib = syscall.NewLazyDLL("Tolk.dll")
var loaded = false


// Load loads the library
//
// This function must be the first to be called when working with with tolk
func Load() {
	if loaded == false {
		lib.NewProc("Tolk_Load").Call()
	}
	loaded = true	
}

// Unload must be called after all the operations with library are finished
func Unload() error {
	if loaded == true {
		lib.NewProc("Tolk_Unload").Call()
	}
	return nil
}

func TrySapi(sapi bool) {
	lib.NewProc("Tolk_TrySAPI").Call(BoolToUintptr(sapi))
}

func PreferSapi(sapi bool) {
	lib.NewProc("Tolk_PreferSAPI").Call(BoolToUintptr(sapi))
}

// DetectScreenReader will try to detect the current screenreader that is running on your computer
func DetectScreenReader() string {
	ret, _, _ := lib.NewProc("Tolk_DetectScreenReader").Call()
	// replacing spaces with nothing because for some reason the letters are returned with a space between them
	return strings.ReplaceAll(CharPToString(ret), string(0), "")
}

func HasSpeech() bool {
	ret, _, _ := lib.NewProc("Tolk_HasSpeech").Call()
	return UintptrToBool(ret)
}

func HasBraille() bool {
	ret, _, _ := lib.NewProc("Tolk_HasBraille").Call()
	return UintptrToBool(ret)
}

func Output(text string, interrupt bool) bool {
	ret, _, _ := lib.NewProc("Tolk_Output").Call(StringToUintptr(text), BoolToUintptr(interrupt))
	return UintptrToBool(ret)
}

func Speak(text string, interrupt bool) bool {
	ret, _, _ := lib.NewProc("Tolk_Speak").Call(StringToUintptr(text), BoolToUintptr(interrupt))
	return UintptrToBool(ret)
}

func Braille(text string) bool {
	ret, _, _ := lib.NewProc("Tolk_Braille").Call(StringToUintptr(text))
	return UintptrToBool(ret)
}

func Silence() bool {
	ret, _, _ := lib.NewProc("Tolk_Output").Call()
	return UintptrToBool(ret)
}

func CharPToString(u uintptr) string {
	return string((*[unsafe.Sizeof(u)*2]byte)(unsafe.Pointer(u))[:])
}

func StringToUintptr(s string) uintptr {
	ret, err  := syscall.UTF16PtrFromString(s)
	if err != nil {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(ret))
}

func BoolToUintptr(v bool) uintptr {
	var ret int
	if v == true {
		ret = 1
	}	else {
		ret = 0
	}
	return uintptr(ret)
}

func UintptrToBool(v uintptr) bool {
	return v != 0
}

func main() {
	Load()
	sr := DetectScreenReader()
	fmt.Println(sr)
	//Output(sr, false)
	Unload()
}