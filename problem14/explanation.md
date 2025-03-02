# Problem 14 explanation

Very simple. There's a function called collatz which computes the collatz sequence of a number recursivley and returns the length the sequence "times". 
The solve function takes in a parameter "target" and returns the number with the longest collatz seuquence under "target". It simply loops over every number under "target" computes it's collatz sequence and keeps track of the longest one.