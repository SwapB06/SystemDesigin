package main

//abstract class logprocessor
INFO := 1
DEBUG := 2
ERROR := 3

//constructor
LogProcessor(LogProcessor logProcessor){
	this.nextLogProcessor = logProcessor
}

func log(int logLevel, message string){
	if nextLogProcessor != nil{
		nextLogProcessor.log(logLevel, message)
	}
}

//InfoLogProcessor extends LogProcessor
InfoLogProcessor(LogProcessor nextLogProcessor) { super(nextLogProcessor)}\
func log(logLevel int, message string){
	if logLevel == INFO{
		fmt.Println("INFO:", message)
	} else{
		super.log(logLevel, message)
	}
}

//ErrorLogProcessor extends LogPRocessor
ErrorLogProcessor(LogProcessor nextLogProcessor) { super(nextLogProcessor)}\
func log(logLevel int, message string){
	if logLevel == ERROR{
		fmt.Println("ERROR:", message)
	} else{
		super.log(logLevel, message)
	}
}

//DebugLogProcessor extends LogProcessor
DebugLogProcessor(LogProcessor nextLogProcessor) { super(nextLogProcessor)}\
func log(logLevel int, message string){
	if logLevel == DEBUG{
		fmt.Println("DEBUG:", message)
	} else{
		super.log(logLevel, message)
	}
}

func main() {
	logProcessor logObject = newInfoLogProcessor(new DebugLogProcessor(new ErrorLogProcessor(nil)))

	logObject.log(LogProcessor.ERROR, "Exception Happens")
	logObject.log(LogProcessor.DEBUG, "need to debug this")
	logObject.log(LogProcessor.INFO, "just for info")
}