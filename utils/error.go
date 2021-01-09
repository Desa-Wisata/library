package utils

// ErrorString func to convert error to string
func ErrorString(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
