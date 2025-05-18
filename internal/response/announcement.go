package response

import "time"

type AnnouncementList struct {
	AnnouncementList []AnnouncementInfo
}

type AnnouncementInfo struct {
	AnnouncementID string

	Title       string
	Content     string
	WhoAnnounce string

	CreatedAt time.Time
}
