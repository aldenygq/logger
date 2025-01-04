package zap


func (l *Logger) Debugf(tpl string,args ...interface{}) {
    s,f:= l.GetOtherFileds(tpl, args...)
    l.Log.Debug(s,f...)
}
func (l *Logger) Errorf(tpl string,args ...interface{}) {
    s,f:= l.GetOtherFileds(tpl, args...)
    l.Log.Error(s,f...)
}
func (l *Logger) Infof(format string, args ...interface{}) {
    s,f := l.GetOtherFileds(format, args...)
    l.Log.Info(s,f...)
}
func (l *Logger) Panicf(format string, args ...interface{}) {
    s,f:= l.GetOtherFileds(format, args...)
   l.Log.Panic(s,f...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
    s,f:= l.GetOtherFileds(format, args...)
    l.Log.Fatal(s,f...)
}
