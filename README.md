# Math Expressions parser (REPL)

This is a math expressions parser. It can parse and evaluate simple math expressions.

> **Note:** This repo was created as a submission for [missinsemester.io](https://missingsemester.io) 2024 spring cohort.

## Features

- [x] Parse and evaluate math expressions
- [x] Support for addition, subtraction, multiplication, division, and parentheses
- [x] Support for functions (sin, cos, tan, log, exp, sqrt...etc)

## Usage

The code has no external dependencies other than golang itself.

build it using `go build -o interpreter main.go`

### Using Makefile

`make`

### Using docker

`docker build -t missing-semester:latest .`


If you run the binary without any arguments, it will start a REPL (Read-Eval-Print-Loop) where you can enter math expressions and get the result.

Or you can give it a single argument, which should be the math expression you want to evaluate.

`./interpreter "3 + 4 * 2 / ( 1 - 5 ) ** 2"`

or if using docker

`docker run -it missing-semester:latest "3 + 4 * 2 / ( 1 - 5 ) ** 2"`

## Explanations

3 distinct parts are implemented, and each has its purpose:

1. **Lexer**: This part is responsible for tokenizing the input string. 
    It takes the input string and returns a list of tokens. Each token
    generated has a type and a value.
2. **Parser**: This part is responsible for parsing the list of tokens, making sense of them
     and building an abstract syntax tree (AST). 
     The AST is a data structure that represents the input expression.
     A Pratt Parser implementation is used to parse Math expressions, taking into
     consideration the operator precedence, including using parentheses to change precedence.
3. **Evaluator**: This part is responsible for evaluating the AST and returning the result. 
       It uses a recursive approach to traverse the AST and evaluate the expression.
       functions like `cos`, `sin`...etc are implemented as builtin functions.


### Lexer

The lexer will traverse the input one character at a time and generate tokens.

The parser is implemented using the following rules: 

- It will ignore whitespace characters.
- It will generate tokens for numbers, identifiers, and operators.
- It will generate tokens for parentheses, commas, and semicolons.
- It will generate an `EOF` token when it reaches the end of the input.
- It will generate an `ILLEGAL` token if it encounters an unknown character.

Tokens are predefined [here](./token/token.go).

> **Note:** Calculating powers is both supported through the `**` operator as well as by calling `pow(a, b)`.

```golang

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// user defined identifiers + literals
	IDENT  = "IDENT"
	NUMBER = "NUMBER"

	// Delimiters
	COMMA     = ","

	LPAREN = "("
	RPAREN = ")"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	// Exponentiation
	EXPONENT = "**"
)
```

When reading numbers, the lexer will convert all numbers as floats internally, this 
is to make dealing with numbers simpler. Because in our use case, integers and floats
can be treated the same way when evaluating a math expression.


### Parser

The parser is what makes sense of the tokens generated. It iterate the tokens
and build an AST as it goes. A Pratt Parser is used to parse the tokens.

The parser is implemented in [parser/parser.go](./parser/parser.go).

The AST is defined in [ast/ast.go](./ast/ast.go).

The top-level node of the AST is the `Program` node. It contains a list of statements.

A math expression is represented by the `ExpressionStatement` node. It contains an `Expression` node.

Math function calls are represented as `CallExpression` nodes. They contain the function name and a list of arguments,
which are expressions themselves. This allows us to also evaluate `cos(3 + 4)` for example.


`-` and `+` operators are also implemented as prefix operators, meaning `-5` `+5`, `--5` are all valid expressions.


### Evaluator

The evaluator is responsible for traversing the AST and evaluating the expression.

The evaluator is implemented in [evaluator/evaluator.go](./evaluator/evaluator.go), It is
a recursive approach to traverse the AST and evaluate the expression.

The evaluator returns objects defined in [object/object.go](./object/object.go).

We have three objects defined:

- `Number`: Represents a number.
- `Function`: Represents a builtin function like `cos`, `sin`, `tan`...etc.
- `Error`: Represents an error.

builtin functions are implemented in [evaluator/builtin.go](./evaluator/builtin.go).

the Builtin object defines a function that takes a list of arguments and returns an object after validating the arguments.

The bulitin function implemented are `cos`, `acos`, `sin`, `asin`, `tan`, `atan`, `log`, `exp`, `abs`, `sqrt`, `pow`.

the `math` package from the standard library is used to implement these functions.

### Error Handling

Both the parser and the evaluator can return errors. 
The parser returns errors when it encounters an unexpected token or when it encounters a syntax error. 
The evaluator returns errors when it encounters an error while evaluating the AST.

Parsing and Evaluation is stopped as soon as an error occurs, then the errors are printed to the user.
