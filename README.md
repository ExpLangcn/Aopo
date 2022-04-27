# Aopo

## [Telegram 频道](https://t.me/+zMfQF-k20ZE0NTdl) | [Twitter](https://twitter.com/ExpLangcn)

### 内网自动化快速打点工具｜资产探测｜漏洞扫描｜服务扫描｜弱口令爆破

### 当前功能

- [x] ICMP并发探测存活IP
- [x] 对A段只进行网关探测
- [x] 集成Xray POC扫描
- [x] 解析Xray POC V1模版
- [x] 并发探测主机端口服务
- [x] 并发探测主机服务弱口令
- [x] 服务弱口令探测模块 支持 SSH
- [x] 服务弱口令探测模块 支持 Mysql
- [x] 服务弱口令探测模块 支持 Mssql
- [x] 服务弱口令探测模块 支持 Mongodb
- [x] 服务弱口令探测模块 支持 Oracle
- [x] 服务弱口令探测模块 支持 FTP
- [x] 服务弱口令探测模块 支持 SMB

### 使用教程

**扫描内网全部资产** （默认探测：10 A段｜172 A段｜192 B段）

> ./Aopo -all

**指定网段扫描** （支持格式：192.168.1.1｜192.168.1.1/24｜192.168.1.1/16｜192.168.1.1/8｜192.168.1.1,192.168.1.2）

> ./Aopo -ip 192.168.1.1/24

**从文件读取目标** （从文件读取目标，支持格式如上 注意：一行一个）

> ./Aopo -ipf target.txt

### 使用截图

![image-20220427185229364](https://tva1.sinaimg.cn/large/e6c9d24egy1h1oh61jojdj217m0u00ya.jpg)

![image-20220427185332450](https://tva1.sinaimg.cn/large/e6c9d24egy1h1oh74w37dj21dq0u0wl1.jpg)

![image-20220427185341238](https://tva1.sinaimg.cn/large/e6c9d24egy1h1oh7b9lb6j215k0u0tfj.jpg)

![image-20220427185251992](https://tva1.sinaimg.cn/large/e6c9d24egy1h1oh6f0j89j210g0u0tdh.jpg)

![image-20220427185255857](https://tva1.sinaimg.cn/large/e6c9d24egy1h1oh6hyfutj21d20poq8a.jpg)

![image-20220427185259556](https://tva1.sinaimg.cn/large/e6c9d24egy1h1oh6k3f0vj20ki0ti0vb.jpg)

![image-20220427185303196](https://tva1.sinaimg.cn/large/e6c9d24egy1h1oh6mlxlyj20u017bwld.jpg)

### 联系方式

#### 微信

![img](https://github.com/ExpLangcn/FuYao/raw/master/img/WechatIMG408.jpeg)

# 知识星球介绍：

【**一次付费 永久免费**，到期联系运营即可免费加入】

星球面向群体：主要面向信息安全研究人员.

更新周期：最晚每两天更新一次.

内容方向：`原创安全工具`｜`安全开发`｜`WEB安全`｜`内网渗透`｜`Bypass`｜`代码审计`｜`CTF`｜`免杀`｜`思路技巧`｜`实战分享`｜`最新漏洞`｜`安全资讯`

[![img](https://camo.githubusercontent.com/19d5b75f5f041159fc3031b157f55872532b7a4f24d79a23fbb254c2877dd6fd/68747470733a2f2f6d6d62697a2e717069632e636e2f6d6d62697a5f6a70672f3977566b37505357496a4a517a4c79524e6844757877506f764c4b7a593878714f71415a6e696356357564395862696338386b657259643349797135307772326b4553756652595952396239565043674463313063644c512f3634303f77785f666d743d6a70656726777866726f6d3d352677785f6c617a793d312677785f636f3d31)](https://camo.githubusercontent.com/19d5b75f5f041159fc3031b157f55872532b7a4f24d79a23fbb254c2877dd6fd/68747470733a2f2f6d6d62697a2e717069632e636e2f6d6d62697a5f6a70672f3977566b37505357496a4a517a4c79524e6844757877506f764c4b7a593878714f71415a6e696356357564395862696338386b657259643349797135307772326b4553756652595952396239565043674463313063644c512f3634303f77785f666d743d6a70656726777866726f6d3d352677785f6c617a793d312677785f636f3d31)
