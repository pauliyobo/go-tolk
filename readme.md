# go-tolk
[godoc](https://godoc.org/github.com/pauliyobo/go-tolk)
go-tolk is a wrapper over the DLL of the Tolk library which allows an application to communicate with the user's screenreader
## Notes
* This wrapper does not statically link with tolk.dll but uses syscall to call the library functions.
* this wrapper makes call to a dynamic link library (DLL) that is compiled for 32bit systems
If you have a 64bit go compiler you'll have to do the following for successfully using the wrapper
```batch
set GOARCH=386
set GOOS=windows even though it should already be
```
## Usage
```go
package main

import (
	"tolk"
	"fmt"
)

func main() {
	err : = Tolk.Load()
	if err != nil {
		fmt.Println("Tolk loaded successfully")
	}
	sr, e := Tolk.DetectScreenReader()
	if e != nil {
		fmt.Printf("Screenreader found: %s", sr)
	}
	Tolk.Output("hello", false)
	Tolk.Braille("hello from braille")
	Tolk.Speak("hello from speech only", true)
	e2 := Tolk.Unload()
	if e2 !=  nil {
		fmt.Println("Tolk Unloaded")
	}
}
```
## bugs
The function Tolk.DetectScreenReader() returns null characters when the name of the screenreader is returned. This will result on having an output such as
N V D A
## todo
* Improving documentation
* Additionally generating documentation with godoc
## contributing
Pull requests are welcome