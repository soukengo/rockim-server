package socket

import "github.com/zhenjl/cityhash"

func (m *server) Bucket(channelId string) *Bucket {
	idx := cityhash.CityHash32([]byte(channelId), uint32(len(channelId))) % m.opt.BucketSize
	return m.buckets[idx]
}

func (m *server) Channel(channelId string) (Channel, bool) {
	return m.Bucket(channelId).Channel(channelId)
}
func (m *server) Group(groupId string) (Group, bool) {
	return m.Bucket(groupId).Group(groupId)
}
func (m *server) JoinGroup(groupId string, channel Channel) error {
	return m.Bucket(groupId).JoinGroup(groupId, channel)
}
func (m *server) QuitGroup(groupId string, channel Channel) error {
	m.Bucket(groupId).QuitGroup(groupId, channel)
	return nil
}
