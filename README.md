# Coding Challenges

Here are summaries of the challenges:

## 1. Narcissistic Number

A narcissistic number is a number that returns the sum of its own digits, each raised to the power of the number of digits in a given base. (Assume there won’t be any decimal)

### Example:

- 1634 = 4 Digits
  - \(1^4 + 6^4 + 3^4 + 4^4 = 1 + 1296 + 81 + 256 = 1634\)
- 153 = 3 Digits
  - \(1^3 + 5^3 + 3^3 = 3 + 125 + 27 = 153\)

### Challenge

Create a function that returns true or false depending on whether the given number in the parameter is narcissistic or not.

### Example:

- `narcissistic(153)` // this will return true
- `narcissistic(111)` // this will return false

## 2. Parity Outlier

Given an array of integers (minimum length of 3), the array either entirely contains whole odd integers with one outlier even integer or the whole even integers with one outlier odd integer.

### Challenge

Write a method that takes an array as an argument and returns the outlier.

### Example:

- `[2, 4, 0, 100, 4, 11, 2602, 36]` Should return: 11 (the only odd number)
- `[160, 3, 1719, 19, 11, 13, -21]` Should return: 160 (the only even number)
- `[11, 13, 15, 19, 9, 13, -21]` Should return: false (all odd integer, no outlier)

## 3. Needle in the Haystack

### Challenge

Write a function that takes two arguments: the first one takes an array of strings (as a haystack) and the second one is a single string (as the needle). This function should return the index of the needle’s position. Rules: Using array_search() function in PHP is prohibited.

### Example:

- `findNeedle(["red", "blue", "yellow", "black", "grey"], "blue")` This function should return 1 as the index of the needle (blue)

## 4. The Blue Ocean Reverse

Blue Ocean Strategy is very famous in marketing strategy; it tries the business to differ from other competitors with a new product/business model. In this case, we'll try the reverse of it!

### Challenge

Create a function that takes two arguments, both should be arrays of integers. This function should subtract one list from another and return the result. It should remove all values from the first list which are present in the second list.

### Example:

- `blueOcean([1, 2, 3, 4, 6, 10], [1])` Should return [2, 3, 4, 6, 10] because 1 is present in the second list
- `blueOcean([1, 5, 5, 5, 5, 3], [5])` Should return [1, 3] because 5 is present in the second list
