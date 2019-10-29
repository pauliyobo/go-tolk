package Tolk
import (
	"syscall"
	"unsafe"
)

var 		lib = syscall.NewLazyDLL("Tolk.dll")
var loaded = false

func Load() error {
	if loaded == false {
		_, _, err := lib.NewProc("Tolk_Load").Call()
		if err != nil {
			return err
		}
		loaded = true
	}
	return nil
}

func Unload() error {
	if loaded == true {
		_, _, err := lib.NewProc("Tolk_Unload").Call()
		if err != nil {
			return err
		}
	}
	return nil
}

func TrySapi(sapi bool) error {
	_, _, err := lib.NewProc("Tolk_TrySAPI").Call(BoolToUintptr(sapi))
	if err != nil {
		return err
	}
	return nil
}

func PreferSapi(sapi bool) error {
	_, _, err := lib.NewProc("Tolk_PreferSAPI").Call(BoolToUintptr(sapi))
	if err != nil {
		return err
	}
	return nil
}

func DetectScreenReader() (string, error) {
	ret, _, err := lib.NewProc("Tolk_DetectScreenReader").Call()
	if err != nil {
		return "", err
	}
	return CharPToString(ret), nil
}

func HasSpeech() (bool, error) {
	ret, _, err := lib.NewProc("Tolk_HasSpeech").Call()
	if err != nil {
		return false, err
	}
	return UintptrToBool(ret), nil
}

func HasBraille() (bool, error) {
	ret, _, err := lib.NewProc("Tolk_HasBraille").Call()
	if err != nil {
		return false, err
	}
	return UintptrToBool(ret), nil
}

func Output(text string, interrupt bool) (bool, error) {
	ret, _, err := lib.NewProc("Tolk_Output").Call(StringToUintptr(text), BoolToUintptr(interrupt))
	if err != nil {
		return false, err
	}
	return UintptrToBool(ret), nil
}

func Speak(text string, interrupt bool) (bool, error) {
	ret, _, err := lib.NewProc("Tolk_Speak").Call(StringToUintptr(text), BoolToUintptr(interrupt))
	if err != nil {
		return false, err
	}
	return UintptrToBool(ret), nil
}

func Braille(text string) (bool, error) {
	ret, _, err := lib.NewProc("Tolk_Braille").Call(StringToUintptr(text))
	if err != nil {
		return false, err
	}
	return UintptrToBool(ret), nil
}

func Silence() (bool, error) {
	ret, _, err := lib.NewProc("Tolk_Output").Call()
	if err != nil {
		return false, err
	}
	return UintptrToBool(ret), nil
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
