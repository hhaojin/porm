package porm

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"text/template"
)

var DebugMode = false

type F map[string]interface{}

type H map[string]interface{}

func DebugSql(sql string) {
	if DebugMode {
		log.Println(sql)
	}
}

func parseSql(name string, sql *string, funs F, data interface{}) error {
	tpl := template.New(name)
	if funs != nil {
		tpl = tpl.Funcs(template.FuncMap(funs))
	}
	sqlTpl, err := tpl.Parse(*sql)
	if err != nil {
		return fmt.Errorf("sqlFormat error:", err)
	}
	buf := new(bytes.Buffer)
	err = sqlTpl.Execute(buf, data)
	if err != nil {
		return fmt.Errorf("sqlParse error:", err)
	}
	*sql = buf.String()
	return nil
}

func Model(obj interface{}) {
	r := reflect.ValueOf(obj)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	bm := NewBaseModel()
	for i := 0; i < r.NumField(); i++ {
		if r.Field(i).Type() == reflect.TypeOf(bm) {
			r.Field(i).Set(reflect.ValueOf(bm))
			bm.Init(obj)
			break
		}
	}
}
