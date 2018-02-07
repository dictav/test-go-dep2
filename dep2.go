package dep2

import "github.com/dictav/test-go-dep1"

type Dep2 struct {
	Name string
	Dep1 dep1
}
