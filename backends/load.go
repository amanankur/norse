package backends

import (
	"github.com/amanankur/norse/config"
	"sync"
)

var doOnce *sync.Once

func init() {
	doOnce = &sync.Once{}
}

func Configure() {
	doOnce.Do(doConfigure)
}

func doConfigure() {

	if config.IsRedisConfigured != false {
		configureRedis()
	}
	if config.IsMemcacheConfigured {
		configureMemcache()
	}
	if config.IsMysqlConfigured {
		configureMySql()
	}
}
