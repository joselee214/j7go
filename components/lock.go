package components

import (
	"github.com/joselee214/j7f/components/lock"
	"sync"
)

type locks struct {
	redisLocks *sync.Map
}

var Locks *locks

func init() {
	var once sync.Once
	once.Do(func() {
		Locks = &locks{
			redisLocks: new(sync.Map),
		}
	})
}

//获取一个新的redis锁
func NewRedisLock(c *lock.RedisLockConfig) *lock.RedisLock {
	if len(c.Pools) == 0 {
		c.Pools = append(c.Pools, R)
	}
	l, ok := Locks.redisLocks.Load(c.Name)
	if !ok {
		rl := lock.NewRedsync(c)

		Locks.redisLocks.Store(c.Name, rl)
		return rl
	}

	return l.(*lock.RedisLock)
}

func GetRedisLock(name string) *lock.RedisLock {
	l, ok := Locks.redisLocks.Load(name)
	if !ok {
		c := &lock.RedisLockConfig{
			Name: name,
		}
		return NewRedisLock(c)
	}

	l, _ = Locks.redisLocks.Load(name)
	return l.(*lock.RedisLock)
}
