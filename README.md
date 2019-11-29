# log-agent
小型日志收集系统

# 项目背景
a、收集系统日志,后续对系统日志分析,数据挖掘,可以发现用户的一些规律等等...\
b、系统日志收集对测试人员,开发人员,定位问题等等..\
c、日志收集,对后续接口做一些监控等等

# 解决方案
1、把机器上的日志实时收集，统一存储到中心系统\
2、针对这些日志建立索引，通过索引即可以找到对应日志\
3、通过提供友好的web界面,通过web完成日志搜索

# 面临的问题
1、实时日志量非常大、每天几十亿条\
2、日志实时收集，延迟控制在分钟级别\
3、能够水平扩展

# 业界的解决方案(ELK)
1、什么是ELK\
Elasticsearch是个开源分布式搜索引擎，它的特点有：分布式，零配置，自动发现，索引自动分片，索引副本机制，restful风格接口，多数据源，自动搜索负载等。\
Logstash是一个完全开源的工具，他可以对你的日志进行收集、过滤，并将其存储供以后使用（如，搜索）。\
Kibana 也是一个开源和免费的工具，它Kibana可以为 Logstash 和 ElasticSearch 提供的日志分析友好的 Web 界面，可以帮助您汇总、分析和搜索重要数据日志\
2、ELK的简单架构\
![](https://github.com/duanlei123/log-agent/blob/master/IMG/ELK%E7%AE%80%E5%8D%95%E6%9E%B6%E6%9E%84.png)\
3、ELK方案问题\
    a、运维成本高,每增加一个日志收集，都需要手动修改配置\
    b、监控缺失，无法准确获取logstash的状态\
    c、无法做定制化开发或维护\
    d、logstash消耗资源大，运行占用cpu和内存高，另外没有消息队列缓存，存在数据丢失隐患。\

4、ELK的优点是非常多的。\
5、ELK架构也有很多种。可根据实际情况选择。网上有很多关于ELK的架构说明，这里不再阐述。

# 系统架构(不完成)
![](https://github.com/duanlei123/log-agent/blob/master/IMG/%E6%97%A5%E5%BF%97%E6%94%B6%E9%9B%86%E7%B3%BB%E7%BB%9F%E6%9E%B6%E6%9E%84%E5%9B%BE.png)\
log Agent : 需要部署到没一台业务应用机器,负责根据配置读取应用日志。\
kafka(http://kafka.apachecn.org/quickstart.html): 接收log Agent发送的日志。\ kafka 设计为了水平扩展日志系统。\
ES(https://es.xiaoleilu.com/): 存kafak获取应用日志,创建索引.\
hadoop:离线计算(不做)、\
storm:日志分析(不做)、\
除上面hadoop 和 storm 还可以水平扩展很多系统平台\

欢迎大家前来指导.......

