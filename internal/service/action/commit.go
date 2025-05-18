package action

func (s *Service) CommitFromActionRequest(req *ActionRequest) {
	s.waitingActions.Enqueue(req)
}

func (s *Service) Commit(fromApplicationID string, message string, meta *MetaChain) {
	s.CommitFromActionRequest(&ActionRequest{
		FromApplicationID: fromApplicationID,
		Message:           message,
		Meta:              meta.Build(),
	})
}
