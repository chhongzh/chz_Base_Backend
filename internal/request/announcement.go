package request

type AnnoucnementAnnounce struct {
	AuthToken     string
	ApplicationID string

	Title   string
	Content string
}

type AnnouncementList struct {
	AuthToken     string
	ApplicationID string
	Page          int
}

type AnnouncementLast struct {
	ApplicationID string
}

type AnnouncementDelete struct {
	AuthToken      string
	AnnouncementID string
}
