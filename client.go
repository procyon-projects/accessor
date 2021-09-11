package accessor

import "sync"

var accessorClientMutex = sync.RWMutex{}
var accessorClient Client

type Client interface {
	Execute(request Request) Response
}

func EnableClient(client Client)  {
	if client == nil {
		panic("Client cannot be nil")
	}

	defer accessorClientMutex.Unlock()
	accessorClientMutex.Lock()
	accessorClient = client
}

func GetClient() Client {
	defer accessorClientMutex.RUnlock()
	accessorClientMutex.RLock()
	return accessorClient
}