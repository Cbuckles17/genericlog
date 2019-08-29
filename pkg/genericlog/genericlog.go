package genericlog

// g_gLog GenericLog provides access to all functions. 
var g_gLog GenericLog

// Const for log levels.
const (
    DEBUGLEVEL  = "debug"
    INFOLEVEL   = "info"
    WARNLEVEL   = "warning"
    ERRORLEVEL  = "error"
    FATALLEVEL  = "fatal"
    PANICLEVEL  = "panic"
)

// Conf struct stores the configuration for the GenericLog and allows the user to
// customize the GenericLog to their needs.
type Conf struct {
    LogFormat           string
    LogLevel            string
    LogTimeUTC          bool
    LogToStdout         bool
    LogToFile           bool
    FilePath            string
    UseDefaultFields    bool
    DefaultFields       Fields
}

// Fields map is a map of strings that will be passed to the .WithFields function.
type Fields map[string]interface{}

// GenericLog interface generalizes log types and their functions.
type GenericLog interface {
    WithField(key string, value interface{}) GenericLog
    WithFields(fields Fields) GenericLog
    WithError(err error) GenericLog

    Debugf(format string, args ...interface{})
    Infof(format string, args ...interface{})
    Warnf(format string, args ...interface{})
    Errorf(format string, args ...interface{})
    Fatalf(format string, args ...interface{})
    Panicf(format string, args ...interface{})

    Debug(args ...interface{})
    Info(args ...interface{})
    Warn(args ...interface{})
    Error(args ...interface{})
    Fatal(args ...interface{})
    Panic(args ...interface{})

    Debugln(args ...interface{})
    Infoln(args ...interface{})
    Warnln(args ...interface{})
    Errorln(args ...interface{})
    Fatalln(args ...interface{})
    Panicln(args ...interface{})
}

// Create will make a log instance of a specified type and then set up the 
// g_gLog global.
func Create(conf Conf) (GenericLog, error) {
    gLog, err := createLogrusLogger(conf)
    if err != nil {
        return nil, err
    }

    g_gLog = gLog

    return g_gLog, nil
}

// GenericLog interface functions

func WithField(key string, value interface{}) GenericLog {
    return g_gLog.WithField(key, value)
}

func WithFields(fields Fields) GenericLog {
    return g_gLog.WithFields(fields)
}

func WithError(err error) GenericLog {
    return g_gLog.WithError(err)
}

func Debugf(format string, args ...interface{}) {
    g_gLog.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
    g_gLog.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
    g_gLog.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
    g_gLog.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
    g_gLog.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
    g_gLog.Panicf(format, args...)
}

func Debug(args ...interface{}) {
    g_gLog.Debug(args...)
}

func Info(args ...interface{}) {
    g_gLog.Info(args...)
}

func Warn(args ...interface{}) {
    g_gLog.Warn(args...)
}

func Error(args ...interface{}) {
    g_gLog.Error(args...)
}

func Fatal(args ...interface{}) {
    g_gLog.Fatal(args...)
}

func Panic(args ...interface{}) {
    g_gLog.Panic(args...)
}

func Debugln(args ...interface{}) {
    g_gLog.Debugln(args...)
}

func Infoln(args ...interface{}) {
    g_gLog.Infoln(args...)
}

func Warnln(args ...interface{}) {
    g_gLog.Warnln(args...)
}

func Errorln(args ...interface{}) {
    g_gLog.Errorln(args...)
}

func Fatalln(args ...interface{}) {
    g_gLog.Fatalln(args...)
}

func Panicln(args ...interface{}) {
    g_gLog.Panicln(args...)
}
