package web

import (
	"flag"
	"fmt"
	"github.com/oceango/di"
	"github.com/oceango/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Application struct {
	port string
	workDir string
	router *router.Router
}

func NewApplication(router *router.Router) *Application {
	workDir,_ := os.Getwd()
	var port string

	port = viper.GetString("server.port")

	return &Application{
		port:        ":"+port,
		workDir:     workDir,
		router:      router,
	}
}

func (a *Application) Run() {
	printBanner(a.workDir)
	log.Print("application starting...")
	initContainer(a)

	log.Println("application started and listen on port:" + a.port)
	panic(http.ListenAndServe(a.port, a.router))
}

func BuildConfiguration()  {
	workDir, _ := os.Getwd()
	mode := flag.String("mode", "dev", "application environment mode")

	flag.Int("port", 1016, "application port")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigName("application") // name of config file (without extension)
	if mode !=nil {
		viper.SetConfigName("application-" + *mode)
	}
	viper.SetConfigType("yml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(workDir +"/config") // optionally look for config in the working directory

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func initContainer(a *Application)  {
	// inject configuration

	//di.Singleton(func() {
	//
	//	return
	//})

	// inject application
	di.Singleton(func() *Application{
		return a
	})

	// inject db
	//di.Singleton(func() *gorm.DB {
	//	return db.GetDb()
	//})
}

func printBanner(workDir string)  {
	banner := `
        ___     ___     ___     ___     ___     ___     ___   
   /\  \   /\  \   /\  \   /\  \   /\__\   /\  \   /\  \  
  /::\  \ /::\  \ /::\  \ /::\  \ /:| _|_ /::\  \ /::\  \ 
 /:/\:\__/:/\:\__/::\:\__/::\:\__/::|/\__/:/\:\__/:/\:\__\
 \:\/:/  \:\ \/__\:\:\/  \/\::/  \/|::/  \:\:\/__\:\/:/  /       http://www.oceango.tech
  \::/  / \:\__\  \:\/  /  /:/  /  |:/  / \::/  / \::/  / 
   \/__/   \/__/   \/__/   \/__/   \/__/   \/__/   \/__/  

	`
	filename := workDir + "/config/banner.txt"

	if fileExists(filename) {
		file, err := ioutil.ReadFile(filename)
		if err == nil {
			banner = string(file)
		}
	}

	fmt.Println(banner)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

