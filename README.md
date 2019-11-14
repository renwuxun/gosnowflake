#### 分布式ID生成器
分布式ID生成器，参考Twitter的雪花ID生成

#### ID结构
ttttttttttttttttttttttttttttttttttttttttttccccccccccccmmmmmmmmmm
* t:毫秒时间戳（42个bit）
* c:时间点计数器（14个bit）
* m:生成器实例标识（8个bit）

#### 特性
* 唯一性是基本保障
* 最多可存在256个实例
* 相同实例生成的id可保证有序性
* 没有时间点计数上限
* 高负载时时间戳可能会超前，但不会回退（假设系统时间不被修改），负载下降后会自动恢复正常
