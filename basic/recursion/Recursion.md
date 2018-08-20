# Recursion

Recursion occurs when a method calls itself. It helps to make some problems easier to understand

## Two Features (Base case and recursive case)

1. The recursion must eventually stop. This is the base case when the function doesn't call itself again. 
2. Each time the method executes, it reduces the current problem to a smaller instance of the same problem and then calls itself to solve the smaller problem. This is the recursive case when the function calls itself.

## Issues

The recursion may make the solution clearer to developers, but there is no performance benefit of using it. In fact, it may bring performance issues to the solution, such as the deep call stack may use a lot of memory, or even cause the stack overflow problem.

To address the performance problems caused by the recursion, we can rewrite the code with a loop. There are 3 ways to remove the recursion.

1. Use tail recursion removal. Tail recursion occurs when the last thing a singly recursive algorithm does before returning is call itself.
2. Store intermediate values.
3. Mimic what the program does when it performs recursion. Push and pop the state to stack.

## Some problems that can be solved with recusion

- Tower of Hanoi