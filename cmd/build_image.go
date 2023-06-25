package cmd

import (
	"tektonctl/model"
	"tektonctl/server"
	
	"github.com/spf13/cobra"
)

// buildImageCmd represents the build command
var buildImageCmd = &cobra.Command{
	Use:   "buildimage",
	Short: "build container image",
	Run: func(cmd *cobra.Command, args []string) {
		//gitUrl, _ := cmd.Flags().GetString("giturl")
		output ,_ := cmd.Flags().GetString("output")
		pipelineId, _ := cmd.Flags().GetString("pipelineid")
		pid,_ := cmd.Flags().GetString("id")
		appName, _ := cmd.Flags().GetString("appname")
		productName, _ := cmd.Flags().GetString("productname")
		operator, _ := cmd.Flags().GetString("operator")
		extEnv, _ := cmd.Flags().GetString("extenv")
		
		var parambuildImage model.ParamBuildImage
		//parambuildImage.GitUrl = gitUrl
		parambuildImage.Operator = operator
		parambuildImage.ExtEnv = extEnv
		parambuildImage.ProductName = productName
		parambuildImage.AppName = appName
		parambuildImage.PipelinerunId = pipelineId
		parambuildImage.OutPut = output
		parambuildImage.PId = pid
		server.StartBuildImage(parambuildImage)
	},
}

func init() {
	rootCmd.AddCommand(buildImageCmd)
	var (
		//gitUrl string
		outPut string
		pipelineId string
 		appName  string
		productName string
		operator string
		extEnv string
		pid string
	)
	//buildImageCmd.Flags().StringVarP(&gitUrl, "giturl", "", "", "project code info")
	//_ = buildImageCmd.MarkFlagRequired("giturl")
	buildImageCmd.Flags().StringVarP(&outPut, "output", "", "", "git dir path")
	_ = buildImageCmd.MarkFlagRequired("output")
	buildImageCmd.Flags().StringVarP(&pipelineId, "pipelineid", "", "", "tekton pipeline id")
	_ = buildImageCmd.MarkFlagRequired("pipelineid")
	buildImageCmd.Flags().StringVarP(&appName, "appname", "", "", "app name")
	_ = buildImageCmd.MarkFlagRequired("appname")
	buildImageCmd.Flags().StringVarP(&productName, "productname", "", "", "project name")
	_ = buildImageCmd.MarkFlagRequired("productname")
	buildImageCmd.Flags().StringVarP(&operator, "operator", "", "", "operator")
	_ = buildImageCmd.MarkFlagRequired("operator")
	buildImageCmd.Flags().StringVarP(&extEnv, "extenv", "", "", "operate env")
	_ = buildImageCmd.MarkFlagRequired("extenv")
	buildImageCmd.Flags().StringVarP(&pid, "id", "", "", "short id")
	_ = buildImageCmd.MarkFlagRequired("pid")
}

