# 182. Duplicate Emails

https://leetcode.com/problems/duplicate-emails/

# Description

SQL Schema

Write a SQL query to find all duplicate emails in a table named `Person`.

```
+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
```

For example, your query should return the following for the above table:

```
+---------+
| Email   |
+---------+
| a@b.com |
+---------+
```



# Solution

**Algorithm**

没什么可说的`GROUP BY` & `HAVING` 的用法；

**MySQL**

```sql
select Email
from Person
group by Email
having count(Email) > 1;
```


