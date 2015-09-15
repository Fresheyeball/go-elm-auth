package main

import "log"

func attempt(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func attemptGet(x interface{}, err error) interface{} {
	attempt(err)
	return x
}
