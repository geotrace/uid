package uid

import (
	"fmt"
	"testing"
)

func TestUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		uid := New()
		info := Parse(uid)
		if info == nil {
			t.Fatal("bad uid", uid)
		}
		fmt.Println(uid, info)
	}
}
