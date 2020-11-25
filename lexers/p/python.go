package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Python lexer.
var Python = internal.Register(MustNewLazyLexer(
	&Config{
		Name:    "Python",
		Aliases: []string{"python", "py", "sage", "python3", "py3"},
		Filenames: []string{
			"*.py",
			"*.pyi",
			"*.pyw",
			// Jython
			"*.jy",
			// Sage
			"*.sage",
			// SCons
			"*.sc",
			"SConstruct",
			"SConscript",
			// Skylark/Starlark (used by Bazel, Buck, and Pants)
			"*.bzl",
			"BUCK",
			"BUILD",
			"BUILD.bazel",
			"WORKSPACE",
			// Twisted Application infrastructure
			"*.tac",
		},
		MimeTypes: []string{"text/x-python", "application/x-python", "text/x-python3", "application/x-python3"},
	},
	pythonRules,
))

func pythonRules() Rules {
	const pythonIdentifier = `[A-Z_a-zªµºÀ-ÖØ-öø-ˁˆ-ˑˠ-ˤˬˮͰ-ʹͶ-ͷͻ-ͽͿΆΈ-ΊΌΎ-ΡΣ-ϵϷ-ҁҊ-ԯԱ-Ֆՙՠ-ֈא-תׯ-ײؠ-يٮ-ٯٱ-ۓەۥ-ۦۮ-ۯۺ-ۼۿܐܒ-ܯݍ-ޥޱߊ-ߪߴ-ߵߺࠀ-ࠕࠚࠤࠨࡀ-ࡘࡠ-ࡪࢠ-ࢴࢶ-ࢽऄ-हऽॐक़-ॡॱ-ঀঅ-ঌএ-ঐও-নপ-রলশ-হঽৎড়-ঢ়য়-ৡৰ-ৱৼਅ-ਊਏ-ਐਓ-ਨਪ-ਰਲ-ਲ਼ਵ-ਸ਼ਸ-ਹਖ਼-ੜਫ਼ੲ-ੴઅ-ઍએ-ઑઓ-નપ-રલ-ળવ-હઽૐૠ-ૡૹଅ-ଌଏ-ଐଓ-ନପ-ରଲ-ଳଵ-ହଽଡ଼-ଢ଼ୟ-ୡୱஃஅ-ஊஎ-ஐஒ-கங-சஜஞ-டண-தந-பம-ஹௐఅ-ఌఎ-ఐఒ-నప-హఽౘ-ౚౠ-ౡಀಅ-ಌಎ-ಐಒ-ನಪ-ಳವ-ಹಽೞೠ-ೡೱ-ೲഅ-ഌഎ-ഐഒ-ഺഽൎൔ-ൖൟ-ൡൺ-ൿඅ-ඖක-නඳ-රලව-ෆก-ะาเ-ๆກ-ຂຄງ-ຈຊຍດ-ທນ-ຟມ-ຣລວສ-ຫອ-ະາຽເ-ໄໆໜ-ໟༀཀ-ཇཉ-ཬྈ-ྌက-ဪဿၐ-ၕၚ-ၝၡၥ-ၦၮ-ၰၵ-ႁႎႠ-ჅჇჍა-ჺჼ-ቈቊ-ቍቐ-ቖቘቚ-ቝበ-ኈኊ-ኍነ-ኰኲ-ኵኸ-ኾዀዂ-ዅወ-ዖዘ-ጐጒ-ጕጘ-ፚᎀ-ᎏᎠ-Ᏽᏸ-ᏽᐁ-ᙬᙯ-ᙿᚁ-ᚚᚠ-ᛪᛮ-ᛸᜀ-ᜌᜎ-ᜑᜠ-ᜱᝀ-ᝑᝠ-ᝬᝮ-ᝰក-ឳៗៜᠠ-ᡸᢀ-ᢨᢪᢰ-ᣵᤀ-ᤞᥐ-ᥭᥰ-ᥴᦀ-ᦫᦰ-ᧉᨀ-ᨖᨠ-ᩔᪧᬅ-ᬳᭅ-ᭋᮃ-ᮠᮮ-ᮯᮺ-ᯥᰀ-ᰣᱍ-ᱏᱚ-ᱽᲀ-ᲈᲐ-ᲺᲽ-Ჿᳩ-ᳬᳮ-ᳱᳵ-ᳶᴀ-ᶿḀ-ἕἘ-Ἕἠ-ὅὈ-Ὅὐ-ὗὙὛὝὟ-ώᾀ-ᾴᾶ-ᾼιῂ-ῄῆ-ῌῐ-ΐῖ-Ίῠ-Ῥῲ-ῴῶ-ῼⁱⁿₐ-ₜℂℇℊ-ℓℕ℘-ℝℤΩℨK-ℹℼ-ℿⅅ-ⅉⅎⅠ-ↈⰀ-Ⱞⰰ-ⱞⱠ-ⳤⳫ-ⳮⳲ-ⳳⴀ-ⴥⴧⴭⴰ-ⵧⵯⶀ-ⶖⶠ-ⶦⶨ-ⶮⶰ-ⶶⶸ-ⶾⷀ-ⷆⷈ-ⷎⷐ-ⷖⷘ-ⷞ々-〇〡-〩〱-〵〸-〼ぁ-ゖゝ-ゟァ-ヺー-ヿㄅ-ㄯㄱ-ㆎㆠ-ㆺㇰ-ㇿ㐀-䶵一-鿯ꀀ-ꒌꓐ-ꓽꔀ-ꘌꘐ-ꘟꘪ-ꘫꙀ-ꙮꙿ-ꚝꚠ-ꛯꜗ-ꜟꜢ-ꞈꞋ-ꞹꟷ-ꠁꠃ-ꠅꠇ-ꠊꠌ-ꠢꡀ-ꡳꢂ-ꢳꣲ-ꣷꣻꣽ-ꣾꤊ-ꤥꤰ-ꥆꥠ-ꥼꦄ-ꦲꧏꧠ-ꧤꧦ-ꧯꧺ-ꧾꨀ-ꨨꩀ-ꩂꩄ-ꩋꩠ-ꩶꩺꩾ-ꪯꪱꪵ-ꪶꪹ-ꪽꫀꫂꫛ-ꫝꫠ-ꫪꫲ-ꫴꬁ-ꬆꬉ-ꬎꬑ-ꬖꬠ-ꬦꬨ-ꬮꬰ-ꭚꭜ-ꭥꭰ-ꯢ가-힣ힰ-ퟆퟋ-ퟻ豈-舘並-龎ﬀ-ﬆﬓ-ﬗיִײַ-ﬨשׁ-זּטּ-לּמּנּ-סּףּ-פּצּ-ﮱﯓ-ﱝﱤ-ﴽﵐ-ﶏﶒ-ﷇﷰ-ﷹﹱﹳﹷﹹﹻﹽﹿ-ﻼＡ-Ｚａ-ｚｦ-ﾝﾠ-ﾾￂ-ￇￊ-ￏￒ-ￗￚ-ￜ𐀀-𐀋𐀍-𐀦𐀨-𐀺𐀼-𐀽𐀿-𐁍𐁐-𐁝𐂀-𐃺𐅀-𐅴𐊀-𐊜𐊠-𐋐𐌀-𐌟𐌭-𐍊𐍐-𐍵𐎀-𐎝𐎠-𐏃𐏈-𐏏𐏑-𐏕𐐀-𐒝𐒰-𐓓𐓘-𐓻𐔀-𐔧𐔰-𐕣𐘀-𐜶𐝀-𐝕𐝠-𐝧𐠀-𐠅𐠈𐠊-𐠵𐠷-𐠸𐠼𐠿-𐡕𐡠-𐡶𐢀-𐢞𐣠-𐣲𐣴-𐣵𐤀-𐤕𐤠-𐤹𐦀-𐦷𐦾-𐦿𐨀𐨐-𐨓𐨕-𐨗𐨙-𐨵𐩠-𐩼𐪀-𐪜𐫀-𐫇𐫉-𐫤𐬀-𐬵𐭀-𐭕𐭠-𐭲𐮀-𐮑𐰀-𐱈𐲀-𐲲𐳀-𐳲𐴀-𐴣𐼀-𐼜𐼧𐼰-𐽅𑀃-𑀷𑂃-𑂯𑃐-𑃨𑄃-𑄦𑅄𑅐-𑅲𑅶𑆃-𑆲𑇁-𑇄𑇚𑇜𑈀-𑈑𑈓-𑈫𑊀-𑊆𑊈𑊊-𑊍𑊏-𑊝𑊟-𑊨𑊰-𑋞𑌅-𑌌𑌏-𑌐𑌓-𑌨𑌪-𑌰𑌲-𑌳𑌵-𑌹𑌽𑍐𑍝-𑍡𑐀-𑐴𑑇-𑑊𑒀-𑒯𑓄-𑓅𑓇𑖀-𑖮𑗘-𑗛𑘀-𑘯𑙄𑚀-𑚪𑜀-𑜚𑠀-𑠫𑢠-𑣟𑣿𑨀𑨋-𑨲𑨺𑩐𑩜-𑪃𑪆-𑪉𑪝𑫀-𑫸𑰀-𑰈𑰊-𑰮𑱀𑱲-𑲏𑴀-𑴆𑴈-𑴉𑴋-𑴰𑵆𑵠-𑵥𑵧-𑵨𑵪-𑶉𑶘𑻠-𑻲𒀀-𒎙𒐀-𒑮𒒀-𒕃𓀀-𓐮𔐀-𔙆𖠀-𖨸𖩀-𖩞𖫐-𖫭𖬀-𖬯𖭀-𖭃𖭣-𖭷𖭽-𖮏𖹀-𖹿𖼀-𖽄𖽐𖾓-𖾟𖿠-𖿡𗀀-𘟱𘠀-𘫲𛀀-𛄞𛅰-𛋻𛰀-𛱪𛱰-𛱼𛲀-𛲈𛲐-𛲙𝐀-𝑔𝑖-𝒜𝒞-𝒟𝒢𝒥-𝒦𝒩-𝒬𝒮-𝒹𝒻𝒽-𝓃𝓅-𝔅𝔇-𝔊𝔍-𝔔𝔖-𝔜𝔞-𝔹𝔻-𝔾𝕀-𝕄𝕆𝕊-𝕐𝕒-𝚥𝚨-𝛀𝛂-𝛚𝛜-𝛺𝛼-𝜔𝜖-𝜴𝜶-𝝎𝝐-𝝮𝝰-𝞈𝞊-𝞨𝞪-𝟂𝟄-𝟋𞠀-𞣄𞤀-𞥃𞸀-𞸃𞸅-𞸟𞸡-𞸢𞸤𞸧𞸩-𞸲𞸴-𞸷𞸹𞸻𞹂𞹇𞹉𞹋𞹍-𞹏𞹑-𞹒𞹔𞹗𞹙𞹛𞹝𞹟𞹡-𞹢𞹤𞹧-𞹪𞹬-𞹲𞹴-𞹷𞹹-𞹼𞹾𞺀-𞺉𞺋-𞺛𞺡-𞺣𞺥-𞺩𞺫-𞺻𠀀-𪛖𪜀-𫜴𫝀-𫠝𫠠-𬺡𬺰-𮯠丽-𪘀][0-9A-Z_a-zªµ·ºÀ-ÖØ-öø-ˁˆ-ˑˠ-ˤˬˮ̀-ʹͶ-ͷͻ-ͽͿΆ-ΊΌΎ-ΡΣ-ϵϷ-ҁ҃-҇Ҋ-ԯԱ-Ֆՙՠ-ֈ֑-ֽֿׁ-ׂׄ-ׇׅא-תׯ-ײؐ-ؚؠ-٩ٮ-ۓە-ۜ۟-۪ۨ-ۼۿܐ-݊ݍ-ޱ߀-ߵߺ߽ࠀ-࠭ࡀ-࡛ࡠ-ࡪࢠ-ࢴࢶ-ࢽ࣓-ࣣ࣡-ॣ०-९ॱ-ঃঅ-ঌএ-ঐও-নপ-রলশ-হ়-ৄে-ৈো-ৎৗড়-ঢ়য়-ৣ০-ৱৼ৾ਁ-ਃਅ-ਊਏ-ਐਓ-ਨਪ-ਰਲ-ਲ਼ਵ-ਸ਼ਸ-ਹ਼ਾ-ੂੇ-ੈੋ-੍ੑਖ਼-ੜਫ਼੦-ੵઁ-ઃઅ-ઍએ-ઑઓ-નપ-રલ-ળવ-હ઼-ૅે-ૉો-્ૐૠ-ૣ૦-૯ૹ-૿ଁ-ଃଅ-ଌଏ-ଐଓ-ନପ-ରଲ-ଳଵ-ହ଼-ୄେ-ୈୋ-୍ୖ-ୗଡ଼-ଢ଼ୟ-ୣ୦-୯ୱஂ-ஃஅ-ஊஎ-ஐஒ-கங-சஜஞ-டண-தந-பம-ஹா-ூெ-ைொ-்ௐௗ௦-௯ఀ-ఌఎ-ఐఒ-నప-హఽ-ౄె-ైొ-్ౕ-ౖౘ-ౚౠ-ౣ౦-౯ಀ-ಃಅ-ಌಎ-ಐಒ-ನಪ-ಳವ-ಹ಼-ೄೆ-ೈೊ-್ೕ-ೖೞೠ-ೣ೦-೯ೱ-ೲഀ-ഃഅ-ഌഎ-ഐഒ-ൄെ-ൈൊ-ൎൔ-ൗൟ-ൣ൦-൯ൺ-ൿං-ඃඅ-ඖක-නඳ-රලව-ෆ්ා-ුූෘ-ෟ෦-෯ෲ-ෳก-ฺเ-๎๐-๙ກ-ຂຄງ-ຈຊຍດ-ທນ-ຟມ-ຣລວສ-ຫອ-ູົ-ຽເ-ໄໆ່-ໍ໐-໙ໜ-ໟༀ༘-༙༠-༩༹༵༷༾-ཇཉ-ཬཱ-྄྆-ྗྙ-ྼ࿆က-၉ၐ-ႝႠ-ჅჇჍა-ჺჼ-ቈቊ-ቍቐ-ቖቘቚ-ቝበ-ኈኊ-ኍነ-ኰኲ-ኵኸ-ኾዀዂ-ዅወ-ዖዘ-ጐጒ-ጕጘ-ፚ፝-፟፩-፱ᎀ-ᎏᎠ-Ᏽᏸ-ᏽᐁ-ᙬᙯ-ᙿᚁ-ᚚᚠ-ᛪᛮ-ᛸᜀ-ᜌᜎ-᜔ᜠ-᜴ᝀ-ᝓᝠ-ᝬᝮ-ᝰᝲ-ᝳក-៓ៗៜ-៝០-៩᠋-᠍᠐-᠙ᠠ-ᡸᢀ-ᢪᢰ-ᣵᤀ-ᤞᤠ-ᤫᤰ-᤻᥆-ᥭᥰ-ᥴᦀ-ᦫᦰ-ᧉ᧐-᧚ᨀ-ᨛᨠ-ᩞ᩠-᩿᩼-᪉᪐-᪙ᪧ᪰-᪽ᬀ-ᭋ᭐-᭙᭫-᭳ᮀ-᯳ᰀ-᰷᱀-᱉ᱍ-ᱽᲀ-ᲈᲐ-ᲺᲽ-Ჿ᳐-᳔᳒-᳹ᴀ-᷹᷻-ἕἘ-Ἕἠ-ὅὈ-Ὅὐ-ὗὙὛὝὟ-ώᾀ-ᾴᾶ-ᾼιῂ-ῄῆ-ῌῐ-ΐῖ-Ίῠ-Ῥῲ-ῴῶ-ῼ‿-⁀⁔ⁱⁿₐ-ₜ⃐-⃥⃜⃡-⃰ℂℇℊ-ℓℕ℘-ℝℤΩℨK-ℹℼ-ℿⅅ-ⅉⅎⅠ-ↈⰀ-Ⱞⰰ-ⱞⱠ-ⳤⳫ-ⳳⴀ-ⴥⴧⴭⴰ-ⵧⵯ⵿-ⶖⶠ-ⶦⶨ-ⶮⶰ-ⶶⶸ-ⶾⷀ-ⷆⷈ-ⷎⷐ-ⷖⷘ-ⷞⷠ-ⷿ々-〇〡-〯〱-〵〸-〼ぁ-ゖ゙-゚ゝ-ゟァ-ヺー-ヿㄅ-ㄯㄱ-ㆎㆠ-ㆺㇰ-ㇿ㐀-䶵一-鿯ꀀ-ꒌꓐ-ꓽꔀ-ꘌꘐ-ꘫꙀ-꙯ꙴ-꙽ꙿ-꛱ꜗ-ꜟꜢ-ꞈꞋ-ꞹꟷ-ꠧꡀ-ꡳꢀ-ꣅ꣐-꣙꣠-ꣷꣻꣽ-꤭ꤰ-꥓ꥠ-ꥼꦀ-꧀ꧏ-꧙ꧠ-ꧾꨀ-ꨶꩀ-ꩍ꩐-꩙ꩠ-ꩶꩺ-ꫂꫛ-ꫝꫠ-ꫯꫲ-꫶ꬁ-ꬆꬉ-ꬎꬑ-ꬖꬠ-ꬦꬨ-ꬮꬰ-ꭚꭜ-ꭥꭰ-ꯪ꯬-꯭꯰-꯹가-힣ힰ-ퟆퟋ-ퟻ豈-舘並-龎ﬀ-ﬆﬓ-ﬗיִ-ﬨשׁ-זּטּ-לּמּנּ-סּףּ-פּצּ-ﮱﯓ-ﱝﱤ-ﴽﵐ-ﶏﶒ-ﷇﷰ-ﷹ︀-️︠-︯︳-︴﹍-﹏ﹱﹳﹷﹹﹻﹽﹿ-ﻼ０-９Ａ-Ｚ＿ａ-ｚｦ-ﾾￂ-ￇￊ-ￏￒ-ￗￚ-ￜ𐀀-𐀋𐀍-𐀦𐀨-𐀺𐀼-𐀽𐀿-𐁍𐁐-𐁝𐂀-𐃺𐅀-𐅴𐇽𐊀-𐊜𐊠-𐋐𐋠𐌀-𐌟𐌭-𐍊𐍐-𐍺𐎀-𐎝𐎠-𐏃𐏈-𐏏𐏑-𐏕𐐀-𐒝𐒠-𐒩𐒰-𐓓𐓘-𐓻𐔀-𐔧𐔰-𐕣𐘀-𐜶𐝀-𐝕𐝠-𐝧𐠀-𐠅𐠈𐠊-𐠵𐠷-𐠸𐠼𐠿-𐡕𐡠-𐡶𐢀-𐢞𐣠-𐣲𐣴-𐣵𐤀-𐤕𐤠-𐤹𐦀-𐦷𐦾-𐦿𐨀-𐨃𐨅-𐨆𐨌-𐨓𐨕-𐨗𐨙-𐨵𐨸-𐨿𐨺𐩠-𐩼𐪀-𐪜𐫀-𐫇𐫉-𐫦𐬀-𐬵𐭀-𐭕𐭠-𐭲𐮀-𐮑𐰀-𐱈𐲀-𐲲𐳀-𐳲𐴀-𐴧𐴰-𐴹𐼀-𐼜𐼧𐼰-𐽐𑀀-𑁆𑁦-𑁯𑁿-𑂺𑃐-𑃨𑃰-𑃹𑄀-𑄴𑄶-𑄿𑅄-𑅆𑅐-𑅳𑅶𑆀-𑇄𑇉-𑇌𑇐-𑇚𑇜𑈀-𑈑𑈓-𑈷𑈾𑊀-𑊆𑊈𑊊-𑊍𑊏-𑊝𑊟-𑊨𑊰-𑋪𑋰-𑋹𑌀-𑌃𑌅-𑌌𑌏-𑌐𑌓-𑌨𑌪-𑌰𑌲-𑌳𑌵-𑌹𑌻-𑍄𑍇-𑍈𑍋-𑍍𑍐𑍗𑍝-𑍣𑍦-𑍬𑍰-𑍴𑐀-𑑊𑑐-𑑙𑑞𑒀-𑓅𑓇𑓐-𑓙𑖀-𑖵𑖸-𑗀𑗘-𑗝𑘀-𑙀𑙄𑙐-𑙙𑚀-𑚷𑛀-𑛉𑜀-𑜚𑜝-𑜫𑜰-𑜹𑠀-𑠺𑢠-𑣩𑣿𑨀-𑨾𑩇𑩐-𑪃𑪆-𑪙𑪝𑫀-𑫸𑰀-𑰈𑰊-𑰶𑰸-𑱀𑱐-𑱙𑱲-𑲏𑲒-𑲧𑲩-𑲶𑴀-𑴆𑴈-𑴉𑴋-𑴶𑴺𑴼-𑴽𑴿-𑵇𑵐-𑵙𑵠-𑵥𑵧-𑵨𑵪-𑶎𑶐-𑶑𑶓-𑶘𑶠-𑶩𑻠-𑻶𒀀-𒎙𒐀-𒑮𒒀-𒕃𓀀-𓐮𔐀-𔙆𖠀-𖨸𖩀-𖩞𖩠-𖩩𖫐-𖫭𖫰-𖫴𖬀-𖬶𖭀-𖭃𖭐-𖭙𖭣-𖭷𖭽-𖮏𖹀-𖹿𖼀-𖽄𖽐-𖽾𖾏-𖾟𖿠-𖿡𗀀-𘟱𘠀-𘫲𛀀-𛄞𛅰-𛋻𛰀-𛱪𛱰-𛱼𛲀-𛲈𛲐-𛲙𛲝-𛲞𝅥-𝅩𝅭-𝅲𝅻-𝆂𝆅-𝆋𝆪-𝆭𝉂-𝉄𝐀-𝑔𝑖-𝒜𝒞-𝒟𝒢𝒥-𝒦𝒩-𝒬𝒮-𝒹𝒻𝒽-𝓃𝓅-𝔅𝔇-𝔊𝔍-𝔔𝔖-𝔜𝔞-𝔹𝔻-𝔾𝕀-𝕄𝕆𝕊-𝕐𝕒-𝚥𝚨-𝛀𝛂-𝛚𝛜-𝛺𝛼-𝜔𝜖-𝜴𝜶-𝝎𝝐-𝝮𝝰-𝞈𝞊-𝞨𝞪-𝟂𝟄-𝟋𝟎-𝟿𝨀-𝨶𝨻-𝩬𝩵𝪄𝪛-𝪟𝪡-𝪯𞀀-𞀆𞀈-𞀘𞀛-𞀡𞀣-𞀤𞀦-𞀪𞠀-𞣄𞣐-𞣖𞤀-𞥊𞥐-𞥙𞸀-𞸃𞸅-𞸟𞸡-𞸢𞸤𞸧𞸩-𞸲𞸴-𞸷𞸹𞸻𞹂𞹇𞹉𞹋𞹍-𞹏𞹑-𞹒𞹔𞹗𞹙𞹛𞹝𞹟𞹡-𞹢𞹤𞹧-𞹪𞹬-𞹲𞹴-𞹷𞹹-𞹼𞹾𞺀-𞺉𞺋-𞺛𞺡-𞺣𞺥-𞺩𞺫-𞺻𠀀-𪛖𪜀-𫜴𫝀-𫠝𫠠-𬺡𬺰-𮯠丽-𪘀󠄀-󠇯]*`

	return Rules{
		"root": {
			{`\n`, Text, nil},
			{`^(\s*)([rRuUbB]{,2})("""(?:.|\n)*?""")`, ByGroups(Text, LiteralStringAffix, LiteralStringDoc), nil},
			{`^(\s*)([rRuUbB]{,2})('''(?:.|\n)*?''')`, ByGroups(Text, LiteralStringAffix, LiteralStringDoc), nil},
			{`\A#!.+$`, CommentHashbang, nil},
			{`#.*$`, CommentSingle, nil},
			{`\\\n`, Text, nil},
			{`\\`, Text, nil},
			Include("keywords"),
			{`(def)((?:\s|\\\s)+)`, ByGroups(Keyword, Text), Push("funcname")},
			{`(class)((?:\s|\\\s)+)`, ByGroups(Keyword, Text), Push("classname")},
			{`(from)((?:\s|\\\s)+)`, ByGroups(KeywordNamespace, Text), Push("fromimport")},
			{`(import)((?:\s|\\\s)+)`, ByGroups(KeywordNamespace, Text), Push("import")},
			Include("expr"),
		},
		"expr": {
			{`(?i)(rf|fr)(""")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Combined("rfstringescape", "tdqf")},
			{`(?i)(rf|fr)(''')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Combined("rfstringescape", "tsqf")},
			{`(?i)(rf|fr)(")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Combined("rfstringescape", "dqf")},
			{`(?i)(rf|fr)(')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Combined("rfstringescape", "sqf")},
			{`([fF])(""")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Combined("fstringescape", "tdqf")},
			{`([fF])(''')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Combined("fstringescape", "tsqf")},
			{`([fF])(")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Combined("fstringescape", "dqf")},
			{`([fF])(')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Combined("fstringescape", "sqf")},
			{`(?i)(rb|br|r)(""")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Push("tdqs")},
			{`(?i)(rb|br|r)(''')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Push("tsqs")},
			{`(?i)(rb|br|r)(")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Push("dqs")},
			{`(?i)(rb|br|r)(')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Push("sqs")},
			{`([uUbB]?)(""")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Combined("stringescape", "tdqs")},
			{`([uUbB]?)(''')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Combined("stringescape", "tsqs")},
			{`([uUbB]?)(")`, ByGroups(LiteralStringAffix, LiteralStringDouble), Combined("stringescape", "dqs")},
			{`([uUbB]?)(')`, ByGroups(LiteralStringAffix, LiteralStringSingle), Combined("stringescape", "sqs")},
			{`[^\S\n]+`, Text, nil},
			Include("numbers"),
			{`!=|==|<<|>>|:=|[-~+/*%=<>&^|.]`, Operator, nil},
			{`[]{}:(),;[]`, Punctuation, nil},
			{`(in|is|and|or|not)\b`, OperatorWord, nil},
			Include("expr-keywords"),
			Include("builtins"),
			Include("magicfuncs"),
			Include("magicvars"),
			Include("name"),
		},
		"expr-inside-fstring": {
			{`[{([]`, Punctuation, Push("expr-inside-fstring-inner")},
			{`(=\s*)?(\![sraf])?\}`, LiteralStringInterpol, Pop(1)},
			{`(=\s*)?(\![sraf])?:`, LiteralStringInterpol, Pop(1)},
			{`\s+`, Text, nil},
			Include("expr"),
		},
		"expr-inside-fstring-inner": {
			{`[{([]`, Punctuation, Push("expr-inside-fstring-inner")},
			{`[])}]`, Punctuation, Pop(1)},
			{`\s+`, Text, nil},
			Include("expr"),
		},
		"expr-keywords": {
			{Words(``, `\b`, `async for`, `await`, `else`, `for`, `if`, `lambda`, `yield`, `yield from`), Keyword, nil},
			{Words(``, `\b`, `True`, `False`, `None`), KeywordConstant, nil},
		},
		"keywords": {
			{Words(``, `\b`, `assert`, `async`, `await`, `break`, `continue`, `del`, `elif`, `else`, `except`, `finally`, `for`, `global`, `if`, `lambda`, `pass`, `raise`, `nonlocal`, `return`, `try`, `while`, `yield`, `yield from`, `as`, `with`), Keyword, nil},
			{Words(``, `\b`, `True`, `False`, `None`), KeywordConstant, nil},
		},
		"builtins": {
			{Words(`(?<!\.)`, `\b`, `__import__`, `abs`, `all`, `any`, `bin`, `bool`, `bytearray`, `bytes`, `chr`, `classmethod`, `compile`, `complex`, `delattr`, `dict`, `dir`, `divmod`, `enumerate`, `eval`, `filter`, `float`, `format`, `frozenset`, `getattr`, `globals`, `hasattr`, `hash`, `hex`, `id`, `input`, `int`, `isinstance`, `issubclass`, `iter`, `len`, `list`, `locals`, `map`, `max`, `memoryview`, `min`, `next`, `object`, `oct`, `open`, `ord`, `pow`, `print`, `property`, `range`, `repr`, `reversed`, `round`, `set`, `setattr`, `slice`, `sorted`, `staticmethod`, `str`, `sum`, `super`, `tuple`, `type`, `vars`, `zip`), NameBuiltin, nil},
			{`(?<!\.)(self|Ellipsis|NotImplemented|cls)\b`, NameBuiltinPseudo, nil},
			{Words(`(?<!\.)`, `\b`, `ArithmeticError`, `AssertionError`, `AttributeError`, `BaseException`, `BufferError`, `BytesWarning`, `DeprecationWarning`, `EOFError`, `EnvironmentError`, `Exception`, `FloatingPointError`, `FutureWarning`, `GeneratorExit`, `IOError`, `ImportError`, `ImportWarning`, `IndentationError`, `IndexError`, `KeyError`, `KeyboardInterrupt`, `LookupError`, `MemoryError`, `NameError`, `NotImplementedError`, `OSError`, `OverflowError`, `PendingDeprecationWarning`, `ReferenceError`, `ResourceWarning`, `RuntimeError`, `RuntimeWarning`, `StopIteration`, `SyntaxError`, `SyntaxWarning`, `SystemError`, `SystemExit`, `TabError`, `TypeError`, `UnboundLocalError`, `UnicodeDecodeError`, `UnicodeEncodeError`, `UnicodeError`, `UnicodeTranslateError`, `UnicodeWarning`, `UserWarning`, `ValueError`, `VMSError`, `Warning`, `WindowsError`, `ZeroDivisionError`, `BlockingIOError`, `ChildProcessError`, `ConnectionError`, `BrokenPipeError`, `ConnectionAbortedError`, `ConnectionRefusedError`, `ConnectionResetError`, `FileExistsError`, `FileNotFoundError`, `InterruptedError`, `IsADirectoryError`, `NotADirectoryError`, `PermissionError`, `ProcessLookupError`, `TimeoutError`, `StopAsyncIteration`, `ModuleNotFoundError`, `RecursionError`), NameException, nil},
		},
		"magicfuncs": {
			{Words(``, `\b`, `__abs__`, `__add__`, `__aenter__`, `__aexit__`, `__aiter__`, `__and__`, `__anext__`, `__await__`, `__bool__`, `__bytes__`, `__call__`, `__complex__`, `__contains__`, `__del__`, `__delattr__`, `__delete__`, `__delitem__`, `__dir__`, `__divmod__`, `__enter__`, `__eq__`, `__exit__`, `__float__`, `__floordiv__`, `__format__`, `__ge__`, `__get__`, `__getattr__`, `__getattribute__`, `__getitem__`, `__gt__`, `__hash__`, `__iadd__`, `__iand__`, `__ifloordiv__`, `__ilshift__`, `__imatmul__`, `__imod__`, `__imul__`, `__index__`, `__init__`, `__instancecheck__`, `__int__`, `__invert__`, `__ior__`, `__ipow__`, `__irshift__`, `__isub__`, `__iter__`, `__itruediv__`, `__ixor__`, `__le__`, `__len__`, `__length_hint__`, `__lshift__`, `__lt__`, `__matmul__`, `__missing__`, `__mod__`, `__mul__`, `__ne__`, `__neg__`, `__new__`, `__next__`, `__or__`, `__pos__`, `__pow__`, `__prepare__`, `__radd__`, `__rand__`, `__rdivmod__`, `__repr__`, `__reversed__`, `__rfloordiv__`, `__rlshift__`, `__rmatmul__`, `__rmod__`, `__rmul__`, `__ror__`, `__round__`, `__rpow__`, `__rrshift__`, `__rshift__`, `__rsub__`, `__rtruediv__`, `__rxor__`, `__set__`, `__setattr__`, `__setitem__`, `__str__`, `__sub__`, `__subclasscheck__`, `__truediv__`, `__xor__`), NameFunctionMagic, nil},
		},
		"magicvars": {
			{Words(``, `\b`, `__annotations__`, `__bases__`, `__class__`, `__closure__`, `__code__`, `__defaults__`, `__dict__`, `__doc__`, `__file__`, `__func__`, `__globals__`, `__kwdefaults__`, `__module__`, `__mro__`, `__name__`, `__objclass__`, `__qualname__`, `__self__`, `__slots__`, `__weakref__`), NameVariableMagic, nil},
		},
		"numbers": {
			{`(\d(?:_?\d)*\.(?:\d(?:_?\d)*)?|(?:\d(?:_?\d)*)?\.\d(?:_?\d)*)([eE][+-]?\d(?:_?\d)*)?`, LiteralNumberFloat, nil},
			{`\d(?:_?\d)*[eE][+-]?\d(?:_?\d)*j?`, LiteralNumberFloat, nil},
			{`0[oO](?:_?[0-7])+`, LiteralNumberOct, nil},
			{`0[bB](?:_?[01])+`, LiteralNumberBin, nil},
			{`0[xX](?:_?[a-fA-F0-9])+`, LiteralNumberHex, nil},
			{`\d(?:_?\d)*`, LiteralNumberInteger, nil},
		},
		"name": {
			{`@` + pythonIdentifier, NameDecorator, nil},
			{`@`, Operator, nil},
			{pythonIdentifier, Name, nil},
		},
		"funcname": {
			Include("magicfuncs"),
			{pythonIdentifier, NameFunction, Pop(1)},
			Default(Pop(1)),
		},
		"classname": {
			{pythonIdentifier, NameClass, Pop(1)},
		},
		"import": {
			{`(\s+)(as)(\s+)`, ByGroups(Text, Keyword, Text), nil},
			{`\.`, NameNamespace, nil},
			{pythonIdentifier, NameNamespace, nil},
			{`(\s*)(,)(\s*)`, ByGroups(Text, Operator, Text), nil},
			Default(Pop(1)),
		},
		"fromimport": {
			{`(\s+)(import)\b`, ByGroups(Text, KeywordNamespace), Pop(1)},
			{`\.`, NameNamespace, nil},
			{`None\b`, NameBuiltinPseudo, Pop(1)},
			{pythonIdentifier, NameNamespace, nil},
			Default(Pop(1)),
		},
		"rfstringescape": {
			{`\{\{`, LiteralStringEscape, nil},
			{`\}\}`, LiteralStringEscape, nil},
		},
		"fstringescape": {
			Include("rfstringescape"),
			Include("stringescape"),
		},
		"stringescape": {
			{`\\([\\abfnrtv"\']|\n|N\{.*?\}|u[a-fA-F0-9]{4}|U[a-fA-F0-9]{8}|x[a-fA-F0-9]{2}|[0-7]{1,3})`, LiteralStringEscape, nil},
		},
		"fstrings-single": {
			{`\}`, LiteralStringInterpol, nil},
			{`\{`, LiteralStringInterpol, Push("expr-inside-fstring")},
			{`[^\\\'"{}\n]+`, LiteralStringSingle, nil},
			{`[\'"\\]`, LiteralStringSingle, nil},
		},
		"fstrings-double": {
			{`\}`, LiteralStringInterpol, nil},
			{`\{`, LiteralStringInterpol, Push("expr-inside-fstring")},
			{`[^\\\'"{}\n]+`, LiteralStringDouble, nil},
			{`[\'"\\]`, LiteralStringDouble, nil},
		},
		"strings-single": {
			{`%(\(\w+\))?[-#0 +]*([0-9]+|[*])?(\.([0-9]+|[*]))?[hlL]?[E-GXc-giorsaux%]`, LiteralStringInterpol, nil},
			{`\{((\w+)((\.\w+)|(\[[^\]]+\]))*)?(\![sra])?(\:(.?[<>=\^])?[-+ ]?#?0?(\d+)?,?(\.\d+)?[E-GXb-gnosx%]?)?\}`, LiteralStringInterpol, nil},
			{`[^\\\'"%{\n]+`, LiteralStringSingle, nil},
			{`[\'"\\]`, LiteralStringSingle, nil},
			{`%|(\{{1,2})`, LiteralStringSingle, nil},
		},
		"strings-double": {
			{`%(\(\w+\))?[-#0 +]*([0-9]+|[*])?(\.([0-9]+|[*]))?[hlL]?[E-GXc-giorsaux%]`, LiteralStringInterpol, nil},
			{`\{((\w+)((\.\w+)|(\[[^\]]+\]))*)?(\![sra])?(\:(.?[<>=\^])?[-+ ]?#?0?(\d+)?,?(\.\d+)?[E-GXb-gnosx%]?)?\}`, LiteralStringInterpol, nil},
			{`[^\\\'"%{\n]+`, LiteralStringDouble, nil},
			{`[\'"\\]`, LiteralStringDouble, nil},
			{`%|(\{{1,2})`, LiteralStringDouble, nil},
		},
		"dqf": {
			{`"`, LiteralStringDouble, Pop(1)},
			{`\\\\|\\"|\\\n`, LiteralStringEscape, nil},
			Include("fstrings-double"),
		},
		"sqf": {
			{`'`, LiteralStringSingle, Pop(1)},
			{`\\\\|\\'|\\\n`, LiteralStringEscape, nil},
			Include("fstrings-single"),
		},
		"dqs": {
			{`"`, LiteralStringDouble, Pop(1)},
			{`\\\\|\\"|\\\n`, LiteralStringEscape, nil},
			Include("strings-double"),
		},
		"sqs": {
			{`'`, LiteralStringSingle, Pop(1)},
			{`\\\\|\\'|\\\n`, LiteralStringEscape, nil},
			Include("strings-single"),
		},
		"tdqf": {
			{`"""`, LiteralStringDouble, Pop(1)},
			Include("fstrings-double"),
			{`\n`, LiteralStringDouble, nil},
		},
		"tsqf": {
			{`'''`, LiteralStringSingle, Pop(1)},
			Include("fstrings-single"),
			{`\n`, LiteralStringSingle, nil},
		},
		"tdqs": {
			{`"""`, LiteralStringDouble, Pop(1)},
			Include("strings-double"),
			{`\n`, LiteralStringDouble, nil},
		},
		"tsqs": {
			{`'''`, LiteralStringSingle, Pop(1)},
			Include("strings-single"),
			{`\n`, LiteralStringSingle, nil},
		},
	}
}
