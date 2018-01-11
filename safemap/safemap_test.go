package safemap

import (
	"strconv"
	"testing"
)

func Test_New(t *testing.T) {
	sm := New()
	if sm == nil {
		t.Error("except non nil")
	}

	if sm.Size() != 0 {
		t.Error("except zero")
	}
}

func Test_Set(t *testing.T) {
	sm := New()
	sm.Set("k", 1)

	if sm.Size() != 1 {
		t.Error("except 1")
	}
}

func Test_Get(t *testing.T) {
	sm := New()
	v, exist := sm.Get("k")

	if exist != false {
		t.Error("except false")
	}

	if v != nil {
		t.Error("except nil")
	}

	sm.Set("k", 1)
	v, exist = sm.Get("k")

	if exist != true {
		t.Error("except true")
	}

	if v == nil {
		t.Error("except non nil")
	}
}

func Test_safe(t *testing.T) {
	sm := New()
	ch := make(chan struct{}, 3)
	itemsCount := 100

	// write
	go func() {
		for i := 0; i < itemsCount; i++ {
			sm.Set(strconv.Itoa(i), 1)
		}
		ch <- struct{}{}
	}()
	// write
	go func() {
		for i := 100; i < itemsCount*2; i++ {
			sm.Set(strconv.Itoa(i), 1)
		}
		ch <- struct{}{}
	}()
	// read
	go func() {
		for i := 0; i < itemsCount*2; i++ {
			sm.Get(strconv.Itoa(i))
		}
		ch <- struct{}{}
	}()

	<-ch
	<-ch
	<-ch

	if sm.Size() != itemsCount*2 {
		t.Errorf("except %v\n", itemsCount*2)
	}
}
