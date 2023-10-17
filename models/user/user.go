package models

type User struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Password   []byte `json:"-"`
	Email      string `gorm:"unique" json:"email"`
	UData_id   int    `json:"udata_id"`
	Lastlogin  int64  `json:"lastlogin"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
	Deleted_at int64  `json:"deleted_at"`
}

type UserData struct {
	Id          uint    `json:"id"`
	Nickname    string  `json:"nickname"`
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	Sex         string  `json:"sex"`
	Address     string  `json:"address"`
	Dateofbirth int64   `json:"dateofbirth"`
	Nationality string  `json:"nationality"`
	Saldo       int64   `json:"saldo"`
	Wishlist    []int64 `json:"wishlist"`
	Purchased   []int64 `json:"purchased"`
	Created_at  int64   `json:"created_at"`
	Updated_at  int64   `json:"updated_at"`
	Deleted_at  int64   `json:"deleted_at"`
}

// func (user *User) RegisterUser() {

// }

// func (userdata *UserData) RegisterUserData() {

// }