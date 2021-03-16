# antlr-compiler

this repo was created for academic use only

## Comands

Build parser for .g4 grammers

```bash
antlr -Dlanguage=Go -o parser <name-of-file>.g4
```

Run exercice
```bash
go run cmd/path-to-exercice/main.go
```

## Explicações

### Expr

dada a granática Expr definida no arquivo Expr.g4

```
grammar Expr;		
// rules
prog:	(expr NEWLINE)* ;
expr:	expr ('*'|'/') expr
    |	expr ('+'|'-') expr
    |	INT
    |	'(' expr ')'
    ;
//Tokens
NEWLINE : [\r\n]+ ;
INT     : [0-9]+ ;
SUM: '+';
SUB: '-';
MULTI: '*';
DIV: '/';
// define que quando encontrar o withespace o mesmo será ignorado
WHITESPACE: [ \t]+ -> skip;

```

trata-se de uma gramatica que consegue identificar o padrão de uma operação matemática com um numero a esquerda, um sinal de operação e um número a direita, sso pode ser visto no arquivo definido pelos tokens que permitem uma quebra de lnha, um espaço em branco, sinais de operações e valores inteiros

a regra de escrita que define que temos um número, um sinal matemático, uma outro numero

