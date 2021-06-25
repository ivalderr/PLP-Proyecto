datatype --> char;int;float;double.
value --> natural;integer;decimal;string;character.

relational_operator(OP) :-
  OP=='==' ; OP=='!=' ;
  OP=='>'  ; OP=='<'  ;
  OP=='=>' ; OP=='=<' .

logical_operator --> [&&];['|'];[!].
operator --> [+];[-];[*];[/];['%'].

variable([H|T]) :-
  (letter([H],[]) ; H=='_') ,
  maplist(alfNum,T).

alfNum(V) :- digit([V],[]); letter([V],[]).

digit --> [0];[1];[2];[3];[4];[5];[6];[7];[8];[9].
letter -->
  [a];[b];[c];[d];[e];[f];[g];[h];[i];[j];[k];[l];[m];
  [n];[o];[p];[q];[r];[s];[t];[u];[v];[w];[x];[y];[z].
