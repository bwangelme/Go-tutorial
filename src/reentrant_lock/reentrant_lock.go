package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

//ReentrantLock 是一个可重入的锁，并不是一个公平的锁。即不是按照上锁顺序来解锁的
type ReentrantLock struct {
	mu        *sync.Mutex
	cond      *sync.Cond
	owner     int
	holdCount int
}

func NewReentrantLock() sync.Locker {
	rl := &ReentrantLock{}
	rl.mu = new(sync.Mutex)
	rl.cond = sync.NewCond(rl.mu)
	return rl
}

func GetGoroutineId() int {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		log.Fatalf("cannot get goroutine id: %v\n", err)
	}
	return id
}

func (rl *ReentrantLock) Lock() {
	rl.mu.Lock()
	me := GetGoroutineId()
	defer rl.mu.Unlock()

	if rl.owner == me {
		rl.holdCount++
		return
	}
	for rl.holdCount != 0 {
		rl.cond.Wait()
	}
	rl.owner = me
	rl.holdCount = 1
}

func (rl *ReentrantLock) Unlock() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.holdCount == 0 || rl.owner != GetGoroutineId() {
		log.Fatalln("illegalMonitorStateError")
	}
	rl.holdCount--
	if rl.holdCount == 0 {
		rl.cond.Signal()
	}
}

type LockStruct struct {
	Mu   sync.Locker
	name string
	id   int
}

func (s *LockStruct) setName(name string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.name = name
}

func (s *LockStruct) setId(id int) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.id = id
}

func (s LockStruct) PrintName() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.setName("goroutine id : ")
	s.setId(GetGoroutineId())
	fmt.Println(s.name, s.id)
}

func TestWithSingleGoroutine() {
	fmt.Println("reentrant lock single goroutine test start")
	ls := &LockStruct{Mu: NewReentrantLock()}
	ls.PrintName()
	fmt.Println("reentrant lock single goroutine test end")
}

func TestWithMultiGoroutine() {
	fmt.Println("reentrant lock multi goroutine test start")
	ls := &LockStruct{Mu: NewReentrantLock()}
	for i := 0; i < 5; i++ {
		go ls.PrintName()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("reentrant lock multi goroutine test end")
}

func main() {
	TestWithSingleGoroutine()
	TestWithMultiGoroutine()
}
