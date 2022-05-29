package ginsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/example/admin-permission/service/dbsvc"

	"github.com/gin-gonic/gin"
)

type BgApiSession struct {
	MongoDbFactory contract.IDbFactory

	dbFactory contract.IDbFactory
	adminID   string
}

func (m *BgApiSession) GetDbFactory() contract.IDbFactory {
	if m.dbFactory == nil && m.adminID != "" {
		m.dbFactory = dbsvc.NewDbFactory(m.MongoDbFactory, m.adminID)
	}

	return m.dbFactory
}

func (m *BgApiSession) InitSession(req interface{}) (err error) {
	if ctx, ok := req.(*gin.Context); ok {
		if v, ok := ctx.Get("admin"); ok {
			m.adminID = v.(string)
		}
		return
	}

	return
}
