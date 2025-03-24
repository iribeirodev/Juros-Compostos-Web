package logger

import (
	"log"
	"os"
)

// Provê um item de log formatado
func Logger() *log.Logger {
	return log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
