package algcode_lib

import (
	"sync"
	"testing"
)

func TestConcurrentModify(t *testing.T) {
	ss := NewStackSync(100)

	wg := new(sync.WaitGroup)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			ss.Push(i)
			wg.Done()
		}()
	}

	wg.Wait()

	if ss.Size() != 100 {
		t.Error("sync stack size != 100")
	}
}

func TestConcurrentModifyAndRead(t *testing.T) {
	ss := NewStackSync(100)

	wg := new(sync.WaitGroup)
	wg.Add(200)

	go func() {
		for i := 0; i < 100; i++ {
			for ss.Size() == 0 {
			}
			wg.Done()
		}
	}()
	for i := 0; i < 100; i++ {
		go func() {
			ss.Push(i)
			wg.Done()
		}()
	}

	wg.Wait()

	if ss.Size() != 100 {
		t.Error("sync stack size != 100")
	}
}

func TestCocurrentPushAndPop(t *testing.T) {
	ss := NewStackSync(100)

	wg := new(sync.WaitGroup)
	wg.Add(200)

	// push 100 elems
	for i := 0; i < 100; i++ {
		go func() {
			ss.Push(i)
			wg.Done()
		}()
	}

	// pop 100 elems
	for i := 0; i < 100; i++ {
		go func() {
			ss.Pop()
			wg.Done()
		}()
	}

	wg.Wait()
}
