package server

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	//"os"
	"strconv"
	
	"github.com/go-playground/validator/v10"
	//"fmt"
	//"github.com/go-playground/validator/v10"
	"gitlab.sftcwl.com/sf-op-public/golang/fabu"
	"tektonctl/model"
	"tektonctl/pkg"
	
	//"tektonctl/pkg"
)


//func StartPreDeploy(parampredeploy model.ParamPreDeploy) model.RespData {
func StartPreDeploy(parampredeploy model.ParamPreDeploy) {
	var (
		err            error
		msg string
		param *fabu.ParamHookStart = &fabu.ParamHookStart{}
		pipelineparam *PipelineParam = &PipelineParam{}
		respbuildimage model.RespBuildImage
	)
	
	reqparam ,_ := json.Marshal(parampredeploy)
	msg = fmt.Sprintf("请求参数:%v",string(reqparam))
	model.PrintInfo(msg)
	
	err = valid.Struct(parampredeploy)
	if err != nil {
		msg = fmt.Sprintf("参数校验失败,失败原因:%v",pkg.Translate(err.(validator.ValidationErrors)))
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	msg = fmt.Sprintf("请求参数检验通过")
	model.PrintInfo(msg)
	
	pid,_ := strconv.ParseInt(parampredeploy.PId, 10, 64)
	pipelineparam.Step = "build-image"
	pipelineparam.PId = pid
	pipelineparam.Appname = parampredeploy.AppName
	pipelineparam.ProductName = parampredeploy.ProductName
	pipelineparam.Env = parampredeploy.ExtEnv
	pa,err := pipelineparam.Get()
	if err != nil {
		msg = fmt.Sprintf("获取继承参数失败:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	msg = fmt.Sprintf("镜像信息:%v",string(pa))
	model.PrintInfo(msg)
	
	err = json.Unmarshal([]byte(string(pa)),&respbuildimage)
	if err != nil {
		msg = fmt.Sprintf("解析镜像信息失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	im,_ := json.Marshal(respbuildimage.ImageInfo)
	msg = fmt.Sprintf("镜像信息解析成功,镜像信息:%v",string(im))
	model.PrintInfo(msg)
	

	/*
	fClient, err := fabu.NewFaBuClient(HTTP_FABU_SERVER_URL)
	if err != nil {
		msg = fmt.Sprintf("初始化发布系统client失败，失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	 */
	param.ImageName = respbuildimage.ImageInfo
	param.DeployEnv = "dev"
	param.AppName = parampredeploy.AppName
	param.ProductName = parampredeploy.ProductName
	param.Uname = parampredeploy.Uname
	param.Emp = parampredeploy.Emp
	
	reqhookstartparam,_ := json.Marshal(param)
	msg = fmt.Sprintf("hook start请求参数:%v",string(reqhookstartparam))
	model.PrintInfo(msg)
	
	err = fClient.HookStart(param)
	if err != nil {
		msg = fmt.Sprintf("预发布环境部署失败，失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	msg = fmt.Sprintf("hook start执行成功")
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("预发布部署成功")
	model.PrintInfo(msg)
	
	os.Exit(0)
	
	//return
}
