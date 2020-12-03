package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// 顶层
func main()  {
	res , err := GetInfo("111")
	if err != nil {
		fmt.Printf("error info %v", err)
	}
	fmt.Println(res)
}

//biz
func GetInfo(p string) (*INFO, error){
	Info, err := getInfo(p)
	if err != nil {
		return nil, err
	}
	return  Info, err
}

//dao
type INFO struct {
	name string
	id   string
}

var reply INFO
var Db *sql.DB

func getInfo(p string) (*INFO, error) {
	res := INFO{}
	rows, err := Db.Query("select * from data where Info = ?", p)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("GetInfo error:" + sql.ErrNoRows.Error())
		}
		return nil, errors.Wrap(err,"GetInfo error ...")
	}

	if err = rows.Scan(&res); err != nil {
		fmt.Println(err) // Handle scan error
		return nil, errors.Wrap(err,"GetInfo error ...")
	}

	return &reply, nil
}