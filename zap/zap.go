package zap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
    "strings"
    "context"
    "fmt"
)
type Logger struct {
    Log     *zap.Logger
    Opt     *ZapConfig
}
type ZapConfig struct {
    Level string
    Format string
    LogPath string
    LogFile string
    MaxSize int
    MaxBackups int
    MaxAge int
    Compress bool
    CtxKey string
}

func parseLevel(lvl string) zapcore.Level {
    switch strings.ToLower(lvl) {
    case "panic", "dpanic":
        return zapcore.PanicLevel
    case "fatal":
        return zapcore.FatalLevel
    case "error":
        return zapcore.ErrorLevel
    case "warn", "warning":
        return zapcore.WarnLevel
    case "info":
        return zapcore.InfoLevel
    case "debug":
        return zapcore.DebugLevel
    default:
        return zapcore.DebugLevel
    }
}

func NewZapLogger(cnf *ZapConfig) *Logger{
    var logger *Logger = &Logger{}
    logger.Opt = cnf
    writeSyncer := []zapcore.WriteSyncer{NewRollingFile(cnf.LogPath, cnf.LogFile, cnf.MaxSize, cnf.MaxAge,cnf.MaxBackups,cnf.Compress)}
    //writeSyncer := logger.GetLogWriter()
	encoder := logger.GetEncoder()
    dyn := zap.NewAtomicLevel()
    dyn.SetLevel(parseLevel(cnf.Level))
    //core := zapcore.NewCore(encoder, writeSyncer, dyn)
    core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer...), dyn)
	log := zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1))
	defer log.Sync()
    logger.Log = log
    return logger
}

// getEncoder zapcore.Encoder
func (l *Logger) GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = l.CustomTimeEncoder
    encoderConfig.MessageKey  = "msg"
    encoderConfig.LineEnding = zapcore.DefaultLineEnding
    encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
    encoderConfig.TimeKey = "time"
    encoderConfig.FunctionKey = "func"
    encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
    encoderConfig.LineEnding = zapcore.DefaultLineEnding
    encoderConfig.NameKey = "logger"
    return zapcore.NewJSONEncoder(encoderConfig)
}

// CustomTimeEncoder 自定义日志输出时间格式
func (l *Logger) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func (l *Logger) GetTraceField(ctx context.Context) []zap.Field {
    zf := make([]zap.Field,0)
    zf = append(zf,zap.String(l.Opt.CtxKey,ctx.Value(l.Opt.CtxKey).(string)))
    return zf
}

//判断其他类型--start
func (l *Logger) GetOtherFileds(format string, args ...interface{}) (string,[]zap.Field) {
    //判断是否有context
    num := len(args)
    if num > 0 {
        if ctx, ok := args[num-1].(context.Context); ok {
            return fmt.Sprintf(format, args[:num-1]...),l.GetTraceField(ctx)
        } else {
            return fmt.Sprintf(format, args[:num]...),[]zap.Field{}
        }
    }
    return format,[]zap.Field{}
}
