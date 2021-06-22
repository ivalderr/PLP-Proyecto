function ::=
  <datatype> function_name '(' {<arg>} ')' '{'
    {<statement>}
    return ( <variable> | <value> | <function_call_statement> ) ;
  '}'

arg ::= <datatype> arg_name ,

statement ::=
  <assignment_statement> |
  <loop_statement> |
  <conditional_statement> |
  <function_call_statement>

<assignment_statement> ::=

<loop_statement> ::=
  while ( conditional ) '{' <statement> '}' |
  do '{' <statement> '}' while ( conditional ) ;

<conditional_statement> ::=
  if ( conditional ) '{' <statement> '}' |
  if ( conditional ) '{' <statement> '}'
  else '{' <statement> '}' |
  if ( conditional ) '{' <statement> '}'
  { else if ( conditional ) '{' <statement> '}' }
  else '{' <statement> '}'

datatype ::= void | char | int | short | long | float | double
relational_operator  == | > | < | => | =< | !=

conditional :: =
function_call_statement ::=
variable ::=
value :: =
