package model

import "time"

type RemoteConfig interface {
	GetByName(string) (string, error)
	GetInt64(name string) (int64, bool)
	GetInt(name string) (int, bool)
	GetFloat32(name string) (float32, bool)
	GetTime(name string, format string) (time.Time, bool)
	GetBool(name string) (bool, bool)
}
