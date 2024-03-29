package models

type SpeechInput struct {
	Context  []string `json:"context,omitempty"`
	Language string   `json:"language,omitempty"`
	UUID     []string `json:"uuid"`
}

type NCCO struct {
	Text         string       `json:"text,omitempty"`
	Action       string       `json:"action"`
	EventURL     []string     `json:"eventUrl,omitempty"`
	EndOnKey     string       `json:"endOnKey,omitempty"`
	Timeout      string       `json:"timeout,omitempty"`
	EndOnSilence string       `json:"endOnSilence,omitempty"`
	Speech       *SpeechInput `json:"speech,omitempty"`
	Format       string       `json:"format,omitempty"`
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

type ASRResult struct {
	Confidence string `json:"confidence"`
	Text       string `json:"text"`
}

type ASRResponse struct {
	Speech struct {
		TimeoutReason string      `json:"timeout_reason"`
		Results       []ASRResult `json:"results"`
	}
	DTMF struct {
		// "digits": null,
		// "timed_out": false
	} `json:"dtmf"`
	UUID             string `json:"uuid"`
	ConversationUUID string `json:"conversation_uuid"`
	Timestamp        string `json:"timestamp"`
}
