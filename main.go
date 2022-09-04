package main

import (
	"backend_github_trending/db"
	"backend_github_trending/handler"
	"backend_github_trending/helper"
	"backend_github_trending/log"
	"backend_github_trending/repository/repo_impl"
	"backend_github_trending/router"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
	"time"
)

func init() {
	//fmt.Println("INIT")
	//os.Setenv("APP_NAME", "github")
	fmt.Println(">>>>", os.Getenv("APP_NAME"))
	log.InitLogger(false)
}

// @title Github Trending API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
func main() {
	fmt.Println("Main function")
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "123987654",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}
	api.SetUpRouter()
	go scheduleUpdateTrending(360*time.Second, repoHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}

//func logErr(errMsg string) {
//	_, file, line, _ := runtime.Caller(1)
//	log.WithFields(log.Fields{
//		"file": file,
//		"line": line,
//	}).Fatal(errMsg)
//}
