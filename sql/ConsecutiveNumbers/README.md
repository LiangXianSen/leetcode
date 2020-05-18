# 180. Consecutive Numbers

https://leetcode.com/problems/consecutive-numbers/

# Description

SQL Schema

Write a SQL query to find all numbers that appear at least three times consecutively.

```
+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
```

For example, given the above `Logs` table, `1` is the only number that appears consecutively for at least three times.

```
+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
```



# Solution

**Algorithm**

遍历`Logs` 表，Num相同则`@count` +1, 容易理解

**MySQL**

```sql
select distinct Num as ConsecutiveNums
from (
  select Num, 
    case 
      when @prev = Num then @count := @count + 1
      when (@prev := Num) is not null then @count := 1
    end as CNT
  from Logs, (select @prev := null,@count := null) as t
) as temp
where temp.CNT >= 3
```

> NOTE: 仔细审题，连续出现次数！ （相邻的行）



**Official Solution**

Consecutive appearing means the Id of the Num are next to each others. Since this problem asks for numbers appearing at least three times consecutively, we can use 3 aliases for this table **Logs**, and then check whether 3 consecutive numbers are all the same.

```sql
SELECT *
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1
    AND l1.Num = l2.Num
    AND l2.Num = l3.Num
;
```

| Id   | Num  | Id   | Num  | Id   | Num  |
| ---- | ---- | ---- | ---- | ---- | ---- |
| 1    | 1    | 2    | 1    | 3    | 1    |

> Note: The first two columns are from l1, then the next two are from l2, and the last two are from l3.

Then we can select any *Num* column from the above table to get the target data. However, we need to add a keyword `DISTINCT` because it will display a duplicated number if one number appears more than 3 times consecutively.

**MySQL**

```sql
SELECT DISTINCT
    l1.Num AS ConsecutiveNums
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1
    AND l1.Num = l2.Num
    AND l2.Num = l3.Num
;
```