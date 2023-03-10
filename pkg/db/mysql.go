package db

import "helloworld/internal/conf"

type Mysql struct {
	bs *conf.Bootstrap
}

func (m *Mysql) GetBs() *conf.Bootstrap {
	return m.bs
}

func (m *Mysql) SetBs(bs *conf.Bootstrap) {
	m.bs = bs
}

func newDb(bs *conf.Bootstrap) *Mysql {
	return &Mysql{bs}
}

func (m *Mysql) GetMaster() {

}

func (m *Mysql) GetSlave() {

}

func (m *Mysql) GetCleanup() func() {

}
