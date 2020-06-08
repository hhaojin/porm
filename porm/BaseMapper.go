package porm

import (
	"github.com/jmoiron/sqlx"
)

type Values []interface{}

type Builder struct { //sql 相关执行的封装
	stmt *sqlx.Stmt
	err  error
	args Values
}

//普通sql查询
func Build(sql string) (builder *Builder) {
	builder = &Builder{}
	stmt, err := MySql().Preparex(sql)
	if err != nil {
		builder.err = err
		return
	}
	DebugSql(sql)
	builder.stmt = stmt
	return
}

//支持模板sql查询
func BuildX(sql string, funs F, data interface{}) (builder *Builder) {
	err := parseSql("buildx", &sql, funs, data)
	if err != nil {
		builder.err = err
		return
	}
	builder = Build(sql)
	return
}

//参数值
func (this *Builder) Args(values ...interface{}) *Builder {
	this.args = values
	return this
}

//查询一条
func (this *Builder) First(m interface{}) error {
	row := this.stmt.QueryRowx(this.args...)
	return row.StructScan(m)
}

//查询全部
func (this *Builder) Find(m interface{}) error {

	rows, err := this.stmt.Queryx(this.args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	err = sqlx.StructScan(rows, m)

	return err

}
