package main

import (
  "os"
  "gopkg.in/urfave/cli.v2" // imports as package "cli"
  "log"
  "encoding/csv"
  "io"
  "fmt"
  "strconv"
  "strings"
  "gopkg.in/yaml.v2"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Product struct {
  Category string
  ID string
  Desc string
  Qty int
  Price float64
}

// how to read a csv
// http://stackoverflow.com/q/26437796/4126114
func readPims(filename string) *map[string][]Product {
  //log.Println("Reading csv file: ", filename)
  handle, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }

  defer handle.Close()

  reader := csv.NewReader(handle)
  data := map[string][]Product{}
  current := ""
  for i := 0; i < 10000; i++ { // 10k lines max
      record, err := reader.Read()
      if err == io.EOF {
        break
      } else if err != nil {
        log.Fatal(err)
      }

      // skip header
      if i==0 {
        continue
      }

      // drop empty fields
      var filtered []string
      for _, field := range record {
        if field!="" {
          filtered = append(filtered,field)
        }
      }

      // category
      if len(filtered)==1 {
        current = filtered[0]
        data[current] = []Product{}
        continue
      }

      if len(filtered)==5 {
        if strings.Contains(filtered[2]," Pcs") {
          qty,err := strconv.Atoi(strings.Replace(filtered[2]," Pcs","",1))
          if err!= nil {
            log.Fatalf("Invalid qty field found in %s / %s: %s",current, filtered[0], filtered[2])
          }
          price,err := strconv.ParseFloat(filtered[3],32)
          if err!= nil {
            log.Fatalf("Invalid price field found in %s / %s: %s",current, filtered[0], filtered[2])
          }
          pr := Product{Category: current, ID: filtered[0], Desc: filtered[1], Qty: qty, Price: price}
          data[current] = append(data[current], pr)
          continue
        }
      }

      if len(filtered)==0 {
        continue
      }

      fmt.Printf("Unparsed Line %d (%d): %s\n", i, len(filtered), filtered)
      //for _, field := range filtered {
      //  fmt.Println(field)
      //}
  }

  //log.Println("Done Reading csv file: ", filename)
  return &data
}

func readOc(dsn string) {
  log.Printf("Reading opencart data from: %s\n", dsn)
  db, err := sql.Open("mysql", dsn) // "root:password@/opencart"
  if err != nil {
    log.Fatal("connect fails:", err)
  }
  db.Close()
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
            log.Fatal("No csv file passed")
          }
          csvFile := c.Args().First() // Get(0)
          data := readPims(csvFile)

          // Golang - How to print struct variables in console?
          // http://stackoverflow.com/a/24512194/4126114
          // fmt.Printf("%+v\n", data)

          // convert to yaml
          // https://github.com/go-yaml/yaml#example
          // Note:
          // struct fields are important to be uppercased
          // documented here: https://godoc.org/gopkg.in/yaml.v2#Marshal
          y, err := yaml.Marshal(data)
          if err != nil {
              return err
          }
          fmt.Println(string(y))

          return nil
        },
      },
      {
        Name:    "read:oc",
        Usage:   "Read data from opencart mysql database",
        ArgsUsage: "dsn",
        Action:  func(c *cli.Context) error {
          if c.NArg() == 0 {
            log.Fatal("No DSN passed")
          }
          dsn := c.Args().First() // Get(0)
          readOc(dsn)
          return nil
        },
      },
    },

  }

  app.Name = "pims-oc"
  app.Version = "0.0.0.0"
  app.Run(os.Args)
}
