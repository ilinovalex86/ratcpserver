package ratcpserver

import (
	"log"
	"os"
	"time"
)

func ToLog(logFile string, data string, flag bool) {
	var createLogFile = func() {
		file, err := os.Create(logsDir + "/" + logFile)
		if err != nil {
			log.Fatal("err create log file: " + logsDir + "/" + logFile)
		}
		_, err = file.WriteString(data)
		if err != nil {
			log.Fatal("write data to log: ", logFile)
		}
		file.Close()
	}
	var toLog = func() {
		file, err := os.OpenFile(logsDir+"/"+logFile, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("open log file: ", logsDir+"/"+logFile)
		}
		_, err = file.WriteString(data)
		if err != nil {
			file.Close()
			log.Fatal("write data to log", logsDir+"/"+logFile)
		}
		file.Close()
		if flag {
			log.Fatal(data)
		}
	}
	data = time.Now().Format("02.01.2006 15:04:05") + " " + data + "\n"
	for {
		logs.mu.Lock()
		if status, ok := logs.m[logFile]; ok && status {
			logs.mu.Unlock()
			time.Sleep(time.Millisecond)
			continue
		} else if ok && !status {
			logs.m[logFile] = true
			logs.mu.Unlock()
			toLog()
			logs.mu.Lock()
			logs.m[logFile] = false
			logs.mu.Unlock()
			return
		} else {
			logs.m[logFile] = true
			logs.mu.Unlock()
			createLogFile()
			logs.mu.Lock()
			logs.m[logFile] = false
			logs.mu.Unlock()
			return
		}
	}
}
