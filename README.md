# reconcile-pims-opencart
Reconcile pims export (csv) with opencart export (sql)

## Usage
Run recon: `go run recon.go pims.csv DSN`

where
* `pims.csv` is the excel file of inventory by item exported from pims 2, and then saved as csv manually
* DSN is as exemplified on
[go mysql driver](https://github.com/Go-SQL-Driver/MySQL/#examples):
e.g. `user:password@/dbname`

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
go get github.com/go-sql-driver/mysql
go run recon.go
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


