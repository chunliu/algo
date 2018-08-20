# Big O Notation

Big O notation is the notation that tells you how fast an algorithm is. It uses a function to describe how the algorithm's worst-case performance relates to the problem size as the size grows very large. The function is written within parentheses after a capital letter O. For example, O(n^2).

## 5 Rules

There are five basic rules for calculating an algorithm's Big O notation:

1. If an algorithm performs a certain sequence of steps f(N) times for a mathematical function f, it takes O(f(N)) steps.
2. If an algorithm performs an operation that takes O(f(N)) steps and then performs a second operation that takes O(g(N)) steps for functions f and g, the algorithm's total performance is O(f(N) + g(N)).
3. If an algorithm takes O(f(N) + g(N)) and the function f(N) is greater than g(N) for large N, the algorithm's performance can be simplified to O(f(N)).
4. If an algorithm performs an operation that takes O(f(N)) steps, and for every step in that operation it performs another O(g(N)) steps, the algorithm's total performance is O(f(N) × g(N)).
5. Ignore constant multiples. If C is a constant, O(C × f(N)) is the same as O(f(N)), and O(f(C × N)) is the same as O(f(N)).

## Some common Big O runtime

Runtime   | Example
-----------|----------
O(log n)   | Binary search
O(n)       | Simple search
O(n*log n) | Quick sort
O(n^2)     | Selection sort
O(n!)      | Traveling salesperson

Input | log n | n | n^2 | 2^n | n!
------|-------|---|-----|-----|-------
1     | 0     | 1 | 1   | 2   | 1
5     | 2.32  | 5 | 25  | 32  | 625
10    | 3.32  | 10 | 100 | 1024 | 1.0 x 10^9
15    | 3.90  | 15 | 225 | 3.3 x 10^4 | 2.9 x 10^16