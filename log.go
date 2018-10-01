// Package logging 日志的统一封装,这个只是应急用的简单封装可能存在性能或者其他问题，后期继续改进.
package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 默认的全局日志变量
var logN Loger

//默认的日志全局日志输出
func init() {
	logN = NewLogger(SetCallDepth(3))
}

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
	// LevelAll 全部日志
	LevelAll LogLevel = iota
	// LevelDebug 调试日志
	LevelDebug
	// LevelInfo 运行信息
	LevelInfo
	//LevelWarning 需要特别注意的信息
	LevelWarning
	// LevelErr 错误日志
	LevelErr
	// LevelFatal 致命的错误日志
	LevelFatal
	// LevelNon 不打印任何日志
	LevelNon
)

// Options 日志选项
type Options struct {
	logLevel  LogLevel
	output    io.Writer
	calldepth int
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

// SetCallDepth 设置调用深度，以确保打印的日志输出代码所在文件的正确性
func SetCallDepth(d int) Option {
	return func(o *Options) {
		o.calldepth = d
	}
}

// NewLogger 获取日志打印实例
func NewLogger(opt ...Option) Loger {
	opts := new(Options)
	for _, o := range opt {
		o(opts)
	}
	if opts.logLevel == 0 {
		opts.logLevel = LevelDebug
	}
	if opts.output == nil {
		opts.output = os.Stdout
	}
	if opts.calldepth == 0 {
		opts.calldepth = 2
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

// 设置日志格式
func logFormat(out io.Writer, prex string) *log.Logger {
	return log.New(out, prex, log.Ldate|log.Ltime|log.Lshortfile)
}

//Debug 打印调试日志
func (l *Logging) Debug(args ...interface{}) {
	if l.logLevel > LevelDebug || l.logLevel == LevelNon {
		return
	}
	l.debug.Output(l.calldepth, format(args...))
}

//Info 打印提示信息日志
func (l *Logging) Info(args ...interface{}) {
	if l.logLevel > LevelInfo || l.logLevel == LevelNon {
		return
	}
	l.info.Output(l.calldepth, format(args...))
}

//Warning 打印错误日志
func (l *Logging) Warning(args ...interface{}) {
	if l.logLevel > LevelWarning || l.logLevel == LevelNon {
		return
	}
	l.warning.Output(l.calldepth, format(args...))
}

//Error 打印严重的错误日志
func (l *Logging) Error(args ...interface{}) {
	if l.logLevel > LevelErr || l.logLevel == LevelNon {
		return
	}
	l.err.Output(l.calldepth, format(args...))
}

//Fatal 打印致命错误日志，并中断程序执行
func (l *Logging) Fatal(args ...interface{}) {
	if l.logLevel > LevelFatal || l.logLevel == LevelNon {
		return
	}
	l.fatal.Output(l.calldepth, format(args...))
	os.Exit(1)
}

// 日志内容格式化
func format(v ...interface{}) string {
	return fmt.Sprintln(v...)
}
