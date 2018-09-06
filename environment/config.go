package environment

import (
	"bitbucket.org/blockchain/util"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var envVar *env
var once sync.Once

type env struct {
	prop *viper.Viper
	E    *echo.Echo
}

func Instance() *env {

	once.Do(func() {
		envVar = &env{}
		envVar.init()
		envVar.setupMiddleWare()
	})
	return envVar

}

func (env *env) setupMiddleWare() {

	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(setRequestID()) //Attach request ID's
	e.Use(middleware.Recover())
	e.Use(middleware.AddTrailingSlash())
	env.E = e
}

func Port(portVal string) string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = portVal
	}
	log.Infoln("[BlockChain] Server listening on: ", port)
	return ":" + port
}

func (e *env) init() {

	color.Print(color.Cyan(`
 _|_|_|     _|                         _|               _|_|_|   _|                    _|             
 _|    _|   _|     _|_|       _|_|_|   _|  _|         _|         _|_|_|       _|_|_|        _|_|_|    
 _|_|_|     _|   _|    _|   _|         _|_|           _|         _|    _|   _|    _|   _|   _|    _|  
 _|    _|   _|   _|    _|   _|         _|  _|         _|         _|    _|   _|    _|   _|   _|    _|  
 _|_|_|     _|     _|_|       _|_|_|   _|    _|         _|_|_|   _|    _|     _|_|_|   _|   _|    _|  							
	  `, color.B))
	println("")
	color.Println(color.Yellow("Block chain : API Server", color.B, color.U))
	color.Println(color.Red("Developer: Vaibhav Agrawal <029vaibhav@gmail.com>"))
	e.profileSetUp()
	e.setLogLevel()

}

func (e *env) setLogLevel() {

	if cmp.Equal(util.LOG_LEVEL, "debug") {
		log.SetLevel(log.DebugLevel)
		log.Debugln("Debug enabled")
	} else if cmp.Equal(util.LOG_LEVEL, "info") {
		log.SetLevel(log.InfoLevel)
		log.Debugln("info enabled")
	}

}
func (e *env) profileSetUp() {

	profile := os.Getenv("GOENV")
	if profile == "" {
		profile = "dev"
	}
	e.prop = envSpecificConfig("application-" + profile)
}

func (e *env) Get(key string) interface{} {
	value := os.Getenv(key)
	if value == "" {
		get := e.prop.Get(key)
		return get
	}
	return value

}

func (e *env) Set(key string, val interface{}) {
	e.prop.Set(key, val)
}

func envSpecificConfig(file string) *viper.Viper {

	viper.SetConfigName(file) // name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		e, ok := err.(viper.ConfigParseError)
		if ok {
			log.Fatalf("Error parsing config file: %v", e)
		}
		log.Debugf("No config file used")
	} else {
		log.Debugf("Using config file: %v", viper.ConfigFileUsed())
	}
	return viper.GetViper()
}

//SetRequestID assigns every request with a randomly unique request id
func setRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			rid := util.RandomIdGen()
			res := ctx.Response()
			ctx.Request().Header.Set(echo.HeaderXRequestID, rid)
			res.Header().Set(echo.HeaderXRequestID, rid)
			return next(ctx)
		}
	}
}
