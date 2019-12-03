package models

type NNCO struct {
	Text         string   `json:"text,omitempty"`
	Action       string   `json:"action"`
	EventURL     []string `json:"eventUrl,omitempty"`
	EndOnKey     string   `json:"endOnKey,omitempty"`
	Timeout      string   `json:"timeout,omitempty"`
	EndOnSilence string   `json:"endOnSilence,omitempty"`
}

type RecordingMessage struct {
	StartTime        string `json:"start_time"`
	RecordingURL     string `json:"recording_url"`
	Size             int    `json:"size"`
	RecordingUUID    string `json:"recording_uuid"`
	EndTime          string `json:"end_time"`
	ConversationUUID string `json:"conversation_uuid"`
	Timestamp        string `json:"timestamp"`
}
