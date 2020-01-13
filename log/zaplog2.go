package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

//在定制化zap logger的基础上进行日志切割与归档
// zap 自身不支持切割归档
//定制化我们的zap logger,到此为止，我们的zap logger应该就能在应用中使用了
var CusZLogger1 *zap.Logger
var CusSLogger1 *zap.SugaredLogger

//之前只是返回具体的写入对象,这里加入Lumberjack支持进行切割与归档
func GetLogWriter1() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test1.log", //文件路径
		MaxSize:    1,             //进行切割之前,日志最大大小MB
		MaxBackups: 5,             //旧文件保留个数
		MaxAge:     30,            //旧文件最大保留天数
		Compress:   true,          //是否压缩/归档旧文件
	}

	return zapcore.AddSync(lumberJackLogger)
}


//普通的Encoder而不是json,同时覆盖一些默认配置
func GetNormalEncoder1() zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderCfg)
}


func InitZapLogger2() {
	writeSyncer := GetLogWriter1()
	encoder := GetNormalEncoder1()

	zcore := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel) //第三个参数达到那个级别才会写入
	CusZLogger1 = zap.New(zcore, zap.AddCaller())                      //同时添加调用函数信息
	CusSLogger1 = CusZLogger1.Sugar()

}

func CustomizeZapUse1(url string) {
	resp, err := http.Get(url)
	if err != nil {
		CusSLogger1.Errorf("fetching url:%s error:%s", url, err.Error()) //sugar支持格式化
	} else {
		CusSLogger1.Infof("fetching url:%s success! resp code:%d", url, resp.StatusCode)
		resp.Body.Close()
	}
}

func main() {
	InitZapLogger2()
	for i := 0; i < 1000000; i++ {
		CusSLogger1.Infof("i am %d test msg", i+1)
	}
}
