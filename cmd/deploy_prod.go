package cmd

import (
	"tektonctl/model"
	"tektonctl/server"
	
	"github.com/spf13/cobra"
)

// deployprodCmd represents the build command
var deployprodCmd = &cobra.Command{
	Use:   "deployprod",
	Short: "deploy prod",
	Run: func(cmd *cobra.Command, args []string) {
		step, _ := cmd.Flags().GetString("step")
		uname, _ := cmd.Flags().GetString("uname")
		emp, _ := cmd.Flags().GetString("emp")
		pipelinerunId, _ := cmd.Flags().GetString("pipelinerunid")
		pid, _ := cmd.Flags().GetString("pid")
		productname, _ := cmd.Flags().GetString("productname")
		appname, _ := cmd.Flags().GetString("appname")
		extenv, _ := cmd.Flags().GetString("extenv")
		
		var paramdeployprod model.ParamDeployProd
		//paramdeployprod.MenuId = menuId
		paramdeployprod.Step = step
		paramdeployprod.Emp = emp
		paramdeployprod.PipelinerunId = pipelinerunId
		paramdeployprod.Uname = uname
		paramdeployprod.PId = pid
		paramdeployprod.ExtEnv = extenv
		paramdeployprod.AppName = appname
		paramdeployprod.ProductName = productname
		
		server.StartDeployProd(paramdeployprod)
	},
}

func init() {
	rootCmd.AddCommand(deployprodCmd)
	var (
		//menuId        string
		step          string
		uname         string
		emp           string
		pipelinerunId string
		pid string
		productname string
		appname string
		extenv string
	)
	deployprodCmd.Flags().StringVarP(&pipelinerunId, "pipelinerunid", "", "", "pipelinerun id")
	_ = deployprodCmd.MarkFlagRequired("pipelinerunid")
	deployprodCmd.Flags().StringVarP(&step, "step", "", "", "step")
	_ = deployprodCmd.MarkFlagRequired("step")
	deployprodCmd.Flags().StringVarP(&uname, "uname", "", "", "uname")
	_ = deployprodCmd.MarkFlagRequired("uname")
	deployprodCmd.Flags().StringVarP(&emp, "emp", "", "", "emp")
	_ = deployprodCmd.MarkFlagRequired("emp")
	deployprodCmd.Flags().StringVarP(&pid, "pid", "", "", "short pipeline run id")
	_ = deployprodCmd.MarkFlagRequired("pid")
	deployprodCmd.Flags().StringVarP(&productname, "productname", "", "", "product name")
	_ = deployprodCmd.MarkFlagRequired("productname")
	deployprodCmd.Flags().StringVarP(&appname, "appname", "", "", "app name")
	_ = deployprodCmd.MarkFlagRequired("app name")
	deployprodCmd.Flags().StringVarP(&extenv, "extenv", "", "", "ext env")
	_ = deployprodCmd.MarkFlagRequired("extenv")
}




