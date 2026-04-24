package cron

import (
	"fmt"
	"sync"
)

// Factory 任务构造工厂
type Factory func(meta map[string]interface{}) (Job, error)

var (
	registry = make(map[string]Factory)
	mu       sync.RWMutex
)

// Register 注册任务工厂
func Register(method string, f Factory) {
	mu.Lock()
	defer mu.Unlock()
	registry[method] = f
}

// Get 获取任务工厂
func Get(method string) (Factory, bool) {
	mu.RLock()
	defer mu.RUnlock()
	f, ok := registry[method]
	return f, ok
}

// List 列出已注册的任务方法
func List() []string {
	mu.RLock()
	defer mu.RUnlock()
	methods := make([]string, 0, len(registry))
	for m := range registry {
		methods = append(methods, m)
	}
	return methods
}

// Build 根据方法名和元数据构建 Job
func Build(method string, meta map[string]interface{}) (Job, error) {
	f, ok := Get(method)
	if !ok {
		return nil, fmt.Errorf("cron method %q not registered", method)
	}
	return f(meta)
}
