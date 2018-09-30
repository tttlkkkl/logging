// Package logging 日志的统一封装
package logging

import (
	"errors"
	"io"
	"log"
	"os"
)

//Log 供外部使用的全局日志变量
var Log *Logging

// Loger 日志接口 如果要对日志进行扩展都应该遵循这个接口
type Loger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// Logging 日志打印统一封装
type Logging struct {
	//Debug 调试日志
	debug *log.Logger
	//Info 重要提示
	info *log.Logger
	//Warning 错误日志
	warning *log.Logger
	//Error 严重的错误日志
	err *log.Logger
	//Fatal 致命的错误日志
	fatal *log.Logger
	Options
}

// LogLevel 日志级别
type LogLevel uint8

const (
	// All 全部日志
	All LogLevel = iota
	// Debug 调试日志
	Debug
	// Info 运行信息
	Info
	// Warning 需要特别注意的信息
	Warning
	// Err 错误日志
	Err
	// Fatal 致命的错误日志
	Fatal
	// Non 不打印任何日志
	Non
)

// Options 日志选项
type Options struct {
	logLevel LogLevel
	output   io.Writer
}

// Option 日志选项设置方法
type Option func(*Options)

// SetLogLevel 设置日志等级
func SetLogLevel(lev LogLevel) Option {
	return func(o *Options) {
		o.logLevel = lev
	}
}

// SetOutPut 设置日志输出路径 这个使用默认的输出到标准输入输出即可，后续用 docker 统一采集
func SetOutPut(out io.Writer) Option {
	return func(o *Options) {
		o.output = out
	}
}

// NewLogger 获取日志打印实例
func NewLogger(opt ...Option) *Logging {
	opts := new(Options)
	for _, o := range opt {
		o(opts)
	}
	if opts.logLevel == 0 {
		opts.logLevel = Debug
	}
	if opts.output == nil {
		opts.output = os.Stdout
	}
	log := Logging{
		debug:   logFormat(opts.output, "DEBUG: "),
		info:    logFormat(opts.output, "INFO: "),
		warning: logFormat(opts.output, "WARNING: "),
		err:     logFormat(opts.output, "ERROR: "),
		fatal:   logFormat(opts.output, "FATAL: "),
		Options: *opts,
	}
	return &log
}
func logFormat(out io.Writer, prex string) *log.Logger {
	return log.New(out, prex, log.Ldate|log.Ltime|log.Lshortfile)
}

// SetLogOutput 设置日志输出
func SetLogOutput(w io.Writer, lv LogLevel) error {
	if Log == nil {
		return errors.New("日志模块未初始化")
	}
	Log.debug.SetOutput(w)
	Log.info.SetOutput(w)
	Log.warning.SetOutput(w)
	Log.err.SetOutput(w)
	Log.fatal.SetOutput(w)
	return nil
}

//Debug 打印调试日志
func (l *Logging) Debug(args ...interface{}) {
	if l.logLevel > Debug || l.logLevel == Non {
		return
	}
	l.debug.Println(args...)
}

//Info 打印提示信息日志
func (l *Logging) Info(args ...interface{}) {
	if l.logLevel > Info || l.logLevel == Non {
		return
	}
	l.info.Println(args...)
}

//Warning 打印错误日志
func (l *Logging) Warning(args ...interface{}) {
	if l.logLevel > Warning || l.logLevel == Non {
		return
	}
	l.warning.Println(args...)
}

//Error 打印严重的错误日志
func (l *Logging) Error(args ...interface{}) {
	if l.logLevel > Err || l.logLevel == Non {
		return
	}
	l.err.Println(args...)
}

//Fatal 打印致命错误日志，并中断程序执行
func (l *Logging) Fatal(args ...interface{}) {
	if l.logLevel > Fatal || l.logLevel == Non {
		return
	}
	l.fatal.Fatalln(args...)
}
