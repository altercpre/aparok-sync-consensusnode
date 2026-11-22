package config

import (
	"os"

	commonaccount "github.com/altercpre/aparok-sync/accountservice"
	"github.com/altercpre/aparok-sync/app"
	"github.com/altercpre/aparok-sync/app/logger"
	"github.com/altercpre/aparok-sync/metric"
	"github.com/altercpre/aparok-sync/net/rpc"
	"github.com/altercpre/aparok-sync/net/secureservice"
	"github.com/altercpre/aparok-sync/net/transport/quic"
	"github.com/altercpre/aparok-sync/net/transport/yamux"
	"github.com/altercpre/aparok-sync/nodeconf"
	"gopkg.in/yaml.v3"

	"github.com/altercpre/aparok-sync-consensusnode/db"
	"github.com/altercpre/aparok-sync-consensusnode/deletelog"
)

const CName = "config"

func NewFromFile(path string) (c *Config, err error) {
	c = &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return
}

type Config struct {
	Drpc                     rpc.Config             `yaml:"drpc"`
	Account                  commonaccount.Config   `yaml:"account"`
	Network                  nodeconf.Configuration `yaml:"network"`
	NetworkStorePath         string                 `yaml:"networkStorePath"`
	NetworkUpdateIntervalSec int                    `yaml:"networkUpdateIntervalSec"`
	Mongo                    db.Config              `yaml:"mongo"`
	Metric                   metric.Config          `yaml:"metric"`
	Log                      logger.Config          `yaml:"log"`
	Yamux                    yamux.Config           `yaml:"yamux"`
	Quic                     quic.Config            `yaml:"quic"`
	Deletion                 deletelog.Config       `yaml:"deletion"`
	Secure                   secureservice.Config   `yaml:"secure"`
}

func (c *Config) Init(a *app.App) (err error) {
	return
}

func (c Config) Name() (name string) {
	return CName
}

func (c Config) GetMongo() db.Config {
	return c.Mongo
}

func (c Config) GetDrpc() rpc.Config {
	return c.Drpc
}

func (c Config) GetAccount() commonaccount.Config {
	return c.Account
}

func (c Config) GetMetric() metric.Config {
	return c.Metric
}

func (c Config) GetNodeConf() nodeconf.Configuration {
	return c.Network
}

func (c Config) GetNodeConfStorePath() string {
	return c.NetworkStorePath
}

func (c Config) GetNodeConfUpdateInterval() int {
	return c.NetworkUpdateIntervalSec
}

func (c Config) GetYamux() yamux.Config {
	return c.Yamux
}

func (c Config) GetQuic() quic.Config {
	return c.Quic
}

func (c Config) GetDeletion() deletelog.Config {
	return c.Deletion
}

func (c Config) GetDB() db.Config {
	return c.Mongo
}

func (c Config) GetSecureService() secureservice.Config {
	return c.Secure
}
