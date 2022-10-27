package discovery

type Config struct {
	Engine    string
	AppName   string           `yaml:"appName,omitempty"`
	Zookeeper *ZookeeperConfig `yaml:"zookeeper,omitempty"`
	Etcd      *EtcdConfig      `yaml:"etcd,omitempty"`
}

type ZookeeperConfig struct {
	Nodes    []string
	RootPath string
}

type EtcdConfig struct {
	Nodes    []string
	Username string
	Password string
}
