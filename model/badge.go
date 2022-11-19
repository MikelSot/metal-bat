package model

type Badge string

const (
	BadgeEmployee Badge = "empleado"
)

type Shift string

const (
	LateShift Shift = ""
)

type Position string

type Identification string

const (
	IdentificationDNI Identification = "DNI"
	IdentificationRUC Identification = "RUC"
)
