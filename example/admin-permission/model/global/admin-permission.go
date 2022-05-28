package global

import dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"

// 管理员权限模型
type AdminPermission struct {
	ID         string `bson:"_id" db:"_id"`
	Permission map[string]map[dbop.Value]interface{}
}

func (m AdminPermission) GetID() string {
	return m.ID
}
