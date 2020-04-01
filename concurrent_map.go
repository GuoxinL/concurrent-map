/*
   Created by guoxin in 2020/1/10 11:16 上午
*/
package tools

import "sync"

type ConcurrentMap struct {
	data map[string]interface{}
	sync.RWMutex
}

//func (this *ConcurrentMap) put(key string, value interface{}) {
//	this.data[key] = value
//}
//func (this *ConcurrentMap) get(key string) interface{} {
//	value, ok := this.data[key]
//	if ok {
//		return value
//	}
//	return nil
//}

func (this *ConcurrentMap) Put(key string, value interface{}) {
	this.Lock()
	this.data[key] = value
	this.Unlock()
}

func (this *ConcurrentMap) Get(key string) interface{} {
	this.RLock()
	value := this.data[key]
	this.RUnlock()
	return value
}

func (this *ConcurrentMap) Size() int {
	this.RLock()
	defer this.RUnlock()
	return len(this.data)
}

func (this *ConcurrentMap) Merge(source map[string]interface{}) {
	this.Lock()
	defer this.Unlock()
	this.data = CopyMap(source)
}

func CopyMap(source map[string]interface{}) map[string]interface{} {
	maps := make(map[string]interface{})
	for k, v := range source {
		maps[k] = v
	}
	return maps
}
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data:    make(map[string]interface{}),
		RWMutex: sync.RWMutex{},
	}
}
func NewConcurrentMapData(source map[string]interface{}) *ConcurrentMap {
	return &ConcurrentMap{
		data:    CopyMap(source),
		RWMutex: sync.RWMutex{},
	}
}
