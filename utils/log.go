package utils

import (
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"os"
)

// 可以定义任意数量的实例
var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&logrus.TextFormatter{
		// 设置不同日志级别显示不同的颜色
		ForceColors: true,
		// 显示时间戳
		FullTimestamp: true,
		// 时间戳格式
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// fix:解决logrus在windows终端下输出无颜色区别的问题
	Logger.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	// 设置日志级别
	// 当使用Fatal及以上级别打印日志消息时，调用Fatal或Panic函数，在调用位置之后的程序将不会执行
	// 设置什么日志级别，就只会打印该日志级别及以上级别的日志
	Logger.SetLevel(logrus.DebugLevel)
	// 设置打印信息中显示调用打印语句的函数所在的文件路径及函数名和在文件中的哪一行
	Logger.SetReportCaller(true)
}
