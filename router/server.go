package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func StartServer(port int) {
	r := gin.Default()
	initRouter(r)
	if err := r.Run("127.0.0.1:" + strconv.Itoa(port)); err != nil {
		fmt.Println(err)
	}
}
