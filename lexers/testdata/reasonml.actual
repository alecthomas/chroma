// This is a single line comment.
/* A multi line comment, in a single line... */
/* This is a multi line comment
   Second Line...
*/
/* This is a nested comment
  /* Nested Line... */
*/

module Process_Bind_Without_Do =
       (W: Monad_Bind with type m('a) = writer(string, 'a)) => {
  let process = s => W.(up_case(s) >>= (up_str => to_words(up_str)));
};

let a = 1 or 2;
let b = 1 || 2;
let c = 1 && 2;

let str = "Hello, Lexer!";

let chr = 'a';

type test;

open Belt;

include Pervasives;

let test: unit => Map.String.t(string) = () => Map.String.empty;

let tup = (1: int, 2: int);

let myRec = {x: 0, y: 10};

let myFuncs = {
  myFun: (x) => x + 1,
  your: (a, b) => a + b
};

let lst = [1, 2, 3];

let logRest = (lst) => 
  switch (lst) {
    | [] => Js.log("no entry")
    | [hd, ...rest] => Js.log2("Rest: ", rest);
  };

let arr = [|1, 2, 3|];

let res = (x) =>
  switch (x) {
  | HasNothing => 0
  | HasSingleInt(x) => 0
  | HasSingleTuple((x, y)) => 0
  | HasMultipleInts(x, y) => 0
  | HasMultipleTuples((x, y), (q, r)) => 0
  };

module View = {
  [@react.component]
  let make = () => {
    <div className="view"> 
      <ul>
      <li> React.string("Hello, World!") </li>
      <li> "pipe"->React.string </li>
      <li> <span> "nested"->React.string </span> </li>
      </ul>
    </div>
  }
}