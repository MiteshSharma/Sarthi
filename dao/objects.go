package dao

type Task struct {
	Id              string `json:"id" bson:"_id"`
	State           string `json:"state" bson:"state"`
	CallbackUrl     string `json:"callback_url" bson:"callback_url"`
	CallbackMethod  string `json:"callback_method" bson:"callback_method"`
	CallbackPayload string `json:"callback_payload" bson:"callback_payload"`
	Schedule        string `json:"schedule" bson:"schedule"`
	ScheduledAt     int64  `json:"scheduled_at" bson:"scheduled_at"`
}
