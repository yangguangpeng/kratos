package db

import (
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/conf"
)

type Mysql struct {
	bs  *conf.Bootstrap
	log *log.Helper
}

type Option func(*options)

type options struct {
	config *conf.Bootstrap
	log    *log.Helper
}

func WithConfig(config *conf.Bootstrap) Option {
	return func(opts *options) {
		opts.config = config
	}
}

func WithLog(log *log.Helper) Option {
	return func(opts *options) {
		opts.log = log
	}
}

func New(opts ...Option) *Mysql {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	return &Mysql{o.config, o.log}
}

func (m *Mysql) GetMaster() {

}

func (m *Mysql) GetSlave() {

}

func (m *Mysql) GetCleanup() func() {
	return func() {

	}
}
