package models

type Items struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Likes       int    `json:"likes"`
	Itemdata_id int    `json:"itemdata_id"`
	User_id     int    `json:"user_id"`
	Created_at  int64  `json:"created_at"`
	Updated_at  int64  `json:"updated_at"`
	Deleted_at  int64  `json:"deleted_at"`
}

type ItemData struct {
	Id         uint   `json:"id"`
	Type       string `json:"type"`
	Stock      int    `json:"stock"`
	Desc       string `json:"desc"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
	Deleted_at int64  `json:"deleted_at"`
}