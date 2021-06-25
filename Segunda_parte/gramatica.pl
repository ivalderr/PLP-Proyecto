variable --> "," , (digit ; letter).
variable --> (letter ; "_"), variable.

letter(C) :- 96<C,C<123;
              64<C, C<91;
              47<C, C<58.
digit(D) --> [D], {47 < D, D < 58}.
