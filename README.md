# Closure-Test
Understanding Closure in Go(lang) better

From a course by Todd McLeod @GoesToEleven

Wanted to understand how the variables were handled.
Very interesting: 

The anonymous func that gets returned gets created only once, while its state gets created anew every time the func is assigned to a new variable.
