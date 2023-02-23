package cache

import (
	"sync"
)

type Config struct {
	Default    Category
	Categories []*Category
	cateMap    map[string]*Category
	lock       sync.RWMutex
}

func NewConfig(defaultCategory Category, categories ...*Category) *Config {
	return &Config{Default: defaultCategory, Categories: categories, cateMap: make(map[string]*Category)}
}

func (c *Config) Parse() (err error) {
	if c.Default.Expire == 0 {
		c.Default.Expire = defaultExpire
	}
	if c.Default.EmptyExpire == 0 {
		c.Default.EmptyExpire = emptyExpire
	}
	c.cateMap = make(map[string]*Category)
	if c.Categories == nil {
		c.Categories = make([]*Category, 0)
	}
	c.lock.Lock()
	for _, category := range c.Categories {
		c.cateMap[category.Category] = c.Default.CopyFrom(category)
	}
	c.lock.Unlock()
	return
}

func (c *Config) Category(category string) *Category {
	c.lock.RLock()
	cate, ok := c.cateMap[category]
	c.lock.RUnlock()
	if ok {
		return cate
	}
	return c.new(category)
}

func (c *Config) new(category string) *Category {
	c.lock.Lock()
	cate := c.Default.CopyFrom(&Category{Category: category})
	if c.cateMap == nil {
		c.cateMap = make(map[string]*Category)
	}
	c.cateMap[category] = cate
	c.lock.Unlock()
	return cate
}
