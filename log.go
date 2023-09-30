package backend_utils

import (
	"fmt"
	"log"
	"runtime"
)

const (
	redColor    = "\033[31m"
	yellowColor = "\033[33m"
	blueColor   = "\033[34m"
	resetColor  = "\033[0m"
)

type printer struct{}

var Debug printer

func (printer) Debug(message string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()
	additional_info := fmt.Sprintf(message, a...)

	log.Printf(
		"%s== DEBUG == %s --FILE--:%s --FUNCTION--: %s --LINE--: %v --MESSAGE-- : %s%s\n",
		blueColor, resetColor, file, functionName, line, additional_info, resetColor)
}

func (printer) Error(message string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()
	additional_info := fmt.Sprintf(message, a...)

	log.Printf(
		"%s== ERROR == %s --FILE--:%s --FUNCTION--: %s --LINE--: %v --MESSAGE-- : %s%s\n",
		redColor, resetColor, file, functionName, line, additional_info, resetColor)
}
func (printer) Info(message string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()
	additional_info := fmt.Sprintf(message, a...)

	log.Printf(
		"%s== INFO == %s --FILE--:%s --FUNCTION--: %s --LINE--: %v --MESSAGE-- : %s%s\n",
		yellowColor, resetColor, file, functionName, line, additional_info, resetColor)
}
