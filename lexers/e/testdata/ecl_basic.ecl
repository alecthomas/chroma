DATA132 ColumnMap(UNICODE str) := BEGINC++
    size_t pos;
    unsigned char col = 0;
    memset(__result, '\0', 132);  // init to no column
    for(pos=0; pos<132; pos++) {
      if (pos<lenStr && str[pos] == (UChar)'<') col++;
      ((unsigned char *)__result)[pos] = col;
    }
ENDC++;
