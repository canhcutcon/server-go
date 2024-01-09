package models

type User struct {
	BaseModel
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Username string `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Role     string `gorm:"type:varchar(255);not null" json:"role"`
	Phone    string `gorm:"type:varchar(255);not null" json:"phone"`
}

type Role struct {
	BaseModel
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Permission *[]Permission
}

type Permission struct {
	BaseModel
	User   User `gorm:"foreignKey:UserID"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:"user"`
	Role   Role `gorm:"foreignKey:RoleID"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:"role"`
	UserId uint `gorm:"type:integer;not null" json:"user_id"`
	RoleId uint `gorm:"type:integer;not null" json:"role_id"`
}
