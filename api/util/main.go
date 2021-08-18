package util

import "github.com/Davidmnj91/MyExpenses/util/error"

// Interface Util interface
type Interface interface {
}

// Util provide utilities
type Util struct {
	Error *error.Error
}

// Initialize initialize utilities
func Initialize() *Util {
	err := error.New()
	return &Util{Error: err}
}