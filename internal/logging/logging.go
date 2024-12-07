package logging

import (
	"io"
	"log"
	"os"
)

func Println(msg string) {
	file, err := os.OpenFile("./log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	multiOutput := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multiOutput)

	log.Println(msg)
}
