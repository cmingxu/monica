package models

type DbInterface interface {
	DbSave() *DbInterface
	DbUpdate() *DbInterface
	DbDestroy() bool
	Persisted() bool
}
