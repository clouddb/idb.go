package idb

import (
    "os"
    "errors"
    //"xattr"
    //"github.com/davecheney/xattr"	
)

const Version = 0.1

const(
	_    = iota
	FILE
	XATTR
)

type Idb struct{
	Path          string
	Store_type    int
}

func existsDir(path string) (bool, error) {
    fileInfo, err := os.Stat(path)
    if err == nil && fileInfo.IsDir() { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func (p *Idb) Setup( conn string ) (bool, error){
	if (len(conn) == 0) && (len(p.Path) == 0) {
		return false, errors.New("Must specify a directory of DB.")
	}
	if r0, _ := existsDir(conn); r0{
		p.Path = conn
		return true, nil
	}
	if r1, _ := existsDir(p.Path); r1{
		return true, nil
	}
	return false, errors.New("Can not setup connection.")
}

func (p *Idb) Set(){

}



