package server

import (
	//"encoding/base64"
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	//"os"
	//"os"
	"strconv"
	"strings"
	
	"github.com/go-playground/validator/v10"
	//"github.com/go-playground/validator/v10"
	"gitlab.sftcwl.com/sf-op-public/golang/release"
	"gitlab.sftcwl.com/sf-op-public/golang/tektonapi"
	"tektonctl/model"
	"tektonctl/pkg"
)

func StartPublish(paramprerelease model.ParamPrerelease)  {
	var (
		err       error
		msg string
		newimages []*model.RespProjectInfo = make([]*model.RespProjectInfo ,0)
		imagemap map[string]string = make(map[string]string,0)
		imagenames []string = make([]string,0)
		respbuildimage model.RespBuildImage
		pipelineparam *PipelineParam = &PipelineParam{}
		pipelinerunresult *PipelineRunResult = &PipelineRunResult{}
		output string
	)
	reqparam ,_ := json.Marshal(paramprerelease)
	msg = fmt.Sprintf("请求参数:%v",string(reqparam))
	model.PrintInfo(msg)
	
	err = valid.Struct(paramprerelease)
	if err != nil {
		msg = fmt.Sprintf("参数校验失败,失败原因:%v",pkg.Translate(err.(validator.ValidationErrors)))
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	msg = fmt.Sprintf("请求参数检验通过")
	model.PrintInfo(msg)
	
	pid,_ := strconv.ParseInt(paramprerelease.PId, 10, 64)
	pipelineparam.Step = "build-image"
	pipelineparam.PId = pid
	pipelineparam.Appname = paramprerelease.AppName
	pipelineparam.ProductName = paramprerelease.ProductName
	pipelineparam.Env = paramprerelease.ExtEnv
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
	if len(respbuildimage.ProImage) <= 0 {
		msg = fmt.Sprintf("当前app无可用镜像")
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	im,_ := json.Marshal(respbuildimage.ProImage)
	msg = fmt.Sprintf("镜像信息解析成功,镜像信息:%v",string(im))
	model.PrintInfo(msg)
	
	//rClient, err := release.NewReleaseClient(HTTP_RELEASE_SERVER_URL)
	//if err != nil {
	//	msg = fmt.Sprintf("初始化release client失败,失败原因:%v",err)
	//	model.PrintError(msg)
	//	os.Exit(-1)
		//return
	//}
	for _,v := range respbuildimage.ProImage{
		if paramprerelease.ScriptName == "prerelease" {
			var param *release.ParamPrelease = &release.ParamPrelease{}
			param.ModName = v.ModuleName
			param.PipelineId = paramprerelease.PId
			reqprereleaseparam,_ := json.Marshal(param)
			msg = fmt.Sprintf("预发布请求参数:%v",string(reqprereleaseparam))
			model.PrintInfo(msg)
			image,err := rClient.Prerelease(param)
			if err != nil {
				msg = fmt.Sprintf("预发布失败,失败原因:%v",err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			msg = fmt.Sprintf("模块:%v预发布成功",v.ModuleName)
			model.PrintInfo(msg)
			
			respprerelease,_ := json.Marshal(param)
			msg = fmt.Sprintf("预发布结果:%v",string(respprerelease))
			model.PrintInfo(msg)
			
			if _,ok := imagemap[v.ModuleName]; ok {
				continue
			}
			imagemap[v.ModuleName] = v.ModuleName
			imagename := fmt.Sprintf("模块:%s 镜像:%s",v.ModuleName,image.Image)
			imagenames = append(imagenames,imagename)
			newimages = append(newimages,v)
			dockerPullCommand:= fmt.Sprintf("docker  pull %s",v.CommitImage)
			dockerTagCommand := fmt.Sprintf("docker tag %s %s",v.CommitImage,image.Image)
			dockerPushCommand := fmt.Sprintf("docker push %v",image.Image)
			v.CommitImage = image.Image
			output,err = pkg.RunCmd(dockerPullCommand)
			if err != nil {
				msg  = fmt.Sprintf("执行命令失败，命令：%v，失败原因：%v,err:%v",dockerPullCommand,output,err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			msg = fmt.Sprintf("模块:%v执行docker pull成功,执行输出:%v",v.ModuleName,output)
			model.PrintInfo(msg)
			output,err = pkg.RunCmd(dockerTagCommand)
			if err != nil {
				msg  = fmt.Sprintf("执行命令失败，命令：%v，失败原因：%v,err:%v",dockerTagCommand,output,err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			msg = fmt.Sprintf("模块:%v执行docker tag成功,执行输出:%v",v.ModuleName,output)
			model.PrintInfo(msg)
			output,err = pkg.RunCmd(dockerPushCommand)
			if err != nil {
				msg  = fmt.Sprintf("执行命令失败，命令：%v，失败原因：%v,err:%v",dockerPushCommand,output,err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			msg = fmt.Sprintf("模块:%v执行docker push成功,执行输出:%v",v.ModuleName,output)
			model.PrintInfo(msg)
		}else {
			var param *release.ParamPrelease = &release.ParamPrelease{}
			param.ModName = v.ModuleName
			param.PipelineId = paramprerelease.PId
			
			reqreleaseparam,_ := json.Marshal(param)
			msg = fmt.Sprintf("发布请求参数:%v",string(reqreleaseparam))
			model.PrintInfo(msg)
			
			err := rClient.Release(param)
			if err != nil {
				msg = fmt.Sprintf("发布失败,失败原因:%v",err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			
			msg = fmt.Sprintf("模块:%v发布成功",v.ModuleName)
			model.PrintInfo(msg)
			
			//tClient, err := tektonapi.NewTektonClient(HTTP_TEKTONAPI_SERVER_URL)
			//if err != nil {
			//	msg = fmt.Sprintf("初始化tektonapi client失败,失败原因:%v",err)
			//	model.PrintError(msg)
			//	os.Exit(-1)
				//return
			//}
			var tParam *tektonapi.ParamUpdateStatus = &tektonapi.ParamUpdateStatus{}
			tParam.PipelineRunId = paramprerelease.PipelinerunId
			tParam.Status = "SUCCESS"
			
			requpdatestatusparam,_ := json.Marshal(tParam)
			msg = fmt.Sprintf("tekton update pipelinerun status请求参数:%v",string(requpdatestatusparam))
			model.PrintInfo(msg)
			
			err = tClient.UpdateStatus(tParam)
			if err != nil {
				msg = fmt.Sprintf("更新pipelinerun 状态失败,失败原因:%",err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			
			msg = fmt.Sprintf("%v更新状态成功",paramprerelease.PipelinerunId)
			model.PrintInfo(msg)
		}
	}
	var result bool
	if paramprerelease.ScriptName == "prerelease" {
		imageinfo, _ := json.Marshal(newimages)
		pipelineparam.Param = string(imageinfo)
		pipelineparam.Step = "prerelease"
		pipelineparam.PipelinerunId = paramprerelease.PipelinerunId
		pipelineparam.PId = pid
		pipelineparam.Env = paramprerelease.ExtEnv
		pipelineparam.Appname = paramprerelease.AppName
		pipelineparam.ProductName = paramprerelease.ProductName
		result, err = pipelineparam.Exist()
		if err != nil {
			msg = fmt.Sprintf("参数继承失败,失败原因:%v", err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
		if !result {
			err = pipelineparam.Insert()
			if err != nil {
				msg = fmt.Sprintf("参数继承失败,失败原因:%v", err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
		} else {
			err = pipelineparam.Update()
			if err != nil {
				msg = fmt.Sprintf("参数继承失败,失败原因:%v", err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
		}
		msg = fmt.Sprintf("继承参数写入数据库成功,参数内容:%v",string(imageinfo))
		model.PrintInfo(msg)
	}

	
	pipelinerunresult.PipelinerunId = paramprerelease.PipelinerunId
	pipelinerunresult.Env = paramprerelease.ExtEnv
	pipelinerunresult.PId = pid
	pipelinerunresult.Appname = paramprerelease.AppName
	pipelinerunresult.ExecResult = strings.Join(imagenames,"\n")
	pipelinerunresult.ProductName = paramprerelease.ProductName
	pipelinerunresult.Step = "prerelease"
	result,err = pipelinerunresult.Exist()
	if err != nil {
		msg = fmt.Sprintf("参数继承失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	if !result {
		err = pipelinerunresult.Insert()
		if err != nil {
			msg  = fmt.Sprintf("参数继承失败,失败原因:%v", err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
	}else {
		err = pipelinerunresult.Update()
		if err != nil {
			msg  = fmt.Sprintf("参数继承失败,失败原因:%v", err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
	}
	msg = fmt.Sprintf("任务执行结果输出写入数据库成功,输出内容:%v",strings.Join(imagenames,"\n"))
	model.PrintInfo(msg)
	switch paramprerelease.ScriptName {
	case "prerelease":
		msg = fmt.Sprintf("预发布成功")
	case "release":
		msg = fmt.Sprintf("发布成功")
	}
	model.PrintInfo(msg)
	os.Exit(0)
	
	//return
}
