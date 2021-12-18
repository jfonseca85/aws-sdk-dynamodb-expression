package app

import "strings"

func ApplyPlaceholderMap(flatmap map[string]string) {

	// Apply placeholder
	for k, v := range flatmap {
		flatmap[k] = ApplyPlaceholderString(v, flatmap)
	}
}

func ApplyPlaceholderString(body string, flatmap map[string]string) string {
	s := len(body)
	var sb1 strings.Builder
	// Iterate over body char by char
	for i := 0; i < s; i++ {
		c := body[i]
		// Placeholder start
		if c == '$' {
			var sb strings.Builder
			// Iterate while the curr char is not a delimiter AND is not the end of the string
			for i2 := i; i2 < s && (body[i2] != '\r' && body[i2] != '\n' && body[i2] != '"' && body[i2] != ',' && body[i2] != ':' && body[i2] != '=' && body[i2] != ' '); i2++ {
				// Write to buffer
				sb.WriteByte(body[i2])
				// Increment the global counter to skip
				i = i2
			}
			// Normalize to use the flatmap lookup approach and also skips the first char (which is a dollar sign $)
			key := strings.Replace(sb.String(), ".", "-", -1)[1:]
			// If the flatmap contains the key, insert the value at the current buffer position
			if value, ok := flatmap[key]; ok {
				sb1.WriteString(value)
			} else {
				// Otherwise just insert the placeholder as-is
				sb1.WriteString(sb.String())
			}
		} else {
			// Keep inserting the chars
			sb1.WriteByte(c)
		}
	}

	// Return the interpolated body
	return sb1.String()
}
