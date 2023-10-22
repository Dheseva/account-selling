package models

type Wishlist struct {
	Id         uint  `json:"id"`
	Items_id   int   `json:"items_id"`
	User_id    int   `json:"user_id"`
	Created_at int64 `json:"created_at"`
	Updated_at int64 `json:"updated_at"`
	Deleted_at int64 `json:"deleted_at"`
}