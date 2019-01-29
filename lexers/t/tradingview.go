package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TradingView lexer.
var TradingView = internal.Register(MustNewLexer(
	&Config{
		Name:      "TradingView",
		Aliases:   []string{"tradingview", "tv"},
		Filenames: []string{"*.tv"},
		MimeTypes: []string{"text/x-tradingview"},
		DotAll:    true,
		EnsureNL:  true,
	},
	Rules{
		"root": {
			{`[^\S\n]+|\n|[()]`, Text, nil},
			{`(//.*?)(\n)`, ByGroups(CommentSingle, Text), nil},
			{`>=|<=|==|!=|>|<|\?|-|\+|\*|\/|%|\[|\]`, Operator, nil},
			{`[:,.]`, Punctuation, nil},
			{`=`, KeywordPseudo, nil},
			{`"(\\\\|\\"|[^"\n])*["\n]`, LiteralString, nil},
			{`'\\.'|'[^\\]'`, LiteralString, nil},
			{`[0-9](\.[0-9]*)?([eE][+-][0-9]+)?`, LiteralNumber, nil},
			{`#[a-fA-F0-9]{8}|#[a-fA-F0-9]{6}|#[a-fA-F0-9]{3}`, LiteralStringOther, nil},
			{`(abs|acos|alertcondition|alma|asin|atan|atr|avg|barcolor|barssince|bgcolor|cci|ceil|change|cog|correlation|cos|crossover|crossunder|cum|dev|ema|exp|falling|fill|fixnan|floor|heikinashi|highest|highestbars|hline|iff|input|kagi|linebreak|linreg|log|log10|lowest|lowestbars|macd|max|min|mom|nz|percentile_(linear_interpolation|nearest_rank)|percentrank|pivothigh|pivotlow|plot|plotarrow|plotbar|plotcandle|plotchar|plotshape|pointfigure|pow|renko|rising|rma|roc|round|rsi|sar|security|sign|sin|sma|sqrt|stdev|stoch|study|sum|swma|tan|timestamp|tostring|tsi|valuewhen|variance|vwma|wma|strategy\.(cancel|cancel_all|close|close_all|entry|exit|order|risk\.(allow_entry_in|max_cons_loss_days|max_drawdown|max_intraday_filled_orders|max_intraday_loss|max_position_size)))\b`, NameFunction, nil},
			{`\b(cross|dayofmonth|dayofweek|hour|minute|month|na|offset|second|strategy|tickerid|time|tr|vwap|weekofyear|year)(\()`, ByGroups(NameFunction, Text), nil}, // functions that can also be variable
			{`(accdist|adjustment\.(dividends|none|splits)|aqua|area|areabr|black|blue|bool|circles|close|columns|currency\.(AUD|CAD|CHF|EUR|GBP|HKD|JPY|NOK|NONE|NZD|RUB|SEK|SGD|TRY|USD|ZAR)|dashed|dotted|float|friday|fuchsia|gray|green|high|histogram|hl2|hlc3|integer|interval|isdaily|isdwm|isintraday|ismonthly|isweekly|lime|line|linebr|location\.(abovebar|absolute|belowbar|bottom|top)|low|maroon|monday|n|navy|ohlc4|olive|open|orange|period|purple|red|resolution|saturday|scale\.(left|none|right)|session|session\.(extended|regular)|silver|size\.(auto|huge|large|normal|small|tiny)|solid|source|stepline|string|sunday|symbol|syminfo\.(mintick|pointvalue|prefix|root|session|timezone)|teal|thursday|ticker|timenow|tuesday|volume|wednesday|white|yellow|strategy\.(cash|closedtrades|commission\.(cash_per_contract|cash_per_order|percent)|direction\.(all|long|short)|equity|eventrades|fixed|grossloss|grossprofit|initial_capital|long|losstrades|max_contracts_held_(all|long|short)|max_drawdown|netprofit|oca\.(cancel|none|reduce)|openprofit|opentrades|percent_of_equity|position_avg_price|position_entry_name|position_size|short|wintrades)|shape\.(arrowdown|arrowup|circle|cross|diamond|flag|labeldown|labelup|square|triangledown|triangleup|xcross)|barstate\.is(first|history|last|new|realtime)|barmerge\.(gaps_on|gaps_off|lookahead_on|lookahead_off))\b`, NameVariable, nil},
			{`(cross|dayofmonth|dayofweek|hour|minute|month|na|second|tickerid|time|tr|vwap|weekofyear|year)(\b[^\(])`, ByGroups(NameVariable, Text), nil}, // variables that can also be function
			{`(true|false)\b`, KeywordConstant, nil},
			{`(and|or|not|if|else|for|to)\b`, OperatorWord, nil},
			{`@?[_a-zA-Z]\w*`, Text, nil},
		},
	},
))
