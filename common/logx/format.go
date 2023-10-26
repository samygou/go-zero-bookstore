package logx

import (
	"fmt"
)

type OutPut struct {
	Level  string      `json:"level"`
	Time   string      `json:"time"`
	Caller interface{} `json:"caller"`
	Msg    string      `json:"msg"`
	Fields Fields      `json:"detail,omitempty"`
}

// JSONFormatter formats logs into parsable json
type Formatter struct{}

// Format renders a single log entry
func (f *Formatter) Format(entry *Entry) ([]byte, error) {
	data := make(Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	var caller interface{}
	if v, exist := data["internal-caller"]; exist {
		caller = v
		delete(data, "internal-caller")
	}

	level := "unknown"
	switch entry.Level {
	case TraceLevel:
		level = "[trace]"
	case DebugLevel:
		level = "[debug]"
	case InfoLevel:
		level = "[info] "
	case WarnLevel:
		level = "[warn] "
	case ErrorLevel:
		level = "[error]"
	case FatalLevel:
		level = "[fatal]"
	case PanicLevel:
		level = "[panic]"
	}

	detail := ""
	for k, v := range data {
		if detail == "" {
			detail += "detail:{ "
		} else {
			detail += ", "
		}
		detail += fmt.Sprintf("%v: %v", k, v)
	}
	if detail != "" {
		detail += " }"
	}

	out := fmt.Sprintf("%v  %v  %v  %v  %v\n", level, entry.Time.Format("2006/01/02 15:04:05.000"), caller, entry.Message, detail)

	return []byte(out), nil
}
