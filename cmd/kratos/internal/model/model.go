package model

import (
	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/model/mongo"
	"github.com/spf13/cobra"
)

// CmdProto represents the proto command.
var CmdModel = &cobra.Command{
	Use:   "model",
	Short: "Generate the model files",
	Long:  "Generate the model files.",
	Run:   run,
}

func init() {
	CmdModel.AddCommand(mongo.CmdMongo)
}

func run(cmd *cobra.Command, args []string) {
}
