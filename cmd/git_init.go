package cmd

import (
	"tektonctl/model"
	"tektonctl/server"
	
	"github.com/spf13/cobra"
)

// gitinitCmd represents the build command
var gitinitCmd = &cobra.Command{
	Use:   "gitinit ",
	Short: "git code and code init",
	Run: func(cmd *cobra.Command, args []string) {
		gitUrl, _ := cmd.Flags().GetString("giturl")
		subModules, _ := cmd.Flags().GetString("submodules")
		depth, _ := cmd.Flags().GetString("depth")
		sslVerify, _ := cmd.Flags().GetString("sslverify")
		deleteExisting, _ := cmd.Flags().GetString("deleteexisting")
		userHome, _ := cmd.Flags().GetString("userhome")
		outPut, _ := cmd.Flags().GetString("output")
		pipelinerunid, _ := cmd.Flags().GetString("pipelinerunid")
		pid, _ := cmd.Flags().GetString("pid")
		productname, _ := cmd.Flags().GetString("productname")
		appname, _ := cmd.Flags().GetString("appname")
		extenv, _ := cmd.Flags().GetString("extenv")
		
		var paramgitinit model.ParamGitInit
		paramgitinit.GitUrl = gitUrl
		paramgitinit.Depth = depth
		paramgitinit.DeleteExisting = deleteExisting
		paramgitinit.OutPut = outPut
		paramgitinit.SubModules = subModules
		paramgitinit.UserHome = userHome
		paramgitinit.SslVerify = sslVerify
		paramgitinit.ProductName = productname
		paramgitinit.AppName = appname
		paramgitinit.ExtEnv = extenv
		paramgitinit.PipelinerunId = pipelinerunid
		paramgitinit.PId = pid
		server.StartGitInit(paramgitinit)
	},
}

func init() {
	rootCmd.AddCommand(gitinitCmd)
	var (
		gitUrl string
		subModules string
		depth string
		sslVerify string
		deleteExisting  string
		userHome string
		outPut string
		pipelinerunid string
		pid string
		productname string
		appname string
		extenv string
	)
	gitinitCmd.Flags().StringVarP(&gitUrl, "giturl", "", "", "project code info")
	_ = gitinitCmd.MarkFlagRequired("giturl")
	gitinitCmd.Flags().StringVarP(&subModules, "submodules", "", "", "initialize and fetch git submodules.")
	gitinitCmd.Flags().StringVarP(&depth, "depth", "", "", "perform a shallow clone, fetching only the most recent N commits")
	_ = gitinitCmd.MarkFlagRequired("depth")
	gitinitCmd.Flags().StringVarP(&sslVerify, "sslverify", "", "", "")
	gitinitCmd.Flags().StringVarP(&deleteExisting, "deleteexisting", "", "", "clean out the contents of the destination directory if it already exists before cloning.")
	_ = gitinitCmd.MarkFlagRequired("deleteexisting")
	gitinitCmd.Flags().StringVarP(&userHome, "userhome", "", "/tekton/home", "absolute path to the user's home directory. Set this explicitly if you are running the image as a non-root user or have overridden the gitInitImage param with an image containing custom user configuration.")
	_ = gitinitCmd.MarkFlagRequired("userhome")
	gitinitCmd.Flags().StringVarP(&outPut, "output", "", "", "the git repo will be cloned onto the volume backing this Workspace.")
	_ = gitinitCmd.MarkFlagRequired("output")
	gitinitCmd.Flags().StringVarP(&pipelinerunid, "pipelinerunid", "", "", "pipelinerun id")
	_ = gitinitCmd.MarkFlagRequired("pipelinerunid")
	gitinitCmd.Flags().StringVarP(&pid, "pid", "", "", "short pipelinerun id ")
	_ = gitinitCmd.MarkFlagRequired("pid")
	gitinitCmd.Flags().StringVarP(&productname, "productname", "", "", "product name")
	_ = gitinitCmd.MarkFlagRequired("productname")
	gitinitCmd.Flags().StringVarP(&appname, "appname", "", "", "app name")
	_ = gitinitCmd.MarkFlagRequired("appname")
	gitinitCmd.Flags().StringVarP(&extenv, "extenv", "", "", "extenv")
	_ = gitinitCmd.MarkFlagRequired("extenv")
	
}
