package service

type Database interface {
	Create(data interface{}) error
	Where(query interface{}, args ...interface{}) Database
	First(out interface{}, where ...interface{}) Database
}
