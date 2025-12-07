package filestore

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	vmapi "github.com/reiott/vm-api"
)

const (
	DefaultVMDir = "data/vm"
)

var _ vmapi.VMStore = (*VMStore)(nil)

type VMStore struct {
	dir string
	now func() time.Time
	mu  *sync.RWMutex
}

type Option func(*VMStore)

func WithDir(s string) Option {
	return func(v *VMStore) {
		if len(strings.TrimSpace(s)) == 0 {
			return
		}
		v.dir = s
	}
}

func NewVMStore(opts ...Option) vmapi.VMStore {
	store := &VMStore{
		dir: DefaultVMDir,
		now: time.Now,
		mu:  &sync.RWMutex{},
	}

	for _, opt := range opts {
		opt(store)
	}

	store.init()
	return store
}

func (s *VMStore) init() {
	os.MkdirAll(s.dir, 0755)
}

func (s *VMStore) filePath(id string) string {
	return filepath.Join(s.dir, id+".json")
}

func (s *VMStore) IDs(context.Context) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return []string{}
}

func (s *VMStore) All(ctx context.Context) ([]*vmapi.VM, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	files, err := os.ReadDir(s.dir)
	if err != nil {
		return nil, err
	}

	var vms []*vmapi.VM
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !file.Type().IsRegular() || filepath.Ext(file.Name()) != ".json" {
			continue
		}
		id := file.Name()[:len(file.Name())-5]
		vm, err := s.Get(ctx, id)
		if err != nil {
			continue
		}
		vms = append(vms, vm)
	}
	return vms, nil
}

func (s *VMStore) Get(ctx context.Context, id string) (*vmapi.VM, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	b, err := os.ReadFile(s.filePath(id))
	if os.IsNotExist(err) {
		return nil, vmapi.ErrVMNotFound
	}
	if err != nil {
		return nil, err
	}

	var vm vmapi.VM
	if err := json.Unmarshal(b, &vm); err != nil {
		return nil, err
	}
	return &vm, nil
}

func (s *VMStore) Update(ctx context.Context, vm *vmapi.VM) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	b, err := json.MarshalIndent(vm, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath(vm.ID), b, 0644)
}

func (s *VMStore) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return os.Remove(s.filePath(id))
}

func (s *VMStore) SetStatus(ctx context.Context, id string, status vmapi.Status) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	vm, err := s.Get(ctx, id)
	if err != nil {
		return err
	}
	vm.Status = status
	vm.UpdatedAt = s.now().UTC()
	return s.Update(ctx, vm)
}
