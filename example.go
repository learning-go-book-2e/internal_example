package internal_example

import "github.com/learning-go-book-2e/internal_example/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}
