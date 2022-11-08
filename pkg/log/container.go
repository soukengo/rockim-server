package log

import "sync"

var (
	std      = newHelper(newAdapter(LevelInfo, nil))
	stdProxy = proxy(proxy(std))
)
var global = &loggerContainer{
	root:      std,
	rootProxy: proxy(std),
	loggers:   map[string]Helper{},
	empty:     &emptyHelper{}}

type loggerContainer struct {
	lock      sync.RWMutex
	root      Helper
	rootProxy Helper
	loggers   map[string]Helper
	empty     Helper
}

func (c *loggerContainer) setLogger(root Helper) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.root = root
	c.rootProxy = proxy(root)
}
func (c *loggerContainer) register(name string, logger Helper) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.loggers[name] = proxy(logger)
}

func (c *loggerContainer) use(name string) Helper {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if l, ok := c.loggers[name]; ok {
		return l
	}
	return c.empty
}

func (c *loggerContainer) has(name string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if _, ok := c.loggers[name]; ok {
		return true
	}
	return false
}
