package model

type UserDB struct {
	UserID   int64  `gorm:"primaryKey column:user_id"`
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

func (UserDAO) QueryUserByName(name string) UserDB {
	var ans UserDB
	DB.Where("user_name = ?", name).Find(&ans)
	return ans
}

func (UserDAO) QueryNameExists(name string) bool {
	var count int64
	DB.Model(&UserDB{}).Where("user_name = ?", name).Count(&count)
	return count > 0
}

func (UserDAO) AddNewUser(newUser *UserDB) {
	DB.Omit("user_id").Create(newUser)
}