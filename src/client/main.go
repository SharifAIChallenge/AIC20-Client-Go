package main

import (
	"os"
	"strconv"
)

const GlobalVerboseFlag = false

var ArgNames = [...]string{"AICHostIP", "AICHostPort", "AICToken", "AICRetryDelay"}
var ArgDefaults = [...]string{"localhost", "7099", "00000000000000000000000000000000", "1000"}

func main() {
	run(getArgs())
}
func run(args []string) {
	port, _ := strconv.Atoi(args[1])
	retryDelay, _ := strconv.ParseInt(args[3], 10, 64)
	controller := Controller{host: args[0], port: port, token: args[2], retryDelay: retryDelay}
	controller.Start()
}

func getArgs() []string {
	args := make([]string, len(ArgNames))
	for i := range ArgNames {
		var ok bool
		args[i], ok = os.LookupEnv(ArgNames[i])
		if !ok {
			args[i] = ArgDefaults[i]
		}
	}
	return args
}
