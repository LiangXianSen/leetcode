# 178. Rank Scores

https://leetcode.com/problems/rank-scores/

# Description

SQL Schema

Write a SQL query to rank scores. If there is a tie between two scores, both should have the same ranking. Note that after a tie, the next ranking number should be the next consecutive integer value. In other words, there should be no "holes" between ranks.

```
+----+-------+
| Id | Score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
```

For example, given the above `Scores` table, your query should generate the following report (order by highest score):

```
+-------+------+
| Score | Rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+
```



# Solution

Approach: SQL `Self Join` (join itself)

Self joins are commonly used in the following areas:

- Hierarchical relationships
- Sequential relationships
- Graph data

> [ref] https://learnsql.com/blog/illustrated-guide-sql-self-join/

**Algorithm**

从两张相同的表scores分别命名为s1，s2， 同一个表自己join自己，相同字段相连得到这两个字段的乘积

```sql
SELECT 
    s1.score,
    s2.score
FROM
    Scores AS s1,
    Scores AS s2
WHERE
    s1.score = s2.score
```

>SQL 自连接，它解决了列与列之间的逻辑关系问题，准确的讲是列与列之间的层次关系

通过修改 `WHERE` 子句，s1中的score与s2中的score比较大小。在输出s1.score的前提下，有多少个s2.score大于等于它。比如当s1.score=3.65的时候，s2.score中[4.00,4.00,3.85,3.65,3.65]有5个成绩大于等于他

**MySQL**

```sql
SELECT 
    s1.score,
    s2.score
FROM
    Scores AS s1,
    Scores AS s2
WHERE
    s1.score <= s2.score
```

利用count(distinct s2.score)去重可得s1.score 3.65的Rank为 3 , 然后group by s1.id，最后根据Rank排序 desc

**MySQL**

```sql
SELECT 
    s1.score AS Score,
    COUNT(DISTINCT s2.score) AS 'Rank'
FROM
    Scores AS s1,
    Scores AS s2
WHERE
    s1.score <= s2.score
GROUP BY s1.id
ORDER BY Rank DESC
```

> NOTE: alias Rank 加上引用，避免与关键字冲突



Approach: window function `DENSE_RANK`

**MySQL**

```sql
SELECT 
    Score,
    DENSE_RANK() OVER (ORDER BY Score DESC) 'Rank'
FROM Scores;
```

> NOTE: 仅支持MySQL 8.0 以上