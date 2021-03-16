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
