package logs

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, char := range log {
		if char == '❗' {
			return "recommendation"
		} else if char == '🔍' {
			return "search"
		} else if char == '☀' {
			return "weather"
		}

	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var output string
	for _, char := range log {
		if char == oldRune {
			output += string(newRune)
		} else {
			output += string(char)
		}
	}
	return output
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
