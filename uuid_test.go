package util

import (
	"testing"
)

func TestGenerateUuid(t *testing.T) {
	println(Uuid(0, 1))
	println(Uuid(1110, 1))

}