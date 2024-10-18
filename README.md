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

> The parser is highly coupled with this format.

```txt
((age > 30 AND department = 'Sales') OR (age < 25 AND
department = 'Marketing')) AND (salary > 50000 OR experience > 5)
```

Source:

- [internal/ast](internal/ast) holds the implementation of `AST`
- [internal/parser](internal/parser) holds the implementation of `Parser`

They are tested extensively in [internal/ast/ast_test.go](internal/ast/ast_test.go) and [internal/parser/parser_test.go](internal/parser/parser_test.go).

### Application (API + Postgres + UI)

![Application Architecture demonstration image](https://i.imgur.com/nh1lKvx.png)

## Setup

Dependencies:

- Docker (optional)
- GNU Make (optional)

Clone the repository, and go to root of project.

### Docker Setup (recommended)

This will start DB and APP 

```bash
docker compose up
```

Go to [http://localhost:5000](http://localhost:5000)

### Manual Setup

1. Run postgres (using docker). 

> if you are on windows or macOS (i think sh works on macOS)
> Execute the docker command for Postgres manually (see the [rundb.sh](rundb.sh) file)

```bash
sudo bash ./rundb.sh
```

2. Download Dependencies

```bash
go mod download
```

3. Run application
 
```bash
make run

# or 
# go run cmd/api/*.go
```

Go to [http://localhost:5000](http://localhost:5000)

#### Environment Variables

| ENV      | DEFAULT VALUE                                   | USE CASE                   |
|----------|-------------------------------------------------|----------------------------|
| PORT     | 5000                                            | application port           |
| HOST     | "0.0.0.0"                                       | application host           |
| POSTGRES | "postgresql://admin:admin@localhost:5432/rules" | postgres connection string |

Here, **we use** `rules` database, so make sure your postgres instance has this.

## Stack

- UI : with React and Vite
- API and Backend : with GoLang
- DB is PostgreSQL

## UI Demo

![UI Demo](https://i.imgur.com/gylOn7I.jpeg)

## Results

#### API Design Code:

1. `create_rule(rule_string)` is in [internal/parser/parser.go](internal/parser/parser.go)
2. `combine_rules(rules)` is also in [internal/parser/parser.go](internal/parser/parser.go)
3. `evaluate_rule(JSON data)` is in [internal/ast/ast.go](internal/ast/ast.go)

#### Test Case result

Run Test Case

```bash
make test

# or 
# go test ./...
```

1. First test case is tested in [tests/parse_ast_test.go](tests/parse_ast_test.go) in `TestParseAST` function
2. (and 4th) Test case is tested in [tests/parse_ast_test.go](tests/parse_ast_test.go) in `TestParser_CombineRules_AND` and `TestParser_CombineRules_OR`
3. Third test case is tested in [internal/ast/ast_test.go](internal/ast/ast_test.go) in multiple function. (all cases covered)

#### DB Schema and Queries

Database schema is in [schema/schema.sql](schema/schema.sql).

All the SQL queries used are in [schema/query.sql](schema/query.sql). I have used `sqlc` in golang as ORM.

## Decision Choice reasons

- Use of `PostgresSQL` for DB.

I have used postgres for DB, specifically for storing AST in the JSONB (Json in Binary) form.

I have decided to store the AST in DB after Parsing since, it will reduce the PARSING COST on every execution.


- Use of `sqlc` for SQL ORM.

I have used `sqlc` it an ORM for golang, it generated code with sql queries.

- Use of `React` for UI.

I have used React in UI. so that I can quickly create a frontend. It is compiled in [ui/dist](ui/dist) folder, which go
uses to serve statically.

- Use of `net/http` for creating API.

In [cmd/api](cmd/api), for creating API endpoints, I have used go's standard `net/http` package. It's pretty powerful, but
can cause some repetitive code (epically for extracting and send json data). I could have used some abstractions like `Echo` or `Fiber`. 

