package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
	HEAD    = "HEAD"
	SQL     = "SQL"
)

var l *Logger
var once sync.Once

type LogInterface interface {
	Write(c interface{})
	Output(c interface{})
	All(c interface{})
	FULL(t string , c interface{}, write bool)
}

type Logger struct {
	dir    string
	format string
	suffix string
}

func (l Logger) Write(c interface{}) {
	file := findFile()
	defer file.Close()
	log.SetOutput(file)
	log.Println(c)
}

func (l Logger) Output(c interface{}) {
	fmt.Println(c)
}

func (l Logger) All(c interface{}) {
	l.Write(c)
	l.Output(c)
}

func (l Logger) FULL(t string, c interface{}, write bool) {
	if write {
		l.Write("[" + t + "] " + fmt.Sprintln(c))
	}
	l.Output("[" + t + "] " + fmt.Sprintln(c))
}

func GetLogger() *Logger {
	// Singleton
	once.Do(func() {
		l = &Logger{
			"logs", "2006-01-02", "log",
		}
	})
	return l
}

func findFile() *os.File {
	findDir(l.dir)
	file, err := os.OpenFile(l.dir+"/"+time.Now().Format(l.format)+"."+l.suffix, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return file
}

// TODO Refactoring to Utils oceanGO package
func findDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
}