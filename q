#!/bin/sh
rm -rf beego-admin-demo
go build
./beego-admin-demo -syncdb

