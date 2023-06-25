package model

type ParamGitInit struct {
	GitUrl string `json:"giturl" validate:"required,min=1" label:"giturl"`
	SubModules string `json:"submodules" validate:"required,min=1" label:"submodules"`
	Depth string `json:"depth" validate:"required,min=1" label:"depth"`
	SslVerify string `json:"sslverify" validate:"required,min=1" label:"sslverify"`
	DeleteExisting string `json:"deleteexisting" validate:"required,min=1" label:"deleteexisting"`
	UserHome string `json:"userhome" validate:"required,min=1" label:"userhome"`
	OutPut string `json:"output" validate:"required,min=1" label:"output"`
	PipelinerunId string `json:"pipelinerunid" validate:"required,min=1" label:"pipelinerunid"`
	AppName  string `json:"appname" validate:"required,min=1" label:"appname"`
	PId string `json:"pid" validate:"required,min=1" label:"pid"`
	ProductName string `json:"productname" validate:"required,min=1" label:"productname"`
	ExtEnv  string `json:"env" validate:"required,min=1" label:"env"`
}


type ParamCodeScan struct {
	//GitUrl string `json:"giturl" validate:"required,min=1" label:"giturl"`
	OutPut string `json:"output" validate:"required,min=1" label:"output"`
	PipelinerunId string `json:"pipelinerunid" validate:"required,min=1" label:"pipelinerunid"`
	AppName  string `json:"appname" validate:"required,min=1" label:"appname"`
	PId string `json:"pid" validate:"required,min=1" label:"pid"`
	ProductName string `json:"productname" validate:"required,min=1" label:"productname"`
	ExtEnv  string `json:"env" validate:"required,min=1" label:"env"`
}
type ParamBuildImage struct {
	//GitUrl string `json:"giturl" validate:"required,min=1" label:"giturl"`
	OutPut string `json:"output" validate:"required,min=1" label:"output"`
	PipelinerunId string `json:"pipelinerunid" validate:"required,min=1" label:"pipelinerunid"`
	AppName string `json:"appname" validate:"required,min=1" label:"appname"`
	ProductName string `json:"productname" validate:"required,min=1" label:"productname"`
	Operator string `json:"operator" validate:"required,min=1" label:"operator"`
	ExtEnv string `json:"extenv" validate:"required,min=1" label:"extenv"`
	PId string `json:"pid" validate:"required,min=1" label:"pid"`
}

type ParamPreDeploy struct {
	AppName string `json:"appname" validate:"required,min=1" label:"appname"`
	ProductName string `json:"productname" validate:"required,min=1" label:"productname"`
	Uname string `json:"uname" validate:"required,min=1" label:"uname"`
	Emp string `json:"emp" validate:"required,min=1" label:"emp"`
	ExtEnv string `json:"extenv" validate:"required,min=1" label:"extenv"`
	PipelinerunId string `json:"pipelinerunid" validate:"required,min=1" label:"pipelinerunid"`
	PId string `json:"pid" validate:"required,min=1" label:"pid"`
}

type ParamPrerelease struct {
	AppName string `json:"appname" validate:"required,min=1" label:"appname"`
	ProductName string `json:"productname" validate:"required,min=1" label:"productname"`
	PipelinerunId string `json:"pipelinerunid" validate:"required,min=1" label:"pipelinerunid"`
	ScriptName string `json:"scriptname" validate:"required,min=1" label:"scriptname"`
	PId string `json:"pid" validate:"required,min=1" label:"pid"`
	ExtEnv string `json:"extenv" validate:"required,min=1" label:"extenv"`
}


type ParamCreateOrder struct {
	ExtLevel string `json:"extlevel" validate:"required,min=1" label:"extlevel"`
	ExtIllustrate string `json:"extillustrate" validate:"required,min=1" label:"extillustrate"`
	ExtEnv string `json:"extenv" validate:"required,min=1" label:"extenv"`
	Operator  string `json:"operator" validate:"required,min=1" label:"operator"`
	AppName string `json:"appname" validate:"required,min=1" label:"appname"`
	ProductName string `json:"productname" validate:"required,min=1" label:"productname"`
	//ImageName string `json:"imagename" validate:"required,min=1" label:"imagename"`
	PipelinerunId string `json:"pipelinerunid" validate:"required,min=1" label:"pipelinerunid"`
	PId string `json:"pid" validate:"required,min=1" label:"pid"`
	ExtAffect string `json:"extaffect" validate:"required,min=1" label:"extaffect"`
}

type ParamDeployProd struct {
	Step  string `json:"step" validate:"required,min=1" label:"step"`
	Uname string `json:"uname" validate:"required,min=1" label:"uname"`
	Emp string `json:"emp" validate:"required,min=1" label:"emp"`
 	PipelinerunId string `json:"pipelinerunid" validate:"required,min=1" label:"pipelinerunid"`
	PId string `json:"pid" validate:"required,min=1" label:"pid" `
	AppName string `json:"appname" validate:"required,min=1" label:"appname"`
	ProductName string `json:"productname" validate:"required,min=1" label:"productname"`
	ExtEnv string `json:"extenv" validate:"required,min=1" label:"extenv"`
}