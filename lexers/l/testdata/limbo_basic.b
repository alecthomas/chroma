implement Values;

include "sys.m";
include "draw.m";

sys: Sys;
print, sprint: import sys;

Values: module {
	init: fn(nil: ref Draw->Context, nil: list of string);
};

init(nil: ref Draw->Context, nil: list of string) {
	sys = load Sys Sys->PATH;

	n := 7;
	b := big 8;
	f := real 3.2;
	str := "String!";

	print("%d\n", 0 || 1);
	print("%d\n", 0 && 1);

	print("%d\n", n / int f);
	print("%f\n", real n / f);
	print("%bd\n", b / big 8);

	print("%s\n", str[:len str-1]);
	print("%s\n", str[2:]);

	print("%s", "inferno " + "os " + sprint("%c", '\n'));
	print("limbo" + " " + "lang\n");

	exit;
}