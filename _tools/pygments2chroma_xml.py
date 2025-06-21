#!/usr/bin/env -S uv run --script
import functools
import importlib
import re
import sys
import types
import html

import pystache
from pygments import lexer as pygments_lexer
from pygments.token import _TokenType

TEMPLATE = r'''
<lexer>
  <config>
    <name>{{name}}</name>
    {{#aliases}}
    <alias>{{alias}}</alias>
    {{/aliases}}
    {{#filenames}}
    <filename>{{filename}}</filename>
    {{/filenames}}
    {{#mimetypes}}
    <mime_type>{{mimetype}}</mime_type>
    {{/mimetypes}}
    {{#re_ignorecase}}
    <case_insensitive>true</case_insensitive>
    {{/re_ignorecase}}
    {{#re_dotall}}
    <dot_all>true</dot_all>
    {{/re_dotall}}
    {{#re_not_multiline}}
    <not_multiline>true</not_multiline>
    {{/re_not_multiline}}
  </config>
  <rules>
    {{#tokens}}
    <state name="{{state}}">
      {{#rules}}
      {{{.}}}
      {{/rules}}
    </state>
    {{/tokens}}
  </rules>
</lexer>
'''


def xml_regex(s):
    return xml_string(s)

def xml_string(s):
    s = html.escape(s)
    return '"' + s + '"'


def to_camel_case(snake_str):
    components = snake_str.split('_')
    return ''.join(x.title() for x in components)


def warning(message):
    print('warning: ' + message, file=sys.stderr)


def resolve_emitter(emitter):
    if isinstance(emitter, types.FunctionType):
        if repr(emitter).startswith('<function bygroups.'):
            args = emitter.__closure__[0].cell_contents
            emitter = '<bygroups>%s</bygroups>' % ''.join(resolve_emitter(e) for e in args)
        elif repr(emitter).startswith('<function using.'):
            args = emitter.__closure__[0].cell_contents
            if isinstance(args, dict):
                state = 'root'
                if 'stack' in args:
                    state = args['stack'][1]
                    args.pop('stack')
                assert args == {}, args
                emitter = '<usingself state="%s"/>' % state
            elif issubclass(args, pygments_lexer.Lexer):
                name = args.__name__
                if name.endswith('Lexer'):
                    name = name[:-5]
                emitter = '<using lexer="%s"/>' % state
            else:
                raise ValueError('only support "using" with lexer classes, not %r' % args)
        else:
            warning('unsupported emitter function %r' % emitter)
            emitter = '?? %r ??' % emitter
    elif isinstance(emitter, _TokenType):
        emitter = '<token type="%s"/>' % str(emitter).replace('.', '')[5:]
    elif emitter is None:
        return 'None'
    else:
        raise ValueError('unsupported emitter type %r' % emitter)
    assert isinstance(emitter, str)
    return emitter


def process_state_action(action):
    if isinstance(action, tuple):
        return functools.reduce(lambda a, b: a + b, (process_state_action(a) for a in action))
    if action.startswith('#'):
        action = action[1:]
        if action== 'pop':
            action = '<pop depth="1"/>'
        elif action.startswith('pop:'):
            action = '<pop depth="%s"/>' % action[4:]
        elif action == 'push':
            action = '<push/>'
        elif action.startswith('push:'):
            action = '<push state="%s"/>' % action[5:]
        else:
            raise ValueError('unsupported action %r' % (action,))
    else:
        action = '<push state="%s"/>' % action
    return (action,)


def translate_rules(rules):
    out = []
    for rule in rules:
        if isinstance(rule, tuple):
            regex = rule[0]
            if isinstance(regex, str):
                regex = xml_regex(regex)
            elif isinstance(regex, pygments_lexer.words):
                regex = xml_string('%s(%s)%s' % (regex.prefix,
                                      '|'.join(re.escape(w) for w in regex.words),
                                      regex.suffix))
            else:
                raise ValueError('expected regex string but got %r' % regex)
            emitter = resolve_emitter(rule[1])
            if len(rule) == 2:
                modifier = ''
            elif type(rule[2]) is str:
                modifier = process_state_action(rule[2])[0]
            elif isinstance(rule[2], pygments_lexer.combined):
                modifier = '<combined state="%s"/>' % '" state="'.join(rule[2])
            elif type(rule[2]) is tuple:
                modifier = '<push state="%s"/>' % '" state="'.join(rule[2])
            else:
                raise ValueError('unsupported modifier %r' % (rule[2],))
            out.append('<rule pattern={}>{}{}</rule>'.format(regex, emitter, modifier))
        elif isinstance(rule, pygments_lexer.include):
            out.append('<rule><include state="{}"/></rule>'.format(rule))
        elif isinstance(rule, pygments_lexer.default):
            process_state_action(rule.state)
            out.append('<rule>{}</rule>'.format(''.join(process_state_action(rule.state))))
        else:
            raise ValueError('unsupported rule %r' % (rule,))
    return out


class TemplateView(object):
    def __init__(self, **kwargs):
        for key, value in kwargs.items():
            setattr(self, key, value)

    def re_not_multiline(self):
        return not (self.regex_flags & re.MULTILINE)

    def re_dotall(self):
        return self.regex_flags & re.DOTALL

    def re_ignorecase(self):
        return self.regex_flags & re.IGNORECASE


def main():
    package_name, symbol_name = sys.argv[1].rsplit(sep=".", maxsplit=1)

    package = importlib.import_module(package_name)

    lexer_cls = getattr(package, symbol_name)

    assert issubclass(lexer_cls, pygments_lexer.RegexLexer), 'can only translate from RegexLexer'

    print(pystache.render(TEMPLATE, TemplateView(
        name=lexer_cls.name,
        regex_flags=lexer_cls.flags,
        aliases=[{'alias': alias} for alias in lexer_cls.aliases],
        filenames=[{'filename': filename} for filename in lexer_cls.filenames],
        mimetypes=[{'mimetype': mimetype} for mimetype in lexer_cls.mimetypes],
        tokens=[{'state': state, 'rules': translate_rules(rules)} for (state, rules) in lexer_cls.get_tokendefs().items()],
    )))


if __name__ == '__main__':
    main()
