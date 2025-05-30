package action

import "github.com/chhongzh/chz_Base_Backend/pkg/shortcuts"

func (s *Service) CommitFromActionRequest(req *ActionRequest) {
	s.waitingActions.Enqueue(req)
}

func (s *Service) Commit(fromApplicationID string, message string, meta *shortcuts.MetaChain) {
	s.CommitFromActionRequest(&ActionRequest{
		FromApplicationID: fromApplicationID,
		Message:           message,
		Meta:              meta.Build(),
	})
}
