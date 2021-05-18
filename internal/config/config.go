package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"

)

type Config interface {
	comfig.Listenerer
	comfig.Logger
}

type config struct {
	comfig.Listenerer
	comfig.Logger

	getter kv.Getter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		getter:     getter,
	}
}

