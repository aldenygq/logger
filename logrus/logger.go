package logrus

func (l *Logger) Debugf(format string, args ...interface{}) {
    s, entry := l.GetOtherFileds(format, args...)
    entry.Debug(s)
}
func (l *Logger) Infof(format string, args ...interface{}) {
    s, entry := l.GetOtherFileds(format, args...)
    entry.Info(s)
}
func (l *Logger) Warnf(format string, args ...interface{}) {
    s, entry := l.GetOtherFileds(format, args...)
    entry.Warn(s)
}
func (l *Logger) Errorf(format string, args ...interface{}) {
    s, entry := l.GetOtherFileds(format, args...)
    entry.Error(s)
}
func (l *Logger) Panicf(format string, args ...interface{}) {
    s, entry := l.GetOtherFileds(format, args...)
    entry.Panic(s)
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
    s, entry := l.GetOtherFileds(format, args...)
    entry.Fatal(s)
}
