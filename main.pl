% <assignStmt> --> <id> = <expr>
% <expr> --> <expr> <op1> <expr1> | <expr1>
% <expr1> --> <expr1> <op2> <expr0> | <expr0>
% <expr0> --> <id> | <entero> | <numDecimal> | (<expr>)
% <op2> --> * | /
% <op1> --> + | -

assignStmt --> id, [=] , expr.
expr --> expr1 ; expr, op1, expr1.
expr1 --> expr0 ; expr1, op2, expr0.  
expr0 --> id ; entero ; numDecimal ; ['('] , expr , [')'].
op2 --> [*] ; [/].
op1 --> [+] ; [-].

id --> [a].
entero --> [2].
entero --> [3].
numDecimal --> [5].
