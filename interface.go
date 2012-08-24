package idb

const(
	_  = iota
	IDB
)

type idbInterface interface {
	Setup( conn string ) (bool, error)
    Get() (int, error)
    Set(key string, value string) error
    getDBtype()
}