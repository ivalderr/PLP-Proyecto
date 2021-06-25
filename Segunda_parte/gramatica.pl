variable([H|T]) :-
  (letter([H],[]) ; H=='_') ,
  maplist(alfNum,T).

alfNum(V) :- digit(V,[]); letter(V,[]).



digit --> [0];[1];[2];[3];[4];[5];[6];[7];[8];[9].
letter --> 'a';'b';'c'.
