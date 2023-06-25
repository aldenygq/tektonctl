
package server

import (
	"encoding/json"
	"fmt"
	"os"
	//"fmt"
	"strconv"
	
	"github.com/go-playground/validator/v10"
	//"github.com/go-playground/validator/v10"
	"gitlab.sftcwl.com/sf-op-public/golang/fabu"
	"gitlab.sftcwl.com/sf-op-public/golang/tektonapi"
	"tektonctl/model"
	"tektonctl/pkg"
	
	//"tektonctl/pkg"
)

func StartDeployProd(paramdeployprod model.ParamDeployProd) {
	var (
		pipelineparam *PipelineParam = &PipelineParam{}
		err  error
		msg string
	)
	
	reqparam ,_ := json.Marshal(paramdeployprod)
	msg = fmt.Sprintf("请求参数:%v",string(reqparam))
	model.PrintInfo(msg)
	
	err = valid.Struct(paramdeployprod)
	if err != nil {
		msg = fmt.Sprintf("参数校验失败,失败原因:%v",pkg.Translate(err.(validator.ValidationErrors)))
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	msg = fmt.Sprintf("请求参数检验通过")
	model.PrintInfo(msg)
	
	pid,_ := strconv.ParseInt(paramdeployprod.PId, 10, 64)
	pipelineparam.Step = "create-order"
	pipelineparam.PId = pid
	pipelineparam.Appname = paramdeployprod.AppName
	pipelineparam.ProductName = paramdeployprod.ProductName
	pipelineparam.Env = paramdeployprod.ExtEnv
	param,err := pipelineparam.Get()
	if err != nil {
		msg = fmt.Sprintf("获取继承参数失败:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	var deployStep string
	switch paramdeployprod.Step {
	case "mirror":
		deployStep = "deploymirror-deploy"
	case "pause":
		deployStep = "deploypause-deploy"
	case "all":
		deployStep = "deployall-deploy"
	}
	
	var tParam *tektonapi.ParamTaskInfo = &tektonapi.ParamTaskInfo{}
	tParam.PipelinerunId = paramdeployprod.PipelinerunId
	tParam.TaskName = deployStep
	
	/*
	tClient, err := tektonapi.NewTektonClient(HTTP_TEKTONAPI_SERVER_URL)
	if err != nil {
		msg = fmt.Sprintf("初始化tektonapi client失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	*/
	
	
	reqtaskinfoparam,_ := json.Marshal(tParam)
	msg = fmt.Sprintf("获取tekton task信息请求参数:%v",string(reqtaskinfoparam))
	model.PrintInfo(msg)
	
	taskinfo,err := tClient.TaskInfo(tParam)
	if err != nil {
		msg = fmt.Sprintf("获取任务信息失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	msg = fmt.Sprintf("获取任务%v信息成功",tParam.TaskName)
	model.PrintInfo(msg)
	
	resptaskinfo,_ := json.Marshal(taskinfo)
	msg = fmt.Sprintf("tekton task信息:%v",string(resptaskinfo))
	model.PrintInfo(msg)
	
	/*
	fClient, err := fabu.NewFaBuClient(HTTP_FABU_SERVER_URL)
	if err != nil {
		msg = fmt.Sprintf("初始化发布系统client失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	*/
	
	var fParam *fabu.ParamStart = &fabu.ParamStart{}
	menuid ,_ := strconv.ParseInt(param,10,64)
	fParam.MenuId = menuid
	fParam.Step = paramdeployprod.Step
	fParam.IsPipelineId = 0
	
	reqstartparam,_ := json.Marshal(tParam)
	msg = fmt.Sprintf("获取发布start请求参数:%v",string(reqstartparam))
	model.PrintInfo(msg)
	
	err = fClient.Start(fParam,taskinfo.Uname,taskinfo.Emp)
	if err != nil {
		msg = fmt.Sprintf("上线失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	msg = fmt.Sprintf("%v上线成功,上线单ID:%v",paramdeployprod.Step,param)
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("上线完成")
	model.PrintInfo(msg)
	
	//return
	os.Exit(0)
}
