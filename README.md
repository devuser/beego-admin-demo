beego-admin-demo
================

根据 `Gopher` 的邀请，编写演示如何在项目中集成 `Beego-Admin` 。

我没有前端工作的经验，为了演示需要把 `beego` 官网源代码和 `Beego-Admin` 集成。

仅供 `Gopher` 参考。

如下版权属于 `Astaxie` 。

-	Beego [https://github.com/astaxie/beego](https://github.com/astaxie/beego)
-	Beego-Admin [https://github.com/beego/admin](https://github.com/beego/admin)



环境要求
---------

1. 安装 `MySQL` 数据库
2. 创建数据库 `dp`
3. 使用 `sql/beego-admin-demo.sql` 初始化数据库

配置文件conf/app.conf
---------------------

运行
----------------

```
cd $GOPATH/src
bee run github.com/devuser/beego-admin-demo
```

访问
------------------

1. 访问 `http://127.0.0.1:8080/public/login`
2. 输入用户名密码 `admin/admin`


Beego Web
=========

[![Requirement >= Go 1.2rc1](http://b.repl.ca/v1/Requirement-%3E%3D_Go%201.2rc1-blue.png)]() [![Requirement >= beego 0.9.9](http://b.repl.ca/v1/Requirement-%3E%3D_beego%200.9.9-blue.png)]()

An open source project for official documentation website of beego app framework.

Install site locally
--------------------

Beego Web is a `go get` able project:

```
$ go get github.com/beego/beeweb
```

Switch to project root path:

```
$ cd $GOPATH/src/github.com/beego/beeweb
```

Build and run with Go tools:

```
$ go build
$ ./beeweb
```

Or build with bee tool:

```
$ bee run
```

Open your browser and visit [http://localhost:8090](http://localhost:8090).

Build as your site
------------------

This project can be easily transferred as your own documentation site, there are some tips that you may want to know:

-	In the file `conf/app.ini`:

	-	`lang -> types`: languages that you want to support
	-	`lang -> names`: user-friendly name of languages.
	-	It's **NOT** necessary but if you want to you can use GitHub app keys as following format:

		```
		[github]
		client_id=1862bcb2******f36c
		client_secret=308d71ab53ccd858416cfceaed52******53c5f
		```

-	In the file `conf/docTree.json`:

	-	This file saves the file tree(with file name and commit) of your project that is hosted in GitHub. About how to use documentation project please see [beedoc](http://github.com/beego/beedoc). Note that if you added new section to documentation list and you do not want to wait auto-refresh, simple delete this file and restart.
	-	To change the documentation project URL, you need to change it in function `checkDocUpdates` in file `models/models.go`, as well as somewhere in `views`.
