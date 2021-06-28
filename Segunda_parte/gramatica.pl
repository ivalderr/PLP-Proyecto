% TODO: void y otras palabras reservadas
% TODO: manejar parentesis

function -->
  [void], variable, ['('], ([void];[];argf), [')'].
function -->
  datatype, variable, ['('], ([void];[];argf), [')'].

datatype --> [char];[int];[float];[double].

variable -->
  [V],
  {atom_codes(V,[H|T])},
  {(letter(H,[H|_],_) ; H==95)},
  {maplist(alphaNumeric,T)}.

letter(L) --> [L], {96<L,L<123 ; 64<L, L<91}.

% argf --> [].
argf --> datatype, variable.
argf --> datatype, variable, [','] , argf.

statement -->
  assignment_statement ;
  loop_statement ;
  conditional_statement ;
  expression [';'] ;
  [return] expression';'
  [return] ';'

expression -->
  variable ; conditional.
expression -->
  (vairiable ; conditional), operator, expression.

conditional -->
  ([];[!]), variable.
conditional -->
  ([];[!]), variable, relational_operator, conditional.
conditional -->
  ([];[!]), variable, logical_operator, conditional.

operator --> [+];[-];[*];[/];['%'].
logical_operator --> [&&];['|'];[!].
relational_operator(OP) :-
  OP=='==' ; OP=='!=' ;
  OP=='>'  ; OP=='<'  ;
  OP=='=>' ; OP=='=<' .
  
% ===================================================================
args([')'|[]]) :- true.

args([H|T]) :-
  variable(H);
  function_call_statement(H);
  args(T).

function_call_statement([H,'(',V|T]) :-
  atom_codes(H,X),
  variable(X),
  (varible(V);V==[]),
  args(T).

value --> natural;integer;float;string;character.

% function -->
%   datatype, variable, ['('], argf, [')'].

% variable(V) :-
%   atom_codes(V,[H|T]),
%   (letter(H) ; H=='_') ,
%   maplist(alphaNumeric,T).
%
% letter(L) :- 96<L,L<123;
%              64<L, L<91.

% variable([H|T]) :-
%   (letter([H],[]) ; H=='_') ,
%   maplist(alfNum,T).

% alfNum(V) :- digit([V],[]); letter([V],[]).

% digit --> [0];[1];[2];[3];[4];[5];[6];[7];[8];[9].
% letter -->
%   [a];[b];[c];[d];[e];[f];[g];[h];[i];[j];[k];[l];[m];
%   [n];[o];[p];[q];[r];[s];[t];[u];[v];[w];[x];[y];[z].
