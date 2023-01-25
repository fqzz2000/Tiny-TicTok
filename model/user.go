package model

type UserDB struct {
	UserID   uint64 `gorm:"primaryKey column:user_id"`
	UserName string `gorm:"column:user_name"`
	UserPswd string `gorm:"column:user_pswd"`
}

func (*UserDB) TableName() string {
	return "users"
}

type UserDAO struct{}

func NewUserDAO() UserDAO {
	return UserDAO{}
}

func (UserDAO) QueryUserById(id int64) UserDB {
	var ans UserDB
	DB.Where("user_id = ?", id).Find(&ans)
	return ans
}

func (UserDAO) AddNewUser(newUser *UserDB) {
	DB.Omit("user_id").Create(newUser)
}