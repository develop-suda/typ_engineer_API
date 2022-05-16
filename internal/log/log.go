package logs

import (
	"log"
	"os"
	"fmt"
	"runtime"
	"strconv"
	"unicode/utf8"

	def "github.com/develop-suda/typ_engineer_API/common"
	carbon "github.com/golang-module/carbon"
)

func WriteLog(message interface{}, logType string) {

	year,month,day := carbon.Now().Date()
	monthStr := strconv.Itoa(month)
	dayStr := strconv.Itoa(day)

	// 月、日が一桁の場合最初の0が付かないため、一桁の場合は0を付ける
	if utf8.RuneCountInString(monthStr) == 1 {
		monthStr = "0" + monthStr
	}
	if utf8.RuneCountInString(dayStr) == 1 {
		dayStr = "0" + dayStr
	}

	today := strconv.Itoa(year) + monthStr + dayStr
	WriteLogPath := settingLogPath(logType)

	loggingSettings(today+".log", logType)

	log.Println(message)

	if logType == def.ERROR {
		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		value := fmt.Sprintf("\ncall:%s\ndata:%s\nfile:%s:%d\n", f.Name(), "test", file, line)
		log.Println(value)
	}

	_, err := os.Open(WriteLogPath + today + ".log")
	if err != nil {
		log.Fatalln("Exit", err)
	}
}

func loggingSettings(fileName string, logType string) {

	WriteLogPath := settingLogPath(logType)
	logfile, _ := os.OpenFile(WriteLogPath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(logfile)

}

func settingLogPath(logType string) string {

	var WriteLogPath string

	if logType == def.ERROR {
		WriteLogPath = def.ERROR_LOGS_PATH
	} else if logType == def.NORMAL {
		WriteLogPath = def.NORMAL_LOGS_PATH
	} else {
		return "error"
	}
	return WriteLogPath

}
