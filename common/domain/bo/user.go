package bo

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	Msg   string
	Token string
}

type RegistryInput struct {
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
}

type RegistryOutput struct {
	Msg string
}
