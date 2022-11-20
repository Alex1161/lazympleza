package memoize

import (
	"sync"
)

type MemoizedFunction func(args ...string) string

func Memoized(f func(args ...string) string) MemoizedFunction {
	var v string
	var once sync.Once
	return func(args ...string) string {
		once.Do(func() {
			v = f()
			f = nil
		})
		return v
	}
}
