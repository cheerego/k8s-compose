package compose

import (
	"github.com/cheerego/k8s-compose/cmd"
	"github.com/cheerego/k8s-compose/cmd/compose"
	"github.com/spf13/cobra"
)

type KCompose struct {
	RootCommand *cobra.Command
	Compose     *compose.Compose
}

func NewK() *KCompose {
	return &KCompose{
		Compose: compose.NewC(),
	}
}

func (k KCompose) Run() {
	k.RootCommand = cmd.NewRootCommand(k.Compose)
	k.RootCommand.Execute()
}
