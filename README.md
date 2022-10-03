#### 概念

cola 框架是 Alibaba/Cola 的 Go 版本，以下介绍是来自Alibaba/Cola 作者的文章 https://blog.csdn.net/significantfrank/article/details/110934799

#### 架构设计

​	对于一个典型的业务应用系统来说，COLA会做如下层次定义，每一层都有明确的职责定义：
<img width="849" alt="image" src="https://user-images.githubusercontent.com/20272951/193545109-ef29595d-7da8-4948-8ce0-c5df5c5a6a8c.png">

1. 适配层（Adapter Layer）：负责对前端展示（web，wireless，wap）的路由和适配，对于传统B/S系统而言，adapter就相当于MVC中的controller；
2. 应用层（Application Layer）：主要负责获取输入，组装上下文，参数校验，调用领域层做业务处理，如果需要的话，发送消息通知等。层次是开放的，应用层也可以绕过领域层，直接访问基础实施层；
3. 领域层（Domain Layer）：主要是封装了核心业务逻辑，并通过领域服务（Domain Service）和领域对象（Domain Entity）的方法对App层提供业务实体和业务逻辑计算。领域是应用的核心，不依赖任何其他层次；
4. 基础实施层（Infrastructure Layer）：主要负责技术细节问题的处理，比如数据库的CRUD、搜索引擎、文件系统、分布式服务的RPC等。此外，领域防腐的重任也落在这里，外部依赖需要通过gateway的转义处理，才能被上面的App层和Domain层使用。

#### 包结构

<img width="844" alt="image" src="https://user-images.githubusercontent.com/20272951/193545238-b2f63b23-5738-4241-b68b-ce8154a9a36b.png">


各个包结构的简要功能描述，如下表所示：

<img width="836" alt="image" src="https://user-images.githubusercontent.com/20272951/193545274-5b8ed154-90a5-4968-94ab-f7db7e758f61.png">


你可能会有疑问，为什么Domain的model是可选的？**因为COLA是应用架构，不是DDD架构**。在工作中，很多同学问我领域模型要怎么设计，我的回答通常是：无有必要勿增实体。领域模型对设计能力要求很高，没把握用好，一个错误的抽象还不如不抽象，宁可不要用，也不要滥用，不要为了DDD而DDD。

问题的关键是要看，新增的模型没有给你带来收益。比如有没有帮助系统解耦，有没有提升业务语义表达能力的提升，有没有提升系统的可维护性和可测性等等。

模型虽然可选，但DDD的思想是一定要去学习和贯彻的，特别是统一语言、边界上下文、防腐层的思想，值得深入学习，仔细体会。实际上，COLA里面的很多设计思想都来自于DDD。其中就包括领域包的设计。

前面的包定义，都是功能维度的定义。为了兼顾领域维度的内聚性，我们有必要对包结构进行一下微调，即顶层包结构应该是按照领域划分，让领域内聚。

也就是说，我们要综合考虑功能和领域两个维度包结构定义。按照领域和功能两个维度分包策略，最后呈现出来的，是如下图所示的顶层包节点是领域名称，领域之下，再按功能划分包结构。

<img width="842" alt="image" src="https://user-images.githubusercontent.com/20272951/193545313-f19fbdac-dcb7-47a6-abee-7da3814b169a.png">
