package logging

//Debug 打印调试日志
func Debug(msg string, err error, args ...interface{}) {
	logN.Debug(msg, err, args...)
}

//Info 打印提示信息日志
func Info(msg string, err error, args ...interface{}) {
	logN.Info(msg, err, args...)
}

//Warning 打印错误日志
func Warning(msg string, err error, args ...interface{}) {
	logN.Warning(msg, err, args...)
}

//Error 打印严重的错误日志
func Error(msg string, err error, args ...interface{}) {
	logN.Error(msg, err, args...)
}

//Fatal 打印致命错误日志，并中断程序执行
func Fatal(msg string, err error, args ...interface{}) {
	logN.Fatal(msg, err, args...)
}
