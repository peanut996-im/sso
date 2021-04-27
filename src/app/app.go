package app

import (
	"sso/handler"
	"sync"

	"framework/cfgargs"
	"framework/logger"
	"framework/net"
	"framework/net/http"

	"github.com/gin-gonic/gin"
)

var (
	once sync.Once
	app  *App
)

type App struct {
	srvCfg  *cfgargs.SrvConfig
	httpSrv *http.Server
}

func GetApp() *App {
	once.Do(func() {
		a := &App{}
		app = a
	})
	return app
}

func (a *App) Init(cfg *cfgargs.SrvConfig) {
	a.srvCfg = cfg

	gin.DefaultWriter = logger.MultiWriter(logger.DefLogger().GetLogWriters()...)

	a.httpSrv = http.NewServer(a.srvCfg)

	a.httpSrv.Use(http.CheckSign(cfg))
	a.httpSrv.AddNodeRoute(a.GetNodeRoute()...)
	go a.httpSrv.Serve(a.srvCfg) //nolint: errcheck

}

func (a *App) GetNodeRoute() []*http.NodeRoute {
	routers := []*http.Route{}
	routers = append(routers, http.NewRoute(net.HTTP_METHOD_POST, "login", handler.SignIn))
	routers = append(routers, http.NewRoute(net.HTTP_METHOD_POST, "register", handler.SignUp))
	routers = append(routers, http.NewRoute(net.HTTP_METHOD_POST, "logout", handler.SignOut))
	routers = append(routers, http.NewRoute(net.HTTP_METHOD_POST, "login/user", handler.SignUp))
	node := http.NewNodeRoute("", routers...)
	return []*http.NodeRoute{node}
}
