grammar sqlQuery;

sqlExpr: SELECT fields FROM tNames SEMI?;

fields: START | ID (COMM ID)*;

SELECT: S E L E C T;
START: '*';
FROM: F R O M;

tNames: tName tNameAlias? (COMM tName tNameAlias?)*;
tName: ID;
tNameAlias: ID;
SEMI:  ';';
COMM:  ',';
ID: [a-zA-Z_] [a-zA-Z0-9_]*;
// LEXER HERE
fragment A                  : [aA] ;
fragment B                  : [bB] ;
fragment C                  : [cC] ;
fragment D                  : [dD] ;
fragment E                  : [eE] ;
fragment F                  : [fF] ;
fragment G                  : [gG] ;
fragment H                  : [hH] ;
fragment I                  : [iI] ;
fragment J                  : [jJ] ;
fragment K                  : [kK] ;
fragment L                  : [lL] ;
fragment M                  : [mM] ;
fragment N                  : [nN] ;
fragment O                  : [oO] ;
fragment P                  : [pP] ;
fragment Q                  : [qQ] ;
fragment R                  : [rR] ;
fragment S                  : [sS] ;
fragment T                  : [tT] ;
fragment U                  : [uU] ;
fragment V                  : [vV] ;
fragment W                  : [wW] ;
fragment X                  : [xX] ;
fragment Y                  : [yY] ;
fragment Z                  : [zZ] ;
fragment DEC_DIGIT          : [0-9];
fragment DEC_DIGITS         : DEC_DIGIT+;

// IGNORED TOKENS
SPACE                       : [ \t\r\n]+    -> skip;