package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	//查看其内存地址
	t.Logf("%p", instance1)

	instance2 := GetInstance()
	t.Logf("%p", instance2)

	t.Logf("instance2==instance1:%v", instance1 == instance2)
}

func TestGetInstance_(t *testing.T) {
	instance1 := GetInstance_()
	//查看其内存地址
	t.Logf("%p", instance1)

	instance2 := GetInstance_()
	t.Logf("%p", instance2)

	t.Logf("instance2==instance1:%v", instance1 == instance2)
}
