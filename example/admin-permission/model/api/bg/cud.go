package bg

import (
	dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"
	"github.com/ahl5esoft/lite-go/example/admin-permission/service/ginsvc"
)

type CUDApi struct {
	ginsvc.BgApiSession

	DbOp dbop.Value
}

func (m CUDApi) Call() (interface{}, error) {
	entry := global.Admin{}
	db := m.GetDbFactory().Db(entry)
	var err error
	switch m.DbOp {
	case dbop.Delete:
		err = db.Delete(entry)
		break
	case dbop.Insert:
		err = db.Insert(entry)
		break
	case dbop.Update:
		err = db.Update(entry)
		break
	}
	return nil, err
}
