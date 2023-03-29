# BlueBird

BlueBird（简称BB）是通过Twitter API抓取Tweet的系统。用户可以添加推特username作为seed，BB可以自动安排抓取。

## 系统构成

BB基于[go-zero](https://github.com/zeromicro/go-zero)开发，由以下模块组成：

`api` api portal

`rpc/seed` seed管理，获取Tweets

`rpc/scheduler` 定时任务Scheduler模块
