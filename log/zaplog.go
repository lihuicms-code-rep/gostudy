package main

import (
	"go.uber.org/zap"
	"net/http"
)

var (
	zLogger     *zap.Logger        //zap Logger
	sugarLogger *zap.SugaredLogger //SugaredLogger
)

func InitZapLogger() {
	zLogger, _ = zap.NewProduction()
	sugarLogger = zLogger.Sugar()
}

func simpleZapUse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		zLogger.Error("fetching url error....", zap.String("url", url), zap.Error(err))
		sugarLogger.Errorf("fetching url:%s error:%s", url, err.Error())  //sugar支持格式化
	} else {
		zLogger.Info("fetching url success...", zap.String("url", url), zap.String("StatusCode", resp.Status))
		sugarLogger.Infof("fetching url:%s success! resp code:%d", url, resp.StatusCode)
		resp.Body.Close()
	}
}

func main() {
	InitZapLogger()
	defer func() {
		zLogger.Sync()
		sugarLogger.Sync()
	}()

	simpleZapUse("https://www.google.com")
	simpleZapUse("https://www.baidu.com")
}
