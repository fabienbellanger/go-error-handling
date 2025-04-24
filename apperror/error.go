package apperror

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type Err struct {
	Next      *Err      `json:"next"`
	Value     error     `json:"value"`
	Msg       string    `json:"msg"`
	Details   any       `json:"details"`
	File      string    `json:"file"`
	Line      int       `json:"line"`
	Timestamp time.Time `json:"timestamp"`
}

func NewErr(value error, msg string, details any, next *Err) Err {
	_, file, line, _ := runtime.Caller(1)

	return Err{
		Value:     value,
		Msg:       msg,
		Details:   details,
		File:      file,
		Line:      line,
		Timestamp: time.Now(),
		Next:      next,
	}
}

// TODO: To better
func (e Err) Error() string {
	return fmt.Sprintf("value=%v, msg=%s, details=%v, file=%s, line=%d, timestamp=%s, next=%s",
		e.Value,
		e.Msg,
		e.Details,
		e.File,
		e.Line,
		e.Timestamp,
		e.Next,
	)
}

func EmptyErr() Err {
	return Err{}
}

func (e Err) IsErr() bool {
	return e.Value != nil
}

func (e Err) JSON() ([]byte, Err) {
	s, err := json.Marshal(e)
	if err != nil {
		return []byte{}, NewErr(err, "Error when converting Err into JSON", nil, nil)
	}
	return s, EmptyErr()
}
