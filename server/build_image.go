package server

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	
	"github.com/go-playground/validator/v10"
	//"github.com/go-playground/validator/v10"
	"gitlab.sftcwl.com/sf-op-public/golang/release"
	"tektonctl/model"
	"tektonctl/pkg"
)



func StartBuildImage(parambuildimage model.ParamBuildImage) {
	var (
		err          error
		msg string
		images      []string = make([]string, 0)
		imagenames      map[string]string = make(map[string]string, 0)
		pipelineparam *PipelineParam = &PipelineParam{}
		pipelinerunresult *PipelineRunResult = &PipelineRunResult{}
		respbuildimage model.RespBuildImage
		pros []*model.RespProjectInfo = make([]*model.RespProjectInfo,0)
		output string
		dir  string
	)
	
	reqparam ,_ := json.Marshal(parambuildimage)
	msg = fmt.Sprintf("请求参数:%v",string(reqparam))
	model.PrintInfo(msg)
	
	err = valid.Struct(parambuildimage)
	if err != nil {
		msg = fmt.Sprintf("参数校验失败,失败原因:%v",pkg.Translate(err.(validator.ValidationErrors)))
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	msg = fmt.Sprintf("请求参数检验通过")
	model.PrintInfo(msg)
	
	pid,_ := strconv.ParseInt(parambuildimage.PId, 10, 64)
	pipelineparam.Step = "git-init"
	pipelineparam.PId = pid
	pipelineparam.Appname = parambuildimage.AppName
	pipelineparam.ProductName = parambuildimage.ProductName
	pipelineparam.Env = parambuildimage.ExtEnv
	
	p,_ := json.Marshal(pipelineparam)
	msg = fmt.Sprintf("查询参数:%v",string(p))
	model.PrintInfo(msg)
	
	param,err := pipelineparam.Get()
	if err != nil {
		msg = fmt.Sprintf("获取继承参数失败:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	
	var codeinfo []*model.RespGitInit = make([]*model.RespGitInit,0)
	
	msg = fmt.Sprintf("代码信息:%v",param)
	model.PrintInfo(msg)
	
	err = json.Unmarshal([]byte(string(param)),&codeinfo)
	if err != nil {
		msg = fmt.Sprintf("解析代码信息失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
	}
	if len(codeinfo) <= 0 {
		msg = fmt.Sprintf("当前app无合法代码")
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	msg = fmt.Sprintf("代码信息解析成功")
	model.PrintInfo(msg)
	
		for _, code := range codeinfo {
			if _,ok := imagenames[code.ModName];ok {
				continue
			}
			
			dir,_ =  os.Getwd()
			msg  = fmt.Sprintf("当前目录:%v",dir)
			model.PrintInfo(msg)
			
			checkdir := fmt.Sprintf("%s/%s", parambuildimage.OutPut,code.GitProjectName)
			if err := os.Chdir(checkdir); err != nil {
				msg  = fmt.Sprintf("切换目录失败,目标目录:%v",checkdir)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			
			projPath := strings.TrimPrefix(code.GitUrl, "http://gitlab.sftcwl.com/")
			projPath = strings.TrimSuffix(projPath, ".git")
			var p *release.ParamBuildImage = &release.ParamBuildImage{}
			p.AppName = parambuildimage.AppName
			p.Product = parambuildimage.ProductName
			p.BuildType = parambuildimage.ExtEnv
			p.Branch = code.GitRevision
			p.ModName = code.ModName
			p.ModType = code.ModType
			p.Project = projPath
			p.PipelineId = parambuildimage.PId
			p.Operator = parambuildimage.Operator
			
			//rClient, err := release.NewReleaseClient(HTTP_RELEASE_SERVER_URL)
			//if err != nil {
			//	msg  = fmt.Sprintf("release client初始化失败，失败原因:%v",err)
			//	model.PrintError(msg)
			//	os.Exit(-1)
				//return
			//}
			
			reqbuildimageparam,_ := json.Marshal(p)
			msg = fmt.Sprintf("release build image请求参数:%v",string(reqbuildimageparam))
			model.PrintInfo(msg)
			
			image,err := rClient.BuildImage(p)
			if err != nil {
				msg  = fmt.Sprintf("构建镜像失败，失败原因:%v",err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			
			msg = fmt.Sprintf("模块:%v执行release build 完成",code.ModName)
			model.PrintInfo(msg)
			
			respbuildimage,_ := json.Marshal(image)
			msg = fmt.Sprintf("release 构建结果:%v",string(respbuildimage))
			model.PrintInfo(msg)
			
			im := fmt.Sprintf("模块:%s,镜像:%s",p.ModName,image.Image)
			images = append(images,im)
			imagenames[code.ModName] = image.Image
			
			var pro *model.RespProjectInfo = &model.RespProjectInfo{}
			pro.ModuleName = code.ModName
			pro.ModuleType = code.ModType
			pro.CodeUrl = code.GitUrl
			pro.CommitImage = image.Image
			pros = append(pros,pro)
			
			//构建镜像
			buildEnvCommand := fmt.Sprintf("sh build.sh -p paas%s",parambuildimage.ExtEnv)
			dockerbuildCommand := fmt.Sprintf("docker build --network=host -t %s .",image.Image)
			dockerpushCommand := fmt.Sprintf("docker push %s",image.Image)
			output,err = pkg.RunCmd(buildEnvCommand)
			if err != nil {
				msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",buildEnvCommand,output,err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			msg = fmt.Sprintf("模块:%v执行build.sh完成,执行输出:%v",code.ModName,output)
			model.PrintInfo(msg)
			
			output,err = pkg.RunCmd(dockerbuildCommand)
			if err != nil {
				msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",dockerbuildCommand,output,err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			
			msg = fmt.Sprintf("模块:%v执行docker build完成,执行输出:%v",code.ModName,output)
			model.PrintInfo(msg)
			
			output,err = pkg.RunCmd(dockerpushCommand)
			if err != nil {
				msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",dockerpushCommand,output,err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			
			msg = fmt.Sprintf("模块:%v执行docker push完成,执行输出:%v",code.ModName,output)
			model.PrintInfo(msg)
		}
	respbuildimage.ImageInfo = imagenames
	respbuildimage.ProImage = pros
	respimages,_ := json.Marshal(respbuildimage)
	pipelineparam.PipelinerunId = parambuildimage.PipelinerunId
	pipelineparam.Env = parambuildimage.ExtEnv
	pipelineparam.PId = pid
	pipelineparam.Appname = parambuildimage.AppName
	pipelineparam.Param = string(respimages)
	pipelineparam.ProductName = parambuildimage.ProductName
	pipelineparam.Step = "build-image"
	result,err := pipelineparam.Exist()
	if err != nil {
		msg = fmt.Sprintf("参数继承失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	if !result {
		err = pipelineparam.Insert()
		if err != nil {
			msg  = fmt.Sprintf("参数继承失败,失败原因:%v", err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
	}else {
		err = pipelineparam.Update()
		if err != nil {
			msg  = fmt.Sprintf("参数继承失败,失败原因:%v", err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
	}
	msg = fmt.Sprintf("继承参数写入数据库成功,参数内容:%v",string(respimages))
	model.PrintInfo(msg)
	
	pipelinerunresult.PipelinerunId = parambuildimage.PipelinerunId
	pipelinerunresult.Env = parambuildimage.ExtEnv
	pipelinerunresult.PId = pid
	pipelinerunresult.Appname = parambuildimage.AppName
	pipelinerunresult.ExecResult = strings.Join(images,"\n")
	pipelinerunresult.ProductName = parambuildimage.ProductName
	pipelinerunresult.Step = "build-image"
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
	msg = fmt.Sprintf("任务执行结果输出写入数据库成功,输出内容:%v",strings.Join(images,"\n"))
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("镜像信息:%v",strings.Join(images,"\n"))
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("镜像构建完成")
	model.PrintInfo(msg)
	
	os.Exit(0)
}
