package models

type RecordingMessage struct {
	StartTime         string `json:"start_time"`
	RecordingURL      string `json:"recording_url"`
	Size              int    `json:"size"`
	RecordingUUID     string `json:"recording_uuid"`
	EndTime           string `json:"end_time"`
	Conversation_uuid string `json:"conversation_uuid"`
	Timestamp         string `json:"timestamp"`
}
