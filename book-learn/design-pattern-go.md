# 1.单例模式

## 饿汉模式

> 借助go的init函数来实现

```go
// 注意定义非导出类型
type databaseConn struct {
	//变量
}

var dbConn *databaseConn

func init() {
	dbConn = &databaseConn{}
}

// GetInstance 获取实例
func Db() *databaseConn {
	return dbConn
}

```

## 懒汉模式

> 通俗点说就是延迟加载，不过这块特别注意，要考虑并发环境下，你的判断实例是否已经创建时，是不是用的当前读。

 Java 双重锁实现线程安全的单例模式，双重锁指的是`volatile`和`synchronized`。

```java
class Singleton {
    private volatile static Singleton instance = null;

    private Singleton() {}

    public static Singleton getInstance() {
        if(instance == null) {
            synchronized (Singleton.class) {
                if(instance == null)
                    instance = new Singleton();
            }
        }
        return instance;
    }
}
```

如果不给`instance`属性加上 `volatile`修饰符，那么虽说创建的过程已经用`synchronized`给类加了锁，但是有可能读到的`instance`是线程缓存是滞后的，有可能属性此时已经被其他线程初始化了，所以就必须加上`volatile`保证当前读。

### 1.使用

1. Go 里边没有`volatile`这种机制，需用原子操作`atomic.Load`、`atomic.Store`去读写这个状态变量。

```go
import (
	"sync"
	"sync/atomic"
)

type Singleton struct {
	Host string
	Port int
}

var initialized uint32
var instance *Singleton
var mu sync.Locker

func GetInstance() *Singleton {

	if atomic.LoadUint32(&initialized) == 1 { // 原子操作
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &Singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}
```

### 2.`Go native`

还有另外一种更`Go` native 的写法，比这种写法更简练。如果用 Go 更惯用的写法，我们可以借助其`sync`库中自带的并发同步原语`Once`来实现

```go
import (
	"sync"
)

type Singleton struct {
	Name string
	Port int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Name:"192.168.0.130",
			Port: 3306,
		}
	})
	return instance
}

```

# 2.工厂模式

> 设计模式中的工厂模式是我们编写代码时常用的一种建造型模式，用于创建指定类的实例。
>
> 把一些执行流程明确、但流程细节有差异的业务处理逻辑抽象成公共类库。

在不使用设计模式的时候，我们是怎么创建类的实例的呢？ Java 语言里是通过 new 关键字直接调用类的构造方法，完成实例的创建。

```java
class  Person {}

Person p1 = new Person();
```

Go 语言这类，虽说是非面向对象语言，但也提供了创建类型实例指针的 new 方法。

```go
type Person struct{}

var p1 Person 
p1 = new(Person)
```

## 简单工厂

> Go 语言没有构造函数一说，所以一般会定义 NewXXX 函数来初始化相关类。NewXXX 函数返回接口时就是简单工厂模式。
>
> - 简单工厂：是简单工厂模式的核心，负责实现创建所有实例的内部逻辑。工厂类的创建产品类的方法可以被外界直接调用，创建所需的产品对象。
> - 抽象产品：是简单工厂创建的所有对象的抽象父类/接口，负责描述所有实例的行为。
> - 具体产品：是简单工厂模式的创建目标。

<img src="https://myresou.oss-cn-shanghai.aliyuncs.com/img/640-20231017103709276.png" alt="图片" style="zoom:50%;" />

应用场景里会提供很多语言的打印机，他们都源于一个 Printer 接口。

```go
// Printer 简单工厂要返回的接口类型
type Printer interface {
	Print(name string) string
}
```

程序通过简单工厂向客户端提供需要的语种的打印机。

```go
func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}
```

目前这个场景里先提供两个语种的打印机，他们都是 Printer 接口的具体实现类型。

```go
// CnPrinter 是 Printer 接口的实现，它说中文
type CnPrinter struct {}

func (*CnPrinter) Print(name string) string {
 return fmt.Sprintf("你好, %s", name)
}

// EnPrinter 是 Printer 接口的实现，它说中文
type EnPrinter struct {}

func (*EnPrinter) Print(name string) string {
 return fmt.Sprintf("Hello, %s", name)
}

```

