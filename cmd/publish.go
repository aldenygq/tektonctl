

package cmd

import (
	"tektonctl/model"
	"tektonctl/server"
	
	"github.com/spf13/cobra"
)

// publishCmd represents the build command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "code publish，contain prerelease and release",
	Run: func(cmd *cobra.Command, args []string) {
		pipelinerunId, _ := cmd.Flags().GetString("pipelinerunid")
		scriptName, _ := cmd.Flags().GetString("scriptname")
		pid, _ := cmd.Flags().GetString("pid")
		appName, _ := cmd.Flags().GetString("appname")
		productName, _ := cmd.Flags().GetString("productname")
		extEnv, _ := cmd.Flags().GetString("extenv")
		
		var paramprerelease model.ParamPrerelease
		paramprerelease.PipelinerunId = pipelinerunId
		paramprerelease.PId = pid
		paramprerelease.ScriptName = scriptName
		paramprerelease.ExtEnv = extEnv
		paramprerelease.ProductName = productName
		paramprerelease.AppName = appName

		server.StartPublish(paramprerelease)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
	var (
		pipelinerunId string
		scriptName string
		//buildImageName string
		pid string
		appName string
		productName string
		extEnv string
	)
	publishCmd.Flags().StringVarP(&pipelinerunId, "pipelinerunid", "", "", "pipelinerun id")
	_ = publishCmd.MarkFlagRequired("pipelinerunid")
	publishCmd.Flags().StringVarP(&scriptName, "scriptname", "", "", "script name")
	_ = publishCmd.MarkFlagRequired("scriptname")
	//publishCmd.Flags().StringVarP(&buildImageName, "buildimagename", "", "", "build image name")
	//_ = publishCmd.MarkFlagRequired("buildimagename")
	publishCmd.Flags().StringVarP(&pid, "pid", "", "", "pipelune run id")
	_ = publishCmd.MarkFlagRequired("pid")
	publishCmd.Flags().StringVarP(&appName, "appname", "", "", "app name")
	_ = publishCmd.MarkFlagRequired("appname")
	publishCmd.Flags().StringVarP(&productName, "productname", "", "", "project name")
	_ = publishCmd.MarkFlagRequired("productname")
	publishCmd.Flags().StringVarP(&extEnv, "extenv", "", "", "deploy env ")
	_ = publishCmd.MarkFlagRequired("extenv")
}


