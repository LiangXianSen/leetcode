# 196. Delete Duplicate Emails

https://leetcode.com/problems/delete-duplicate-emails/

# Description

编写一个 SQL 查询，来删除 `Person` 表中所有重复的电子邮箱，重复的邮箱里只保留 **Id** *最小* 的那个。

```
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
Id 是这个表的主键。
```

例如，在运行你的查询语句之后，上面的 `Person` 表应返回以下几行:

```
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
```

 

**提示：**

- 执行 SQL 之后，输出是整个 `Person` 表。
- 使用 `delete` 语句。



# Solution

**MySQL**

```mysql
DELETE from Person 
Where Id not in 
(
    select t.id from   
    (
        Select MIN(Id) as id
        From Person 
        Group by Email
    ) t
)
```

