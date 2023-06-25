package cmd

import (
	//"fmt"
	
	"tektonctl/model"
	"tektonctl/server"
	
	"github.com/spf13/cobra"
)

// createorderCmd represents the build command
var createorderCmd = &cobra.Command{
	Use:   "createorder",
	Short: "create order",
	Run: func(cmd *cobra.Command, args []string) {
		extLevel, _ := cmd.Flags().GetString("extlevel")
		extIllustrate, _ := cmd.Flags().GetString("extillustrate")
		extEnv, _ := cmd.Flags().GetString("extenv")
		operator, _ := cmd.Flags().GetString("operator")
		appName, _ := cmd.Flags().GetString("appname")
		productName, _ := cmd.Flags().GetString("productname")
		pipelinerunId, _ := cmd.Flags().GetString("pipelinerunid")
		extAffect, _ := cmd.Flags().GetString("extaffect")
		pid, _ := cmd.Flags().GetString("pid")
		
		
		
		var paramcreateorder model.ParamCreateOrder
		paramcreateorder.AppName = appName
		paramcreateorder.ProductName = productName
		paramcreateorder.PId =pid
		paramcreateorder.PipelinerunId = pipelinerunId
		paramcreateorder.ExtEnv = extEnv
		paramcreateorder.ExtLevel = extLevel
		paramcreateorder.Operator = operator
		paramcreateorder.ExtIllustrate = extIllustrate
		paramcreateorder.ExtAffect = extAffect

		
		server.StartCreateOrder(paramcreateorder)
	},
}

func init() {
	rootCmd.AddCommand(createorderCmd)
	var (
		extLevel string
		extIllustrate string
		extEnv string
		operator string
		appName string
		productName string
		pipelinerunId string
		extAffect string
		pid string
	)
	createorderCmd.Flags().StringVarP(&pipelinerunId, "pipelinerunid", "", "", "pipelinerun id")
	_ = createorderCmd.MarkFlagRequired("pipelinerunid")
	createorderCmd.Flags().StringVarP(&extLevel, "extlevel", "", "", "level")
	_ =createorderCmd.MarkFlagRequired("extlevel")
	createorderCmd.Flags().StringVarP(&extIllustrate, "extillustrate", "", "", "")
	_ = createorderCmd.MarkFlagRequired("extillustrate")
	createorderCmd.Flags().StringVarP(&extEnv, "extenv", "", "", "env")
	_ = createorderCmd.MarkFlagRequired("extenv")
	createorderCmd.Flags().StringVarP(&operator, "operator", "", "", "operator")
	_ = createorderCmd.MarkFlagRequired("operator")
	createorderCmd.Flags().StringVarP(&appName, "appname", "", "", "app name")
	_ = createorderCmd.MarkFlagRequired("appname")
	createorderCmd.Flags().StringVarP(&productName, "productname", "", "", "product name")
	_ = createorderCmd.MarkFlagRequired("productname")
	createorderCmd.Flags().StringVarP(&extAffect, "extaffect", "", "", "")
	_ = createorderCmd.MarkFlagRequired("extaffect")
	createorderCmd.Flags().StringVarP(&pid, "pid", "", "", "short pipeline run id")
	_ = createorderCmd.MarkFlagRequired("pid")
}



