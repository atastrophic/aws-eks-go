package exception

import "fmt"

func WithMessageAndError(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s -- %w", msg, err))
	}
}

func WithError(err error) {
	if err != nil {
		panic(err)
	}
}
