package flowcontrol_test

import (
	"testing"
	"time"

	"github.com/misphix/cryptocurrencyserver/flowcontrol"
)

func TestHasPermission(t *testing.T) {
	f := flowcontrol.New(1, 1)
	if f.AcquirePermission() != true {
		t.Errorf("First time should be true")
	}

	if f.AcquirePermission() != false {
		t.Errorf("Second time should be false")
	}

	time.Sleep(2 * time.Second)

	if f.AcquirePermission() != true {
		t.Errorf("Third time should be true")
	}
}

func TestUnlimitPermission(t *testing.T) {
	f := flowcontrol.New(0, 0)
	for i := 0; i < 1000; i++ {
		if f.AcquirePermission() != true {
			t.Errorf("It should be true")
		}
	}
}
