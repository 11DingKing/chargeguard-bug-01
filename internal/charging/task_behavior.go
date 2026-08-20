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
	next := s.state
	next.Status = "assigned"
	next.Owner = owner
	if s.FailAudit {
		return ErrDispatchAudit
	}
	next.Audits++
	s.state = next
	return nil
}
func (s *DispatchStore) Snapshot() DispatchState { return s.state }
