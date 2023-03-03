package socket

import (
	"github.com/google/uuid"
	"sync"
)

type Bucket struct {
	channelSize uint32
	id          string
	chs         map[string]Channel
	cLock       sync.RWMutex
	gLock       sync.RWMutex
	groups      map[string]Group
}

func newBucket(channelSize uint32) *Bucket {
	b := &Bucket{
		id:  uuid.New().String(),
		chs: make(map[string]Channel, channelSize),
	}
	return b
}

func (b *Bucket) PutChannel(ch Channel) {
	b.cLock.Lock()
	b.chs[ch.Id()] = ch
	b.cLock.Unlock()
}
func (b *Bucket) DelChannel(ch Channel) {
	b.cLock.Lock()
	delete(b.chs, ch.Id())
	groups := ch.Groups()
	for _, groupId := range groups {
		b.QuitGroup(groupId, ch)
	}
	b.cLock.Unlock()
}

func (b *Bucket) Channel(channelId string) (ch Channel, ok bool) {
	b.cLock.RLock()
	ch, ok = b.chs[channelId]
	b.cLock.RUnlock()
	return
}

func (b *Bucket) Group(groupId string) (group Group, ok bool) {
	b.gLock.RLock()
	group, ok = b.groups[groupId]
	b.gLock.RUnlock()
	return
}
func (b *Bucket) DelGroup(group Group) {
	b.gLock.Lock()
	delete(b.groups, group.ID())
	b.gLock.Unlock()
	group.Close()
}

func (b *Bucket) JoinGroup(groupId string, ch Channel) (err error) {
	var (
		g  Group
		ok bool
	)
	b.gLock.Lock()
	if g, ok = b.groups[groupId]; !ok {
		g = NewGroup(groupId)
		b.groups[groupId] = g
	}
	ch.AddGroup(groupId)
	b.gLock.Unlock()
	if g != nil {
		err = g.Put(ch)
	}
	return
}

func (b *Bucket) QuitGroup(groupId string, ch Channel) {
	b.gLock.RLock()
	g, ok := b.groups[groupId]
	b.gLock.RUnlock()
	if ok {
		if g.Del(ch) {
			b.DelGroup(g)
		}
	}
	ch.DelGroup(groupId)
	return
}
