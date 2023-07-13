package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/host"
	"rockimserver/apis/rockim/service"
	"time"
)

type Watcher struct {
	serviceId string
	manager   *CometManager
}

func newWatcher(manager *CometManager, dis registry.Discovery) (res *Watcher, err error) {
	serviceId := service.AppComet
	w := &Watcher{manager: manager, serviceId: serviceId}
	w.watch(dis)
	return w, nil
}

func (w *Watcher) watch(dis registry.Discovery) {
	err := discovery.Watch(dis, w.serviceId, time.Second*30, w.update)
	if err != nil {
		log.Errorf("discovery.Watch error: %v", err)
		time.Sleep(time.Second)
		w.watch(dis)
		return
	}
}
func (w *Watcher) update(ins []*registry.ServiceInstance) {
	var insMap = make(map[string]*registry.ServiceInstance)
	for _, in := range ins {
		var instanceId = in.ID
		insMap[instanceId] = in
		endpoint, err := host.FindEndpoint(in.Endpoints, host.SchemeGrpc)
		if err != nil {
			log.Errorf("FindEndpoint Failed to parse discovery endpoint: %v", err)
			continue
		}
		if len(endpoint) == 0 {
			continue
		}
		if _, ok := w.manager.cometServers[instanceId]; ok {
			continue
		}
		c, err := newComet(in.ID, endpoint, w.manager.config.Comet)
		if err != nil {
			log.Errorf("watchComet newComet(%+v) error(%v)", in, err)
			continue
		}

		w.manager.cometServers[instanceId] = c
	}
	for instanceId := range w.manager.cometServers {
		if _, ok := insMap[instanceId]; !ok {
			delete(w.manager.cometServers, instanceId)
		}
	}
}
