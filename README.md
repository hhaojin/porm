# go-porm

#### 介绍
基于github.com/jmoiron/sqlx的半自动化orm


#### 安装
```
go get github.com/hhaojin/porm
```

使用示例
```go

package main

import (
	"fmt"
	"log"
	"github.com/hhaojin/porm/porm"
)

type UsersModel struct {
	User_id  string `db:"user_id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
func NewUsersModel() *UsersModel {
	m := &UsersModel{}
	return m
}

func main() {
	user := NewUsersModel()
	porm.Model(user)
	{
		porm.DebugMode = true
		users := make([]UsersModel, 0)
		err := porm.Build(`select user_id,name,password from users where name=?`).
			Args("user5").
			First(user)
		if err != nil {
			log.Println("ERR---", err)
		}
		fmt.Println(user)
		fmt.Println(users)
	}
}

```

