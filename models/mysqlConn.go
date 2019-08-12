package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConnInfo struct {
	USRNAME    string
	USRPASSWRD string
	DBHOST     string
	DBNAME     string
}

func (connInfo DBConnInfo) Conn() (*gorm.DB, error) {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", connInfo.USRNAME, connInfo.USRPASSWRD, connInfo.DBHOST, connInfo.DBNAME)
	return gorm.Open("mysql", connectString)
}

func MysqlConnect() (*gorm.DB, error) {
	const (
		db_userName = db_tUserName
		db_password = db_tPassword
		db_host     = db_tHost
		db_name     = db_tName
	)

	return DBConnInfo{USRNAME: db_userName, USRPASSWRD: db_password, DBHOST: db_host, DBNAME: db_name}.Conn()
}
