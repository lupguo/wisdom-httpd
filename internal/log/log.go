package log

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/lupguo/wisdom-httpd/internal/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// OutputType 日志输出类型
type OutputType string

const (
	OutputTypeConsole OutputType = `console`
)

// 日志字段定义
const (
	FieldTime    = `time`     // 时间
	FieldPath    = `path`     // 路径
	FieldMethod  = `method`   // 方法
	FieldSrcAddr = `src_addr` // 来源地址
	FieldDstAddr = `dst_addr` // 目标地址
	FieldFile    = `file`     // 日志文件
	FieldTraceId = `trace_id` // TraceId
	FieldLevel   = `level`    // 日志等级
	FieldElapsed = `elapsed`  // 请求耗时
	FieldMsg     = `msg`      // 日志消息
	FieldError   = `err`      // 错误信息
	FieldReq     = `req`      // 请求参数
	FieldRsp     = `rsp`      // 响应参数
)

// FieldSort 排序字段
var FieldSort = []string{
	FieldTime, FieldElapsed, FieldLevel,
	FieldTraceId, FieldSrcAddr, FieldDstAddr, FieldMethod, FieldPath,
	FieldFile, FieldReq, FieldRsp,
	FieldMsg, FieldError,
}

// Config 日志配置
type Config struct {
	Output        OutputType `json:"output" yaml:"output"`
	OutputFile    string     `json:"output_file" yaml:"output_file,omitempty"` // omitempty表示如果没有设置该字段则不输出
	LogLevel      string     `json:"log_level" yaml:"log_level"`
	LogFormat     string     `json:"log_format" yaml:"log_format"`
	LogTimeFormat string     `json:"log_time_format" yaml:"log_time_format"`
}

type IServerLogger interface {
	// WithContext WithCtxFile日志携带上ctx内关键信息
	WithContext(ctx context.Context) *logrus.Entry
}

// 服务日志
var srvLog *SrvLogger

type SrvLogger struct {
	*logrus.Logger
}

// InitServerLog 初始化日志输出
func InitServerLog(cfg *Config) error {
	// 注入logrus
	srvLog = &SrvLogger{
		Logger: &logrus.Logger{
			Out:          os.Stderr,
			Hooks:        make(logrus.LevelHooks),
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		},
	}

	// 设置log等级
	level, err := logrus.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		return errors.Wrap(err, "parse logrus level")
	}
	srvLog.SetLevel(level)
	srvLog.SetFormatter(NewCustomTextFormatter(FieldSort, cfg.LogTimeFormat))
	// std.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// 设置log输出位置
	output := os.Stdout
	if cfg.Output != OutputTypeConsole {
		output, err = os.OpenFile(cfg.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
	}
	srvLog.SetOutput(output)

	return nil
}

// WithContext `uuid=%s|method=%s|path=%s|status=%v|src_addr=%s|req=>%s|elapsed=%s`
func (l *SrvLogger) WithContext(ctx context.Context) *logrus.Entry {
	utilCtx := ctx.(*util.Context)

	// 日志点
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	return l.WithFields(logrus.Fields{
		FieldTraceId: utilCtx.TraceId(),
		FieldPath:    utilCtx.Path(),
		FieldFile:    fmt.Sprintf("%s:%d", file, line),
		FieldMethod:  utilCtx.Request().Method,
		FieldSrcAddr: utilCtx.RealIP(),
		FieldDstAddr: utilCtx.Request().Host,
	})
}

func WithContext(ctx context.Context) *logrus.Entry {
	return srvLog.WithContext(ctx)
}

// Infof 上下文打印信息
func Infof(format string, v ...any) {
	srvLog.Infof(format, v...)
}

// InfoContextf 上下文打印信息
func InfoContextf(ctx context.Context, format string, v ...any) {
	srvLog.WithContext(ctx).Infof(format, v...)
}

// Errorf 错误输出
func Errorf(format string, v ...any) {
	srvLog.Errorf(format, v...)
}

// ErrorContextf 错误附带上下文打印信息
func ErrorContextf(ctx context.Context, format string, v ...any) {
	srvLog.WithContext(ctx).Errorf(format, v...)
}

// Fatalf 严重错误，直接退出
func Fatalf(format string, v ...any) {
	srvLog.Fatalf(format, v...)
}

// WrapAndLogErrorf 打印日志并返回wrap的错误码信息
func WrapAndLogErrorf(ctx context.Context, err error, format string, v ...any) error {
	// warp err
	e := errors.Wrapf(err, format, v...)

	// log err
	srvLog.WithContext(ctx).Error(e)

	return e
}
