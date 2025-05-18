package request

type ApplicationDelete struct {
	AuthToken     string
	ApplicationID string
}

type ApplicationCreate struct {
	AuthToken string

	Name string
	Desc string
}

type ApplicationList struct {
	AuthToken string

	Page int
}
