package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	
	//"os"
	
	"github.com/go-playground/validator/v10"
	"tektonctl/model"
	"tektonctl/pkg"
)
func StartGitInit(paramgitinit model.ParamGitInit) {
	var (
		err error
		respgitinits []*model.RespGitInit = make([]*model.RespGitInit,0)
		pipelineparam *PipelineParam = &PipelineParam{}
		msg string
	)
	reqparam ,_ := json.Marshal(paramgitinit)
	msg = fmt.Sprintf("请求参数:%v",string(reqparam))
	model.PrintInfo(msg)
	
	err = valid.Struct(paramgitinit)
	if err != nil {
		msg = fmt.Sprintf("参数校验失败,失败原因:%v",pkg.Translate(err.(validator.ValidationErrors)))
		model.PrintError(msg)
		os.Exit(-1)
		return
	}
	
	msg = fmt.Sprintf("请求参数检验通过")
	model.PrintInfo(msg)
	
	var codeinfo []*model.CodeInfo = make([]*model.CodeInfo,0)
	decoded, _:= base64.StdEncoding.DecodeString(paramgitinit.GitUrl)
	
	msg = fmt.Sprintf("代码信息:%v",string(decoded))
	model.PrintInfo(msg)
	
	err = json.Unmarshal([]byte(string(decoded)),&codeinfo)
	if err != nil {
		msg = fmt.Sprintf("解析代码信息失败,失败原因:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	if len(codeinfo) <= 0 {
		msg = fmt.Sprintf("当前app无合法代码")
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	msg = fmt.Sprintf("代码信息解析成功")
	model.PrintInfo(msg)
	
	for _,code := range codeinfo {
		var respgitinit *model.RespGitInit = &model.RespGitInit{}
		var output string
		var  dir string
		subDirectory := code.GitProjectName
		//清理checkout 目录下的资源
		checkdir := fmt.Sprintf("%s/%s",paramgitinit.OutPut,subDirectory)
		
		dir,_ =  os.Getwd()
		msg  = fmt.Sprintf("当前目录:%v",dir)
		model.PrintInfo(msg)
		
		if paramgitinit.DeleteExisting == "true" {
			deletepath := fmt.Sprintf(" rm -fr %s/*",checkdir)
			output,err := pkg.RunCmd(deletepath)
			if err != nil {
				msg  = fmt.Sprintf("执行命令失败，命令：%v，失败原因：%v,err:%v",deletepath,output,err)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
		}
		
		
		gitinitcmd := fmt.Sprintf("/ko-app/git-init -url=%s  -revision=%s  -path=%s  -sslVerify=%s   -submodules=%s  -depth=%s",
			code.GitUrl,
			code.GitRevision,
			checkdir,
			paramgitinit.SslVerify,
			paramgitinit.SubModules,
			paramgitinit.Depth)
		output,err = pkg.RunCmd(gitinitcmd)
		if err != nil {
			msg = fmt.Sprintf("执行命令失败，命令：%v，失败原因：%v,err:%v",gitinitcmd,output,err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
		
		
		msg = fmt.Sprintf("模块:%v执行git init 成功,执行输出:%v",code.ModName,output)
		model.PrintInfo(msg)
		
		if err := os.Chdir(checkdir); err != nil {
			msg = fmt.Sprintf("切换目录失败，目录：%v，失败原因：%v",checkdir,err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
		
		dir,_ =  os.Getwd()
		msg = fmt.Sprintf("当前目录:%v",dir)
		model.PrintInfo(msg)
		
		gitrevcmd := fmt.Sprintf("git rev-parse HEAD")
		output,err = pkg.RunCmd(gitrevcmd)
		if err != nil {
			msg = fmt.Sprintf("执行命令失败，命令：%v，失败原因：%v,err:%v",gitrevcmd,output,err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
		
		msg = fmt.Sprintf("模块:%v执行git rev-parse HEAD 成功,执行输出:%v",code.ModName,output)
		model.PrintInfo(msg)
		
		if err := os.Chdir(paramgitinit.OutPut); err != nil {
			msg = fmt.Sprintf("执行命令失败，命令：%v，失败原因：%v,err:%v",gitrevcmd,output,err)
			model.PrintError(msg)
			os.Exit(-1)
			//return
		}
		
		dir,_ =  os.Getwd()
		msg = fmt.Sprintf("当前目录:%v",dir)
		model.PrintInfo(msg)
		
		respgitinit.CodeInfo = code
		respgitinit.Path = checkdir
		respgitinit.OutPut = paramgitinit.OutPut
		respgitinits = append(respgitinits,respgitinit)
	}
	
	re,_ := json.Marshal(respgitinits)
	pid,_ := strconv.ParseInt(paramgitinit.PId, 10, 64)
	pipelineparam.PId = pid
	pipelineparam.PipelinerunId = paramgitinit.PipelinerunId
	pipelineparam.Env = paramgitinit.ExtEnv
	pipelineparam.ProductName = paramgitinit.ProductName
	pipelineparam.Appname = paramgitinit.AppName
	pipelineparam.Param = string(re)
	pipelineparam.Step = "git-init"
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
	
	msg = fmt.Sprintf("继承参数写入数据库成功,参数内容:%v",string(re))
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("代码信息:%v",string(re))
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("源码下载完成")
	model.PrintInfo(msg)
	
	os.Exit(0)
}
