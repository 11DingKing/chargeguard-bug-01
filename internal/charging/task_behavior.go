package charging

import "errors"

var ErrDispatchAudit = errors.New("dispatch audit failed")

type DispatchState struct {
	Status string
	Owner  string
	Audits int
}
type DispatchStore struct {
	state     DispatchState
	FailAudit bool
}

func NewDispatchStore() *DispatchStore {
	return &DispatchStore{state: DispatchState{Status: "reported"}}
}
func (s *DispatchStore) Dispatch(owner string) error {
	s.state.Status = "assigned"
	s.state.Owner = owner
	if s.FailAudit {
		return ErrDispatchAudit
	}
	s.state.Audits++
	return nil
}
func (s *DispatchStore) Snapshot() DispatchState { return s.state }
