package mysqldb

import (
	"reflect"
	"strconv"
	"strings"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"github.com/jmoiron/sqlx"
)

// FactoryOption is 工厂选项
type FactoryOption struct {
	DbName   string
	Host     string
	Password string
	Port     int
	Username string
}

type factory struct {
	DB *sqlx.DB
}

func (m factory) Db(entry identity.IIdentity, extra ...interface{}) db.IRepository {
	var uow *unitOfWork
	isTx := true
	underscore.Chain(extra).Each(func(r interface{}, _ int) {
		if v, ok := r.(*unitOfWork); ok {
			uow = v
		}
	})

	if uow == nil {
		isTx = false
		uow = m.Uow().(*unitOfWork)
	}

	return repository{
		DB:        m.DB,
		IsTx:      isTx,
		ModelType: reflect.TypeOf(entry),
		Uow:       uow,
	}
}

func (m factory) Uow() db.IUnitOfWork {
	return &unitOfWork{
		DB:    m.DB,
		Items: make([]unitOfWorkItem, 0),
	}
}

// New is 创建db.IDbFactory
func New(opt FactoryOption) (f db.IFactory, err error) {
	bf := make([]string, 0)
	bf = append(bf, opt.Username, ":", opt.Password, "@tcp(", opt.Host, ":")
	if opt.Port == 0 {
		bf = append(bf, "3306")
	} else {
		bf = append(
			bf,
			strconv.Itoa(opt.Port),
		)
	}
	bf = append(bf, ")/", opt.DbName, "?charset=utf8")
	dataSourceName := strings.Join(bf, "")
	var db *sqlx.DB
	if db, err = sqlx.Open("mysql", dataSourceName); err != nil {
		return
	}

	f = &factory{
		DB: db,
	}
	return
}
