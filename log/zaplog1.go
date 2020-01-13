package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

//定制化我们的zap logger,到此为止，我们的zap logger应该就能在应用中使用了
var CusZLogger *zap.Logger
var CusSLogger *zap.SugaredLogger

func InitZapLogger1() {
	writeSyncer := GetLogWriter()
	encoder := GetNormalEncoder()

	zcore := zapcore.NewCore(encoder, writeSyncer, zapcore.ErrorLevel)  //第三个参数达到那个级别才会写入
	CusZLogger = zap.New(zcore, zap.AddCaller()) //同时添加调用函数信息
	CusSLogger = CusZLogger.Sugar()

}

//Encoder:编码器,怎样写入日志,这里使用NewJSONEncoder,并使用默认的config
//如果以后要改变编码器为普通的Encoder,在这里改变具体的Encoder即可
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

//普通的Encoder而不是json,同时覆盖一些默认配置
func GetNormalEncoder() zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderCfg)
}

//WriterSyncer:指定将日志写到哪里,使用zapcore.AddSync()将打开的文件句柄传递
func GetLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func CustomizeZapUse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		CusSLogger.Errorf("fetching url:%s error:%s", url, err.Error()) //sugar支持格式化
	} else {
		CusSLogger.Infof("fetching url:%s success! resp code:%d", url, resp.StatusCode)
		resp.Body.Close()
	}
}

func main() {
	InitZapLogger1()
	CustomizeZapUse("https://www.google.com")
	CustomizeZapUse("https://www.baidu.com")
}
