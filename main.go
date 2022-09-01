package main

import (
	"backend_github_trending/db"
	"backend_github_trending/handler"
	"backend_github_trending/helper"
	log "backend_github_trending/log"
	"backend_github_trending/repository/repo_impl"
	"backend_github_trending/router"
	"fmt"
	"github.com/labstack/echo"
	"os"
	"time"
)

func init() {
	fmt.Println("INIT")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

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
	}
	api.SetUpRouter()
	go scheduleUpdateTrending(15*time.Second, repoHandler)
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
