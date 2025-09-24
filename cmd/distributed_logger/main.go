package main

import (
	"fmt"
	"sync"
	"time"
)

const WORKERS = 10

var Options = []string{"A", "B", "C", "DEFAULT"}

type LogMessage struct {
	LogLevel string
	TaskName string
	Time     string
	Msg      string
}

type Logger struct {
	Name string
	Chan chan LogMessage
}

func (l *Logger) Error(msg string) {
	l.Chan <- LogMessage{LogLevel: "ERROR", TaskName: l.Name, Msg: msg, Time: time.Now().Format("2006-01-02 15:04:05.0000")}
}
func (l *Logger) Info(msg string) {
	l.Chan <- LogMessage{LogLevel: "INFO", TaskName: l.Name, Msg: msg, Time: time.Now().Format("2006-01-02 15:04:05.0000")}
}

func (l *Logger) Close() {
	close(l.Chan)
}

func NewLogManager() *LogManager {
	return &LogManager{Channels: make([]chan LogMessage, 0)}
}

type LogManager struct {
	Channels []chan LogMessage
}

func (a *LogManager) NewLogger(name string) *Logger {
	ch := make(chan LogMessage, 1)
	a.Channels = append(a.Channels, ch)
	return &Logger{Name: name, Chan: ch}
}

// listens to all the log messages
func (a *LogManager) Listen() {

	var wt sync.WaitGroup
	for _, ch := range a.Channels {
		wt.Add(1)
		go func() {
			defer wt.Done()
			for msg := range ch {
				fmt.Printf("%s %s: time: %s, message: %s\n", msg.TaskName, msg.LogLevel, msg.Time, msg.Msg)
			}
		}()
	}

	wt.Wait()
}

func main() {

	// -----------------------------------------------------------------------------
	// this is where I need to setup the logger stuff
	logMM := NewLogManager()

	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {

		choice := i % len(Options)
		logger := logMM.NewLogger("task_" + Options[choice])
		handler := WorkerFactory(Options[choice])
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer logger.Close()
			handler(logger)
		}()
	}

	logMM.Listen()
	wg.Wait()

}
