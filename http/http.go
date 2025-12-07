package http

import vmapi "github.com/reiott/vm-api"

var _ DataStore = (*Store)(nil)

type DataStore interface {
	VMs() vmapi.VMStore
}

type Store struct {
	VMStore vmapi.VMStore
}

func (s *Store) VMs() vmapi.VMStore {
	return s.VMStore
}
