# <img src="./../img/pics/中介.png" width="50px" height="50px"/>代理模式-Proxy

## :question: 问题 

<img src="./../img/pics/中介 代理1.png">

代理模式是把一个实际存在的对象 ***隐藏*** 在一个与它有相同 ***接口*** 的代理者身后。

## :heavy_check_mark: 解决方法

代理者 ***包含*** 实际对象的接口，通过实际对象要经过代理

## :zap:类图

<img src="./../img/design-patterns-06-proxy.png"/>

## :boy:参与者

  * <img src="./../img/pics/interface.png" width="20px" height="20px">IRealClass（实际类的 ***接口*** ）。确定了实际类中的方法的定义。
  * <img src="./../img/pics/实际司机.png" width="20px" height="20px">RealClass（ ***实际*** 类）。实现了IRealClass中的方法
  * <img src="./../img/pics/中介.png" width="20px" height="20px">Proxy（ ***代理*** ）。实现了IRealClass中的方法，并有指向实际类的 ***引用*** 。

## :sunglasses:评价

### :+1:优点

  * :heavy_plus_sign: 可以通过代理 ***扩展*** 已有应用
  * 可以 ***附加*** 实现实际对象不需要的操作

### :-1:缺点

  * :beetle: 增加代理 ***查错*** 困难
  * :arrow_down: ***降低*** 性能