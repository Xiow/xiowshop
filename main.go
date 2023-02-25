package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" //自己手动加
	"github.com/gogf/gf/v2/os/gctx"
	"gongzhaoweishop/internal/cmd"
	_ "gongzhaoweishop/internal/logic" //自己手动加
	_ "gongzhaoweishop/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
