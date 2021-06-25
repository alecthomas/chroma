for (local i in 0 .. __TADS3)
    word += concat(
        rand(rand('', clusters, consonants)), rand('"h"?'),
        rand(vowels...), rand('','', 'i', 'u', rand(ends)));
