package modulerole

import (
	"database/sql"

	"github.com/MikelSot/metal-bat/model/authorization"
)

type ModuleRole struct {
	db *sql.DB
}

func New(db *sql.DB) ModuleRole {
	return ModuleRole{db}
}

func (m ModuleRole) GetByModuleAndRole(name string, roleID uint16) (authorization.ModuleRole, error) {
	//TODO implement me
	panic("implement me")
}

func (m ModuleRole) GetByUserIDAndModule(userID uint, name string) (authorization.ModuleRoles, error) {
	//TODO implement me
	panic("implement me")
}
