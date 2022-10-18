## Variable
A variable is nothing but a name given to a storage area that the programs can manipulate. Each variable in Go has a specific type, which determines the size and layout of the variable's memory, the range of values that can be stored within that memory, and the set of operations that can be applied to the variable.

The name of a variable can be composed of letters, digits, and the underscore character. It must begin with either a letter or an underscore. Upper and lowercase letters are distinct because Go is case-sensitive.

Go programming language also allows to define various other types of variables such as Enumeration, Pointer, Array, Structure, and Union

## Variable Definition in Go
### Syntax
```go
var variable_list optional_data_type;
```
### Example
```go
var  i, j, k int;
var  c, ch byte;
var  f, salary float32;
d =  42;
```
### Type Declaration in Go
1. [Static Type Declaration](01variable-declaration/main.go)
2. [Dynamic / Mixed Type Declaration](02type-inference/main.go)

#### Default Values of variable
- [zero value](03zero-values/main.go)

### The lvalues and the rvalues in Go
There are two kinds of expressions in Go −

- **lvalue** − Expressions that refer to a memory location is called "lvalue" expression. An lvalue may appear as either the left-hand or right-hand side of an assignment.

- **rvalue** − The term rvalue refers to a data value that is stored at some address in memory. An rvalue is an expression that cannot have a value assigned to it which means an rvalue may appear on the right- but not left-hand side of an assignment.

The following statement is valid −
```go
x = 20.0
```
The following statement is not valid. It would generate compile-time error −
```go
10 = 20
```