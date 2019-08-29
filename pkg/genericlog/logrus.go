package genericlog

import (
    "os"
    "io"
    "strings"

    log "github.com/sirupsen/logrus"
)

// logrusLogger struct is just a wrapper for logrus FieldLogger.
type logrusLogger struct {
    logger log.FieldLogger
}

// utcFormatter struct is used for making a custom formatter.
type utcFormatter struct {
    log.Formatter
}

// Format is a standardized function. In this use case it is
// for creating a logrus formatter that puts the log time in UTC.
func (u utcFormatter) Format(entry *log.Entry) ([]byte, error) {
    entry.Time = entry.Time.UTC()
    return u.Formatter.Format(entry)
}

// getFormatter determines the format of the log dependant on the
// passed in arguments. 
func getFormatter(logFormat string, logTimeUTC bool) log.Formatter {
    if strings.Compare(logFormat, "JSON") == 0 {
    // matched JSON
        if logTimeUTC {
        // change to UTC timestamps
            return utcFormatter{&log.JSONFormatter{}}
        } else {
        // normal logrus JSONFormatter
            return &log.JSONFormatter{}
        }
    } else {
    // default to logrus TextFormatter
        return &log.TextFormatter{}
    }
}

// createLogrusLogger creates a logrusLogger. This logrusLogger can be 
// customized by changing values in the passed in conf.
func createLogrusLogger(conf Conf) (GenericLog, error) {
    var ll logrusLogger
    var logOutput io.Writer

    parsedLogLevel, err := log.ParseLevel(conf.LogLevel)
    if err != nil {
    // checking if ParseLevel returned an error
        return nil, err
    }

    if conf.LogToFile {
    // logging to a file
        logFile, err := os.OpenFile(conf.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
        if err != nil {
        // checking if OpenFile returned an error
             return nil, err
        }

        if conf.LogToStdout {
        // logging to Stdout also
            logOutput = io.MultiWriter(os.Stdout, logFile)

        } else {
        // only logging to a file
            logOutput = logFile
        }
    }  else {
    // logging to Stdout by default
        logOutput = os.Stdout
    }

    ll.logger = &log.Logger{
        Out:       logOutput,
        Formatter: getFormatter(conf.LogFormat, conf.LogTimeUTC),
        Hooks:     make(log.LevelHooks),
        Level:     parsedLogLevel,
    }

    if conf.UseDefaultFields == true {
    // add passed in default fields to logging
        ll.logger = ll.logger.WithFields(convertToLogrusFields(conf.DefaultFields))
    }

    return &ll, nil
}

// convertToLogrusFields simply converts a string map/GenericLog Fields to a logrus Fields object.
func convertToLogrusFields(fields Fields) log.Fields {
    logrusFields := log.Fields{}

    //loop through the GenericLog Fields and add them to a logrus Fields object
    for index, val := range fields {
        logrusFields[index] = val
    }

    return logrusFields
}

// GenericLog implemented functions

func (l *logrusLogger) WithField(key string, value interface{}) GenericLog {
    return &logrusLogger{
        logger: l.logger.WithField(key, value),
    }
}

func (l *logrusLogger) WithFields(fields Fields) GenericLog {
    return &logrusLogger{
        logger: l.logger.WithFields(convertToLogrusFields(fields)),
    }
}

func (l *logrusLogger) WithError(err error) GenericLog {
    return &logrusLogger{
        logger: l.logger.WithError(err),
    }
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
    l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
    l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
    l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
    l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
    l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
    l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
    l.logger.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
    l.logger.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
    l.logger.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
    l.logger.Error(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
    l.logger.Fatal(args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
    l.logger.Fatal(args...)
}

func (l *logrusLogger) Debugln(args ...interface{}) {
    l.logger.Debugln(args...)
}

func (l *logrusLogger) Infoln(args ...interface{}) {
    l.logger.Infoln(args...)
}

func (l *logrusLogger) Warnln(args ...interface{}) {
    l.logger.Warnln(args...)
}

func (l *logrusLogger) Errorln(args ...interface{}) {
    l.logger.Errorln(args...)
}

func (l *logrusLogger) Fatalln(args ...interface{}) {
    l.logger.Fatalln(args...)
}

func (l *logrusLogger) Panicln(args ...interface{}) {
    l.logger.Fatalln(args...)
}
