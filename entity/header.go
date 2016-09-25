package entity

type Header struct {
	StatusCode  uint16  `json:"status_code"`
	ExecuteTime float32 `json:"execute_time"`
}
