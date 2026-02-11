package auth

import (
	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func SetupCasbin(db *gorm.DB) (*casbin.Enforcer, error) {
	// 1. Gunakan GORM sebagai penyimpan policy di Postgres
 	adapter, err := gormadapter.NewAdapterByDB(db)
 	if err != nil {
 		return nil, err
 	}
	// 2. Load model (RBAC) - Bisa dari file .conf atau string
	// Kita gunakan model RBAC standar: sub (user/role), obj (resource), act (action)
	m, _ := model.NewModelFromString(`
    [request_definition]
    r = sub, obj, act

    [policy_definition]
    p = sub, obj, act

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act
    `)

	return casbin.NewEnforcer(m, adapter)
}
