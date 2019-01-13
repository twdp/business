package common

import (
	"fmt"
	"testing"
)

func TestPaging_Num(t *testing.T) {
	paging := NewPaging(999, 1, 20, nil)
	fmt.Println(paging.Num())
	paging.PageNo = 2
	fmt.Println(paging.Num())
	paging.PageNo = 3
	fmt.Println(paging.Num())
	paging.PageNo = 50
	fmt.Println(paging.Num())
}
