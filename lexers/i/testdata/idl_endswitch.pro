pro ex_switch, x
   switch x of
      1: print, 'one'
      2: print, 'two'
      else: begin
         print, 'you entered: ', x
         print, 'please enter a value between 1 and 4'
         end
   endswitch
end
ex_switch, 2
