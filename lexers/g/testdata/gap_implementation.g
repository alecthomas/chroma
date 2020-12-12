InstallMethod( Iterator,
    "method for `Integers'",
    [ IsIntegers ],
    function( Integers )
    return Objectify( NewType( IteratorsFamily,
                                   IsIterator
                               and IsIntegersIteratorCompRep ),
                      rec( counter := 0 ) );
    end );
