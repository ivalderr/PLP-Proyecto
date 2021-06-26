expression([H|T]) :-
  ( (atom_codes(H,X),variable(X)) ;
    function_call_statement(H) 
  ),
  operator, expression().

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

datatype --> char;int;float;double.
value --> natural;integer;decimal;string;character.

relational_operator(OP) :-
  OP=='==' ; OP=='!=' ;
  OP=='>'  ; OP=='<'  ;
  OP=='=>' ; OP=='=<' .

logical_operator --> [&&];['|'];[!].
operator --> [+];[-];[*];[/];['%'].

variable([H|T]) :-
  (letter(H) ; H=='_') ,
  maplist(alphaNumeric,T).

letter(L) :- 96<L,L<123;
             64<L, L<91.

% variable([H|T]) :-
%   (letter([H],[]) ; H=='_') ,
%   maplist(alfNum,T).

% alfNum(V) :- digit([V],[]); letter([V],[]).

% digit --> [0];[1];[2];[3];[4];[5];[6];[7];[8];[9].
% letter -->
%   [a];[b];[c];[d];[e];[f];[g];[h];[i];[j];[k];[l];[m];
%   [n];[o];[p];[q];[r];[s];[t];[u];[v];[w];[x];[y];[z].
