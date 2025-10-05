package utils

// Ternary is a generic function that returns one of two values
// based on a boolean condition.
//
// If condition is true, it returns ifTrue, otherwise it returns ifFalse.
//
// The function is useful for writing concise and readable code
// that needs to return different values based on a condition.
func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
