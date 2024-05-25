package helper

import (
	"log"
	"os"
)

func CheckErr(err error, errStr string) bool {
	if err  != nil {
		println(errStr , "\n error:", err)
		return true
	}
	return false
}

func FatalErr(err error, errStr string) {
	if err  != nil {
		println(errStr)
		log.Fatal("error :", err)
		os.Exit(1)
	}
}