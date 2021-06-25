%% Sorted is a sorted version of List if Sorted is
%% a permutation of List (same elements in possibly
%% different order) and Sorted is sorted (second rule).
sorted(List, Sorted) :-
    perm(List, Sorted),
    sorted(Sorted).
