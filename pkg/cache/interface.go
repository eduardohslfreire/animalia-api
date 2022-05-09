package cache

//ICache is an interface of cache
type ICache interface {
	StartConnection() (*Redis, error)
}

// Cache ...
type Cache struct {
	Address  string
	Password string
}
