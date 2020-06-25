# Problem

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.


# Solution

Use 2 ponters. Left pointer starts at index 0, and slides to the right.
Right pointer starts at (length - 1), and slides to the left, until meets the left pointer.

At each step, we calculate the total amount of water can be contained.

