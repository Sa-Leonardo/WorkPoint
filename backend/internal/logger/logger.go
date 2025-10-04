package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Init inicializa o logger global
func Init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)                 // manda pro console
	log.SetLevel(logrus.DebugLevel)          // nível mínimo de log
	log.SetFormatter(&logrus.JSONFormatter{}) // saída em JSON (bom pra produção)
}

// Info loga mensagens de informação
func Info(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Info(msg)
}

// Error loga mensagens de erro
func Error(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Error(msg)
}

// Debug loga mensagens de debug
func Debug(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Debug(msg)
}
