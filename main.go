package main

import (
	"github-trending-backend/db"
	"github-trending-backend/handler"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		Username: "spring",
		Password: "admin",
		DbName:   "githubtrending",
	}

	sql.Connect()
	defer sql.Close()

	logError("An error has occurred here")

	e := echo.New()
	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)
	e.Logger.Fatal(e.Start(":3000"))
}

// Log: Info, Error, Fatal (show error and exit)
// In which file/line does the error occur? -> runtime.Caller() -> pc, file: string, line:int, ok)
// Why use logs? because in Production we need to record errors, we can not use console mode in Production like Development.

func logError(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Fatal(errMsg)
}
