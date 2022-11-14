package lazy

import (
	"sync"
)

type LazyFunction func() string

func Lazy(f func() string) LazyFunction {
	var v string
	var once sync.Once
	return func() string {
		once.Do(func() {
			v = f()
			f = nil
		})
		return v
	}
}
