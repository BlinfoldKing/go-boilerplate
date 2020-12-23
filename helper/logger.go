package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37
)

func getColorByLevel(level logrus.Level) int {
	switch level {
	case logrus.WarnLevel, logrus.DebugLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorBlue
	}
}

type logger struct {
	*logrus.Logger
}

type formatter struct {
	env string
}

// Format custom formatter
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// output buffer
	b := &bytes.Buffer{}
	levelColor := getColorByLevel(entry.Level)

	now := time.Now().Format(time.RFC3339)
	b.WriteString("[")
	_, _ = fmt.Fprintf(b, "\x1b[%dm", levelColor)
	b.WriteString(now)
	_, _ = fmt.Fprintf(b, "\x1b[%dm", colorGray)
	b.WriteString("]")

	_, _ = fmt.Fprintf(b, "\x1b[%dm", colorGray)
	b.WriteString("[")
	_, _ = fmt.Fprintf(b, "\x1b[%dm", levelColor)
	level := strings.ToUpper(entry.Level.String())
	b.WriteString(level)
	_, _ = fmt.Fprintf(b, "\x1b[%dm", colorGray)
	b.WriteString("]")

	if entry.HasCaller() && f.env == "development" {
		b.WriteString("[")
		_, _ = fmt.Fprintf(b, "\x1b[%dm", levelColor)
		fmt.Fprintf(
			b,
			"%s:%d",
			entry.Caller.Function,
			entry.Caller.Line,
		)
		_, _ = fmt.Fprintf(b, "\x1b[%dm", colorGray)
		b.WriteString("]")
	}

	if entry.Message != "" {
		b.WriteString("[")
		_, _ = fmt.Fprintf(b, "\x1b[%dm", levelColor)
		b.WriteString(entry.Message)
		_, _ = fmt.Fprintf(b, "\x1b[%dm", colorGray)
		b.WriteString("]")
	}

	keys := make([]string, 0, len(entry.Data))

	for key := range entry.Data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		json, _ := json.Marshal(entry.Data[key])
		b.WriteString("[")
		_, _ = fmt.Fprintf(b, "\x1b[%dm", levelColor)
		b.WriteString(key)
		_, _ = fmt.Fprintf(b, "\x1b[%dm", colorGray)
		b.WriteString(":")
		b.WriteString(string(json))
		b.WriteString("]")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

// Logger instance of logger
var Logger logger

// InitLogger create logger
func InitLogger(env string) {
	os.Mkdir(".logs/", os.ModePerm)
	l := logrus.New()
	now := time.Now()
	timestamp := now.Format(time.RFC3339)
	filename := fmt.Sprintf("%s_%s.log", env, timestamp)
	_, _ = os.Create(".logs/" + filename)
	file, _ := os.OpenFile(".logs/"+filename, os.O_RDWR, os.ModePerm)

	if GetEnv("ENV", "development") == "development" {
		l.SetLevel(logrus.DebugLevel)
	}

	l.SetFormatter(&formatter{env})

	l.SetOutput(io.MultiWriter(file, os.Stdout))

	l.SetReportCaller(true)
	Logger = logger{l}
}
