package cmd

import (
	"tektonctl/model"
	"tektonctl/server"
	
	"github.com/spf13/cobra"
)

// codescanCmd represents the build command
var deployCmd = &cobra.Command{
	Use:   "predeploy",
	Short: "pre release env deploy",
	Run: func(cmd *cobra.Command, args []string) {
		//imageName, _ := cmd.Flags().GetString("imagename")
		appName, _ := cmd.Flags().GetString("appname")
		productName, _ := cmd.Flags().GetString("productname")
		uname, _ := cmd.Flags().GetString("uname")
		emp, _ := cmd.Flags().GetString("emp")
		extEnv, _ := cmd.Flags().GetString("extenv")
		pipelinerunid, _ := cmd.Flags().GetString("pipelinerunid")
		pid, _ := cmd.Flags().GetString("pid")
		
		var parampreDeploy model.ParamPreDeploy
		//parampreDeploy.ImageName = imageName
		parampreDeploy.AppName = appName
		parampreDeploy.ProductName = productName
		parampreDeploy.Uname = uname
		parampreDeploy.Emp = emp
		parampreDeploy.ExtEnv = extEnv
		parampreDeploy.PipelinerunId = pipelinerunid
		parampreDeploy.PId = pid
		server.StartPreDeploy(parampreDeploy)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	var (
		//imageName string
		appName string
		productName string
		uname string
		emp string
		extEnv string
		pipelinerunid string
		pid string
	)
	//deployCmd.Flags().StringVarP(&imageName, "imagename", "", "", "image name")
	//_ = deployCmd.MarkFlagRequired("imagename")
	deployCmd.Flags().StringVarP(&uname, "uname", "", "", "uname")
	_ = deployCmd.MarkFlagRequired("uname")
	deployCmd.Flags().StringVarP(&emp, "emp", "", "", "emp")
	_ = deployCmd.MarkFlagRequired("emp")
	deployCmd.Flags().StringVarP(&appName, "appname", "", "", "app name")
	_ = deployCmd.MarkFlagRequired("appname")
	deployCmd.Flags().StringVarP(&productName, "productname", "", "", "project name")
	_ = deployCmd.MarkFlagRequired("productname")
	deployCmd.Flags().StringVarP(&extEnv, "extenv", "", "", "deploy env ")
	_ = deployCmd.MarkFlagRequired("extenv")
	deployCmd.Flags().StringVarP(&pipelinerunid, "pipelinerunid", "", "", "pipelinerun id")
	_ = deployCmd.MarkFlagRequired("pipelinerunid")
	deployCmd.Flags().StringVarP(&pid, "pid", "", "", "short pipelinerun id")
	_ = deployCmd.MarkFlagRequired("pid")
}

