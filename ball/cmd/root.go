package cmd

import (
	"ball/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var helpFlag bool

var RootCmd = &cobra.Command{
	Use:   "ball",
	Short: "Root command",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi")
		//generate.GeneratingCLICommands("/home/wso2/BalWithCobra/config", "health")
	},

	Run: func(cmd *cobra.Command, args []string) {
		executeBallerinaCommand()
	},
}
var (
	javaCmdPass     string
	cmdLineArgsPass []string
)

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	//generate.GeneratingCLICommands("/home/wso2/BalWithCobra/config", "health")
	RootCmd.Flags().BoolP("version", "v", false, "Print version information.")
	RootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "h", false, "Help for ballerina CLI")
	cmdLineArgsPass, javaCmdPass = utils.Setup()

}
