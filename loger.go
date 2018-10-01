package logging

//Debug 打印调试日志
func Debug(args ...interface{}) {
	logN.Debug(args...)
}

//Info 打印提示信息日志
func Info(args ...interface{}) {
	logN.Info(args...)
}

//Warning 打印错误日志
func Warning(args ...interface{}) {
	logN.Warning(args...)
}

//Error 打印严重的错误日志
func Error(args ...interface{}) {
	logN.Error(args...)
}

//Fatal 打印致命错误日志，并中断程序执行
func Fatal(args ...interface{}) {
	logN.Fatal(args...)
}
