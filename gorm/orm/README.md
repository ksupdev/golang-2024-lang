# GORM

## setup porject

```sh
go mod init orm
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/go-sql-driver/mysql
go get -u github.com/go-sql-driver/mysql

```

## Note

- In cases where we are unable to implement all methods of an interface, we can declare properties of the interface we need, and it will be considered as implementing that interface.
```
type MyStruct struct {
	TargetInterface
}

type SqlLogger struct {
	logger.Interface
}
```

- Perform a simulation or test to generate SQL for creating a database without actually executing the creation process. This can be configured as needed. `DryRun: true`

```golang
db, err := gorm.Open(dialector, &gorm.Config{
    Logger: &SqlLogger{},
    DryRun: true,
})
```

- Automatically create tables based on the struct definition.

```golang
db.Migrator().CreateTable(Gender{})
```

- Automatically create or alter tables based on the struct definition.

```golang
db.AutoMigrate(Gender{})
```

- Identify column names of a table. By default, GORM auto-generates tables and fields following the struct's snake_case format.
```
type Gender struct {
	ID   uint
	Code uint   `gorm:"primaryKey"`
	Name string `gorm:"column:myname;type:varchar(50);unique;default:Hello;not null"`
}

CREATE TABLE `genders` (`id` bigint unsigned,`code` bigint unsigned AUTO_INCREMENT,`myname` varchar(50) NOT NULL DEFAULT 'Hello',PRIMARY KEY (`code`),CONSTRAINT `uni_genders_myname` UNIQUE (`myname`)) 
```
- Update by Model function , if you send value is empty or zero ,this fuction don't process
```golang
func UpdateGender2(id uint, name string) {
	gender := Gender{
		Name: name,
	}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}
```