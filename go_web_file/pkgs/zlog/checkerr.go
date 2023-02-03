package zlog

import "go.uber.org/zap"



func Fatalerror( _e error ) {
	if _e != nil {
		Logc.Fatal("严重错误",zap.Error(_e))
	}
  }

func Errorerror( _e error ) {
	if _e != nil {
		Logc.Error("一般错误",zap.Error(_e))
	}
}