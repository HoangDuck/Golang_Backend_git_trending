package main

import (
	"backend_github_trending/db"
	"backend_github_trending/handler"
	log "backend_github_trending/log"
	"backend_github_trending/repository/repo_impl"
	"backend_github_trending/router"
	"fmt"
	"github.com/labstack/echo"
	"os"
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
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetUpRouter()
	e.Logger.Fatal(e.Start(":3000"))
}

//func logErr(errMsg string) {
//	_, file, line, _ := runtime.Caller(1)
//	log.WithFields(log.Fields{
//		"file": file,
//		"line": line,
//	}).Fatal(errMsg)
//}
