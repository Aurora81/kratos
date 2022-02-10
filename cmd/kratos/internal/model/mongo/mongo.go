package mongo

import (
	"fmt"
	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/model/parse"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
	"strings"
)

// CmdServer the service command.
var CmdMongo = &cobra.Command{
	Use:   "mongo",
	Short: "Generate the model mongo implementations",
	Long:  "Generate the model mongo implementations. Example: kratos model mongo internal/model/xxx.go -target-dir=internal/data -model=XXX",
	Run:   run,
}
var targetDir string
var modelName string

func init() {
	CmdMongo.Flags().StringVarP(&targetDir, "target-dir", "t", "internal/data", "generate target directory")
	CmdMongo.Flags().StringVarP(&modelName, "model", "m", "", "data model")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the model file. Example: kratos model mongo internal/model/xxx.go")
		return
	}

	if modelName == "" {
		fmt.Fprintln(os.Stderr, "Please specify the model name")
		return
	}

	fields := parse.Parse(args[0], modelName)
	if fields == nil {
		fmt.Fprintln(os.Stderr, "Can't find type {{modelName}}")
		return
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exsits\n", targetDir)
		return
	}

	s := Repo{
		ModelName: modelName,
		Fields: fields,
	}

	to := path.Join(targetDir, strings.ToLower(s.ModelName)+".go")
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s already exists: %s\n", s.ModelName, to)
	}

	b, err := s.execute()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(to, b, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(to)
}
