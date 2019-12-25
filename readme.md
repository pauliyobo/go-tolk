# go-tolk
[godoc](https://godoc.org/github.com/pauliyobo/go-tolk)

Go-tolk is a wrapper over the DLL of the Tolk library which allows an application to communicate with the user's screenreader
## Notes
* This wrapper does not statically link with tolk.dll but uses syscall to call the library functions.
* this wrapper makes call to a dynamic link library (DLL) that is compiled for 32bit systems
If you want to use a 64bit just replace the DLL with a 64bit one.
If instead you're using the library that's included in the repository, keep in mind that if you  have a 64bit compiler you'll have to do the following:
```batch
set GOARCH=386
set GOOS=windows even though it should already be
```
## Installation
To get go-tolk type the following in  your command prompt

```batch
go get -v github.com/pauliyobo/go-tolk
```

## Usage
```go
package main

import (
	"github.com/pauliyobo/go-tolk"
	"fmt"
)

func main() {
	tolk.Load()
	sr := tolk.DetectScreenReader()
	fmt.Printf("my screenreader is: %s", sr)
	tolk.Output("hello", false)
	tolk.Braille("hello from braille")
	tolk.Speak("hello from speech only", true)
	tolk.Unload()
}
```
## contributing
Pull requests are welcome