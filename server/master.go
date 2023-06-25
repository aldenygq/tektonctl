package server

import (
	"fmt"
	
	"gitlab.sftcwl.com/sf-op-public/golang/fabu"
	"gitlab.sftcwl.com/sf-op-public/golang/release"
	"gitlab.sftcwl.com/sf-op-public/golang/tektonapi"
	"tektonctl/model"
)

var (
	tClient *tektonapi.TektonClient
	fClient *fabu.FaBuClient
	rClient *release.ReleaseClient
	//err error
	
)

func InitMaster() error {
	var msg string
	tClient, err = tektonapi.NewTektonClient(HTTP_TEKTONAPI_SERVER_URL)
	if err != nil {
		msg = fmt.Sprintf("tektonapi client初始化失败,失败原因:%v",err)
		model.PrintError(msg)
		return err
	}
	msg = fmt.Sprintf("tekton client 初始化成功")
	model.PrintInfo(msg)
	
	fClient, err = fabu.NewFaBuClient(HTTP_FABU_SERVER_URL)
	if err != nil {
		msg = fmt.Sprintf("发布系统client初始化失败,失败原因:%v",err)
		model.PrintError(msg)
		return err
	}
	
	msg = fmt.Sprintf("发布 client 初始化成功")
	model.PrintInfo(msg)
	
	rClient, err = release.NewReleaseClient(HTTP_RELEASE_SERVER_URL)
	if err != nil {
		msg = fmt.Sprintf("release client初始化失败,失败原因:%v",err)
		model.PrintError(msg)
		return err
	}
	
	msg = fmt.Sprintf("release client 初始化成功")
	model.PrintInfo(msg)
	
	return nil
}