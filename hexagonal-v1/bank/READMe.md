# bank

```sh
go get github.com/jmoiron/sqlx
go get -u github.com/go-sql-driver/mysql
go get -u github.com/gorilla/mux
go get github.com/spf13/viper
go get -u go.uber.org/zap




#run
go run .

#test
curl localhost:8080/customers -i
curl localhost:8080/customers/2001 -i
curl localhost:8080/customers/2000/accounts -i
curl localhost:8080/customers/2003/accounts -i -X POST -H "content-type:application/json" -d '{"account_type":"saving","amount":5000}'
```

## add project to git

```sh
git init
git config user.email "k.sillapapam@gmail.com"

# check email configuration
git config user.email

git add .
git commit -m "Initial commit"
git remote add origin https://github.com/adcapp/csharp-web-api-basic.git

```

## setup mysql database

```sh
docker build -t my-mysql /resource/Dockerfile-mysql
```
