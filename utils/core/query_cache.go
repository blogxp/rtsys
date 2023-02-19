package core

// ICache common/cache里面约定实现的接口
type ICache interface {
	FlushAll() //清空缓存
	BuildAll() //创建缓存
}

func NewICache(cache ICache) ICache{
	return cache
}
