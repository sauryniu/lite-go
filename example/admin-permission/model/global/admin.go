package global

// 管理员
type Admin struct {
	ID   string `alias:"" bson:"_id" db:"_id"`
	Name string
}

func (m Admin) GetID() string {
	return m.ID
}
