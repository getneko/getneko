package tool

import "testing"

//返回错误
func TestRefal(t *testing.T) {
	t.Log(Refal(-1, "testerr"))
}

//返回成功
func TestReSucc(t *testing.T) {
	t.Log(ReSucc("test succ"))
}
