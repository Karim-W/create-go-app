package contracts

import "github.com/gin-gonic/gin"

type test struct{}

func (t test) SetupRoutes(rg *gin.RouterGroup) {
}

func newTest() RestHandler[gin.RouterGroup] {
	return test{}
}

func TestCreateGinRouterGroup() {
	rg := gin.New().Group("test")
	newTest().SetupRoutes(rg)
}
