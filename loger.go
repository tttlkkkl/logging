package logging

//Debug 打印调试日志
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

//Info 打印提示信息日志
func Info(args ...interface{}) {
	Log.Info(args...)
}

//Warning 打印错误日志
func Warning(args ...interface{}) {
	Log.Warning(args...)
}

//Error 打印严重的错误日志
func Error(args ...interface{}) {
	Log.Error(args...)
}

//Fatal 打印致命错误日志，并中断程序执行
func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}
