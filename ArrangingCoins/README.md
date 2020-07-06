```markdown
You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of full staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1:

n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.
Example 2:

n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

Because the 4th row is incomplete, we return 3.
```

# Solution

The sum of Arithmetic sequence is: S = (n*(a_1 + a_n))/2
So, n * (a_1 + a_n) = 2 * S
as a_n = n, a_1 = 1, then n * (1 + n) = 2 * S.

=> n^2 + n - 2 * S = 0
n = (-1 + sqrt(1^2 - 4 * (- 2 * S) * 1))/ 2 * 1
n = (-1 + sqrt(1 + 8*S)) / 2
n = (-1/2 + sqrt(1/4 + 2*S))
In this problem:
S = n
n = number of rows

n should be the maximum number which smaller than the 
