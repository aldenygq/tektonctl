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
	"tektonctl/model"
	"tektonctl/pkg"
)

func StartCodeScan(paramcodescan model.ParamCodeScan) {
	var (
		msg string
		err  error
		reports []string= make([]string,0)
		pipelineparam *PipelineParam = &PipelineParam{}
		output string
		dir string
		pipelinerunresult *PipelineRunResult = &PipelineRunResult{}
	)
	
	reqparam ,_ := json.Marshal(paramcodescan)
	msg = fmt.Sprintf("请求参数:%v",string(reqparam))
	model.PrintInfo(msg)
	
	err = valid.Struct(paramcodescan)
	if err != nil {
		msg = fmt.Sprintf("参数校验失败,失败原因:%v",pkg.Translate(err.(validator.ValidationErrors)))
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	msg = fmt.Sprintf("请求参数检验通过")
	model.PrintInfo(msg)
	
	pid,_ := strconv.ParseInt(paramcodescan.PId, 10, 64)
	pipelineparam.Step = "git-init"
	pipelineparam.PId = pid
	pipelineparam.Appname = paramcodescan.AppName
	pipelineparam.ProductName = paramcodescan.ProductName
	pipelineparam.Env = paramcodescan.ExtEnv
	param,err := pipelineparam.Get()
	if err != nil {
		msg = fmt.Sprintf("获取继承参数失败:%v",err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	var codeinfo []*model.RespGitInit = make([]*model.RespGitInit,0)
	msg = fmt.Sprintf("代码信息:%v",param)
	model.PrintInfo(msg)
	
	err = json.Unmarshal([]byte(string(param)),&codeinfo)
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
	
	envCommand := fmt.Sprintf("export GRADLE_USER_HOME=/root/.gradle\nexport USER_HOME=/root/\n")
	output,err = pkg.RunCmd(envCommand)
	if err != nil {
		msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",envCommand,output,err)
		model.PrintError(msg)
		os.Exit(-1)
		//return
	}
	
	msg = fmt.Sprintf("环境变量设置成功")
	model.PrintInfo(msg)
	
	for _ ,code := range codeinfo{
			dir,_ =  os.Getwd()
			msg  = fmt.Sprintf("当前目录:%v",dir)
			model.PrintInfo(msg)
			
			checkdir := fmt.Sprintf("%s/%s", paramcodescan.OutPut,code.GitProjectName)
			if err := os.Chdir(checkdir); err != nil {
				msg  = fmt.Sprintf("切换目录失败,目标目录:%v",checkdir)
				model.PrintError(msg)
				os.Exit(-1)
				//return
			}
			if !pkg.CheckFileExist(FILE_SCRIPT_CODE_SCAN_SH) {
				groupidCommand := fmt.Sprintf("grep GROUP_ID gradle.properties | awk -F '=' '{print $2}'")
				artifactidCommand := fmt.Sprintf("grep ARTIFACT_ID gradle.properties | awk -F '=' '{print $2}'")
				groupidOutput,err := pkg.RunCmd(groupidCommand)
				if err != nil {
					msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",groupidCommand,groupidOutput,err)
					model.PrintError(msg)
					os.Exit(-1)
					//return
				}
				
				msg = fmt.Sprintf("模块:%v执行%v完成,执行输出:%v",code.ModName,groupidCommand,output)
				model.PrintInfo(msg)
				
				artifactidOutput,err := pkg.RunCmd(artifactidCommand)
				if err != nil {
					msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",artifactidCommand,artifactidOutput,err)
					model.PrintError(msg)
					os.Exit(-1)
					//return
				}
				
				msg = fmt.Sprintf("模块:%v执行%v完成,执行输出:%v",code.ModName,artifactidCommand,output)
				model.PrintInfo(msg)
				
				reportTmp := fmt.Sprintf("http://sonarqube.sftcwl.com/project/issues?branch=%s&id=%s:%s&resolved=false",code.GitRevision,groupidOutput,artifactidOutput)
				reports = append(reports,reportTmp)
			}else {
				//代码检查
				codescanCommand := fmt.Sprintf("sh %s", FILE_SCRIPT_CODE_SCAN_SH)
				output, err = pkg.RunCmd(codescanCommand)
				if err != nil {
					msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",codescanCommand,output,err)
					model.PrintError(msg)
					os.Exit(-1)
					//return
				}
				msg = fmt.Sprintf("模块:%v执行%v完成,执行输出:%v",code.ModName,codescanCommand,output)
				model.PrintInfo(msg)
				
				groupidCommand := fmt.Sprintf("grep GROUP_ID gradle.properties | awk -F '=' '{print $2}'")
				artifactidCommand := fmt.Sprintf("grep ARTIFACT_ID gradle.properties | awk -F '=' '{print $2}'")
				groupidOutput,err := pkg.RunCmd(groupidCommand)
				if err != nil {
					msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",groupidCommand,groupidOutput,err)
					model.PrintError(msg)
					os.Exit(-1)
					//return
				}
				
				msg = fmt.Sprintf("模块:%v执行%v完成,执行输出:%v",code.ModName,groupidCommand,output)
				model.PrintInfo(msg)
				
				artifactidOutput,err := pkg.RunCmd(artifactidCommand)
				if err != nil {
					msg = fmt.Sprintf("执行命令失败,命令:%v,失败原因,%v,输出:%v",artifactidCommand,artifactidOutput,err)
					model.PrintError(msg)
					os.Exit(-1)
					//return
				}
				
				msg = fmt.Sprintf("模块:%v执行%v完成,执行输出:%v",code.ModName,artifactidCommand,output)
				model.PrintInfo(msg)
				
				reportTmp := fmt.Sprintf("http://sonarqube.sftcwl.com/project/issues?branch=%s&id=%s:%s&resolved=false",code.GitRevision,groupidOutput,artifactidOutput)
				reports = append(reports,reportTmp)
			}
	}
	pipelinerunresult.PipelinerunId = paramcodescan.PipelinerunId
	pipelinerunresult.PId = pid
	pipelinerunresult.Step = "code-scan"
	pipelinerunresult.Env = paramcodescan.ExtEnv
	pipelinerunresult.ProductName = paramcodescan.ProductName
	pipelinerunresult.Appname = paramcodescan.AppName
	pipelinerunresult.ExecResult = strings.Join(reports,"\n")
	result,err := pipelinerunresult.Exist()
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
	
	msg = fmt.Sprintf("任务执行结果输出写入数据库成功,输出内容:%v",strings.Join(reports,"\n"))
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("检测报告:%v",strings.Join(reports,"\n"))
	model.PrintInfo(msg)
	
	msg = fmt.Sprintf("代码扫描完成")
	model.PrintInfo(msg)
	//return
	os.Exit(0)
}
