implement Values;

include "sys.m";

sys: Sys;
print, sprint: import sys;

Values: module {
	init: fn(nil: ref Draw->Context, nil: list of string);
};

init(nil: ref Draw->Context, nil: list of string) {
	sys = load Sys Sys->PATH;

	str := "String!";

	exit;
}
