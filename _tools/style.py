import importlib
import sys

import pystache
from pygments.style import Style
from pygments.token import Token


TEMPLATE = r'''
package styles

import (
    "github.com/alecthomas/chroma"
)

// {{upper_name}} style.
var {{upper_name}} = Register(chroma.MustNewStyle("{{name}}", chroma.StyleEntries{
{{#styles}}
    chroma.{{type}}: "{{style}}",
{{/styles}}
}))
'''


def to_camel_case(snake_str):
    components = snake_str.split('_')
    return ''.join(x.title() for x in components)


def translate_token_type(t):
    if t == Token:
        t = Token.Background
    return "".join(map(str, t))


def main():
    name = sys.argv[1]
    package_name, symbol_name = sys.argv[2].rsplit(sep=".", maxsplit=1)

    package = importlib.import_module(package_name)

    style_cls = getattr(package, symbol_name)

    assert issubclass(style_cls, Style), 'can only translate from Style subclass'

    styles = dict(style_cls.styles)
    bg = "bg:" + style_cls.background_color
    if Token in styles:
        styles[Token] += " " + bg
    else:
        styles[Token] = bg
    context = {
        'upper_name': style_cls.__name__[:-5],
        'name': name,
        'styles': [{'type': translate_token_type(t), 'style': s}
                   for t, s in styles.items() if s],
    }
    print(pystache.render(TEMPLATE, context))


if __name__ == '__main__':
    main()