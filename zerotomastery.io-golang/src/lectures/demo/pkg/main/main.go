// run with `go run ./demo/pkg/main`
package main


// must used full path starting with 
// module coursecontent
//
// since that module name is defined in "go.mod"
// for this project
// 
// coursecontent      <<== from go.mod
//   demo/pkg/display <<== folder path
//   demo/pkg/msg     <<== folder path

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("calling a package function")
}
