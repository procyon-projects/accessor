package accessor

import "sync"

var accessorMutex = sync.RWMutex{}
var accessorMap = make(map[string]interface{}, 0)

func GetAccessor(name string) (interface{}, bool) {
	defer accessorMutex.RUnlock()
	accessorMutex.RLock()
	result, ok := accessorMap[name]
	return result, ok
}

func RegisterAccessor(name string, accessor interface{}) {
	defer accessorMutex.Unlock()
	accessorMutex.Lock()
	accessorMap[name] = accessor
}