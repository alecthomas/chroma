%module swig_example

// Add necessary symbols to generated header
%{
#include "swig-example.h"
%}

// Process symbols in header
%include "swig-example.h"
