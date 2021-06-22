% operation statement / parentesis multiples / return opcional / value en condicional
function ::=
  <datatype> function_name '('
    [<datatype> arg_name>{',' <datatype> arg_name}]
  ')' '{'
    {<statement>}
    return
      <variable> | <value> |
      <conditional> | <function_call_statement> ';'
  '}'

function_call_statement ::=
  Fname '(' [<arg>] ')' ';' |
  <variable> '=' Fname '(' [<arg>] ')' ';'

arg ::=
  <variable>{',' <arg>} |
  <value>{',' <arg>} |
  <function_call_statement>{',' <arg>}

statement ::=
  <assignment_statement> |
  <loop_statement> |
  <conditional_statement> |
  <function_call_statement>

assignment_statement ::=
  

conditional :: =
  [!] (<value> | <variable> | <function_call_statement>)
      {<relational_operator> <conditional>}

<loop_statement> ::=
  while ( <conditional> ) '{' <statement> '}' |
  do '{' <statement> '}' while ( conditional ) ;

<conditional_statement> ::=
  if ( <conditional> ) '{' <statement> '}' |
  if ( <conditional> ) '{' <statement> '}'
  else '{' <statement> '}' |
  if ( <conditional> ) '{' <statement> '}'
  { else if ( <conditional> ) '{' <statement> '}' }
  else '{' <statement> '}'

datatype ::= void | char | int | short | long | float | double
relational_operator ::= == | > | < | => | =< | !=
value ::= integer | decimal | string | character
%%
variable ::=
