package main

import (
	"gin_webSocket_project_IM/router"
	"gin_webSocket_project_IM/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8082")
}
