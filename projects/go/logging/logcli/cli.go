package logcli

import (
   "fmt"
   logging_service "github.com/ca-mmis/ref-logging/projects/go/logging/service"
   "github.com/urfave/cli"
   "log"
   "os"
   "time"
)

var app = cli.NewApp()

func Init() {
   app.Commands = []cli.Command{
      {
         Name:    "metrics",
         Aliases: []string{"m"},
         Usage:   "Display a list of all metrics",
         Action: func(c *cli.Context) {
            fmt.Println(logging_service.Metrics())
         },
      },
   }
}

func Exec(args []string) int {
   err := app.Run(args)
   if err != nil {
      log.Fatal(err)
      return 1
   }
   return 0
}

func Info() {
   app.Usage = "A CLI for the logging service"
   app.Author = "Joe Chavez"
   app.Version = "1.0.0"
   app.Compiled = compileTime()
}

func compileTime() time.Time {
   info, err := os.Stat(os.Args[0])
   if err != nil {
      return time.Now()
   }
   return info.ModTime()
}


