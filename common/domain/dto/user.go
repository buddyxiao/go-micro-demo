package dto

type RegistryRequest struct {
	Username string `gorm:"column:username" json:"username" binding:"required"`
	Password string `gorm:"column:password" json:"password" binding:"required"`
	Email    string `gorm:"column:email" json:"email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
