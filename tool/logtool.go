package tool

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

// 日志写入 INFO
func LogINFO(s string) {
	File, err := os.OpenFile("./logFile/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	if err != nil {
		File, err = os.Create("./logFile/log.txt")
	}
	defer File.Close()
	log.Logger = log.Output(File)
	log.Info().Msg(s)
}

// 日志写入 Error
func LogERR(s interface{}) {
	fmt.Println("aaaa")
	File, err := os.OpenFile("./logFile/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	if err != nil {
		fmt.Println("bbbbbb")
		File, err = os.Create("./logFile/log.txt")
	}
	defer File.Close()
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: File,TimeFormat: "2006-01-02 15:04:05.000"})
	log.Logger = log.Output(File)
	log.Error().Msgf("您的程序出现错误，具体原因为:%v",s)
}
