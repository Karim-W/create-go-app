package factory

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"go.uber.org/zap"
)

func (s *sf) DLock(key string) (*redsync.Mutex, error) {
	s.logger.Info("acquiring lock", zap.String("key", key))
	mtx := deps.RSync.NewMutex(key)

	err := mtx.Lock()
	if err != nil {
		s.logger.Error("failed to acquire lock", zap.Error(err))
		return nil, err
	}

	s.logger.Info("lock acquired", zap.String("key", key))
	return mtx, nil
}

func (s *sf) RDB() *redis.Client {
	return deps.Redis
}
