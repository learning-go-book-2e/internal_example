package foo

import "github.com/learning-go-book-2e/internal_example/foo/internal"

func UseDoubler(i int) int {
	return internal.Doubler(i)
}
