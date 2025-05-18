package profile

type Profile struct {
	Database struct {
		Dsn string
	}
	Api struct {
		Host   string
		Prefix string

		CorsOrigins []string

		MaxUserCount int64

		MaxSignSessionCount        int
		MaxSignSessionWaitingCount int
	}
	IsProd bool
}
