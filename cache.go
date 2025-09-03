package cache

type cache struct {
	value map[string]any
}

func (c cache) Set(key string, value any) {
	c.value[key] = value
}

func (c cache) Get(key string) any {
	return c.value[key]
}

func (c cache) Delete(key string) {
	delete(c.value, key)
}

func New() *cache {
	return &cache{
		value: make(map[string]any),
	}
}
