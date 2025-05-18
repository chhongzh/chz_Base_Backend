package sign

import (
	"sync"
	"time"

	"github.com/chhongzh/chz_Base_Backend/internal/utils"
	"github.com/google/uuid"
)

type SignSession struct {
	ApplicationID string // 非空
	CreatedAt     time.Time
	Waiter        *utils.Waiter

	ResultAccessToken string
	ResultErr         error

	SignSessionID string

	once sync.Once
}

func NewLoginSession(forApplication string) *SignSession {
	return &SignSession{
		ApplicationID: forApplication,
		CreatedAt:     time.Now(),
		SignSessionID: uuid.NewString(),
		Waiter:        utils.NewWaiter(),
	}
}

func (s *SignSession) Wait() (string, error) {
	s.Waiter.Wait()
	return s.ResultAccessToken, s.ResultErr
}

func (s *SignSession) Emit(accessToken string, err error) {
	s.once.Do(func() {
		s.ResultAccessToken = accessToken
		s.ResultErr = err
		s.Waiter.Broadcast()
	})
}
