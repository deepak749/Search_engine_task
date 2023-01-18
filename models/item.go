package models

type Pages struct {
	ID  int    `json:"id" bson:"Page_No"`
	Key string `json:"key" bson:"Contents"`
}
