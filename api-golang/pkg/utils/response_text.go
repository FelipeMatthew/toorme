package utils

func ResponseText(msgType, text string) map[string]string {
	response := map[string]string{
		msgType: text,
	}
	return response
}
func ErrorResponse(text string) map[string]string {
	response := map[string]string{
		"error": text,
	}
	return response
}
