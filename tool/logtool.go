package tool

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var datetime string

func init()  {
	year, month, day := time.Now().Date()
	datetime = fmt.Sprintf("%v", year)+ fmt.Sprintf("%v", month) + fmt.Sprintf("%v", day)
}

// 日志写入 INFO
func LogINFOAdmin(s interface{}) {
	File, err := os.OpenFile("./logFileAdmin/logAdimin"+datetime+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer File.Close()
	log.Logger = log.Output(File)
	fmt.Printf("%v",s)
	log.Info().Msgf("%v",s)
}

// 日志写入 Error
func LogERRAdmin(s interface{}) {
	File, err := os.OpenFile("./logFileAdmin/logAdimint"+datetime+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	if err != nil {
		return
	}
	defer File.Close()
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: File,TimeFormat: "2006-01-02 15:04:05.000"})
	log.Logger = log.Output(File)
	log.Error().Msgf("您的程序出现错误，具体原因为:%v",s)
}

// 访客日志 INfo
func LogINFOVisitor(s interface{}) {
	File, err := os.OpenFile("./logFileVisitor/logvisitor"+datetime+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer File.Close()
	log.Logger = log.Output(File)
	fmt.Printf("%v",s)
	log.Info().Msgf("%v",s)
}

// 访客日志 Err
func LogERRVisitor(s interface{}) {

	File, err := os.OpenFile("./logFileVisitor/logvisitor"+datetime+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer File.Close()
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: File,TimeFormat: "2006-01-02 15:04:05.000"})
	log.Logger = log.Output(File)
	log.Error().Msgf("您的程序出现错误，具体原因为:%v",s)
}
