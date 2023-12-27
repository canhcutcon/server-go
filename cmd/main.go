package main

import "server-go/internal/pkg/log"

func main() {
	log.Init()
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
}
