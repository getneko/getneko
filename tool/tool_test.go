package tool

import (
	"testing"
)

func TestGetRandString(t *testing.T) {
	t.Log(GetRandString(10))
}

func TestGetSha256(t *testing.T) {
	t.Log(GetSha256("getneko"))
}
