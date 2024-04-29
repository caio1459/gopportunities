package config

import (
	"io"
	"log"
	"os"
)

// Logger é uma estrutura para configurar diferentes níveis de log.
type Logger struct {
	debug   *log.Logger // Logger para mensagens de depuração.
	info    *log.Logger // Logger para mensagens informativas.
	warning *log.Logger // Logger para mensagens de aviso.
	err     *log.Logger // Logger para mensagens de erro.
	writer  io.Writer   // Escritor para os loggers.
}

// NewLogger cria e retorna uma nova instância de Logger.
func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),   // Logger para mensagens de depuração.
		info:    log.New(writer, "INFO: ", logger.Flags()),    // Logger para mensagens informativas.
		warning: log.New(writer, "WARNING: ", logger.Flags()), // Logger para mensagens de aviso.
		err:     log.New(writer, "ERR: ", logger.Flags()),     // Logger para mensagens de erro.
		writer:  writer,                                       // Escritor para os loggers.
	}
}

// DebugF registra uma mensagem de depuração com formatação.
func (l *Logger) DebugF(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// InfoF registra uma mensagem informativa com formatação.
func (l *Logger) InfoF(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// WarningF registra uma mensagem de aviso com formatação.
func (l *Logger) WarningF(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

// ErrF registra uma mensagem de erro com formatação.
func (l *Logger) ErrF(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

// GetLogger cria e retorna um novo Logger com o prefixo especificado.
func GetLogger(p string) *Logger {
	// Cria um novo Logger com o prefixo especificado.
	logger := NewLogger(p)
	return logger
}
