package handler

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (h *Handler) userFromAuthToken(authToken string) (*model.User, error) {
	// 加载 UserID
	userID, err := h.securityService.UserIDFromAuthToken(authToken)
	if err != nil {
		return nil, err
	}

	return h.userService.FromUserID(userID)
}
