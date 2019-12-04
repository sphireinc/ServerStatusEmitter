package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// HandleError logs and prints the error to console
func LogErrorAndContinue(err error) {
	if err != nil {
		log.Println(Format(err, "ERROR"))
		fmt.Println(err, "ERROR")
	}
}

// HandleErrorFatal logs and prints the error then exits the application
func HandleErrorFatal(err error) {
	LogErrorAndContinue(err)
	os.Exit(1)
}

// Format allows us to know which file and which function is executing at the moment.
func Format(err error, status string) string {
	status = strings.ToUpper(status)
	return fmt.Sprintf("%s %s", status, err.Error())
}

// Trace allows us to know which file and which function is executing at the moment.
func Trace(err error) string {
	p := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, p)
	f := runtime.FuncForPC(p[0])
	file, line := f.FileLine(p[0])

	file = lastString(file, "/")
	funcName := lastString(f.Name(), "/")

	return fmt.Sprintf("%s <%d> %s(): %s", file, line, funcName, err.Error())
}

func lastString(msg string, separator string) string {
	splitString := strings.Split(msg, separator)
	return splitString[len(splitString) - 1]
}