package lock

import (
	"fmt"
	"testing"
)

func TestGetGoroutineID(t *testing.T) {
	goid := GetGoroutineID()
	if goid == "" {
		t.Errorf("GetGoroutineID() = %v, want not empty", goid)
	}
	fmt.Println(goid)
}

func TestGetOwnerId(t *testing.T) {
	oid := GetOwnerId()
	if oid == "" {
		t.Errorf("GetOwnerId() = %v, want not empty", oid)
	}
	fmt.Println(oid)
}
