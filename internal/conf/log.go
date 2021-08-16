package conf

import (
	"fmt"
	"ginapi/tool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/fs"
	"os"
)

func NewZap(conf *Conf) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	logPath := conf.Conf.GetString("log.file.path")
	if _, err := tool.PathExists(logPath); err == nil {
		if err = os.MkdirAll(logPath, fs.ModePerm); err != nil {
			panic(err)
		}
	}
	filename := logPath + "/" + tool.Date("Y-m-d") + ".log"
	logType := conf.Conf.GetString("log.file.type")
	if logType != "" {
		logType = "json"
	}
	dev := conf.Conf.GetBool("log.file.prod")
	config := zap.Config{
		Level:            atom,                                                                           // 日志级别
		Development:      dev,                                                                            // 开发模式，堆栈跟踪
		Encoding:         logType,                                                                        // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                                                  // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": conf.Conf.GetString("server.http.name")}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout", filename},                                                   // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}
	// 构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
	defer logger.Sync()
	return logger
}
