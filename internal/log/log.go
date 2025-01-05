package log

import (
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

// Config 日志配置
type Config struct {
	Output        OutputType `json:"output" yaml:"output"`
	OutputFile    string     `json:"output_file" yaml:"output_file,omitempty"` // omitempty表示如果没有设置该字段则不输出
	LogLevel      string     `json:"log_level" yaml:"log_level"`
	LogFormat     string     `json:"log_format" yaml:"log_format"`
	LogTimeFormat string     `json:"log_time_format" yaml:"log_time_format"`
}

// 服务日志
var srvLog *logrus.Logger

// NewServerLog 初始化日志输出
func NewServerLog(cfg *Config) error {
	// 注入logrus
	srvLog = &logrus.Logger{
		Out:          os.Stderr,
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
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

// `uuid=%s|method=%s|path=%s|status=%v|src_addr=%s|req=>%s|elapsed=%s`
func withCtxFiles(ctx *util.Context) *logrus.Entry {
	// 日志点
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	return srvLog.WithFields(logrus.Fields{
		FieldTraceId: ctx.TraceId(),
		FieldPath:    ctx.Path(),
		FieldFile:    fmt.Sprintf("%s:%d", file, line),
		FieldMethod:  ctx.Request().Method,
		FieldSrcAddr: ctx.RealIP(),
		FieldDstAddr: ctx.Request().Host,
	})
}

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
	FieldLevel, FieldTime, FieldTraceId, FieldSrcAddr, FieldDstAddr, FieldMethod, FieldPath, FieldElapsed,
	FieldFile, FieldReq, FieldRsp,
	FieldMsg, FieldError,
}

// Infof 上下文打印信息
func Infof(format string, v ...any) {
	srvLog.Infof(format, v...)
}

// InfoContextf 上下文打印信息
func InfoContextf(ctx *util.Context, format string, v ...any) {
	withCtxFiles(ctx).Infof(format, v...)
}

// WithFilesInfoContextf 携带其他参数打印
func WithFilesInfoContextf(fields map[string]any, ctx *util.Context, format string, v ...interface{}) {
	logFields := make(map[string]any)
	for k, v := range fields {
		logFields[k] = v
	}
	withCtxFiles(ctx).WithFields(logFields).Infof(format, v...)
}

// Errorf 错误输出
func Errorf(format string, v ...any) {
	srvLog.Errorf(format, v...)
}

// ErrorContextf 错误附带上下文打印信息
func ErrorContextf(ctx *util.Context, format string, v ...any) {
	withCtxFiles(ctx).Errorf(format, v...)
}

// WrapErrorContextf 打印日志并返回wrap的错误码信息
func WrapErrorContextf(ctx *util.Context, err error, format string, v ...any) error {
	e := errors.Wrapf(err, format, v)
	withCtxFiles(ctx).Errorf(format, v...)
	return e
}
