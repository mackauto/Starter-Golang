package singleton

import "sync"

type singleton struct {
}

var instance *singleton
var once sync.Once

// once.DO user check lock check pattern
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
