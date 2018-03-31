package singleton

import "testing"

func TestSingleTon(t *testing.T) {
	obj1 := Instance()
	obj2 := Instance()
	obj1.name = "obj1"
	obj2.name = "obj2"
	if obj1.name != "obj2" {
		t.Error("instance is not single ", obj1.name, " ", obj2.name)
	}
}
