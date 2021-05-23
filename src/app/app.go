package app

import (
	"framework/api"
	"sso/handler"
	"sync"

	"framework/cfgargs"
	"framework/db"
	"framework/logger"
	"framework/net/http"

	"github.com/gin-gonic/gin"
)

var (
	once sync.Once
	app  *App
)

type App struct {
	cfg     *cfgargs.SrvConfig
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
	a.cfg = cfg
	//db
	db.InitRedisClient(cfg)
	err := db.InitMongoClient(cfg)
	if err != nil {
		logger.Fatal("init mongo db err: %v", err)
		return
	}
	//gin
	gin.DefaultWriter = logger.MultiWriter(logger.DefLogger().GetLogWriters()...)
	a.httpSrv = http.NewServer(cfg)
	a.httpSrv.AddNodeRoute(a.GetNodeRoute()...)
	go a.httpSrv.Serve(cfg) //nolint: errcheck

}

//GetNodeRoute Mount routes to the http server.
func (a *App) GetNodeRoute() []*http.NodeRoute {
	routers := []*http.Route{}

	routers = append(routers, http.NewRoute(api.HTTPMethodPost, "login", handler.SignIn))
	routers = append(routers, http.NewRoute(api.HTTPMethodPost, "register", handler.SignUp))
	routers = append(routers, http.NewRoute(api.HTTPMethodPost, "logout", handler.SignOut))

	node := http.NewNodeRoute("", routers...)
	return []*http.NodeRoute{node}
}

func (a *App) GetSrvCfg() *cfgargs.SrvConfig {
	return a.cfg
}
