package profile

type Profile struct {
	Database struct {
		Dsn         string
		UsingSqlite bool
	}
	Api struct {
		Host   string
		Prefix string

		CorsOrigins []string

		MaxUserCount int64

		MaxSignSessionCount        int
		MaxSignSessionWaitingCount int
	}
	Sdk struct {
		Host string
	}
	IsProd bool
}
