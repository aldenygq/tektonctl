package server

import (
	//"encoding/base64"
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

func StartCreateOrder(paramcreateorder model.ParamCreateOrder) {
	var (
		err  error
		msg string
		respcreateorder *model.RespCreateOrder = &model.RespCreateOrder{}
		imageinfos []*fabu.ModeInfo = make([]*fabu.ModeInfo, 0)
		newimages []*fabu.ModeInfo = make([]*fabu.ModeInfo, 0)
		imagemap map[string]string = make(map[string]string,0)
		pipelineparam *PipelineParam = &PipelineParam{}
		result bool
	)
	
	reqparam ,_ := json.Marshal(paramcreateorder)
	msg = fmt.Sprintf("请求参数:%v",string(reqparam))
	model.PrintInfo(msg)
	
	err = valid.Struct(paramcreateorder)
	if err != nil {
		msg = fmt.Sprintf("参数校验失败,失败原因:%v",pkg.Translate(err.(validator.ValidationErrors)))
		model.PrintError(msg)
		os.Exit(-1)
	}
	msg = fmt.Sprintf("请求参数检验通过")
	model.PrintInfo(msg)
	
	pid,_ := strconv.ParseInt(paramcreateorder.PId, 10, 64)
	pipelineparam.Step = "prerelease"
	pipelineparam.PId = pid
	pipelineparam.Appname = paramcreateorder.AppName
	pipelineparam.ProductName = paramcreateorder.ProductName
	pipelineparam.Env = paramcreateorder.ExtEnv
	param,err := pipelineparam.Get()
	if err != nil {
		msg = fmt.Sprintf("获取继承参数失败:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	msg = fmt.Sprintf("镜像信息:%v",param)
	model.PrintInfo(msg)
	
	err = json.Unmarshal([]byte(string(param)), &imageinfos)
	if err != nil {
		msg = fmt.Sprintf("解析镜像信息失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	if len(imageinfos) <= 0 {
		msg = fmt.Sprintf("当前app可用镜像")
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	msg = fmt.Sprintf("镜像信息解析成功")
	model.PrintInfo(msg)
	
	var tParam *tektonapi.ParamTaskInfo = &tektonapi.ParamTaskInfo{}
	tParam.PipelinerunId = paramcreateorder.PipelinerunId
	tParam.TaskName = "createdeployorder-deploy"
	//tClient, err := tektonapi.NewTektonClient(HTTP_TEKTONAPI_SERVER_URL)
	//if err != nil {
	//	msg = fmt.Sprintf("初始化tektonapi client失败,失败原因:%v",err)
	//	model.PrintError(msg)
	//	os.Exit(-1)
	//}
	reqtaskinfoparam,_ := json.Marshal(tParam)
	msg = fmt.Sprintf("获取tekton task信息请求参数:%v",string(reqtaskinfoparam))
	model.PrintInfo(msg)
	
	taskinfo,err := tClient.TaskInfo(tParam)
	if err != nil {
		msg = fmt.Sprintf("获取任务信息失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	msg = fmt.Sprintf("获取任务:%v信息成功",tParam.TaskName)
	model.PrintInfo(msg)
	
	resptaskinfo,_ := json.Marshal(taskinfo)
	msg = fmt.Sprintf("tekton task信息:%v",string(resptaskinfo))
	model.PrintInfo(msg)

	//fClient, err := fabu.NewFaBuClient(HTTP_FABU_SERVER_URL)
	//if err != nil {
	//		msg = fmt.Sprintf("初始化发布系统client失败,失败原因:%v",err)
	//		model.PrintError(msg)
	//		os.Exit(-1)
	//}
	
	reqcommitinfoparam,_ := json.Marshal(paramcreateorder)
	msg = fmt.Sprintf("获取发布 commit信息请求参数:%v",string(reqcommitinfoparam))
	model.PrintInfo(msg)
	
	commitInfo,err := fClient.CommitInfo(paramcreateorder.AppName,paramcreateorder.ProductName,paramcreateorder.ExtEnv)
	if err != nil {
		msg = fmt.Sprintf("获取commit信息失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	respcommitinfo,_ := json.Marshal(commitInfo)
	msg = fmt.Sprintf("发布 commit信息:%v",string(respcommitinfo))
	model.PrintInfo(msg)
	
	if len(commitInfo.Modules) > 0 {
		for _,i := range commitInfo.Modules {
			if _,ok := imagemap[i.ModuleName];!ok {
				imagemap[i.ModuleName] = i.OnlineImage
			}
		}
	}
	if len(imageinfos) > 0 {
		for _ ,i := range 	imageinfos {
			if _,ok := imagemap[i.ModuleName];ok {
				i.OnlineImage = imagemap[i.ModuleName]
				newimages = append(newimages,i)
			}
		}
	}
	
	
	msg = fmt.Sprintf("获取commit info成功")
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("commit镜像信息:%v",newimages)
	model.PrintInfo(msg)
	
	var cParam *fabu.ParamCommit = &fabu.ParamCommit{}
	cParam.MenuName = paramcreateorder.ExtIllustrate
	cParam.MenuType = int64(0)
	cParam.AppName = paramcreateorder.AppName
	cParam.ProductName = paramcreateorder.ProductName
	cParam.Level = paramcreateorder.ExtLevel
	cParam.DeployEnv = paramcreateorder.ExtEnv
	cParam.CommitInfo = newimages
	cParam.RelateUser = paramcreateorder.Operator
	cParam.Influence = paramcreateorder.ExtAffect
	cParam.Uname = taskinfo.Uname
	cParam.Emp = taskinfo.Emp
	cParam.PipelineId = paramcreateorder.PId
	
	reqcommitparam,_ := json.Marshal(cParam)
	msg = fmt.Sprintf("发布 commit请求参数:%v",string(reqcommitparam))
	model.PrintInfo(msg)
	
	m,err := fClient.Commit(cParam)
	if err != nil {
		msg = fmt.Sprintf("创建订单失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	
	respcommit,_ := json.Marshal(m)
	msg = fmt.Sprintf("发布commit请求结果:%v",string(respcommit))
	model.PrintInfo(msg)

	
	respcreateorder.Status = "SUCCESS"
	respcreateorder.MenuId = strconv.FormatInt(m.MenuId,10)
	pipelineparam.Param = strconv.FormatInt(m.MenuId,10)
	pipelineparam.Step = "create-order"
	pipelineparam.PipelinerunId = paramcreateorder.PipelinerunId
	pipelineparam.PId = pid
	pipelineparam.Env = paramcreateorder.ExtEnv
	pipelineparam.Appname = paramcreateorder.AppName
	pipelineparam.ProductName = paramcreateorder.ProductName
	result, err = pipelineparam.Exist()
	if err != nil {
		msg = fmt.Sprintf("参数继承失败,失败原因:%v", err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	if !result {
		err = pipelineparam.Insert()
		if err != nil {
			msg = fmt.Sprintf("参数继承失败,失败原因:%v", err)
			model.PrintError(msg)
			os.Exit(-1)
		}
	} else {
		err = pipelineparam.Update()
		if err != nil {
			msg = fmt.Sprintf("参数继承失败,失败原因:%v", err)
			model.PrintError(msg)
			os.Exit(-1)
		}
	}
	
	
	msg = fmt.Sprintf("创建上线单成功,上线单ID:%v",m.MenuId)
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("继承参数写入数据库成功,参数内容:%v",m.MenuId)
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("上线单创建完成")
	model.PrintInfo(msg)
	
	os.Exit(0)
}
