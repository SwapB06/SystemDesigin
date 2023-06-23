package main

import "fmt"

const INFO = 1
const DEBUG = 2
const ERROR = 3

const InfoLogMessage string = "just for info"
const DebugLogMessage string = "need to debug this"
const ErrorLogMessage string = "Exception Happens"

// LogProcessor...
type LogProcessor interface {
	log(logLevel int, message string)
	setNext(LogProcessor)
}

// InfoLogProcessor implements LogProcessor
type InfoLogProcessor struct {
	next LogProcessor
}

func (i *InfoLogProcessor) log(logLevel int, message string) {
	if logLevel == INFO {
		fmt.Println("INFO:", message)
	} else {
		i.next.log(DEBUG, DebugLogMessage)
	}
}
func (i *InfoLogProcessor) setNext(next LogProcessor) {
	i.next = next
}

// DebugLogProcessor implements LogProcessor
type DebugLogProcessor struct {
	next LogProcessor
}

func (d *DebugLogProcessor) log(logLevel int, message string) {
	fmt.Println(">>", logLevel)
	if logLevel == DEBUG {
		fmt.Println("DEBUG:", message)
	} else {
		d.next.log(ERROR, ErrorLogMessage)
	}
}
func (d *DebugLogProcessor) setNext(next LogProcessor) {
	d.next = next
}

// ErrorLogProcessor implements LogPRocessor
type ErrorLogProcessor struct {
	next LogProcessor
}

func (e *ErrorLogProcessor) log(logLevel int, message string) {
	if logLevel == ERROR {
		fmt.Println("ERROR:", message)
	} else {
		return
	}
}
func (e *ErrorLogProcessor) setNext(next LogProcessor) {
	e.next = next
}

func main() {
	elm := &ErrorLogProcessor{}

	//Set next for Debug Log Processing
	dlm := &DebugLogProcessor{}
	dlm.setNext(elm)

	//Set next for Info Log Processing
	ilm := &InfoLogProcessor{}
	ilm.setNext(dlm)

	ilm.log(INFO, InfoLogMessage)
	dlm.log(DEBUG, DebugLogMessage)
	elm.log(ERROR, ErrorLogMessage)

	fmt.Println("==============")
	ilm.log(DEBUG, DebugLogMessage)
	ilm.log(ERROR, ErrorLogMessage) // **It should print Error but instead it is printing DEBUG

	fmt.Println("==============")
	dlm.log(ERROR, ErrorLogMessage)

	//dlm.log(INFO, InfoLogMessage)
}
