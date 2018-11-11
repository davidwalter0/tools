package x

import (
	"fmt"
	"sync"
	"testing"
)

type Array []int
type M sync.Mutex

var mutex = M(sync.Mutex{})

// Monitor block scoped mutex block scoped mutex returns unlock
// function suitable for use at the start of a protected (monitored)
// function, function scope of defer

// defer call. Ex:
//     defer M.Monitor()()
func (m *M) Monitor() func() {
	var mtx = (*sync.Mutex)(m)
	mtx.Lock()
	return func() {
		mtx.Unlock()
	}
}

var size = 2

func (a *Array) SafeFunc(elem int, wg *sync.WaitGroup) {
	defer wg.Done()
	// safely access a distinct element in the array
	(*a)[elem] = elem * 2 // some concurrent access safe operation
}

func TestSafeAccess(t *testing.T) {
	var array = make(Array, size)
	var wg = sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		go array.SafeFunc(i, &wg)
	}
	// wait for completion of routines
	wg.Wait()
	fmt.Println(array)
}

func (a *Array) MakeUnsafeSafeFunc(elem int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer mutex.Monitor()()
	// resizing the array is unsafe concurrently
	(*a) = append(*a, elem)
}

func TestMakeUnsafeSafeFunc(t *testing.T) {
	var array = make(Array, 0)
	var wg = sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		go array.MakeUnsafeSafeFunc(i, &wg)
	}
	// wait for completion of routines
	wg.Wait()
	fmt.Println(array)
}

func (a *Array) UnsafeFunc(elem int, wg *sync.WaitGroup) {
	defer wg.Done()
	// resizing the array is unsafe concurrently
	(*a) = append(*a, elem)
}

func TestUnsafeAccess(t *testing.T) {
	var array = make(Array, 0)
	var wg = sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		go array.UnsafeFunc(i, &wg)
	}
	// wait for completion of routines
	wg.Wait()
	fmt.Println(array)
}

func TestAppendSafe(t *testing.T) {
	x := []string{"start"}

	t.Log("log func x", cap(x), len(x))
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		t.Log("log func x", cap(x), len(x))
		y := append(x, "hello", "world")
		t.Log("log func y", cap(y), len(y))
	}()
	go func() {
		defer wg.Done()
		t.Log("log func x", cap(x), len(x))
		z := append(x, "goodbye", "bob")
		t.Log("log func z", cap(z), len(z))
	}()
	wg.Wait()
}

func TestAppendUnsafe(t *testing.T) {
	x := make([]string, 0, 6)

	t.Log("log func x", cap(x), len(x))
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		t.Log("log func x", cap(x), len(x))
		y := append(x, "hello", "world")
		t.Log("log func y", cap(y), len(y))
	}()
	go func() {
		defer wg.Done()
		t.Log("log func x", cap(x), len(x))
		z := append(x, "goodbye", "bob")
		t.Log("log func z", cap(z), len(z))
	}()
	wg.Wait()
}
