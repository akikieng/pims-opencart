# reconcile-pims-opencart
Reconcile pims export (csv) with opencart export (sql)

## Usage
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

4. Run recon: `go run recon.go pims.csv opencart`

## Dev
Dependency mysql requires go version > 1.2 (note on ubuntu 12.04 you get go version 1)

Install dependencies
```bash
sudo apt-get install golang
mkdir ~/.go
export GOPATH=~/.go
go get github.com/go-sql-driver/mysql
go run recon.go
```

# create.sql
* Copied from opencart repo
* Create tables for opencart
* https://github.com/opencart/opencart/blob/master/upload/install/opencart.sql

```bash
wget https://raw.githubusercontent.com/opencart/opencart/master/upload/install/opencart.sql -O test/create.sql
```
