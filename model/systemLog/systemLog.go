package systemLog

import "fmt"

type SystemLog struct {
	From      string `json:"from"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	CreatedAt int64  `json:"created_at"`
}

func New() *SystemLog {
	return &SystemLog{}
}

func (s *SystemLog) ToInfluxSQL() []byte {
	return []byte(fmt.Sprintf("system_log,from=%s type=%s value=%s %d\n", s.From, s.Type, s.Value, s.CreatedAt))
}
