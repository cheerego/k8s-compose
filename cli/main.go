package main

import (
	compose "github.com/cheerego/k8s-compose/cmd/kcompose"
	"github.com/cheerego/k8s-compose/pkg/log"
)

func main() {
	defer log.DefaultLogger().Sync()
	k := compose.NewK()
	k.Run()

}
