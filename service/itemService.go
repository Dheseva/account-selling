package service

type ItemServices interface {
	Create(data interface{}) error
	Where(query interface{}, args ...interface{}) ItemServices
	First(out interface{}, where ...interface{}) ItemServices
}
