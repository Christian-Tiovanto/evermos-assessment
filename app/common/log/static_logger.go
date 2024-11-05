package log

import "sync"

var (
	once         sync.Once
	staticLogger *Logs
)

// GetLogger is to get singleton logger
func GetLogger() *Logs {
	once.Do(func() {
		staticLogger = NewLogger()
	})
	return staticLogger
}
