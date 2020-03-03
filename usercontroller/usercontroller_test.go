package usercontroller_test

import (
	"testing"
	"time"

	"github.com/misphix/cryptocurrencyserver/usercontroller"
)

func TestQuerryAcquireWithMax(t *testing.T) {
	ip := "192.168.0.1"
	usercontroller.MaxTime = 2
	if usercontroller.QuerryAcquire(ip) != true {
		t.Errorf("First time should be true")
	}

	if usercontroller.QuerryAcquire(ip) != true {
		t.Errorf("Second time should be true")
	}

	if usercontroller.QuerryAcquire(ip) != false {
		t.Errorf("Third time should be false")
	}
}

func TestQuerryAcquireWithRecords(t *testing.T) {
	ip := "192.168.0.1"
	usercontroller.MaxTime = 2
	usercontroller.AddQueryRecord(ip, time.Now().Add(-25*time.Hour))
	usercontroller.AddQueryRecord(ip, time.Now().Add(-26*time.Hour))
	if usercontroller.QuerryAcquire(ip) != true {
		t.Errorf("First time should be true")
	}

	if usercontroller.QuerryAcquire(ip) != true {
		t.Errorf("Second time should be true")
	}

	if usercontroller.QuerryAcquire(ip) != false {
		t.Errorf("Third time should be false")
	}
}

func TestQuerryAcquireNoMax(t *testing.T) {
	ip := "192.168.0.1"
	for i := 0; i < 1000; i++ {
		if usercontroller.QuerryAcquire(ip) != true {
			t.Errorf("It should be true")
		}
	}
}
