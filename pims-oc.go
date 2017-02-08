package main

import (
    "os"
    "gopkg.in/urfave/cli.v2" // imports as package "cli"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func readPims(csvFile) {
  log.Println("Reading csv file: ", csvFile)
}

func readOc(dsn) {
  log.Printf("Reading opencart data from: %s\n", dsn)
  db, err := sql.Open("mysql", dsn) // "root:password@/opencart"
}

func main() {
  app := &cli.App{
    Commands: []*cli.Command{
      {
        Name:    "read:pims",
        Usage:   "Read pims csv and display as yml",
        ArgsUsage: "csvFile",
        Action:  func(c *cli.Context) error {
          if c.NArg() == 0 {
            log.error("No csv file passed")
          }
          csvFile := c.Args().First() // Get(0)
          readPims(csvFile)
          return nil
        },
      },
      {
        Name:    "read:oc",
        Usage:   "Read data from opencart mysql database",
        ArgsUsage: "dsn",
        Action:  func(c *cli.Context) error {
          if c.NArg() == 0 {
            log.error("No DSN passed")
          }
          dsn := c.Args().First() // Get(0)
          readOc(dsn)
        },
      },
    },

  }

  app.Name = "pims-oc"
  app.Version = "0.0.0.0"
  app.Run(os.Args)
}
