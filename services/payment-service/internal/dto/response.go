package dto

import "time"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Time    time.Time   `json:"time"`
}

func Ok(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Time:    time.Now(),
	}
}

func Err(message string) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
		Time:    time.Now(),
	}
}
