package response

type SignCreate struct {
	SignSessionID string
}

type SignWait struct {
	AccessToken string
}

type SignInfo struct {
	ApplicationName string
	ApplicationDesc string
}
