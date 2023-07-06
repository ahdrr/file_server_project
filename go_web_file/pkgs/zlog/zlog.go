package zlog

import (
	"filrserver/pkgs/config"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugLog *zap.SugaredLogger
var Log *zap.Logger
var Logc *zap.Logger
var (
	loglevel string
	logfile  string
)

func Init() {
	loglevel = config.ViperConfig.GetString("loglevel")
	logfile = config.ViperConfig.GetString("logfile")
	//创建核心对象
	var coreArr []zapcore.Core
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	//encoderConfig.CallerKey=""
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径

	encoder := zapcore.NewConsoleEncoder(encoderConfig) //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	//配置日志级别
	level := getConfiglevel()
	//info和debug级别,debug级别是最低的
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= level && lev < zap.ErrorLevel
	})
	//error级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoFileWriteSyncer := getInfoFileWriter()
	errorFileWriteSyncer := getErrorFileWriter()
	//info文件writeSyncer
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	//error文件writeSyncer
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	//处理
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	//zap.AddCaller()为显示文件名和行号，可省略
	//log := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(),zap.AddCallerSkip(1))
	Log = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	//获取
	SugLog = Log.Sugar()
	Logc = Log.WithOptions(zap.AddCallerSkip(1))

	//日志
	SugLog.Infof("  **********日志初始化完成 输出级别=[%v]**********", level)

}

// 格式获取当前日志级别
func getConfiglevel() (level zapcore.Level) {
	level, err := zapcore.ParseLevel(loglevel)
	//返回
	if err == nil {
		return level
	}
        level, _ := zapcore.ParseLevel("INFO")
        return level
}

// 自定义时间编码器
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
	//enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}

func getInfoFileWriter() zapcore.WriteSyncer {
	//普通日志输出
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logfile, //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    50,      //文件大小限制,单位MB
		MaxBackups: 10,      //最大保留日志文件数量
		MaxAge:     7,       //日志文件保留天数
		Compress:   false,   //是否压缩处理
	})
	//返回
	return infoFileWriteSyncer
}
func getErrorFileWriter() zapcore.WriteSyncer {
	//错误日志输出
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logfile, //日志文件存放目录
		MaxSize:    100,     //文件大小限制,单位MB
		MaxBackups: 10,      //最大保留日志文件数量
		MaxAge:     30,      //日志文件保留天数
		Compress:   false,   //是否压缩处理
	})
	//返回
	return errorFileWriteSyncer
}
