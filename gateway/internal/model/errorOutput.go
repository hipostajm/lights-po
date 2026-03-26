package model

type ErrorOutput struct{
	Message string `json:"message"`
}

func NewErrorOutput(message string) ErrorOutput{
	return ErrorOutput{Message: message}
}
