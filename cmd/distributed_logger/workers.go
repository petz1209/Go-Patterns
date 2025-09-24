package main

import (
	"time"
)

func WorkerFactory(command string) func(logger *Logger) {
	switch command {

	case "A":
		return WorkerA
	case "B":
		return WorkerB
	case "C":
		return WorkerC

	default:
		return DefaultWorker
	}

}

// just a small worker function that just does some random stuff

func WorkerA(logger *Logger) {
	//fmt.Println("A")
	logger.Info("starting...")
	time.Sleep(2 * time.Second)
	logger.Info("sept for 2")
	logger.Info("I am all Done")
}

func WorkerB(logger *Logger) {
	//fmt.Println("A")
	logger.Info("starting...")
	time.Sleep(3 * time.Second)
	logger.Info("sept for 3")
	logger.Info("I am all Done")

}

func WorkerC(logger *Logger) {
	//fmt.Println("A")
	logger.Info("starting...")
	time.Sleep(4 * time.Second)
	logger.Info("sept for 4")
	logger.Info("I am all Done")

}

func DefaultWorker(logger *Logger) {
	//fmt.Println("A")
	logger.Info("starting...")
	time.Sleep(5 * time.Second)
	logger.Info("sept for 5")
	logger.Info("I am all Done")
}
