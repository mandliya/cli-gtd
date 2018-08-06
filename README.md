# cli-gtd
A toy "Get Things Done" todo manager built in Go and BoltDb.

This is a simple task manager inspired from [Gophercises Exercise 7 CLI Task Manager](https://gophercises.com/exercises/task) built in Go and BoltDb to learn go and have fun.

```
C:\Users\ravi\go\src\github.com\GTD>gtd list
You have no pending tasks! Well done!

C:\Users\ravi\go\src\github.com\GTD>gtd add Get Some Milk
Added "Get Some Milk" to your task list with id 1

C:\Users\ravi\go\src\github.com\GTD>gtd add Get Some Butter
Added "Get Some Butter" to your task list with id 2

C:\Users\ravi\go\src\github.com\GTD>gtd list
You have the following tasks:
1. Get Some Milk
2. Get Some Butter

C:\Users\ravi\go\src\github.com\GTD>gtd do 1
Marked 1 task as completed

C:\Users\ravi\go\src\github.com\GTD>gtd list
You have the following tasks:
2. Get Some Butter
```

