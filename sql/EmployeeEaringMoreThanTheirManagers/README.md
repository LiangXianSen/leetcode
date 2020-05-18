# 181. Employees Earning More Than Their Managers

https://leetcode.com/problems/employees-earning-more-than-their-managers/

# Description

SQL Schema

The `Employee` table holds all employees including their managers. Every employee has an Id, and there is also a column for the manager Id.

```
+----+-------+--------+-----------+
| Id | Name  | Salary | ManagerId |
+----+-------+--------+-----------+
| 1  | Joe   | 70000  | 3         |
| 2  | Henry | 80000  | 4         |
| 3  | Sam   | 60000  | NULL      |
| 4  | Max   | 90000  | NULL      |
+----+-------+--------+-----------+
```

Given the `Employee` table, write a SQL query that finds out employees who earn more than their managers. For the above table, Joe is the only employee who earns more than his manager.

```
+----------+
| Employee |
+----------+
| Joe      |
+----------+
```



# Solution

**Algorithm**

`Self Join` 便可轻松解决，不做多的解释了，前面的题多次应用到了。

**MySQL**

```sql
SELECT
    e.Name AS 'Employee'
FROM
    Employee AS e,
    Employee AS m
WHERE
    e.Id = m.ManagerId AND e.Salary > m.Salary
;
```


