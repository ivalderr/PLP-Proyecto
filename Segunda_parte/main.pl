:-  ensure_loaded([tokenizer,gramatica]).

testFile(FileName,RT):-
    open(FileName, 'read', InputStream),
    read_stream_to_codes(InputStream, ProgramString),
    close(InputStream),
    phrase(tokenize(TSBefore), ProgramString),
    write(TSBefore),
    function_call_statement(TSBefore).
