package lock

import (
	"context"
	"errors"
	"sync"
	"time"
)

type ExpiredLock struct {
	mutex        sync.Mutex
	processMutex sync.Mutex
	owner        string
	stop         context.CancelFunc
}

func NewExpiredLock() *ExpiredLock {
	return &ExpiredLock{}
}

func (e *ExpiredLock) Lock(expireSeconds int) {
	e.mutex.Lock()

	e.processMutex.Lock()
	defer e.processMutex.Unlock()

	token := GetOwnerId()
	e.owner = token

	if expireSeconds <= 0 {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	e.stop = cancel
	go func() {
		select {
		case <-time.After(time.Duration(expireSeconds) * time.Second):
			e.unlock(token)
		case <-ctx.Done():
		}
	}()

}

func (e *ExpiredLock) Unlock() error {
	token := GetOwnerId()
	return e.unlock(token)

}

func (e *ExpiredLock) unlock(token string) error {
	e.processMutex.Lock()
	defer e.processMutex.Unlock()

	if token != e.owner {
		return errors.New("not your lock")
	}

	// 把锁的owner置空
	e.owner = ""
	if e.stop != nil {
		e.stop()
	}
	e.mutex.Unlock()

	return nil
}
