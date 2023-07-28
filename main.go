package main

import (
	_ "github.com/CanftIn/gfmgr/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/CanftIn/gfmgr/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
