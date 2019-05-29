package pubsub

import (
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPubSub(t *testing.T) {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	// all 订阅者订阅所有消息
	all := p.Subscribe()
	// golang 订阅者仅订阅包含 golang 的消息
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for msg := range all {
			_, ok := msg.(string)
			assert.True(t, ok)
		}
		wg.Done()
	}()

	go func() {
		for msg := range golang {
			v, ok := msg.(string)
			assert.True(t, ok)
			assert.True(t, strings.Contains(v, "golang"))
		}
		wg.Done()
	}()

	p.Close()
	wg.Wait()
}

func TestRWLock(t *testing.T) {
	var m sync.RWMutex
	var num int64
	var wg sync.WaitGroup

	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(-1))

	goRoutineNum := 3
	wg.Add(goRoutineNum)
	for i := 0; i < goRoutineNum; i++ {
		go func() {
			runtime.GOMAXPROCS(4)
			checkFlag := false
			for i := 0; i < 10000; i++ {
				m.RLock()
				atomic.AddInt64(&num, 10000)
				for i := 0; i < 100; i++ {
				}
				// TODO: 这里不会稳定复现，需要找一个稳定复现的方法
				if num != 10000 && !checkFlag {
					// 这句出现表明同时有 N 个 Goroutine 在修改 num，如果将锁改成了 Lock/Unlock ，就不会出现上述情况了
					checkFlag = true
				}
				atomic.AddInt64(&num, -10000)
				m.RUnlock()
			}
			assert.True(t, checkFlag)
			wg.Done()
		}()
	}

	wg.Wait()
}
