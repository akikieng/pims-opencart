# pims-opencart
CLI tool linking between pims and opencart

Features (unchecked are TODO)
- [x] Read pims export csv and display as yml
- [x] Read opencart mysql database and display data similar to yml of pims csv export
- [ ] Validate opencart: list products without categories, without qty, without price, etc.
- [ ] Reconcile quantities between pims csv and opencart mysql database
- [ ] Reconcile categories and category products between pims export and opencart database

## Usage

1. Read pims csv export: `go run pims-oc.go read:pims pims.csv`

* `pims.csv` is the excel file of inventory by item exported from pims 2, and then saved as csv manually
* Produces yml output

2. Read opencart database: `go run pims-oc.go read:oc <DSN>`

* DSN is as exemplified on
  [go mysql driver](https://github.com/Go-SQL-Driver/MySQL/#examples):
  e.g. `user:password@/dbname`
* Produces yml output

3. Validate opencart: `go run pims-oc.go validate:oc <DSN>`

4. Reconcile: `go run pims-oc.go recon pims.csv <DSN>`

## opencart hosted on a2hosting
* mysql users on a2hosting are only allowed from `localhost`
* so it is a restriction in order to limit users connecting to the database from the local machine after ssh'ing into it
* It is still possible to access the mysql database from outside of the a2hosting server by using a SSH tunnel
 * `ssh -L [local port]:[database host]:[remote port] [username]@[remote host]`
 * forward the a2hosting 3306 traffic to your local 3307 port for example
 * Ref: https://support.cloud.engineyard.com/hc/en-us/articles/205408088-Access-Your-Database-Remotely-Through-an-SSH-Tunnel

Alternatively, just export your opencart mysql database as documented here below and import it locally

## Dev notes

### Dependencies
Dependency mysql requires go version > 1.2 (note on ubuntu 12.04 you get go version 1)

Install dependencies
```bash
sudo apt-get install golang
mkdir ~/.go
export GOPATH=~/.go

# https://github.com/Go-SQL-Driver/MySQL/#usage
go get github.com/go-sql-driver/mysql

# https://github.com/urfave/cli#using-the-v2-branch
go get gopkg.in/urfave/cli.v2

# https://github.com/go-yaml/yaml
go get gopkg.in/yaml.v2
```

### creating a local copy of the db
1. create an opencart empty db

```bash
sudo apt-get install mysql-server
mysql -u root -p -e 'CREATE DATABASE opencart'
mysql -u root -p opencart < test/create.sql
```

2. export existing opencart database
  * Click on settings .. export
  * download `store.sql`
  * rename the table prefix from `ocko` to `oc`: `sed -i "s/ocko_/oc_/g" test/store.sql`


3. import `test/store.sql` into the opencart db created

```bash
mysql -u root -p opencart < test/store.sql
```

Note on `create.sql`
* Copied from opencart repo
* Create tables for opencart
* https://github.com/opencart/opencart/blob/master/upload/install/opencart.sql

```bash
wget https://raw.githubusercontent.com/opencart/opencart/master/upload/install/opencart.sql -O test/create.sql
```
