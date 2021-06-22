function ::=
  <datatype> function_name ( {<arg>} ) '{' <statement> return <datatype> ; '}'

arg ::= <datatype> arg_name ,

statement ::=
  <assignment_statement> |
  <loop_statement> |
  <conditional_statement> |
  <function_call_statement>

<loop_statement> ::=
  while ( conditional_statement ) '{' <statement> '}' |
  do '{' <statement> '}' while ( conditional_statement )

<conditional_statement> ::=
  if ( conditional_statement ) '{' <statement> '}' |
  if ( conditional_statement ) '{' <statement> '}'
  else '{' <statement> '}' |
  if ( conditional_statement ) '{' <statement> '}'
  { else if ( conditional_statement ) '{' <statement> '}' }
  else '{' <statement> '}'

datatype ::= void | char | int | short | long | float | double
specialsymbol ::= if | else | for | while | do | return
relational_operator  == | > | < | => | =< | !=

conditional_statement ::=
  function_call |
  [!] integer | float  relational_operator


function_call 10

do {
  <statement>
} while(logic_expr);

while(logic_expr) { <statement> }

if(conditional_expression)
  { <statement> }

if(conditional_expression)
  { <statement> }
else { <statement> }

if(conditional_expression)
  { <statement> }
else if(conditional_expression)
  { <statement> }

if(conditional_expression)
  { <statement> }
else if(conditional_expression)
  { <statement> }
else { <statement> }
