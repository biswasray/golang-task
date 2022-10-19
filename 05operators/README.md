## Operators
An operator is a symbol that tells the compiler to perform specific mathematical or logical manipulations.
- [Arithmetic Operators](01arithmatic/main.go)
- [Relational Operators](02relational/main.go)
- [Logical Operators](03logical/main.go)
- [Bitwise Operators](04bitwise/main.go)
- [Assignment Operators](05assignment/main.go)
- [Miscellaneous Operators](06miscellaneous/main.go)

### Arithmetic
| Operator  | Description                                                       |	Example         |
|-----------|-------------------------------------------------------------------|-------------------|
| +         |Adds two operands	                                                | A + B gives 30    |
| -	        |Subtracts second operand from the first	                        | A - B gives -10   |
| *	        |Multiplies both operands	                                        | A * B gives 200   |
| /	        |Divides the numerator by the denominator.	                        | B / A gives 2     |
| %	        |Modulus operator; gives the remainder after an integer division.   | B % A gives 0     |
| ++	    |Increment operator. It increases the integer value by one.         | A++ gives 11      |
| --	    |Decrement operator. It decreases the integer value by one.	        | A-- gives 9       |

### Relational
| Operator  | Description                                                                                                                                   |	Example             |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------------------|-----------------------|
| ==	    | It checks if the values of two operands are equal or not; if yes, the condition becomes true.	                                                | (A == B) is not true. |
| !=	    | It checks if the values of two operands are equal or not; if the values are not equal, then the condition becomes true.	                    | (A != B) is true.     |
| >	        | It checks if the value of left operand is greater than the value of right operand; if yes, the condition becomes true.	                    | (A > B) is not true.  |
| <	        | It checks if the value of left operand is less than the value of the right operand; if yes, the condition becomes true.	                    | (A < B) is true.      |
| >=	    | It checks if the value of the left operand is greater than or equal to the value of the right operand; if yes, the condition becomes true.	| (A >= B) is not true. |
| <=	    | It checks if the value of left operand is less than or equal to the value of right operand; if yes, the condition becomes true.	            | (A <= B) is true.     |

### Logical
| Operator      | Description                                                                                                                                       |	Example                 |
|---------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------|
| &&	        | Called Logical AND operator. If both the operands are non-zero, then condition becomes true.	                                                    | (A && B) is false.        |
| &#124&#124	| Called Logical OR Operator. If any of the two operands is non-zero, then condition becomes true.	                                                | (A &#124&#124 B) is true. |
| !	            | Called Logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true then Logical NOT operator will make false.	| !(A && B) is true.        |

### Bitwise
| Operator  | Description                                                                                                               |	Example                                     |
|-----------|---------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| &	        | Binary AND Operator copies a bit to the result if it exists in both operands.	                                            | (A & B) will give 12, which is 0000 1100      |
| &#124 	| Binary OR Operator copies a bit if it exists in either operand.	                                                        | (A &#124 B) will give 61, which is 0011 1101  |
| ^	        | Binary XOR Operator copies the bit if it is set in one operand but not both.	                                            | (A ^ B) will give 49, which is 0011 0001      |
| <<	    | Binary Left Shift Operator. The left operands value is moved left by the number of bits specified by the right operand.	| A << 2 will give 240 which is 1111 0000       |
| >>	    | Binary Right Shift Operator. The left operands value is moved right by the number of bits specified by the right operand. | A >> 2 will give 15 which is 0000 1111        |

### Assignment
| Operator  | Description                                                                                                               |	Example                                     |
|-----------|---------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| =	        | Simple assignment operator, Assigns values from right side operands to left side operand	                                | C = A + B will assign value of A + B into C   |
| +=	    | Add AND assignment operator, It adds right operand to the left operand and assign the result to left operand	            | C += A is equivalent to C = C + A             |
| -=	    | Subtract AND assignment operator, It subtracts right operand from the left operand and assign the result to left operand	| C -= A is equivalent to C = C - A             |
| *=	    | Multiply AND assignment operator, It multiplies right operand with the left operand and assign the result to left operand | C *= A is equivalent to C = C * A             |
| /=	    | Divide AND assignment operator, It divides left operand with the right operand and assign the result to left operand	    | C /= A is equivalent to C = C / A             |
| %=	    | Modulus AND assignment operator, It takes modulus using two operands and assign the result to left operand	            | C %= A is equivalent to C = C % A             |
| <<=	    | Left shift AND assignment operator	                                                                                    | C <<= 2 is same as C = C << 2                 |
| >>=	    | Right shift AND assignment operator	                                                                                    | C >>= 2 is same as C = C >> 2                 |
| &=	    | Bitwise AND assignment operator	                                                                                        | C &= 2 is same as C = C & 2                   |
| ^=	    | bitwise exclusive OR and assignment operator	                                                                            | C ^= 2 is same as C = C ^ 2                   |
| &#124=	| bitwise inclusive OR and assignment operator	                                                                            | C &#124= 2 is same as C = C &#124 2           |

## Miscellaneous
| Operator  | Description                           |	Example                                     |
|-----------|---------------------------------------|-----------------------------------------------|
| &     	| Returns the address of a variable.	| &a; provides actual address of the variable.  |
| *	        | Pointer to a variable.	            | *a; provides pointer to a variable.           |