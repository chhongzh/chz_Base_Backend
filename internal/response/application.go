package response

type ApplicationInfo struct {
	ApplicationID string

	Name string
	Desc string
}

type ApplicationCreate struct {
	ApplicationID string
}
