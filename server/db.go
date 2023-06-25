package server

import (
	"fmt"
	"os"
	"time"
	
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"tektonctl/model"
)

var db *sqlx.DB
var err error
func DbInit() {
	var msg string
	
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)
	db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		msg = fmt.Sprintf("数据库初始化失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	db.SetConnMaxLifetime(590 * time.Second)
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(100 * time.Second)
	
	msg = fmt.Sprintf("数据库初始化成功")
	model.PrintInfo(msg)
	
}

type PipelineParam struct {
	Id          int    `db:"id"`
	Appname     string `db:"appname"`
	ProductName string `db:"product_name"`
	PipelinerunId string `db:"pipelinerun_id""`
	PId int64 `db:"p_id"`
	Param string `db:"param"`
	Env string `db:"env"`
	Step string `db:"step"`
}



func (p *PipelineParam) Insert() error {
	var msg string
	DbInit()
	insertSql := "insert into pipelinerun_param (appname, product_name,pipelinerun_id,p_id,param,env,step) values (?,?,?,?,?,?,?)"
	_, err = db.Exec(insertSql, p.Appname,p.ProductName,p.PipelinerunId,p.PId,p.Param,p.Env,p.Step)
	if err != nil {
		msg = fmt.Sprintf("数据库写入失败,失败原因:%v",err)
		model.PrintError(msg)
		return err
	}
	
	return nil
}

func (p *PipelineParam) Get() (string,error) {
	var msg string
	DbInit()
	sql := "select * from pipelinerun_param where appname =? and product_name=? and p_id =? and env=? and step=?"
	err = db.Get(p,sql, p.Appname,p.ProductName,p.PId,p.Env,p.Step)
	if err != nil {
		msg = fmt.Sprintf("数据库查询失败,失败原因:%v",err)
		model.PrintError(msg)
		return "",err
	}
	
	return p.Param,nil
}
func (p *PipelineParam) Exist() (bool,error) {
	var count int
	var msg string
	DbInit()
	sql := "select count(*) from pipelinerun_param where appname =? and product_name=? and p_id =? and env=? and step=?"
	err = db.QueryRow(sql, p.Appname, p.ProductName, p.PId, p.Env, p.Step).Scan(&count)
	if err != nil {
		msg = fmt.Sprintf("数据库查询失败,失败原因:%v",err)
		model.PrintError(msg)
		return true,err
	}
	if count != 0 {
		return true,nil
	}
	return false,nil
}
func (p *PipelineParam) Update() error {
	var msg string
	var err error
	DbInit()
	updateSql := "update pipelinerun_param set param = ? where appname =? and product_name=? and p_id =? and env=? and step=?"
	_, err = db.Exec(updateSql,p.Param,p.Appname, p.ProductName, p.PId, p.Env, p.Step)
	if err != nil {
		msg = fmt.Sprintf("数据库更新失败,失败原因:%v",err)
		model.PrintError(msg)
		return err
	}
	return nil
}
type PipelineRunResult struct {
	Id          int    `db:"id"`
	Appname     string `db:"appname"`
	ProductName string `db:"product_name"`
	PipelinerunId string `db:"pipelinerun_id""`
	PId int64 `db:"p_id"`
	ExecResult string `db:"exec_result"`
	Env string `db:"env"`
	Step string `db:"step"`
}

func (p *PipelineRunResult) Insert() error {
	var msg string
	DbInit()
	insertSql := "insert into pipelinerun_result (appname, product_name,pipelinerun_id,p_id,exec_result,env,step) values (?,?,?,?,?,?,?)"
	_, err = db.Exec(insertSql, p.Appname,p.ProductName,p.PipelinerunId,p.PId,p.ExecResult,p.Env,p.Step)
	if err != nil {
		msg = fmt.Sprintf("数据库写入失败,失败原因:%v",err)
		model.PrintError(msg)
		return err
	}
	
	return nil
}

func (p *PipelineRunResult) Exist() (bool,error) {
	var count int
	var msg string
	DbInit()
	sql := "select count(*) from pipelinerun_result where appname =? and product_name=? and p_id =? and env=? and step=?"
	err = db.QueryRow(sql, p.Appname, p.ProductName, p.PId, p.Env, p.Step).Scan(&count)
	if err != nil {
		msg = fmt.Sprintf("数据库查询失败,失败原因:%v",err)
		model.PrintError(msg)
		return true,err
	}
	if count != 0 {
		return true,nil
	}
	return false,nil
}


func (p *PipelineRunResult) Update() error {
	var msg string
	var err error
	DbInit()
	updateSql := "update pipelinerun_result set exec_result = ? where appname =? and product_name=? and p_id =? and env=? and step=?"
	_, err = db.Exec(updateSql,p.ExecResult,p.Appname, p.ProductName, p.PId, p.Env, p.Step)
	if err != nil {
		msg = fmt.Sprintf("数据库更新失败,失败原因:%v",err)
		model.PrintError(msg)
		return err
	}
	return nil
}