// reverse polish

mtype = { operator, value }

chan f = [12] of { mtype, int }

proctype calc(chan you)
{	int s, lft, rgt
	chan me = [0] of { int }

	if
	:: f?operator(s)
		run calc(me); me?lft
		run calc(me); me?rgt
		if
		:: s == '+' -> you!(lft+rgt)
		fi
	:: f?value(s) -> you!s
	fi
}
