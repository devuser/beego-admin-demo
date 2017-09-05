#!/bin/sh
CURPWD=$(pwd)
TODAY=`date '+%Y%m%d%H%M'`
SRCFOLDER=$HOME/tmp/src-v$TODAY
mkdir $SRCFOLDER
echo "备份源代码"
cd $HOME/working/g3
tar zcvf $SRCFOLDER/dp-spring-src-v$TODAY.tar.gz  --exclude .svn  dp-spring
cd $HOME/working
tar zcvf $SRCFOLDER/cpfc-admin-src-v$TODAY.tar.gz  --exclude .svn cpfc-admin-v20161224
cd $HOME/working/goBillProcess/src/vendor/github.com/cpfc
tar zcvf $SRCFOLDER/dp-src-v$TODAY.tar.gz --exclude src/main/vuejs --exclude build --exclude .svn dp
cd $HOME/working
tar zcvf $SRCFOLDER/odssparketl-src-v$TODAY.tar.gz --exclude .svn  --exclude target --exclude spark-assembly-1.6.2-hadoop2.4.0.jar --exclude lib/spark-assembly-1.6.2-hadoop2.4.0.jar odssparketl
cd $CURPWD
