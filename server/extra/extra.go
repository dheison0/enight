package extra

import (
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString returns a random string of the given length, using a set of all
// alphanumeric characters. The result is a string of the given length, with
// characters randomly selected from the set.
func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Dedent remove indentation spaces from a string
func Dedent(text string) string {
	lines := strings.Split(text, "\n")
	indentSize := -1
	for _, line := range lines {
		newLine := strings.TrimSpace(line)
		if newLine == line || newLine == "" {
			continue
		}
		indent := len(line) - len(strings.TrimLeft(line, " "))
		if indentSize == -1 || indent < indentSize {
			indentSize = indent
		}
	}

	if indentSize > 0 {
		for i, line := range lines {
			// TODO: First line o the text is being cut if it's not indented
			if len(line) >= indentSize {
				lines[i] = line[indentSize:]
			}
		}
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}
