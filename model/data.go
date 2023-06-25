package model

import (
	"fmt"
	
	"tektonctl/pkg"
)

type RespData struct {
	Time   string `json:"time"`
	Level string `json:"level"`
	Msg string      `json:"errmsg"`
}
func ResHandle(res RespData) {
	fmt.Printf("%v [%v] %v\n",pkg.NowTimeStr(),res.Level,res.Msg)
}
func PrintInfo(msg string) {
	var resp RespData
	resp.Level = "INFO"
	resp.Msg = msg
	ResHandle(resp)
}
func PrintError(msg string) {
	var resp RespData
	resp.Level = "ERROR"
	resp.Msg = msg
	ResHandle(resp)
}


type CodeInfo struct {
	ModName string `json:"modName"`
	ModType string `json:"modType"`
	GitProject string `json:"gitProject"`
	GitUrl string `json:"gitUrl"`
	GitRevision string `json:"gitRevision"`
	GitProjectName string `json:"gitProjectName"`
	CommitId string `json:"commitId"`
	Comment string `json:"comment"`
}

type RespGitInit struct {
	*CodeInfo
	Path string `json:"path"`
	OutPut string `json:"output"`
}

type RespCodeScan struct {
	Status string `json:"status"`
	Reports string `json:"report"`
}

type RespBuildImage struct {
	//Env string `json:"env"`
	//Status string `json:"status"`
	//ImageName string `json:"image_name"`
	ImageInfo map[string]string `json:"image_info"`
	ProImage []*RespProjectInfo `json:"pro_image"`
}

type RespProjectInfo struct {
	ModuleName string `json:"module_name"`
	ModuleType string `json:"module_type"`
	OnlineImage string `json:"online_image"`
	CommitImage string `json:"commit_image"`
	CodeUrl     string `json:"code_url"`
}
type RespCreateOrder struct {
	MenuId string `json:"menuId"`
	Status  string `json:"status"`
}

type RespPublish struct {
	ImageName string `json:"image_name"`
	ImageInfo string `json:"image_info"`
}

type RespGitInfo struct {
	GitInfo string `json:"git_info"`
}


