package authorization

import (
	"github.com/MikelSot/metal-bat/model"
)

func GenerateToken(m model.User, sessionID, userType uint, IP string, roles []uint) (string, error) {

	return "", nil
}

// para usar los metodos de de otros parametros fijate como usaste el metodoe de invoice en billing para las conciliaciones
// las cuales pueden estar en un relation o un set
