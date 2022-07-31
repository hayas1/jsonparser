package presence

import "fmt"

type InvalidRouteError struct {
	dataType string
	keyType  string
}

func (e *InvalidRouteError) Error() string {
	return fmt.Sprint("cannot index ", e.dataType, " with ", e.keyType)
}

type OutOfRangeError struct {
	index  int
	maxLen int
}

func (e *OutOfRangeError) Error() string {
	return fmt.Sprint("out of range index ", e.index, ", must be ", e.index, "<", e.maxLen)
}

type NotFoundKeyError struct {
	key     string
	keyList []string
}

func (e *NotFoundKeyError) Error() string {
	return fmt.Sprint("key \"", e.key, "\" is not found in ", e.keyList)
}
