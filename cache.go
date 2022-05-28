package cache

import "time"

type Data struct {
	value       string
	hasDeadline bool
	deadline    time.Time
}

type Cache struct {
	data map[string]Data
}

func NewCache() Cache {
	return Cache{
		data: map[string]Data{},
	}
}

func (cache Cache) Get(key string) (string, bool) {
	data, exists := cache.data[key]
	if exists && data.hasDeadline {
		if time.Until(data.deadline) < 0 {
			delete(cache.data, key)
			return "", false
		}
	}
	return data.value, exists
}

func (cache *Cache) Put(key, value string) {
	cache.data[key] = Data{
		value:       value,
		hasDeadline: false,
	}
}

func (cache Cache) Keys() []string {
	keys := make([]string, 0, len(cache.data))
	for k, data := range cache.data {
		if data.hasDeadline {
			if time.Until(data.deadline) < 0 {
				delete(cache.data, k)
				continue
			}
		}
		keys = append(keys, k)
	}
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.data[key] = Data{
		value:       value,
		hasDeadline: false,
		deadline:    deadline,
	}
}
