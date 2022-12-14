# MySQL 规范

- [MySQL 规范](#mysql-规范)
  - [基础规范](#基础规范)
  - [命名规范](#命名规范)
  - [流程规范](#流程规范)
  - [设计规范](#设计规范)
    - [表设计规范](#表设计规范)
    - [字段设计规范](#字段设计规范)
    - [索引规范](#索引规范)
  - [业务使用规范](#业务使用规范)

## 基础规范

- **尽可能使用 InnoDB 作为表的存储引擎**。在 MySQL 5.6 之后，InnoDB 被设置成为默认的存储引擎，支持事务和行级锁。
- **数据库和数据表统一使用 UTF8MB4 字符编码**。
- **所有的表和字段必须添加注释**。这是一个习惯问题，即使做到了顾名思义，也存在歧义或误读的可能性，为了防止这种情况的发生，同时也为了提高维护性，必须使用 comment 设定注释。
- **尽量控制表行数在 500 万以内**。数据量越多，查询效率就越差，同时，为了防止磁盘内存 IO 过高和查询时间过长，对于数据量增长较快的表，应使用分表、合理分区等方案进行控制。
- **尽可能采用冷热数据分离策略**。在 MySQL 中，数据表列数最大限制为 4096 行，每条元数据大小不能超过 65535 字节。根据业务场景，将常用字段和非常用字段分离，减小表宽度，可以提高热数据的内存缓存命中率，降低 CPU 使用率和 IO 流。
- **禁止存储图片、文件等二进制数据**。MySQL 虽然支持文件对象的存储，但在实际开发中，应尽量避免直接存储文件。正确的做法是将文件存储在文件服务器中，将文件服务器的路径存储于数据库中。

## 命名规范

- **对象名称必须使用小写，多单词统一使用下划线分割**。
- **命名的单词必须做到顾名思义、简洁，表名最好不要超过 16 个字符，字段名称最好不要超过 32 个字符**。
- **禁止使用保留字，尽量避免使用关键字**。
- **临时表必须以 tmp* 开头，以日期结尾，备份表必须以 bak* 开头，以日期结尾**。

## 流程规范

- **禁止在线上做数据库压力测试**。
- **环境和数据库强绑定，禁止跨环境使用数据库**，比如，测试环境一定不能使用生产环境的数据库。
- **只有 DBA 拥有 super 权限**。
- **养成查看 SQL 性能的习惯，借助性能分析工具进行分析**。譬如：EXPLAIN 语句 | showprofile | mySQLsla 等。
- **禁止在业务高峰期批量更新、查询数据**。比如，定时任务可以在流量较低的凌晨进行操作。
- **活动推广、系统上线以及平台上新务必对流量进行评估**。防患于未然，否则可能造成数据库服务器流量瓶颈进而影响业务。
- **所有的新建表、修改表操作都需要 DBA 进行确认：字段的类型、长度以及索引**。确保表结构设计最优。

## 设计规范

### 表设计规范

- **每张表的索引数务必控制在 5 个以内**。索引能够提高查询的效率，同时也会降低写操作的效率，占用更多的内存空间，因而应该合理控制索引的数量。
- **每一张 InnoDB 表都必须含有一个主键**。InnoDB 是一种索引组织表：数据的存储的逻辑顺序和索引的顺序是相同的。每个表都可以有多个索引，但是表的存储顺序只能有一种，InnoDB 是按照主键索引的顺序来组织表的。不要使用可能会更新的列作为主键，同时尽量不要使用 UUID、MD5、HASH 等无序的字符串作为主键。在没有特别的情况下，应使用自增的整型或发号器作为主键。
- **避免使用外键约束**。外键可以保证数据的准确性、参照完整性，每次进行写操作时都会走校验数据知否正确的流程，这意味着写操作时会有额外的性能损耗，因此，建议在业务层实现数据的参照完整性。有一种例外情况，**倘若字表的写操作很少，务必使用外键约束**。
- **设置数据表架构应考虑后期扩展型**。此项需要业务开发人员和 DBA 进行沟通，根据实际场景设计数据库表。
- **遵循范式与冗余平衡原则**

  > 第一范式：具有原子性；

  > 第二范式：主键列与非主键列遵循完全函数依赖关系；

  > 第三范式：非主键列之间没有传递函数依赖关系；

  范式设计是数据结构的一种思想，在实际的业务场景中应灵活运用，合理取舍，一味追求三范式无疑会影响程序的性能。适当的冗余是可以提高查询的效率的，前提要保证是主键的冗余。

- **每张数据库表字段的个数应控制在 20 以内**。数据表的宽度与内存占用的大小成正比，在进行读写操作时，表宽度越长，消耗的内存越多、占用的 IO 流越大，操作的效率越低。在实际使用中，应尽可能地根据业务场景细分，区分冷热数据，进行分表设计。

### 字段设计规范

- **禁止在表中使用模棱两可都字段**。比如多处使用 ext、ext_1、extend_n，即使每一个都有 comment，也会使维护者困惑。
- **避免使用 TEXT、BLOB、ENUM 数据类型**。MySQL 内存临时表不支持 TEXT、BLOB 这样的大数据类型，如果查询中包含这样的数据，在排序等操作时，就不能使用内存临时表，必须使用磁盘临时表进行，毋庸置疑会降低查询的效率。枚举类型在数据库中保存的值实际上是整数，容易引起歧义。
- **避免数据列中存在空值**。NULL 列比较特殊，需要额外的空间来保存，同时会造成索引失效。
- **DATETIME 应使用 INT 或 TIMESTAMP 替换**。TIMESTAMP 与 INT 占 4 位字节，而 DATETIME 占 8 位字节。因此，存储时间时，应使用 INT 或 TIMESTAMP 替换。TIMESTAMP 的可读性高，INT 的灵活性高，因而经常需要使用计算操作的应当使用 INT 存储，否则使用 TIMESTAMP。
- **金额相关的数据必须使用 DECIMAL 数据类型**。DECIMAL 类型为精准浮点数，在计算时不会丢失精度，可以自定义其长度，可用于存储比 BIGINT 更大的整型数据。
- **表与表关联的键名保持一致或以关联表名的缩写为前缀**。
- **固定长度的字符串字段务必使用 CHAR**。
- **使用 UNSIGNEG 存储非负整数**。
- **禁止敏感数据以明文形式存储**。如密码、身份证号等。

### 索引规范

- **一张表的索引不能超过 5 个**。
- **重要的 SQL 语句必须带上索引作为条件**。
- **避免冗余和重复索引**。
  > 重复索引： 在相同的列上按照相同的顺序创建的相同类型的索引。
  > 冗余索引： 两个索引按照相同的顺序覆盖了相同的列。
- **禁止在索引列进行数学运算和函数运算**。MySQL 不擅长于运算，需要计算的应该移至代码业务层。
- **区分度高的索引优先**。区分度高的索引优先，可以显著提高查询的效率，特别是在 JOIN 连表查询时。

## 业务使用规范

- **重要的 SQL 语句必须带上索引作为条件**。比如，删、改数据，一定要带上 WHERE。
- **避免使用 SELECT \* 查询字段**。
- **避免数据类型隐式转换**。在 MySQL 中，数据会存在隐式转换，当该字段发生转换时，索引会造成失效。
- **禁止使用相同的账号跨库操作**。
- **禁止在 INSERT 操作时只有数据值，不带字段键名**。
    <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```sql
   INSERT INTO user VALUES ('alicfeng',23);
  ```

    </td><td>

  ```sql
    INSERT INTO user (`username`,`age`) VALUES ('alicfeng',23);
  ```

    </td></tr>
    </tbody>
  </table>

- **使用 JOIN 替代子查询操作**。子查询的结果会被存储到临时表中，临时表不会存索引，因此，子查询的结果集无法使用索引，因此会消耗过多的 CPU 和 IO 资源，发生慢查询。
- **避免使用 JOIN 关联过多的表**。JOIN 多表查询比较耗时间，在业务中，最好不要超过 5 个。
- **使用 IN 代替 OR 语句**。
- **禁止使用 ORDER BY RAND() 随机排序语句**。这会把表中所有符合条件的数据装载到内存中，然后在内存中对所有数据根据随机生成的值进行排序，并且可能会对每一行都生成一个随机值，如果满足条件的数据集非常大，就会消耗大量的 CPU 和 IO 及内存资源。
- **禁止在 WHERE 语句中进行计算**。对列进行函数转换或计算时会导致无法使用索引。
    <table>
    <thead>
      <tr><th>索引失效</th><th>灵活使用（推荐）</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```sql
    SELECT `name` FROM `table` WHERE DATE(create_date)='20190308';
  ```

    </td><td>

  ```sql
    SELECT `name` FROM `table` WHERE create_date>='20190308' AND create_date<'20190309';
  ```

    </td></tr>
    </tbody>
  </table>

- **使用 UNION ALL 而不是使用 UNION**。因为 UNION 需要重复值扫描，降低效率。
- **大批量写操作应分批次处理**。进行合理地分批次处理，能够防止死锁影响业务，同时应尽量将跑批这种大操作至于凌晨操作。
- **禁止使用索引做运算**。索引会失效。
- **使用事务尽量简单化，同时控制事务执行的时间**。比如，在执行事务操作之前，将需要增、删、改的数据全部准备好，只在事务中进行数据更新，不做数据查询。
- **IN 语句参数的个数应控制在 1000 以内**。
- **LIMIT 越大，效率越低**。
    <table>
    <thead>
      <tr><th>效率较低</th><th>推荐使用</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```sql
    SELECT `username` FROM `user` LIMIT 10000,20;
  ```

    </td><td>

  ```sql
    SELECT `username` FROM `user` WHERE id>10000 LIMIT 20;
  ```

    </td></tr>
    </tbody>
  </table>

- **编写 SQL 语句必须全部为大写，每个词只允许有一个空格符**。
- **使用 EXIST|NOT EXIST 替代 IN | NOT IN**。
- **禁止使用 LIK E 添加 % 前缀进行模糊查询**。% 前置会导致索引失效。
- **禁止一条语句同时对多个表进行写操作**。
