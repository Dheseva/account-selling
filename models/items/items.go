package models

type Items struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Price      int64  `json:"price"`
	Stock      int    `json:"stock"`
	Desc       string `json:"desc"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
	Deleted_at int64  `json:"deleted_at"`
}