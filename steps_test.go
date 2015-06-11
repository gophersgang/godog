package godog

import (
	"log"
	"regexp"
)

func init() {
	f := StepHandlerFunc(func(args ...interface{}) error {
		log.Println("step triggered")
		return nil
	})
	Step(regexp.MustCompile("hello"), f)
}
