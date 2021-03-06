# 183. Customers Who Never Order

https://leetcode.com/problems/customers-who-never-order/

# Description

SQL Schema

Suppose that a website contains two tables, the `Customers` table and the `Orders` table. Write a SQL query to find all customers who never order anything.

Table: `Customers`.

```
+----+-------+
| Id | Name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
```

Table: `Orders`.

```
+----+------------+
| Id | CustomerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
```

Using the above tables as example, return the following:

```
+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+
```



# Solution

**Algorithm**

主要掌握`Left Join` ， 根据左边表做join，右边表没有的为NULL

**MySQL**

```sql
select c.Name as Customers from Customers c left join Orders o
on c.Id = o.CustomerId
where o.Id is NULL
```


