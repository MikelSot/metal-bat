package sessiontoken

type Storage interface {
	GetByIDAndExpiresDate(id uint) (uint, error)
}

type UseCase interface {
	GetByIDAndExpiresDate(id uint) (uint, error)
}
