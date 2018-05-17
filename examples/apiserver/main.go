package main

import (
	"github.com/micro-plat/hydra/examples/apiserver/services/order"
	"github.com/micro-plat/hydra/examples/apiserver/services/user"
	"github.com/micro-plat/hydra/hydra"
)

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("hydra_test_lcw"),
		hydra.WithSystemName("collector"),
		hydra.WithServerTypes("api-web-rpc"),
		hydra.WithDebug())
	app.Micro("/user/login", user.NewLoginHandler)
	app.Micro("/order/query", order.NewQueryHandler)
	app.Micro("/order/bind", order.NewBindHandler)

	app.Start()
}