客户端只需要告诉工厂想要哪个语种的打印机产品，工厂就会把产品给返回给客户端。

```go
printer := NewPrinter("en")
fmt.Println(printer.Print("Bob"))
```

### 简单工厂-代码实现

```go
package main

import "fmt"

type Printer interface {
	Print(name string) string
}

// CnPrinter Chinese
type CnPrinter struct {
	name string
}

func (cn *CnPrinter) Print(name string) string {

	return "中文打印机" + name
}

// EnPrinter English
type EnPrinter struct {
	name string
}

func (en *EnPrinter) Print(name string) string {

	return "英文打印机" + name
}

// NewPrinter 简单工厂
func NewPrinter(pName string) Printer {
	switch pName {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}

func main() {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("willy"))
}

```

简单工厂的优点是，简单，缺点如果具体产品扩产，就必须修改工厂内部，增加Case，一旦产品过多就会导致简单工厂过于臃肿，为了解决这个问题，才有了下一级别的工厂模式--工厂方法

## 工厂方法

> 工厂父类（在go中为interface）负责定义创建产品对象的公共接口，子工厂类要实现父工厂中定义的接口，每一个工厂子类则负责生成具体种类的产品对象，**这样做的目的是将产品类的实例化操作延迟到工厂子类中完成。**

工厂方法模式中，不再由单一的工厂类生产产品，而是由工厂类的子类实现具体产品的创建。因此，当增加一个产品时，只需增加一个相应的工厂类的子类, 以解决简单工厂生产太多产品时导致其内部代码臃肿（switch … case分支过多）的问题。

**过程**

在工厂方法中，客户端不需要知道产品的类名，只需要知道所对应的子工厂类即可，具体的产品对象由具体的子工厂类创建，客户端只需要知道创建具体产品的子工厂类对象就行。

当要有新种类的产品对象出现时，我们只需要要新增生产这个产品的子工厂类（这个类去实现父工厂类中定义的接口），就可以通过这个子工厂类就能生产这个新产品对象了，此时我们没有对代码进行修改，只是对代码进行了扩展（新增了一个子工厂类），因此符合开闭原则。

**1.定义产品接口或抽象类**

首先，我们需要定义一个产品接口或抽象类，用于表示具体产品的共性。这个接口或抽象类应该包含产品的属性和方法，以及返回产品信息的方法。

例如，我们可以定义一个产品接口，表示一种水果：定义了两个方法，分别用于获取水果的名称和颜色。

```go
type Fruit interface {
    GetName() string
    GetColor() string
}
```

**2.定义工厂接口或抽象类**

接下来，我们需要定义一个工厂接口或抽象类，用于表示工厂的共性。这个接口或抽象类应该包含创建产品的方法。

例如，我们可以定义一个工厂接口，表示一种水果工厂：定义了一个方法，用于创建一种水果。

```go
type FruitFactory interface {
    CreateFruit() Fruit
}
```

**3.定义具体产品类**

接下来，我们需要定义具体的产品类，实现产品接口或抽象类中定义的方法。

例如，我们可以定义一个具体的苹果类：产品接口中定义的方法，用于获取苹果的名称和颜色。

```go
type Apple struct {
    Name  string
    Color string
}

func (a Apple) GetName() string {
    return a.Name
}

func (a Apple) GetColor() string {
    return a.Color
}
```

**4.定义具体工厂类**

最后，我们需要定义具体的工厂类，实现工厂接口或抽象类中定义的方法。

例如，我们可以定义一个具体的苹果工厂类：现了工厂接口中定义的方法，用于创建苹果产品。

```go
type AppleFactory struct {}

func (f *AppleFactory) CreateFruit() Fruit {
    return &Apple{
        Name:  "苹果",
        Color: "红色",
    }
}
```

**5.使用工厂方法模式**

