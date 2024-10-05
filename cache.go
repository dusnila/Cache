package cache

type Item struct {
	Value interface{}
}

type Cache struct {
	items map[string]Item
}

func New() *Cache {

	items := make(map[string]Item)

	cache := Cache{
		items: items,
	}

	return &cache
}

func (c *Cache) Set(key string, value interface{}) {

	c.items[key] = Item{
		Value: value,
	}

}

func (c *Cache) Get(key string) interface{} {

	item, found := c.items[key]

	// ключ не найден
	if !found {
		return nil
	}

	return item.Value
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}
