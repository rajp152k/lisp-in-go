package main

import (
	"log"
	"os"
)

var l *log.Logger

func init(){
	l = log.New(os.Stdout, "[lisp-in-go]: ", log.LstdFlags )
}
