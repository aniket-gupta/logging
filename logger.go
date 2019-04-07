package logging

import (
	"io"
	"log"
	"os"
	"sync"
)

// Level type
type level int

const (
	// DEBUG level
	DEBUG level = iota
	// INFO level
	INFO
	// WARNING level
	WARNING
	// ERROR level
	ERROR
	// FATAL level
	FATAL

	flag = log.Ldate | log.Ltime
)

var (
	currLvl = INFO
	mu      sync.RWMutex
	prefix  = map[level]string{
		DEBUG:   "DEBUG: ",
		INFO:    "INFO: ",
		WARNING: "WARNING: ",
		ERROR:   "ERROR: ",
		FATAL:   "FATAL: ",
	}
)

// SetLevel set logging level
func SetLevel(lvl level) {
	mu.Lock()
	defer mu.Unlock()
	currLvl = lvl
}

// Logger ...
type Logger struct {
	DEBUG   LoggerInterface
	INFO    LoggerInterface
	WARNING LoggerInterface
	ERROR   LoggerInterface
	FATAL   LoggerInterface
}

// New returns instance of Logger
func New(out, errOut io.Writer, f Formatter) Logger {
	// Fall back to stdout if out not set
	if out == nil {
		out = os.Stdout
	}

	// Fall back to stderr if errOut not set
	if errOut == nil {
		errOut = os.Stderr
	}

	// Fall back to DefaultFormatter if f not set
	if f == nil {
		f = new(DefaultFormatter)
	}

	debug := &Wrapper{lvl: DEBUG, formatter: f, logger: log.New(out, f.GetPrefix(DEBUG)+prefix[DEBUG], flag)}
	info := &Wrapper{lvl: INFO, formatter: f, logger: log.New(out, f.GetPrefix(INFO)+prefix[INFO], flag)}
	warning := &Wrapper{lvl: INFO, formatter: f, logger: log.New(out, f.GetPrefix(WARNING)+prefix[WARNING], flag)}
	errLogger := &Wrapper{lvl: INFO, formatter: f, logger: log.New(errOut, f.GetPrefix(ERROR)+prefix[ERROR], flag)}
	fatal := &Wrapper{lvl: INFO, formatter: f, logger: log.New(errOut, f.GetPrefix(FATAL)+prefix[FATAL], flag)}

	return Logger{
		DEBUG:   debug,
		INFO:    info,
		WARNING: warning,
		ERROR:   errLogger,
		FATAL:   fatal,
	}
}

// Wrapper ...
type Wrapper struct {
	lvl       level
	formatter Formatter
	logger    LoggerInterface
}

// Print ...
func (w *Wrapper) Print(v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		v = w.formatter.Format(w.lvl, v...)
		v = append(v, w.formatter.GetSuffix(w.lvl))
		w.logger.Print(v...)
	}
}

// Printf ...
func (w *Wrapper) Printf(format string, v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		suffix := w.formatter.GetSuffix(w.lvl)
		v = w.formatter.Format(w.lvl, v...)
		w.logger.Printf("%s"+format+suffix, v...)
	}
}

// Println ...
func (w *Wrapper) Println(v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		v = w.formatter.Format(w.lvl, v...)
		v = append(v, w.formatter.GetSuffix(w.lvl))
		w.logger.Println(v...)
	}
}

// Fatal ...
func (w *Wrapper) Fatal(v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		v = w.formatter.Format(w.lvl, v...)
		v = append(v, w.formatter.GetSuffix(w.lvl))
		w.logger.Fatal(v...)
	}
}

// Fatalf ...
func (w *Wrapper) Fatalf(format string, v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		suffix := w.formatter.GetSuffix(w.lvl)
		v = w.formatter.Format(w.lvl, v...)
		w.logger.Fatalf("%s"+format+suffix, v...)
	}
}

// Fatalln ...
func (w *Wrapper) Fatalln(v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		v = w.formatter.Format(w.lvl, v...)
		v = append(v, w.formatter.GetSuffix(w.lvl))
		w.logger.Fatalln(v...)
	}
}

// Panic ...
func (w *Wrapper) Panic(v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		v = w.formatter.Format(w.lvl, v...)
		v = append(v, w.formatter.GetSuffix(w.lvl))
		w.logger.Fatal(v...)
	}
}

// Panicf ...
func (w *Wrapper) Panicf(format string, v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		suffix := w.formatter.GetSuffix(w.lvl)
		v = w.formatter.Format(w.lvl, v...)
		w.logger.Panicf("%s"+format+suffix, v...)
	}
}

// Panicln ...
func (w *Wrapper) Panicln(v ...interface{}) {
	mu.RLock()
	mu.RUnlock()
	if w.lvl >= currLvl {
		v = w.formatter.Format(w.lvl, v...)
		v = append(v, w.formatter.GetSuffix(w.lvl))
		w.logger.Panicln(v...)
	}
}
