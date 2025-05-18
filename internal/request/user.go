package request

type UserRegister struct {
	PublicName            string
	PersonalizedSignature string
	Username              string
	Password              string
}

type UserLogin struct {
	Username string
	Password string
}
