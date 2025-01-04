package logrus
import (
    "testing"
    "context"
    "encoding/json"
    "github.com/satori/go.uuid"
)

//go test -v -test.run Test_NewLogrusLogger
func Test_NewLogrusLogger(t *testing.T) {
    var cnf *LogrusConfig = &LogrusConfig{
        Level: "info",
        Format: "json",
        LogPath: "./logs",
        LogFile: "test.log",
        MaxSize: 1,
        MaxBackups: 10,
        MaxAge: 10,
        Compress: true,
        CtxKey: "trace_id",
    }

    l := NewLogrusLogger(cnf)
    ctx := context.WithValue(context.Background(),"trace_id", uuid.NewV4().String())
    conf,_ := json.Marshal(cnf)
    for i := 1; i <= 10000; i++ {
        l.Infof("info:%v",string(conf),ctx)
    }
}
