package cmd

import (
	"tektonctl/model"
	"tektonctl/server"
	
	"github.com/spf13/cobra"
)

// codescanCmd represents the build command
var codescanCmd = &cobra.Command{
	Use:   "codescan",
	Short: "project code  scan",
	Run: func(cmd *cobra.Command, args []string) {
		outPut, _ := cmd.Flags().GetString("output")
		pipelinerunid, _ := cmd.Flags().GetString("pipelinerunid")
		pid, _ := cmd.Flags().GetString("pid")
		productname, _ := cmd.Flags().GetString("productname")
		appname, _ := cmd.Flags().GetString("appname")
		extenv, _ := cmd.Flags().GetString("extenv")
		
		var paramcodescan model.ParamCodeScan
		//paramcodescan.GitUrl = gitUrl
		paramcodescan.OutPut = outPut
		paramcodescan.ExtEnv = extenv
		paramcodescan.ProductName = productname
		paramcodescan.AppName = appname
		paramcodescan.PipelinerunId = pipelinerunid
		paramcodescan.PId = pid
		server.StartCodeScan(paramcodescan)
	},
}

func init() {
	rootCmd.AddCommand(codescanCmd)
	var (
		//gitUrl string
		outPut string
		pipelinerunid string
		pid string
		productname string
		appname string
		extenv string
	)
	//codescanCmd.Flags().StringVarP(&gitUrl, "giturl", "", "", "project code info")
	//_ = codescanCmd.MarkFlagRequired("giturl")
	codescanCmd.Flags().StringVarP(&outPut, "output", "", "", "git dir path")
	_ = codescanCmd.MarkFlagRequired("output")
	codescanCmd.Flags().StringVarP(&pipelinerunid, "pipelinerunid", "", "", "pipelinerun id")
	_ = codescanCmd.MarkFlagRequired("pipelinerunid")
	codescanCmd.Flags().StringVarP(&pid, "pid", "", "", "short pipelinerun id")
	_ = codescanCmd.MarkFlagRequired("pid")
	codescanCmd.Flags().StringVarP(&productname, "productname", "", "", "product name")
	_ = codescanCmd.MarkFlagRequired("productname")
	codescanCmd.Flags().StringVarP(&appname, "appname", "", "", "app name")
	_ = codescanCmd.MarkFlagRequired("appname")
	codescanCmd.Flags().StringVarP(&extenv, "extenv", "", "", "extenv")
	_ = codescanCmd.MarkFlagRequired("extenv")
}
