package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// Severity tags.
const (
	tagInfo    = "INFO"
	tagWarning = "WARN"
	tagError   = "ERROR"
	tagFatal   = "FATAL"
)

// LogEntry representa una entrada de registro
type LogEntry struct {
	Level     string      `json:"level"`
	Timestamp time.Time   `json:"timestamp"`
	Location  string      `json:"location"`
	Title     string      `json:"title"`
	Keys      interface{} `json:"keys"`
}

type ILogger interface {
	Info(title string, keys ...any)
	Warn(title string, keys ...any)
	Error(title string, keys ...any)
	Fatal(title string, keys ...any)
	Close()
}

type Logger struct {
	filename string
}

// NewLogger Funcion para iniciar el loger y crear la carpeta logs
func NewLogger() *Logger {
	_ = os.MkdirAll("logs", os.ModePerm)
	timeReplace := strings.Replace(time.Now().String(), ":", "-", -1)
	return &Logger{
		filename: "logs/logs " + timeReplace + ".json",
	}
}

func (l *Logger) Info(title string, keys ...any) {
	l.log(tagInfo, title, keys)
}

func (l *Logger) Warn(title string, keys ...any) {
	l.log(tagWarning, title, keys)
}

func (l *Logger) Error(title string, keys ...any) {
	l.log(tagError, title, keys)
}

func (l *Logger) Fatal(title string, keys ...any) {
	l.log(tagFatal, title, keys)
	os.Exit(1)
}

// Log registra un mensaje con título y valor any
func (l *Logger) log(typeLog string, title string, keys ...any) {
	// Obtener la ubicación del archivo que llama al logger
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
	}

	entry := LogEntry{
		Level:     typeLog,
		Timestamp: time.Now(),
		Location:  fmt.Sprintf("%s:%d", file, line),
		Title:     title,
		Keys:      keys,
	}

	// Convertimos el registro a JSON
	jsonData, err := json.Marshal(entry)
	if err != nil {

	}

	// Imprimir el registro en la consola
	fmt.Println(string(jsonData))

	fileHandle, _ := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Verificar si el archivo está vacío
	fileInfo, _ := fileHandle.Stat()
	if fileInfo.Size() == 0 {
		if _, err := fileHandle.WriteString("[\n"); err != nil {
		}
	} else {
		if _, err := fileHandle.WriteString(",\n"); err != nil {
		}
	}

	defer func(fileHandle *os.File) {
		_ = fileHandle.Close()
	}(fileHandle)

	// Escribir el registro JSON en el archivo
	if _, err := fileHandle.WriteString(fmt.Sprintf("%s", jsonData)); err != nil {
	}
}

// Close Funcion para agregar "]" al final del archivo antes de apagar el servicio
func (l *Logger) Close() {
	fileHandle, _ := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if _, err := fileHandle.WriteString("\n]"); err != nil {
	}
	_ = fileHandle.Close()
}
