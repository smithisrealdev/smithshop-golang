package smitlog

type LogMessageBean struct {
	Timestamp       string           `json:"@timestamp"`
	ApplicationName string           `json:"@suffix,omitempty"`
	LogType         LogType          `json:"logType,omitempty"`
	CorrelationID   string           `json:"correlationID,omitempty"`
	Level           LogLevel         `json:"level,omitempty"`
	Message         string           `json:"message,omitempty"`
	MessageType     string           `json:"messageType,omitempty"`
	StackTrace      interface{}      `json:"stackTrace,omitempty"`
	CategoryNames   LogCategoryNames `json:"CategoryNames,omitempty"`
	MonitorType     LogMonitorType   `json:"monitorType,omitempty"`
	SourceSystem    LogSystem        `json:"sourceSystem,omitempty"`
	SourceHostName  string           `json:"sourceHostName,omitempty"`
	TargetSystem    LogSystem        `json:"targetSystem,omitempty"`
	TargetURL       string           `json:"targetURL,omitempty"`
	Action          string           `json:"action,omitempty"`
	ElapsedTime     int64            `json:"elapsedTime,omitempty"`
	ResponseCode    string           `json:"responseCode,omitempty"`
}
