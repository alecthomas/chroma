import importlib
import sys
import types
import json

import pystache
from pygments import lexer as pygments_lexer
from pygments.token import _TokenType


TEMPLATE = r'''
package lexers

import (
    . "github.com/alecthomas/chroma" // nolint
)

// {{upper_name}} lexer.
var {{upper_name}} = Register(NewLexer(
    &Config{
        Name:      "{{name}}",
        Aliases:   []string{ {{#aliases}}"{{.}}", {{/aliases}} },
        Filenames: []string{ {{#filenames}}"{{.}}", {{/filenames}} },
        MimeTypes: []string{ {{#mimetypes}}"{{.}}", {{/mimetypes}} },
    },
    Rules{
{{#tokens}}
        "{{state}}": {
            {{#rules}}
            {{{.}}},
            {{/rules}}
        },
{{/tokens}}
    },
))
'''


def go_string(s):
    if '(?<' in s:
        warning('perl regex found in %r' % s)
    if '`' not in s:
        return '`' + s + '`'
    return json.dumps(s)


def to_camel_case(snake_str):
    components = snake_str.split('_')
    return ''.join(x.title() for x in components)


def warning(message):
    print('warning: ' + message, file=sys.stderr)


def resolve_emitter(emitter):
    if isinstance(emitter, types.FunctionType):
        if repr(emitter).startswith('<function bygroups.'):
            args = emitter.__closure__[0].cell_contents
            emitter = 'ByGroups(%s)' % ', '.join(resolve_emitter(e) for e in args)
        elif repr(emitter).startswith('<function using.'):
            args = emitter.__closure__[0].cell_contents
            if isinstance(args, dict):
                state = 'root'
                if 'stack' in args:
                    state = args['stack'][1]
                    args.pop('stack')
                assert args == {}, args
                emitter = 'UsingSelf("%s")' % state
            elif issubclass(args, pygments_lexer.Lexer):
                emitter = 'Using(%s, nil)' % args.__name__
            else:
                raise ValueError('only support "using" with lexer classes, not %r' % args)
        else:
            warning('unsupported emitter function %r' % emitter)
            emitter = '??'
    elif isinstance(emitter, _TokenType):
        emitter = str(emitter).replace('.', '')[5:]
    elif emitter is None:
        # This generally only occurs when a lookahead/behind assertion is used, so we just allow it
        # through.
        return 'None'
    else:
        raise ValueError('unsupported emitter type %r' % emitter)
    assert isinstance(emitter, str)
    return emitter


def process_state_action(action):
    if action.startswith('#'):
        action = action[1:]
        if action == 'pop':
            action = 'Pop(1)'
        elif action == 'push':
            action = 'Push()'
        elif action.startswith('push:'):
            action = 'Push("%s")' % action[5:]
        else:
            raise ValueError('unsupported action %r' % (action,))
    else:
        action = 'Push("%s")' % action
    return action


def translate_rules(rules):
    out = []
    for rule in rules:
        if isinstance(rule, tuple):
            regex = rule[0]
            if isinstance(regex, str):
                regex = go_string(regex)
            elif isinstance(regex, pygments_lexer.words):
                regex = '`%s(?:%s)%s`' % (regex.prefix, '|'.join(regex.words), regex.suffix)
            else:
                raise ValueError('expected regex string but got %r' % regex)
            emitter = resolve_emitter(rule[1])
            if len(rule) == 2:
                modifier = 'nil'
            elif type(rule[2]) is str:
                modifier = process_state_action(rule[2])
            elif isinstance(rule[2], pygments_lexer.combined):
                modifier = 'Combined("%s")' % '", "'.join(rule[2])
            elif type(rule[2]) is tuple:
                modifier = 'Push("%s")' % '", "'.join(rule[2])
            else:
                raise ValueError('unsupported modifier %r' % (rule[2],))
            out.append('{{ {}, {}, {} }}'.format(regex, emitter, modifier))
        elif isinstance(rule, pygments_lexer.include):
            out.append('Include("{}")'.format(rule))
        elif isinstance(rule, pygments_lexer.default):
            out.append('Default({})'.format(process_state_action(rule.state)))
        else:
            raise ValueError('unsupported rule %r' % (rule,))
    return out


def main():
    package_name, symbol_name = sys.argv[1].rsplit(sep=".", maxsplit=1)

    package = importlib.import_module(package_name)

    lexer_cls = getattr(package, symbol_name)

    assert issubclass(lexer_cls, pygments_lexer.RegexLexer), 'can only translate from RegexLexer'

    print(pystache.render(TEMPLATE, {
        'name': lexer_cls.name,
        'upper_name': to_camel_case(lexer_cls.name),
        'aliases': lexer_cls.aliases,
        'filenames': lexer_cls.filenames,
        'mimetypes': lexer_cls.mimetypes,
        'tokens': [{'state': state, 'rules': translate_rules(rules)} for (state, rules) in lexer_cls.get_tokendefs().items()],
    }))


if __name__ == '__main__':
    main()