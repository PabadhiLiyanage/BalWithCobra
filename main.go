package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type CommandConfig struct {
	Name  string
	Short string
	Long  string
}
type FlagConfig struct {
	Name       string
	Usage      string
	Shorthend  string
	DefaultVal interface{}
}

func main() {

	directory := "/home/wso2/My_go_project/config"
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".toml" {
			fileName := filepath.Base(path)
			name := fileName[:len(fileName)-len(".toml")]
			fmt.Printf("Processing file: %s\n", name)
			generatingCLICommands(directory, name)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path: %s\n", err)
	}
}

func generateFlagLine(flag FlagConfig, cmdName string) string {
	if cmdName == "" {
		cmdName = "RootCmd"
	}

	var flagline string

	switch flag.DefaultVal.(type) {
	case bool:
		if flag.Shorthend == "" {
			flagline = fmt.Sprintf("%s.Flags().Bool(\"%s\", %t, \"%s\")", cmdName, flag.Name, flag.DefaultVal.(bool), flag.Usage)
		} else {
			flagline = fmt.Sprintf("%s.Flags().BoolP(\"%s\", \"%s\", %t, \"%s\")", cmdName, flag.Name, flag.Shorthend, flag.DefaultVal.(bool), flag.Usage)
		}
	case string:
		if flag.Shorthend == "" {
			flagline = fmt.Sprintf("%s.Flags().String(\"%s\", \"%s\", \"%s\")", cmdName, flag.Name, flag.DefaultVal.(string), flag.Usage)
		} else {
			flagline = fmt.Sprintf("%s.Flags().StringP(\"%s\", \"%s\", \"%s\", \"%s\")", cmdName, flag.Name, flag.Shorthend, flag.DefaultVal.(string), flag.Usage)
		}
	case int:
		flagline = fmt.Sprintf("%s.Flags().IntP(\"%s\", \"%s\", %d, \"%s\")", cmdName, flag.Name, flag.Shorthend, flag.DefaultVal.(int), flag.Usage)
	}

	return flagline
}

func generateSubCommands(subcommands []interface{}, config CommandConfig) (string, string) {
	const subCommandTemp = `
	var {{.Name}} = &cobra.Command{
		Use:   "{{.Use}}",
		Short: "{{.Short}}",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("search called")
			executeBallerinaCommand()
		},
	}
	`
	var subCommandsStr string
	var subLinesStr string

	for _, subcmd := range subcommands {
		if m, ok := subcmd.(map[string]interface{}); ok {
			subconfig := CommandConfig{
				Name:  cast.ToString(m["name"]),
				Short: cast.ToString(m["short"]),
				Long:  cast.ToString(m["description"]),
			}
			tmplSubcmd := template.Must(template.New("SubcommandTemplate").Parse(subCommandTemp))
			var subcmdStrBuffer bytes.Buffer
			data := struct {
				Name  string
				Use   string
				Short string
			}{
				Name:  subconfig.Name + config.Name + "Cmd",
				Use:   subconfig.Name,
				Short: subconfig.Short,
			}
			if err := tmplSubcmd.Execute(&subcmdStrBuffer, data); err != nil {
				log.Fatalf("Error executing subcommand template: %s", err)
			}
			subCommandsStr += subcmdStrBuffer.String() + "\n"
			subline := fmt.Sprintf("%sCmd.AddCommand(%s)", config.Name, data.Name)
			subLinesStr += subline + "\n"
		}
	}

	return subCommandsStr, subLinesStr
}

func generateSubFlagLines(subflags []interface{}, config CommandConfig, generateFlagLine func(FlagConfig, string) string) string {
	var subflagLines string

	for _, table := range subflags {
		if m, ok := table.(map[string]interface{}); ok {
			subflag := FlagConfig{
				Name:       cast.ToString(m["name"]),
				Usage:      cast.ToString(m["usage"]),
				DefaultVal: m["default_val"],
				Shorthend:  cast.ToString(m["shorthand"]),
			}
			commandName := cast.ToString(m["command"]) + config.Name + "Cmd"
			subflagline := generateFlagLine(subflag, commandName)
			subflagLines += subflagline + "\n"
		}
	}

	return subflagLines
}

func generatingBaseCommandFlags(flags []interface{}) string {
	var flagLines string
	for _, table := range flags {
		if m, ok := table.(map[string]interface{}); ok {
			flag := FlagConfig{
				Name:       cast.ToString(m["name"]),
				Usage:      cast.ToString(m["usage"]),
				DefaultVal: m["default_val"],
				Shorthend:  cast.ToString(m["shorthand"]),
			}
			flagline := generateFlagLine(flag, "RootCmd")
			flagLines += flagline + "\n"
		}
	}
	return flagLines
}

func generatingCLICommands(path string, name string) {

	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config := CommandConfig{
		Name:  viper.GetString("base_command.name"),
		Short: viper.GetString("base_command.short"),
		Long:  viper.GetString("base_command.long"),
	}
	filename := config.Name + ".go"
	fmt.Println(filename)
	// Create the command.go file
	if err := os.MkdirAll("output", 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}
	file, err := os.Create(filepath.Join("output", filename))
	if err != nil {
		log.Fatalf("Error creating command.go file: %s", err)
	}
	defer file.Close()

	// Define the template content
	const templateContent = `
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var {{.Name}}Cmd = &cobra.Command{
	Use:     "{{.Name}}",
	Short:   "{{.Short}}",
	Long:    "{{.Long}}",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("called")
	},
}

func init() {
	RootCmd.AddCommand({{.Name}}Cmd)
	//Add subcommands
	{{.SubLines}}
	{{ .FlagLines }}
	// Add subcommand flags here
	{{.SubFlagLines}}

}

{{.SubCommands}}
`

	tmpl := template.Must(template.New("commandTemplate").Parse(templateContent))
	flags, _ := viper.Get("base_command.flag").([]interface{})
	flagLines := generatingBaseCommandFlags(flags)
	subcommands, _ := viper.Get("base_command.subcommand").([]interface{})
	subCommandsStr, subLinesStr := generateSubCommands(subcommands, config)
	subflags, _ := viper.Get("base_command.subcommand_flag").([]interface{})
	subflagLines := generateSubFlagLines(subflags, config, generateFlagLine)

	data := struct {
		Name         string
		Short        string
		Long         string
		FlagLines    string
		SubCommands  string
		SubLines     string
		SubFlagLines string
	}{
		Name:         config.Name,
		Short:        config.Short,
		Long:         config.Long,
		FlagLines:    flagLines,
		SubCommands:  subCommandsStr,
		SubLines:     subLinesStr,
		SubFlagLines: subflagLines,
	}

	var commandData bytes.Buffer

	if err := tmpl.Execute(&commandData, data); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}
	output := commandData.String() //format the content of the document
	formattedContent, err := format.Source([]byte(output))
	if err != nil {
		log.Fatalf("Error formatting Go code: %s", err)
	}
	if err := os.WriteFile(filename, formattedContent, 0644); err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}
