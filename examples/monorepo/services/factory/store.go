package factory

import "go.uber.org/zap"

func (f *sf) Get(key string) (any, bool) {
	f.logger.Info("fetching key", zap.String("key", key))
	return f.store.Load(key)
}

func (f *sf) Set(key string, value any) {
	f.logger.Info("setting store key", zap.String("key", key))
	f.store.Store(key, value)
}
