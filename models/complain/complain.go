package models

type Complain struct {
	Id         uint   `json:"id"`
	Items_id   int    `json:"items_id"`
	User_id    int    `json:"user_id"`
	Complain   string `json:"complain"`
	Status     int    `json:"status"` // 1.Review 2.Declined 3.Success
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
	Deleted_at int64  `json:"deleted_at"`
}