使用工厂方法模式时，客户端代码只需要调用工厂接口的方法，而无需关心具体的产品对象是如何创建的。

例如，我们可以定义一个客户端代码，用于创建一个苹果工厂，并使用该工厂创建一个苹果产品：先创建了一个苹果工厂，然后使用该工厂创建了一个苹果产品，并输出了该产品的名称和颜色。

```go
func main() {
    factory := &AppleFactory{}
    fruit := factory.CreateFruit()
    fmt.Printf("名称：%s，颜色：%s\n", fruit.GetName(), fruit.GetColor())
}
```

### 工厂方法-代码实现

<img src="https://myresou.oss-cn-shanghai.aliyuncs.com/img/image-20231017135925353.png" alt="image-20231017135925353" style="zoom:50%;" />

```go
package main

import "fmt"

/**
工厂方法
*/

//CalculateHandle 计算器（产品）接口
//每个计算器都需要输入两个数,计算结果
type CalculateHandle interface {
	setNum1(int)
	setNum2(int)
	calResult() int
}

//BaseCalculate 每个计算操作都需要两个数，将这两个数抽出来
type BaseCalculate struct {
	num1 int
	num2 int
}

//PlusCalculate 加法
type PlusCalculate struct {
	*BaseCalculate
}

// MinCalculate 减法
type MinCalculate struct {
	*BaseCalculate
}

//给num1 num2赋值
func (baseCalculate *BaseCalculate) setNum1(num int) {
	baseCalculate.num1 = num
}

func (baseCalculate *BaseCalculate) setNum2(num int) {
	baseCalculate.num2 = num
}

func (p *PlusCalculate) calResult() int {
	return p.num1 + p.num2
}

func (m *MinCalculate) calResult() int {
	return m.num1 - m.num2
}

//CalculateFactory 计算器工厂生产计算器
type CalculateFactory interface {
	Create() CalculateHandle
}

//PlusFactory 加法工厂结构体
type PlusFactory struct {
}

func (p *PlusFactory) Create() CalculateHandle {
	return &PlusCalculate{
		BaseCalculate: &BaseCalculate{},
	}
}

//MinFactory 减法工厂结构体
type MinFactory struct {
}

func (m *MinFactory) Create() CalculateHandle {
	return &MinCalculate{
		BaseCalculate: &BaseCalculate{},
	}
}

func main() {
	var factory CalculateFactory
	factory = &PlusFactory{}
	plusOp := factory.Create()
	plusOp.setNum1(1)
	plusOp.setNum2(4)
	fmt.Printf("加法计算结果:%d\n", plusOp.calResult())

	factory = &MinFactory{}
	MinOp := factory.Create()
	MinOp.setNum1(99)
	MinOp.setNum2(1)
	fmt.Printf("减法计算结果:%d\n", MinOp.calResult())
}

```

工厂方法模式的优点：

- 灵活性增强，对于新产品的创建，只需多写一个相应的工厂类。
- 典型的解耦框架。

工厂方法模式的缺点

- 类的个数容易过多，增加复杂度。
- 增加了系统的抽象性和理解难度。
- 只能生产一种产品，此弊端可使用抽象工厂模式解决。

## 抽象工厂

> 抽象工厂模式：用于创建一系列相关的或者相互依赖的对象。
>
> 这个抽象工厂类定义了多个接口，每个接口都可以生产一种产品。子类工厂在实现抽象工厂的接口后，可以通过不同的接口生产出不同种类的对象。

