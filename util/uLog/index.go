package uLog

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
)

func Error(v ...interface{}) {
	log.Println(aurora.Red(fmt.Sprintf("[ERROR] %v", v)))
}
func Warn(v ...interface{}) {
	log.Println(aurora.Yellow(fmt.Sprintf("[WARNING] %v", v)))
}

func OK(v ...interface{}) {
	log.Println(aurora.Green(fmt.Sprintf("[OK] %v", v)))
}
