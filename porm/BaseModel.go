package porm

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type BaseModel struct {
	parent interface{}
	fields map[string]reflect.Value
	id     string //默认ID字段名
	Table  string //默认表名
}

func NewBaseModel() *BaseModel {
	return &BaseModel{fields: make(map[string]reflect.Value), id: "id"}
}

//设置表名
func (this *BaseModel) SetTableName(table string) {
	this.Table = table
}

//设置ID字段
func (this *BaseModel) SetID(idField string) {
	this.id = idField
}

func (this *BaseModel) parse(objName string) {
	reg := regexp.MustCompile(`(\w+)([A-Z])(\w+)`)
	tbName := reg.ReplaceAllString(objName, "${1}_${3}")
	this.Table = strings.ToLower(strings.Split(tbName, "_")[0])
}

//初始化方法，把父类的属性字段取到，存入map
func (this *BaseModel) Init(obj interface{}) {
	this.parent = obj
	r := reflect.TypeOf(this.parent)
	rv := reflect.ValueOf(this.parent)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		rv = rv.Elem()
	}
	for i := 0; i < r.NumField(); i++ {
		if r.Field(i).Name == "BaseModel" {
			continue
		}
		tag := r.Field(i).Tag.Get("db")
		if tag == "" {
			tag = r.Field(i).Name
		}
		this.fields[tag] = rv.Field(i)
	}
	this.parse(r.Name()) //解析表名等
}

func (this *BaseModel) getFieldValue(name string) interface{} {
	if v, ok := this.fields[name]; ok {
		return v.Interface()
	}
	return nil
}

func (this *BaseModel) SelectByID() error {
	sql := fmt.Sprintf("select * from %s where %s=? limit 1", this.Table, this.id)
	idv := this.getFieldValue(this.id)
	if idv == nil {
		panic("id value error")
	}
	return Build(sql).Args(idv).First(this.parent)
}

func (this *BaseModel) SelectBy(h H) error {
	var buf bytes.Buffer
	sql := fmt.Sprintf("select * from %s  ", this.Table)
	values := make([]interface{}, 0)
	if len(h) > 0 {
		buf.WriteString(" where 1=1 ")
		for k, v := range h {
			buf.WriteString(fmt.Sprintf(" and %s=?", k))
			values = append(values, v)
		}
		sql = fmt.Sprintf("%s %s limit 1", sql, buf.String())
	}
	return Build(sql).Args(values...).First(this.parent)
}

func (this *BaseModel) Select(sql string, funs F, values ...interface{}) error {
	return BuildX(sql, funs, this).Args(values...).First(this.parent)

}