![image-20231017143732095](https://myresou.oss-cn-shanghai.aliyuncs.com/img/image-20231017143732095.png)

### 抽象工厂-代码实现

```go
package main

import "fmt"

/**
抽象方法
*/

//AbstractFactory 抽象的手机制造商，可以制造手机和手表
type AbstractFactory interface {
	CreateWatch() IWatch
	CreatePhone() ICallPhone
}

// IWatch 手表可以看时间
type IWatch interface {
	WatchTime()
}

// ICallPhone 手机可以打给某个人
type ICallPhone interface {
	CallSomebody()
}

//MIFactory 小米手机制造商可以制造手机和手表
type MIFactory struct {
}

func (mi *MIFactory) CreateWatch() IWatch {
	fmt.Println("制造小米手环")
	return &MIWatch{}
}
func (mi *MIFactory) CreatePhone() ICallPhone {
	fmt.Println("制造小米手机")
	return &MIPhone{}
}

//AppleFactory 苹果手机制造商可以制造手机和手表
type AppleFactory struct {
}

func (apple *AppleFactory) CreateWatch() IWatch {
	fmt.Println("制造Apple Watch")
	return &AppleWatch{}
}
func (apple *AppleFactory) CreatePhone() ICallPhone {
	fmt.Println("制造iPhone")
	return &IPhone{}
}

type MIWatch struct{}

func (miWatch *MIWatch) WatchTime() {
	fmt.Println("小米手环看时间")
}

type MIPhone struct{}

func (miPhone *MIPhone) CallSomebody() {
	fmt.Println("小米手机打电话")
}

type AppleWatch struct{}

func (watch *AppleWatch) WatchTime() {
	fmt.Println("Apple Watch看时间")
}

type IPhone struct{}

func (phone *IPhone) CallSomebody() {
	fmt.Println("iPhone打电话给某人")
}

func main() {
	var factory AbstractFactory
	var watch IWatch
	var phone ICallPhone

	factory = &MIFactory{}
	watch = factory.CreateWatch()
	watch.WatchTime()
	phone = factory.CreatePhone()
	phone.CallSomebody()

	fmt.Println("------------")
	factory = &AppleFactory{}
	appleWatch := factory.CreateWatch()
	appleWatch.WatchTime()
	iPhone := factory.CreatePhone()
	iPhone.CallSomebody()
}

```

同样抽象工厂也具备工厂方法把产品的创建推迟到工厂子类去做的特性，假如未来加入了 VIVO 的产品，我们就可以通过再创建 VIVO 工厂子类来扩展。



抽象工厂模式的优点

- 当需要产品族时，抽象工厂可以保证客户端始终只使用同一个产品的产品族。
- 抽象工厂增强了程序的可扩展性，对于新产品族的增加，只需实现一个新的具体工厂即可，不需要对已有代码进行修改，符合开闭原则。

抽象工厂模式的缺点

- 规定了所有可能被创建的产品集合，产品族中扩展新的产品困难，需要修改抽象工厂的接口。
- 增加了系统的抽象性和理解难度。

# 3.建造者模式

建造者模式(Builder Pattern)使用多个简单的对象一步一步构建成一个复杂的对象

```java
Coffee.builder().name("Latti").price("30").build()
```

## 建造者模式-代码实现



# 4.装饰器模式

先看简单的hello函数

```go
package main
import "fmt"
func hello() {
	fmt.Println("Hello World!")
}
func main() {
	hello()
}
```

现在我们想在打印 `Hello World!` 前后各加一行日志，最直接的实现方式如下：

```go
package main
import "fmt"
func hello() {
    fmt.Println("before")
    fmt.Println("Hello World!")
    fmt.Println("after")
}
func main() {
    hell
```

更好的实现方式是单独编写一个 `logger` 函数，专门用来打印日志：

```go
package main
import "fmt"
func logger(f func()) func() {
    return func() {
        fmt.Println("before")
        f()
        fmt.Println("after")
    }
}
func hello() {
    fmt.Println("Hello World!")
}
func main() {
    hello := logger(hello)
    hello()
}
```

这就相当于装饰器模式，在gin中很常见

```go
r := gin.New()
// 使用中间件
r.Use(gin.Logger(), gin.Recovery())
```

> 装饰器模式（Decorator Pattern）和代理模式（Proxy Pattern）是两种常见的设计模式，它们都属于结构型模式，但在目的和应用上有一些区别。
>
> 装饰器模式旨在动态地给一个对象添加额外的行为，而不需要修改原始对象的结构。它通过创建一个包装器（Wrapper）来包裹原始对象，并在包装器中添加新的功能或修改原有功能。装饰器模式通过组合的方式，可以灵活地添加或删除功能，而不会影响到其他对象。这种模式可以让你在运行时动态地扩展对象的功能。
>
> 代理模式则是在访问对象时引入一层间接层，通过这个间接层控制对原始对象的访问。代理模式常用于控制对敏感对象的访问，或者在访问对象时执行一些额外的操作。代理模式可以提供额外的控制，如权限验证、缓存、延迟加载等。代理模式可以实现对原始对象的保护，使得客户端无需直接与原始对象交互。

![img](https://myresou.oss-cn-shanghai.aliyuncs.com/img/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NoaWRhMjE5,size_16,color_FFFFFF,t_70.png)

## 装饰器模式-代码实现

```go
package main

import "fmt"

type Greeter interface {
	Greet(name string)
}

type SimpleGreeter struct{}

func (s *SimpleGreeter) Greet(name string) {
	fmt.Printf("%s, hello\n", name)
}

type DecoratorGreeter struct {
	Greeter
}

func (d *DecoratorGreeter) Greet(name string) {
	//前置操作
	fmt.Println("装饰器前置")
	d.Greeter.Greet(name)
	//后置操作
	fmt.Println("装饰器后置")
}

func main() {
	greeter := &SimpleGreeter{}
	greeter.Greet("yj")

	decoratorGreeter := DecoratorGreeter{
		Greeter: greeter,
	}

	decoratorGreeter.Greet("whisky")
}
```

# 5.适配器模式

> 适配器模式（Adapter Pattern）又叫作变压器模式，它的功能是将一个类的接口变成客户端所期望的另一种接口，从而使原本因接口不匹配而导致无法在一起工作的两个类能够一起工作，属于结构型设计模式。

![image-20231018142658424](https://myresou.oss-cn-shanghai.aliyuncs.com/img/image-20231018142658424.png)

- 目标 target：是一类含有指定功能的接口
- 使用方 client：需要使用 target 的用户
- 被适配的类 adaptee：和目标类 target 功能类似，但不完全吻合
- 适配器类：adapter：能够将 adaptee 适配转换成 target 的功能类

我们作为用户（client）现在手中持有一个两孔的插头，需要匹配的目标是一个两孔的插座（target），但是现状是我们只找到了三孔的插座（adaptee），于是我们通过在三孔插座上插上一个实现三孔转两孔的适配器（adapter），最终实现了两孔插头与三孔插座之间的适配使用.

```go
package main

import "fmt"

//Target 目标接口，最后需要使用request方法
type Target interface {
	request() string
}

//Adapted 被适配的接口
type Adapted interface {
	translateRequest() string
}

func NewAdapted() Adapted {
	return &AdaptedImpl{}
}

// AdaptedImpl Adapted的一个实现，提供translateRequest具体实现
type AdaptedImpl struct {
}

func (a *AdaptedImpl) translateRequest() string {
	return "adapted method:translateRequest()"
}

type Adapter struct {
	Adapted
}

func NewAdapter(adapted Adapted) *Adapter {
	return &Adapter{
		Adapted: adapted,
	}
}

func (adapter *Adapter) request() string {
	return adapter.translateRequest()
}

func main() {
	adapted := NewAdapted()
	adapter := NewAdapter(adapted)
	request := adapter.request()
	fmt.Println(request)
}
```

# 6.外观模式

> 外观模式：是一种结构型设计模式，提供了一个统一的接口，用于访问子系统中的一组接口。外观模式隐藏了子系统的复杂性，使得客户端可以通过简单的接口与子系统进行交互，而不需要了解子系统的内部实现细节。
>
> 外观模式通过引入一个外观类（Facade Class），将一组子系统的接口封装起来。客户端只需与外观类进行交互，而不需要直接与各个子系统进行通信。外观类负责将客户端的请求转发给适当的子系统对象，并返回结果给客户端。
>
> 外观模式的核心思想是简化客户端与子系统之间的交互，并降低客户端对子系统的依赖。它提供了一种高层次的接口，使得客户端可以更方便地使用子系统。

![图片](https://myresou.oss-cn-shanghai.aliyuncs.com/img/640-20231020134500521.png)

## 外观模式-代码实现

```go
import "fmt"

//OutpatientSystem 门诊系统
type OutpatientSystem struct {
}

func (out *OutpatientSystem) RegistrationOperation() {
	fmt.Println("门诊系统...挂号")
}

// DoctorSystem 医生系统
type DoctorSystem struct {
}

func (doctor *DoctorSystem) treatmentOperation() {
	fmt.Println("医生系统...治疗")
}

// PharmacySystem 药房系统
type PharmacySystem struct {
}

func (p *PharmacySystem) dispensingOperation() {
	fmt.Println("药房系统...发药")
}

type Facade struct {
	outpatient     *OutpatientSystem
	doctorSystem   *DoctorSystem
	pharmacySystem *PharmacySystem
}

func NewFacade() *Facade {
	return &Facade{
		new(OutpatientSystem),
		new(DoctorSystem),
		new(PharmacySystem),
	}
}

func (facade *Facade) start() {
	facade.outpatient.RegistrationOperation()
	facade.doctorSystem.treatmentOperation()
	facade.pharmacySystem.dispensingOperation()
}

func main() {
	facade := NewFacade()
	facade.start()
}
```



# 7.桥接模式

> 桥接模式（Bridge Pattern）又叫作桥梁模式、接口模式或柄体（Handle and Body）模式，指将抽象部分与具体实现部分分离，使它们都可以独立地变化，属于结构型设计模式。



举例：某业务系统, 现需要开发数据库导出工具, 根据SQL语句导出表数据到文件，数据库类型有多种, 目前需要支持MySQL, Orache 未来可能支持 SQLServer。导出格式可能有多种, 目前需要支持CSV和JSON格式

此场景下, 数据库类型是一种维度, 导出格式是另一种维度, 组合可能性是乘法关系，即数据可以从MySQL读出后，导出成CSV 或者JSON格式，对于Oracle也是同样的情况。

![图片](https://myresou.oss-cn-shanghai.aliyuncs.com/img/640.png)

## 桥接模式-代码实现

```go
import (
	"fmt"
	"math/rand"
)

type IDataFetcher interface {
	Fetch(sql string) []interface{}
}

// IDataExport abstraction
type IDataExport interface {
	Fetcher(fetcher IDataFetcher)
	Export(sql string) error
}

type MysqlFetcher struct {
	config string
}

func NewMysqlFetcher(config string) *MysqlFetcher {
	return &MysqlFetcher{
		config: config,
	}
}

func (mysql *MysqlFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from mysql source:" + mysql.config)
	data := make([]interface{}, 0)
	data = append(data, rand.Perm(10), rand.Perm(20))
	return data
}

type OracleFetcher struct {
	config string
}

func NewOracleFetcher(config string) *MysqlFetcher {
	return &MysqlFetcher{
		config: config,
	}
}

func (oracle *OracleFetcher) Fetcher(sql string) []interface{} {
	fmt.Println("Fetch data from mysql source:" + oracle.config)
	data := make([]interface{}, 0)
	data = append(data, rand.Perm(10), rand.Perm(20))
	return data
}

type CsvExporter struct {
	mFetcher IDataFetcher
}

func (ce *CsvExporter) Fetcher(fetcher IDataFetcher) {
	ce.mFetcher = fetcher
}

func (ce *CsvExporter) Export(sql string) error {
	rows := ce.mFetcher.Fetch(sql)
	fmt.Printf("CsvExporter.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}

func NewCsvExporter(fetcher IDataFetcher) IDataExport {
	return &CsvExporter{
		mFetcher: fetcher,
	}
}

type JsonExporter struct {
	mFetcher IDataFetcher
}

func (ce *JsonExporter) Fetcher(fetcher IDataFetcher) {
	ce.mFetcher = fetcher
}

func (ce *JsonExporter) Export(sql string) error {
	rows := ce.mFetcher.Fetch(sql)
	fmt.Printf("Json.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}
func NewJsonExporter(fetcher IDataFetcher) IDataExport {
	return &JsonExporter{
		mFetcher: fetcher,
	}
}

func main() {
	mFetcher := NewMysqlFetcher("mysql://127.0.0.1:3306")
	csvExporter := NewCsvExporter(mFetcher)
	err := csvExporter.Export("select * from xzq")
	if err != nil {
		fmt.Println("导出错误")
	}

	fmt.Printf("\n")
	fetcher := NewOracleFetcher("oracle://192.168.1.1")
	jsonExport := NewJsonExporter(fetcher)
	err = jsonExport.Export("select * from yj")
	if err != nil {
		fmt.Println("导出错误")
	}
}
```



# 8.享元模式

> **享元模式**是一种结构型设计模式， 它的核心思想是通过共享多个对象所共有的相同状态，从而有效的支持在有限的内存中载入大量细粒度的对象。

享元模式的意图是复用对象，节省内存，前提是享元对象是不可变对象，不可变对象指的是初始化之后，对象的状态不会改变了，也就是不会存在被修改的情况。

享元模式的实现思路是，在享元对象的工厂类中，通过一个 Map 来缓存已经创建的享元对象，达到复用的目的。

## 举例：享元模式的应用

假设我们要设计一个多人在线棋牌游戏的平台。在每个牌局里我们会给用户发牌然后进行对战，如果在平台中每创建一个牌局就需要初始化对应的卡牌，这样显然很浪费，因为一套扑克牌里的卡牌是固定的，不管多少个牌局使用的扑克牌都是一样的，只是牌的玩法不一样。

```go
import "fmt"

/**
享元模式
*/

type Card struct {
	Name  string
	Color string
}

var pokerCards = map[int]*Card{
	1: {
		Name:  "A",
		Color: "紅",
	},
	2: {
		Name:  "A",
		Color: "黑",
	},
}

type PokerGame struct {
	Cards map[int]*Card
}

func NewPokerGame() *PokerGame {
	board := &PokerGame{
		Cards: map[int]*Card{},
	}

	for i, val := range pokerCards {
		board.Cards[i] = val
	}
	return board
}

func main() {
	game1 := NewPokerGame()
	game2 := NewPokerGame()
	fmt.Println(game1.Cards[1] == game2.Cards[1])
}
```

## 享元模式-代码实现

![图片](https://myresou.oss-cn-shanghai.aliyuncs.com/img/640-20231023103358174.png)

```go
package main

import "fmt"

/**
享元模式
*/

// Flyweight 结构体表示享元对象，其中包含一个共享的数据 sharedData。
// Operation 方法用于执行操作，它接收一个唯一的数据 uniqueData，并将共享数据和唯一数据打印出来。
type Flyweight struct {
	sharedData string
}

func (f *Flyweight) Operation(uniqueData string) {
	fmt.Printf("Shared data: %s, Unique data: %s\n", f.sharedData, uniqueData)
}

//FlyweightFactory 定义享元工厂
type FlyweightFactory struct {
	flyweights map[string]*Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	flyweightsMap := make(map[string]*Flyweight)
	flyweightsMap["sharedData1"] = &Flyweight{
		"go sharedData1",
	}

	flyweightsMap["sharedData2"] = &Flyweight{
		"go sharedData2",
	}
	return &FlyweightFactory{
		flyweights: make(map[string]*Flyweight),
	}

}

func (ff *FlyweightFactory) GetFlyweight(key string) *Flyweight {
	if flyweight, ok := ff.flyweights[key]; ok {
		return flyweight
	}

	flyweight := &Flyweight{sharedData: key}
	ff.flyweights[key] = flyweight
	return flyweight
}

func main() {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("sharedData1")
	flyweight1.Operation("uniqueData1")

	flyweight2 := factory.GetFlyweight("sharedData2")
	flyweight2.Operation("uniqueData2")
}
```

