package strs

import "strconv"

// AtoiZero converts a string to an integer. If the conversion fails, it returns 0.
// It uses strconv.Atoi for the conversion.
//
// Parameters:
//
//	s - the string to be converted to an integer
//
// Returns:
//
//	int - the converted integer, or 0 if the conversion fails
func AtoiZero(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// AtoiDefault converts a string to an integer. If the conversion fails, it returns the provided default value.
// Parameters:
//   - s: The string to be converted to an integer.
//   - def: The default integer value to return if the conversion fails.
//
// Returns:
//   - The integer value of the string if the conversion is successful, otherwise the default value.
func AtoiDefault(s string, def int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return i
}
