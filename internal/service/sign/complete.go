package sign

func (s *Service) CompleteSignSession(signSessionID string, accessToken string, errInComplete error) error {
	// 检查是否有这个会话并获取
	session, err := s.GetSignSession(signSessionID)
	if err != nil {
		return err
	}

	session.Emit(accessToken, errInComplete)

	// 删除这个会话
	delete(s.loginSessionPool, signSessionID)
	delete(s.loginSessionWaitingCountPool, signSessionID)

	return nil
}
