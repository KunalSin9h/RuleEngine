## RuleEngine

![RuleEngine demonstration image](https://i.imgur.com/nm7KkQ1.png)

`RuleEngine` is a dynamic `rule` evaluation application, implemented in _GoLang_. 
This helps in mitigating drawbacks in more conventional `If-Else` rule evaluation, which are hard coded in application.
This hard coding make difficult to _Add_, _Update_ and _Delete_ rules.

`RuleEngine` will enable us define, parse, and evaluate rules flexibly.

### Architecture

#### Primitives

This project uses `Abstract Syntax Tree (AST)` to store rules. This _AST_ is created using `Rule String` with a `Parser`.

![Parser demonstration image](https://i.imgur.com/UH7L7h2.png)

example of _Rule String_:

```txt
((age > 30 AND department = 'Sales') OR (age < 25 AND
department = 'Marketing')) AND (salary > 50000 OR experience > 5)
```

Source:

- `internal/ast` holds the implementation of `AST`
- `internal/parser` holds the implementation of `Parser`

They are jointly tested in `tests/parse_ast_test.go`.

### Application (API + Postgres + UI)

![Application Architecture demonstration image](https://i.imgur.com/nh1lKvx.png)
