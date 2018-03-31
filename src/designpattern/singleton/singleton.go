package singleton

import "sync"

type object struct {
	name string
}

var once sync.Once

var obj *object

func getObj() {
	if obj == nil {
		obj = new(object)
	}
}

func Instance() *object {
	once.Do(getObj)
	return obj
}
