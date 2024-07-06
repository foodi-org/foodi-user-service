# gormctl

自动创建表工具。方便惯用`gorm`模式写表结构的同学使用.

> 注意：该工具仅作为 struct 自动迁移创建 table

使用方式和`gorm`保持一致.
* 修改 `dsn`
* 编写model，并且注册到 `AutoMigrate`
* 执行 `run.go`