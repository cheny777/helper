package syslog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Clog Clog
var Clog *LogAttributes

const (
	// LogError enable error level
	LogError int = 0x01
	// LogWarn enable warnning level
	LogWarn int = 0x02
	// LogTrace enable trace level
	LogTrace int = 0x03
	// LogInfo enable info level
	LogInfo int = 0x04
	// LogDebug enable debug level
	LogDebug int = 0x05
)

// LogAttributes Log Attributes
type LogAttributes struct {
	ServiceName string
	Flag        int
	LogIO       *lumberjack.Logger
}

func init() {
	logpacker := &lumberjack.Logger{
		Filename:   "maximo.log",
		MaxSize:    10, //MB
		MaxBackups: 10,
		MaxAge:     28, //days
	}

	Clog = &LogAttributes{
		ServiceName: "MAXIMO",
		Flag:        0x05,
		LogIO:       logpacker,
	}
}

// GetLogFileInfo 获取日志打印文件和行数
func GetLogFileInfo() string {
	f := "???"
	n := "0"

	_, file, line, ok := runtime.Caller(2)
	if ok {
		indx := strings.LastIndex(file, "/")
		if indx > 0 {
			f = file[indx+1:]
			n = fmt.Sprint(line)
		}
	}
	return f + ":" + fmt.Sprint(n)
}

// Infoln Info with ln.
func (la *LogAttributes) Infoln(a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogInfo {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgGreen).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgGreen).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgGreen).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgGreen).Fprintln(os.Stdout, a...)
	}
}

// Info Info without ln.
func (la *LogAttributes) Info(a ...interface{}) {
	if la.Flag >= LogInfo {
		color.New(color.FgGreen).Fprint(la.LogIO, a...)
		//控制台
		color.New(color.FgGreen).Fprint(os.Stdout, a...)
	}
}

// Infof Info without ln.
func (la *LogAttributes) Infof(format string, a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogInfo {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgGreen).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgGreen).Fprintf(la.LogIO, format, a...)
		color.New(color.FgGreen).Fprintln(la.LogIO)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgGreen).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgGreen).Fprintf(os.Stdout, format, a...)
		color.New(color.FgGreen).Fprintln(os.Stdout)
	}
}

// Errorln Error with ln.
func (la *LogAttributes) Errorln(a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogError {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgHiRed).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgHiRed).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgHiRed).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgHiRed).Fprintln(os.Stdout, a...)
	}
}

// Error Error without ln.
func (la *LogAttributes) Error(a ...interface{}) {
	if la.Flag >= LogError {
		color.New(color.FgHiRed).Fprint(la.LogIO, a...)
		//控制台
		color.New(color.FgHiRed).Fprint(os.Stdout, a...)
	}
}

// Errorf Error without ln.
func (la *LogAttributes) Errorf(format string, a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogError {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgHiRed).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgHiRed).Fprintf(la.LogIO, format, a...)
		color.New(color.FgHiRed).Fprintln(la.LogIO)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgHiRed).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgHiRed).Fprintf(os.Stdout, format, a...)
		color.New(color.FgHiRed).Fprintln(os.Stdout)
	}
}

// Warnln Warn with ln.
func (la *LogAttributes) Warnln(a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogWarn {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgYellow).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgYellow).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgYellow).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgYellow).Fprintln(os.Stdout, a...)
	}
}

// Warn Warn without ln.
func (la *LogAttributes) Warn(a ...interface{}) {
	if la.Flag >= LogWarn {
		color.New(color.FgYellow).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgYellow).Fprintln(os.Stdout, a...)
	}
}

// Warnf Warn without ln.
func (la *LogAttributes) Warnf(format string, a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogWarn {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgYellow).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgYellow).Fprintf(la.LogIO, format, a...)
		color.New(color.FgYellow).Fprintln(la.LogIO)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgYellow).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgYellow).Fprintf(os.Stdout, format, a...)
		color.New(color.FgYellow).Fprintln(os.Stdout)
	}
}

// Traceln Trace with ln.
func (la *LogAttributes) Traceln(a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogTrace {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgCyan).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgCyan).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgCyan).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgCyan).Fprintln(os.Stdout, a...)
	}
}

// Trace Trace without ln.
func (la *LogAttributes) Trace(a ...interface{}) {
	if la.Flag >= LogTrace {
		color.New(color.FgCyan).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgCyan).Fprintln(os.Stdout, a...)
	}
}

// Tracef Trace without ln.
func (la *LogAttributes) Tracef(format string, a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogTrace {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgCyan).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgCyan).Fprintf(la.LogIO, format, a...)
		color.New(color.FgCyan).Fprintln(la.LogIO)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgCyan).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgCyan).Fprintf(os.Stdout, format, a...)
		color.New(color.FgCyan).Fprintln(os.Stdout)
	}
}

// Debugln Debug with ln.
func (la *LogAttributes) Debugln(a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogDebug {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgMagenta).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgMagenta).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgMagenta).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgMagenta).Fprintln(os.Stdout, a...)
	}
}

// Debug Debug without ln.
func (la *LogAttributes) Debug(a ...interface{}) {
	if la.Flag >= LogDebug {
		color.New(color.FgMagenta).Fprintln(la.LogIO, a...)
		//控制台
		color.New(color.FgMagenta).Fprintln(os.Stdout, a...)
	}
}

// Debugf Debug without ln.
func (la *LogAttributes) Debugf(format string, a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogDebug {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgMagenta).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgMagenta).Fprintf(la.LogIO, format, a...)
		color.New(color.FgMagenta).Fprintln(la.LogIO)
		//控制台
		color.New(color.FgWhite|color.BgBlack).Fprint(os.Stdout, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgMagenta).Fprint(os.Stdout, " "+logLoc+" ")
		color.New(color.FgMagenta).Fprintf(os.Stdout, format, a...)
		color.New(color.FgMagenta).Fprintln(os.Stdout)
	}
}

func (la *LogAttributes) FInfof(format string, a ...interface{}) {
	logLoc := GetLogFileInfo()
	if la.Flag >= LogDebug {
		color.New(color.FgWhite|color.BgBlack).Fprint(la.LogIO, "["+time.Now().Format("2006-01-02 15:04:05.000")+"]")
		color.New(color.FgMagenta).Fprint(la.LogIO, " "+logLoc+" ")
		color.New(color.FgMagenta).Fprintf(la.LogIO, format, a...)
		color.New(color.FgMagenta).Fprintln(la.LogIO)
	}
}
