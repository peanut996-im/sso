package app

import (
	"sso/handler"
	"sync"

	"framework/cfgargs"
	"framework/db"
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

	//db
	db.InitRedisClient(cfg)
	err := db.InitMongoClient(cfg)
	if err != nil {
		logger.Fatal("init mongo db err: %v", err)
		return
	}

	gin.DefaultWriter = logger.MultiWriter(logger.DefLogger().GetLogWriters()...)
	a.httpSrv = http.NewServer(cfg)
	if cfg.HTTP.Sign {
		a.httpSrv.Use(http.CheckSign(cfg))
	}

	//gin
	a.httpSrv.AddNodeRoute(a.GetNodeRoute()...)
	go a.httpSrv.Serve(cfg) //nolint: errcheck

	a.srvCfg = cfg
}

//GetNodeRoute Mount routes to the http server.
func (a *App) GetNodeRoute() []*http.NodeRoute {
	routers := []*http.Route{}

	routers = append(routers, http.NewRoute(net.HTTP_METHOD_POST, "login", handler.SignIn))
	routers = append(routers, http.NewRoute(net.HTTP_METHOD_POST, "register", handler.SignUp))
	routers = append(routers, http.NewRoute(net.HTTP_METHOD_POST, "logout", handler.SignOut))

	node := http.NewNodeRoute("", routers...)
	return []*http.NodeRoute{node}
}
