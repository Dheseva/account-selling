package models

type Transaction struct {
	Id          uint   `json:"id"`
	Items_id    int    `json:"items_id"`
	Selluser_id int    `json:"selluser_id"`
	Buyuser_id  int    `json:"buyuser_id"`
	Price       int64  `json:"price"`
	Comment     string `json:"comment"`
	Status      int    `json:"status"` // 1.Waiting 2.Review 3.Pending 4.Approved 5.Invalid 6.Completed
	Created_at  int64  `json:"created_at"`
	Updated_at  int64  `json:"updated_at"`
}