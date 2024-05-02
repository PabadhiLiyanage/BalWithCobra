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
		if helpFlag {
			executeBallerinaCommand()
		}

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
	fmt.Println("hi")
	RootCmd.Flags().BoolP("version", "v", false, "Print version information.")
	RootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "h", false, "Help for ballerina CLI")
	cmdLineArgsPass, javaCmdPass = utils.Setup()

}

// func Setup() {
// 	ballerinaVersion := "2201.8.5"
// 	balScriptPath, err := os.Executable()
// 	balScriptPath = "/usr/lib/ballerina/distributions/ballerina-2201.8.5/bin/bal"
// 	if err != nil {
// 		log.Fatalln("Error: Cannot find the path of the executable file")
// 	}
// 	for {
// 		link, err := filepath.EvalSymlinks(balScriptPath)
// 		if err != nil {
// 			log.Fatalln("Error resolving symbolic link:", err)

// 		}
// 		if link == balScriptPath {
// 			break
// 		}
// 		balScriptPath = link
// 	}
// 	scriptPathDir := filepath.Dir(balScriptPath)

// 	ballerinaHome, _ := filepath.Abs(filepath.Join(scriptPathDir, ".."))
// 	javaHome, javaCmd := utils.GetJavaSettings(ballerinaHome)

// 	if !utils.CommandExists(javaCmd) {
// 		log.Fatalln("Error: JAVA_HOME is not defined correctly.")
// 	}
// 	if javaHome == "" {
// 		log.Fatalln("You must set the JAVA_HOME variable before running Ballerina.")
// 	}

// 	utils.SetBallerinaCLIWidth()
// 	javaOpts := utils.GetJavaOpts()
// 	ballerinaClasspath := utils.SetBallerinaClassPath(ballerinaHome, javaHome)
// 	//BALLERINA_CLASSPATH_EXT is for outsiders to additionally add
// 	//classpath locations, e.g. AWS Lambda function libraries.
// 	ballerinaClasspathExt := os.Getenv("BALLERINA_CLASSPATH_EXT")
// 	if ballerinaClasspathExt != "" {
// 		ballerinaClasspath = ballerinaClasspath + string(filepath.ListSeparator) + ballerinaClasspathExt
// 	} //filepath.join
// 	//Define Ballerina CLI Arguments
// 	cmdLineArgs := []string{
// 		"-Xms256m",
// 		"-Xmx1024m",
// 		"-XX:+HeapDumpOnOutOfMemoryError",
// 		"-classpath", ballerinaClasspath,
// 		fmt.Sprintf("-Dballerina.home=%s", ballerinaHome),
// 		fmt.Sprintf("-Dballerina.version=%s", ballerinaVersion),
// 		"-Denable.nonblocking=false",
// 		"-Dfile.encoding=UTF8",
// 		"-Dballerina.target=jvm",
// 		fmt.Sprintf("-Djava.command=%s", javaCmd),
// 	}

// 	if runtime.GOOS == "windows" {
// 		cmdLineArgs = append(cmdLineArgs, "-Dcom.sun.management.jmxremote")
// 	} else {
// 		cmdLineArgs = append(cmdLineArgs, "-Djava.security.egd=file:/dev/./urandom")
// 	} //random generator in go

// 	if javaOpts != "" {
// 		JAVA_OPTS := strings.Fields(javaOpts)
// 		cmdLineArgs = append(cmdLineArgs, JAVA_OPTS...)
// 	}
// 	cmdLineArgsPass = cmdLineArgs
// 	javaCmdPass = javaCmd
// }
