package gomsf

import "github.com/deranged0tter/gomsf/rpc"

type HealthManager struct {
	rpc *rpc.RPC
}

func (hm *HealthManager) Check() error {
	return hm.rpc.Health.Check()
}
