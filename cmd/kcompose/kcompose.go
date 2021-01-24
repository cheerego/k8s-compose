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
	c := compose.NewC()
	return &KCompose{
		Compose:     c,
		RootCommand: cmd.NewRootCommand(c),
	}
}

func (k KCompose) Run() error {
	return k.RootCommand.Execute()
}
