package user

import (
	"context"

	"github.com/twelveeee/amis-admin-go/dao/mdb"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `json:"id" gorm:"column:id;type:int(64) UNSIGNED;primary_key;auto_increment:true;not null"`
	UserID   int64  `json:"user_id" gorm:"column:user_id;type:int(64) UNSIGNED;auto_increment:true;not null;index:idx_userid"`
	Username string `json:"username" binding:"required" form:"username" gorm:"column:username;type:varchar(64);not null;unique: uni_username" `
	Password string `json:"password" binding:"required" form:"password" gorm:"column:password;type:varchar(32);not null" `
	Email    string `json:"email" gorm:"column:email;type:varchar(64)"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(64)"`
	Avatar   string `json:"avatar" gorm:"column:avatar;type:varchar(64)"`
	Roles    string `json:"roles" gorm:"column:roles;type:varchar(64)"`
}

type userModel struct {
	db *gorm.DB
}

func NewUserModel(ctx context.Context) *userModel {
	return &userModel{
		db: mdb.GetClient().WithContext(ctx),
	}
}

func (m *userModel) TableName() string {
	return "user"
}

func (m *userModel) CreateTable() error {
	if m.db.Migrator().HasTable(m.TableName()) {
		return nil
	}
	return m.db.Table(m.TableName()).AutoMigrate(&User{})
}

func (m *userModel) Create(user *User) error {
	return m.db.Table(m.TableName()).Create(user).Error
}

func (m *userModel) Update(user *User) error {
	return m.db.Table(m.TableName()).Where("user_id = ?", user.UserID).Updates(user).Error
}

func (m *userModel) GetByUserID(userID int64) (*User, error) {
	user := &User{}
	err := m.db.Table(m.TableName()).Where("user_id = ?", userID).First(user).Error
	return user, err
}

func (m *userModel) GetByUsername(username string) (*User, error) {
	user := &User{}
	err := m.db.Table(m.TableName()).Where("username = ?", username).First(user).Error
	return user, err
}
