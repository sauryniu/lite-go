package bg

import (
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"
	"github.com/ahl5esoft/lite-go/example/admin-permission/service/ginsvc"
)

type QueryApi struct {
	ginsvc.BgApiSession
}

func (m QueryApi) Call() (interface{}, error) {
	var rows []global.Admin
	err := m.GetDbFactory().Db(global.Admin{}).Query().ToArray(&rows)
	return rows, err
}
