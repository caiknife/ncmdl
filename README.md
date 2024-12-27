# ncmdl

**本项目仅供学习，请勿用于其他用途。**

网易云VIP每个下载周期的限额是300首歌曲，我经常因为当月超额导致没办法扩充曲库，所以花了点时间写了这个命令行工具，用来下载更多的歌曲。

但是这个工具是有限制的，如果你下载的链接中包含了VIP歌曲，一般情况下，下载下来的歌曲只有30秒的试听片段。必须提供你在网易云音乐网站的cookies文件，并且你在网易云音乐的账户是VIP，才可以下载VIP歌曲。

使用这个工具，我默认你对于命令行操作、如何导出cookies文件等技能有所了解。

### 安装

本地有Go环境，执行

```shell
go install github.com/caiknife/ncmdl/v2/ncmdl@latest
```
本地没有Go环境，请在[release页面](https://github.com/caiknife/ncmdl/releases)下载最新的二进制文件，请注意区分操作系统和CPU架构。并将二进制文件放在曲库文件夹下。

### 下载

**请自行替换链接**

下载单曲

```shell
ncmdl "https://music.163.com/#/song?id=1824625066"
```

下载专辑

```shell
ncmdl "https://music.163.com/#/album?id=123837273"
```

下载歌单

```shell
ncmdl "https://music.163.com/#/playlist?id=6623715587"
```

### 使用cookies下载完整VIP歌曲

**请自行替换下载链接**

首先请使用浏览器插件，将你在网易云音乐网站的cookies导出成Netscape格式的文件，最好命名为ncm.txt，放在你的个人用户目录下。

这种情况下，直接执行下载命令即可

```shell
ncmdl "https://music.163.com/#/album?id=123837273"
```

如果cookies文件保存在其他路径下，请执行

```shell
ncmdl -c cookies文件所在路径 "https://music.163.com/#/album?id=123837273"
```
