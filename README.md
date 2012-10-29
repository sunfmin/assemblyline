### 取得代码

	$ go get github.com/sunfmin/assemblyline

	$ cd $GOPATH/src/github.com/sunfmin/assemblyline
	$ python -m SimpleHTTPServer
	Serving HTTP on 0.0.0.0 port 8000 ...


### 然后打开新命令行启动进程

	$ assemblyline 
	2012/10/27 09:39:55 Restarting:  map[WashTeaPot:1 MakePotOfTea:1 MakeCupOfTea:1 BoilWater:1 WashWaterPot:1 PickTea:1 WashCup:1]
	2012/10/27 09:39:55 Restarted All Finished
	2012/10/27 09:39:55 Starting WashWaterPot  g1
	2012/10/27 09:39:55 Starting WashCup  g1
	2012/10/27 09:39:55 Starting WashTeaPot  g1
	2012/10/27 09:39:55 Starting PickTea  g1


### 然后访问 http://localhost:8000/

### 演讲稿地址

https://raw.github.com/sunfmin/assemblyline/master/keynote.pdf


演示截屏，图上的圈圈是会转的哟

![screen](https://raw.github.com/sunfmin/assemblyline/master/screenshot.png)