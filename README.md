beego-admin-demo
================

[![Requirement >= Go 1.2rc1](http://b.repl.ca/v1/Requirement-%3E%3D_Go%201.2rc1-blue.png)]() [![Requirement >= beego 0.9.9](http://b.repl.ca/v1/Requirement-%3E%3D_beego%200.9.9-blue.png)]()

在 `go version go1.8.3 darwin/amd64` 调试通过。

根据 `Gopher` 的邀请，编写演示如何在项目中集成 `Beego-Admin` 。

我没有前端工作的经验，为了演示需要把 `beego` 官网源代码和 `Beego-Admin` 集成。

仅供 `Gopher` 参考。

如下版权属于 `Astaxie` 。

-	Beego [https://github.com/astaxie/beego](https://github.com/astaxie/beego)
-	Beego-Admin [https://github.com/beego/admin](https://github.com/beego/admin)

作用
----

去年使用 beego+beegoAdmin 完成数据库的快速构建，初始化数据的构建，梳理不太稳定的需求，还是很方便的。

环境要求
--------

1.	安装 `MySQL` 数据库
2.	创建数据库 `dp`
3.	使用 `sql/beego-admin-demo.sql` 初始化数据库

配置文件conf/app.conf
---------------------

运行
----

```
cd $GOPATH/src
bee run github.com/devuser/beego-admin-demo
```

访问
----

1.	访问 `http://127.0.0.1:8080/public/login`
2.	输入用户名密码 `admin/admin`
