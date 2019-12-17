package sphlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// LogInfo logs and prints a message as INFO
func LogInfo(msg string) {
	log.Println(format("INFO", msg))
}

// LogWarn logs and prints a message as WARN
func LogWarn(msg string) {
	log.Println(format("WARN", msg))
}

// LogError logs and prints the sphlog to console
func LogError(err error) {
	if err != nil {
		log.Println(format("ERROR", err.Error()))
	}
}

// LogFatalError logs and prints the sphlog then exits the application
func LogFatalError(err error) {
	if err != nil {
		log.Println(format("FATAL", err.Error()))
		os.Exit(1)
	}
}

// format allows us to know which file and which function is executing at the moment.
func format(status string, msg string) string {
	status = strings.ToUpper(status)

	switch status {
	case "INFO":
		return fmt.Sprintf("%s %s", status, msg)
	case "WARN":
		return fmt.Sprintf("%s %s", status, msg)
	case "ERROR":
		return fmt.Sprintf("%s %s", status, msg)
	case "FATAL":
		return fmt.Sprintf("%s %s", status, msg)
	}
	return ""
}

// trace allows us to know which file and which function is executing at the moment.
func trace(err error) string {
	p := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, p)
	f := runtime.FuncForPC(p[0])
	file, line := f.FileLine(p[0])

	file = lastString(file, "/")
	funcName := lastString(f.Name(), "/")

	return fmt.Sprintf("%s <%d> %s(): %s", file, line, funcName, err.Error())
}

// lastString returns the last element of a string broken by separators
func lastString(msg string, separator string) string {
	splitString := strings.Split(msg, separator)
	return splitString[len(splitString)-1]
}
