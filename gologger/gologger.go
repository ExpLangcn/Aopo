package gologger

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Level defines all the available levels we can log at
type Level int

// Available logging levels
const (
	Null Level = iota
	Fatal
	Silent
	Label
	Misc
	Error
	Warning
	Info
	Debug
	Verbose
)

var (
	// UseColors can be used to control coloring of the output
	UseColors = true
	// MaxLevel is the maximum level to log at. By default, logging
	// is done at Info level. Using verbose will display all the errors too,
	// Using silent will display only the most relevant information.
	MaxLevel = Info

	labels = map[Level]string{
		Warning: "Warning",
		Error:   "Error",
		Label:   "WRN",
		Fatal:   "Fatal",
		Debug:   "DEBUG",
		Info:    "INF",
	}

	mutex  = &sync.Mutex{}
	output = colorable.NewColorableStdout()
)

var stringBuilderPool = &sync.Pool{New: func() interface{} {
	return new(strings.Builder)
}}

// wrap wraps a given label for a message to a logg-able representation.
// It checks if colors are specified and what level we are logging at.
func wrap(label string, level Level) string {
	// Check if we are not using colors, if not, return
	if false {
		return label
	}

	switch level {
	case Silent:
		return label
	case Info, Verbose:
		return aurora.Blue(label).String()
	case Fatal:
		return aurora.Bold(aurora.Red(label)).String()
	case Error:
		return aurora.Red(label).String()
	case Debug:
		return aurora.Magenta(label).String()
	case Warning, Label:
		return aurora.Yellow(label).String()
	default:
		return label
	}
}

// getLabel generates a label for a given message, depending on the level
// and the label passed.
func getLabel(level Level, label string, sb *strings.Builder) {
	timeStr := time.Now().Format("01-02 15:04:05")
	switch level {
	case Silent, Misc:
		return
	case Error, Fatal, Info, Warning, Debug, Label:
		sb.WriteString("[" + aurora.Yellow(timeStr).String() + "] [")
		sb.WriteString(wrap(labels[level], level))
		sb.WriteString("]")
		sb.WriteString(" ")
		return
	case Verbose:
		sb.WriteString("[" + aurora.Yellow(timeStr).String() + "] [")
		sb.WriteString(wrap(label, level))
		sb.WriteString("]")
		sb.WriteString(" ")
		return
	default:
		return
	}
}

// log logs the actual message to the screen
func log(level Level, label string, format string, args ...interface{}) {
	// Don't log if the level is null
	if level == Null {
		return
	}

	if level <= MaxLevel {
		// Build the log message using the string builder pool
		sb := stringBuilderPool.Get().(*strings.Builder)

		// Get the label and append it to string builder
		getLabel(level, label, sb)

		message := fmt.Sprintf(format, args...)
		sb.WriteString(message)

		if strings.HasSuffix(message, "\n") == false {
			sb.WriteString("\n")
		}

		mutex.Lock()
		switch level {
		case Silent:
			fmt.Fprint(os.Stdout, sb.String())
		default:
			fmt.Fprint(output, sb.String())
		}
		mutex.Unlock()

		sb.Reset()
		stringBuilderPool.Put(sb)
	}
}

// Infof writes a info message on the screen with the default label
func Infof(format string, args ...interface{}) {
	log(Info, "", format, args...)
	WriteFile(format, GetExec()+"/result.txt")
}

// Warningf writes a warning message on the screen with the default label
func Warningf(format string, args ...interface{}) string {
	log(Warning, "", format, args...)
	return format
}

// Errorf writes an error message on the screen with the default label

// Debugf writes an error message on the screen with the default label

// Verbosef writes a verbose message on the screen with a tabel

// Silentf writes a message on the stdout with no label

// Fatalf exits the program if we encounter a fatal error

// Printf prints a string on screen without any extra stuff
func Printf(format string, args ...interface{}) {
	log(Misc, "", format, args...)
	//fmt.Println(GetExec() + "/test.txt")
	WriteFile(format, GetExec()+"/result.txt")
}

var logSync sync.Mutex

// Labelf prints a string on screen with a label interface
func WriteFile(result string, filename string) {
	//fmt.Println("this string is:", add_data)
	logSync.Lock()
	f, _ := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0777) //读写模式打开，写入追加
	defer f.Close()
	f.Write([]byte(result + "\n"))
	//模拟执行时间
	//time.Sleep(1000000000)
	logSync.Unlock()
	//fmt.Println(fmt.Sprintf(result+"%d", num))
	f.Close()
}
func GetExec() string {
	exePath, _ := os.Executable()
	return filepath.Dir(exePath)
}
