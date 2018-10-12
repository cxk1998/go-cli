# go-cli
服务计算 go语言基础cli

### 一、确认目标

这次的资料、背景简介很多，所以看起来有点繁乱，但实际上核心目标只有一个，写一个Unix命令行程序selpg。

在确认目标后就容易确定框架。

## 二、确定逻辑框架

其实可以根据网站上的[C语言selpg](https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c)来得出逻辑框架：

![1539330286904](http://a2.qpic.cn/psb?/V10RTeFd1t8RZE/bPDiTTSXA4ZCxY3Nl43wMQm4khoXrqsKXCZxEi*Z3Fo!/b/dDUBAAAAAAAA&ek=1&kp=1&pt=0&bo=EQFSAQAAAAADF3E!&tl=1&vuin=1285224626&tm=1539327600&sce=60-2-2&rf=viewer_4)



### 三、使用方法

```
-s start_page -e end_page [ -f | -l lines_per_page ][ -d dest ] [ in_filename ]
```

```
Options:
  -d, --dest string   destination
  -e, --end int       end page number (default max-page)
  -f, --formFeed      form feed per page
  -h, --help          help
  -l, --lines int     lines per page (default 72)
  -s, --start int     start page number (default 1)
```

### 四、使用检测

