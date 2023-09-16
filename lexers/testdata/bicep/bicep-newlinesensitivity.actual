@allowed(['abc', 'def', 'ghi'])
param foo string

var singleLineFunction = concat('abc', 'def')

var multiLineFunction = concat(
  'abc',
  'def'
)

var multiLineFunctionUnusualFormatting = concat(
              'abc',          any(['hello']),
'def')

var nestedTest = concat(
concat(
concat(
concat(
concat(
'level',
'one'),
'two'),
'three'),
'four'),
'five')

var singleLineArray = ['abc', 'def']
var singleLineArrayTrailingCommas = ['abc', 'def',]

var multiLineArray = [
  'abc'
  'def'
]

var mixedArray = ['abc', 'def'
'ghi', 'jkl'
'lmn']

var singleLineObject = { abc: 'def', ghi: 'jkl'}
var singleLineObjectTrailingCommas = { abc: 'def', ghi: 'jkl',}
var multiLineObject = {
  abc: 'def'
  ghi: 'jkl'
}
var mixedObject = { abc: 'abc', def: 'def'
ghi: 'ghi', jkl: 'jkl'
lmn: 'lmn' }

var nestedMixed = {
  abc: { 'def': 'ghi', abc: 'def', foo: [
    'bar', 'blah'
  ] }
}

var brokenFormatting = [      /*foo */ 'bar'   /*

hello

*/,        'asdfdsf',             12324,       /*   asdf*/ '',     '''


'''
123,      233535
true
              ]
