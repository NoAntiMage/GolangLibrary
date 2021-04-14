package main

import {
	"os"
	"log"
	"io"
}

func main() {
	logFileName := "./" + "log" + ".txt"
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetPrefix("[write_sth_here]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)	
}