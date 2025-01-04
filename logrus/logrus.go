package logrus

import (
    "io"
    "fmt"
    "context"
    "strings"
    "github.com/rifflock/lfshook"
    "github.com/sirupsen/logrus"
)

type Logger struct {
    Log *logrus.Logger
    Opt *LogrusConfig
}
type LogrusConfig struct {
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

func parseLevel(lvl string) logrus.Level {
    switch strings.ToLower(lvl) {
    case "panic", "dpanic":
        return logrus.PanicLevel
    case "fatal":
        return logrus.FatalLevel
    case "error":
        return logrus.ErrorLevel
    case "warn", "warning":
        return logrus.WarnLevel
    case "info":
        return logrus.InfoLevel
    case "debug":
        return logrus.DebugLevel
    default:
        return logrus.DebugLevel
    }
}

func NewLogrusLogger(cnf *LogrusConfig) *Logger {
    var log * Logger = &Logger{}
    log.Opt = cnf
    log.Log = logrus.New()
    log.Log.SetReportCaller(true)
    log.Log.SetLevel(parseLevel(cnf.Level))
    writer := NewRollingFile(cnf.LogPath, cnf.LogFile, cnf.MaxSize, cnf.MaxAge,cnf.MaxBackups,cnf.Compress)
   var format logrus.Formatter
   if cnf.Format == "json" {
        format = &logrus.JSONFormatter{}
   }else {
        format = &logrus.TextFormatter{}
   }
   log.Log.Out = io.Discard
   log.Log.Hooks.Add(lfshook.NewHook(
       lfshook.WriterMap{
           logrus.PanicLevel: writer,
           logrus.FatalLevel: writer,
           logrus.ErrorLevel: writer,
           logrus.WarnLevel:  writer,
           logrus.InfoLevel:  writer,
           logrus.DebugLevel: writer,
       },
       format,
   ))
  return log
}


func (l *Logger) GetTraceField(ctx context.Context) logrus.Fields {
    zf := make(logrus.Fields,0)
    zf[l.Opt.CtxKey] = ctx.Value(l.Opt.CtxKey).(string)
    return zf
}

//判断其他类型--start
func (l *Logger) GetOtherFileds(format string, args ...interface{}) (string, *logrus.Entry) {
    var lf logrus.Fields
    num := len(args)
    if num > 0 {
        if ctx, ok := args[num-1].(context.Context); ok {
            lf := l.GetTraceField(ctx)
            return fmt.Sprintf(format, args[:num-1]...), l.Log.WithFields(lf)
        }
    }
    return format, l.Log.WithFields(lf)
}
