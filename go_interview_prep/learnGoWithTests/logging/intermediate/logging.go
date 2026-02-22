package intermediate

import (
	"io/fs"
	"log"
	"os"
)

func main(){
	 log.Println("This is a log message")

	 log.SetPrefix("INFO: ")
	 log.Println("This is an info msg")

	 // Log Flags
	 log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	 	 log.Println("This is an new msg")

	//
	infoLogger.Println("This is an info logger")
	warnLogger.Println("This is an warning logger")

	file, err := os.OpenFile("logger.log", os.O_CREATE | os.O_APPEND |os.O_WRONLY, fs.FileMode(0666))
	
	if err != nil{
		errorLogger.Fatalln(err)
	}

	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate | log.Ltime)
	debugLogger.Println("This is a debug message")
	debugLogger.Println("This is a debug message2")


}

var (
infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime)
warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate | log.Ltime)
errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate | log.Ltime)
)