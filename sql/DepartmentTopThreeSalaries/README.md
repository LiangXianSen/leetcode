# 185. Department Top Three Salaries

https://leetcode.com/problems/department-top-three-salaries/

# Description

SQL Schema

The `Employee` table holds all employees. Every employee has an Id, and there is also a column for the department Id.

```
+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 85000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
| 5  | Janet | 69000  | 1            |
| 6  | Randy | 85000  | 1            |
| 7  | Will  | 70000  | 1            |
+----+-------+--------+--------------+
```

The `Department` table holds all departments of the company.

```
+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+
```

Write a SQL query to find employees who earn the top three salaries in each of the department. For the above tables, your SQL query should return the following rows (order of rows does not matter).

```
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| IT         | Randy    | 85000  |
| IT         | Joe      | 85000  |
| IT         | Will     | 70000  |
| Sales      | Henry    | 80000  |
| Sales      | Sam      | 60000  |
+------------+----------+--------+
```

**Explanation:**

In IT department, Max earns the highest salary, both Randy and Joe earn the second highest salary, and Will earns the third highest salary. There are only two employees in the Sales department, Henry earns the highest salary while Sam earns the second highest salary.



# Solution

**Algorithm**

先用`Self Join` 根据部门排序， 参考[Rank Scores](RankScores)，然后 `join` Department 表输出结果。

```sql
SELECT d.Name Department, e.Name Employee, e.Salary
FROM Employee e, Department d, (
    SELECT e1.Id, COUNT(DISTINCT e2.Salary) AS 'rank'
    FROM Employee e1, Employee e2
    WHERE e1.DepartmentId = e2.DepartmentId AND e1.Salary <= e2.Salary
    GROUP BY e1.Id
		) AS r
WHERE r.rank <= 3 AND e.DepartmentId = d.Id AND r.Id = e.Id
ORDER BY Department, Salary DESC;
```



**Official Solution**

Approach: Using `JOIN` and sub-query [Accepted]

**Algorithm**

A top 3 salary in this company means there is no more than 3 salary bigger than itself in the company.

```sql
select e1.Name as 'Employee', e1.Salary
from Employee e1
where 3 >
(
    select count(distinct e2.Salary)
    from Employee e2
    where e2.Salary > e1.Salary
)
;
```

In this code, we count the salary number of which is bigger than e1.Salary. So the output is as below for the sample data.

```
| Employee | Salary |
|----------|--------|
| Henry    | 80000  |
| Max      | 90000  |
| Randy    | 85000  |
```

Then, we need to join the **Employee** table with **Department** in order to retrieve the department information.

**MySQL**

```sql
SELECT
    d.Name AS 'Department', e1.Name AS 'Employee', e1.Salary
FROM
    Employee e1
        JOIN
    Department d ON e1.DepartmentId = d.Id
WHERE
    3 > (SELECT
            COUNT(DISTINCT e2.Salary)
        FROM
            Employee e2
        WHERE
            e2.Salary > e1.Salary
                AND e1.DepartmentId = e2.DepartmentId
        )
;
```

```
| Department | Employee | Salary |
|------------|----------|--------|
| IT         | Joe      | 70000  |
| Sales      | Henry    | 80000  |
| Sales      | Sam      | 60000  |
| IT         | Max      | 90000  |
| IT         | Randy    | 85000  |
```