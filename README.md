# MusicBox

MusicPlayer by go

# 优点

* 闭包很爽，解决C++中各种写回调不方便的问题

# 缺点

* 没有qt中的虚拟成员函数，所以qt中的paintEvent这些virtual成员在go里面无法复写
* 没有qt中的同名函数重载，使用时各种NewQLabel1,NewQLabel2....醉了
* 没有qt中的默认参数，使用时各种麻烦，所有参数都得列出来
* 没有qt中的模板，返回的数据是qlist类型的话，你就不能获得数据了，悲剧
* 没有qt中的运算符重载，会丢失运算符重载的数据

写起来挺麻烦的，实在不好用，建议不用qt widgets写界面，功能不全