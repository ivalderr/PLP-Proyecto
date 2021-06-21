%% Author: Javier Tibau

list([H|T],H,T).

%% Tokenizer code
%Created by Bruno Dufour, Fall 2005
%
% Append
append([ ],A,A).
append([A|B],C,[A|D]) :- append(B,C,D).


gather(Chars) --> [C],  {alphaNumeric(C)}, gather(Rest), {Chars=[C|Rest]}.

gather([]) --> {true}.
alphaNumeric(C):- 96<C,C<123;
                  64<C, C<91;
                  47<C, C<58.

% - Floats ---------------------------------------------------------------------
digit(D) --> [D], {47 < D, D < 58}.
nzdigit(D) --> [D], {48 < D, D < 58}.

floatlit(F) -->
        nzdigit(D0),
        digits(D1),
        ".",
        nedigits(D2),
        {append([D0|D1], [46], T), append(T, D2, D), name(F, D)}.

nedigits([D|T]) -->
        digit(D), !,
        digits(T).

digits(L) --> nedigits(L).
digits([]) --> [].
% ------------------------------------------------------------------------------

% - Strings --------------------------------------------------------------------

quote('"').

gatherString(Chars) --> [C], {C=\=34}, gatherString(Rest), {Chars=[C|Rest]}.
gatherString([]) --> {true}.

stringlit(S) --> "\"", gatherString(Chars), "\"", {string_to_list(S,Chars)}.

% ------------------------------------------------------------------------------

% Tokeinze comparison operators
tokenize(Z) --> "==", tokenize(Y), {Z = [== | Y]}.
tokenize(Z) --> ">=", tokenize(Y), {Z = [>= | Y]}.
tokenize(Z) --> "<=", tokenize(Y), {Z = [<= | Y]}.
tokenize(Z) --> "<>", tokenize(Y), {Z = [<> | Y]}.
tokenize(Z) --> ">", tokenize(Y), {Z = [> | Y]}.
tokenize(Z) --> "<", tokenize(Y), {Z = [< | Y]}.

% Tokenize float
tokenize(Result) --> floatlit(F), tokenize(Rest), {Result=[F|Rest]}.
% Tokenize string
tokenize(Result) --> stringlit(S), tokenize(Rest), {Result=[S|Rest]}.
% Tokenize id / int
tokenize(Result) --> gather(Chars),{\+ Chars =[]},tokenize(RestResult),
                    {name(N,Chars), Result=[N|RestResult]}.
% Discard whitespace
tokenize(R)-->[C],{C<33},tokenize(R).
% Tokenize special character
tokenize([N|R]) --> [C],{C>32},
                        {name(N,[C])},tokenize(R).
tokenize([])-->[].

% prefix/2
prefix(P,L) :- append(P,_,L).
% suffix/2
suffix(S,L) :- append(_,S,L).
sublist(SubL,L) :- suffix(S,L), prefix(SubL,S).

% sublist(['(',T],L) :-

case(['('|T],[H|L],[H|T]) :- L=[], H==')'.
case(['('|T],[H|L],[S|R]) :- case(T,L,[S|R]).

%case([H|T],[H|L],[H|T]) :- L=[], H=='*'.
%case([H|T],[H|L],[H|T]) :- L=[], H=='/'.
%case([H|T],[H|L],[H|T]) :- L=[], H=='+'.
%case([H|T],[H|L],[H|T]) :- L=[], H=='-'.
%case([H|T],[H|L],[S|R]) :- case(T,L,[S|R]).

assignStmt --> id, [=] , expr.
expr --> expr1 ; (expr, op1, expr1).
expr1 --> expr0 ; (expr1, op2, expr0).
expr0 --> id ; entero ; numDecimal ; (['('] , expr , [')']).
op2 --> [*] ; [/].
op1 --> [+] ; [-].

id --> [a].
entero --> [2].
entero --> [3].
numDecimal --> [5].

%% <program> --> <type-decl-stmts> ; <stmts>
program(TSBefore, TSAfter, (declarations(DeclarationTree),statements(StatementList))) :-
     typeDeclarationStatementList(TSBefore, TSAfterDeclaration, DeclarationTree),
     list(TSAfterDeclaration,';',T),
     statementList(T, TSAfter, StatementList).

parseTree(FileName,RT):-
    open(FileName, 'read', InputStream),
    read_stream_to_codes(InputStream, ProgramString),
    close(InputStream),
    phrase(tokenize(TSBefore), ProgramString),
    write(TSBefore),
    assignStmt(TSBefore,[]).
    % program(TSBefore, [], RT).
