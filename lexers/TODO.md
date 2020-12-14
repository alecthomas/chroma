# Wakatime lexer support

## Text analysis

For the following lexers, text analysis capabilities of pygments have to be ported to chroma for the wakatime golang client to work. In general we focus on lexers, which are associated with the same file extension. If the lexer doesn't even exist in chroma yet, it has to be added.

### Top languages

| file extension | lexer          | done               |
| ---            | ---            | ---                |
| `*.as`         | ActionScript   | :heavy_check_mark: |
|                | ActionScript 3 | :heavy_check_mark: |
| `*.asm`        | NASM           | :heavy_check_mark: |
|                | TASM           | :heavy_check_mark: |
| `*.bas`        | QBasic         | :heavy_check_mark: |
|                | VB.net         | :heavy_check_mark: |
| `*.c`          | C              | :heavy_check_mark: |
|                | C++            | :heavy_check_mark: |
| `*.fs`         | Forth          | :heavy_check_mark: |
|                | FSharp         | :heavy_check_mark: |
| `*.h`          | C              | :heavy_check_mark: |
|                | C++            | :heavy_check_mark: |
|                | Objective-C    | :heavy_check_mark: |
| `*.inc`        | POVRay         | :heavy_check_mark: |
|                | PHP            | :heavy_check_mark: |
| `*.m`          | Mason          | :heavy_check_mark: |
|                | Matlab         | :heavy_check_mark: |
|                | Objective-C    | :heavy_check_mark: |
|                | Octave         | :heavy_check_mark: |
| `*.mc`         | Mason          | :heavy_check_mark: |
|                | MonkeyC        |                    |
| `*.pl`         | Perl           | :heavy_check_mark: |
|                | Prolog         | :heavy_check_mark: |
| `*.s`          | GAS            | :heavy_check_mark: |
|                | R              | :heavy_check_mark: |
| `*.sql`        | MySQL          | :heavy_check_mark: |
|                | SQL            | :heavy_check_mark: |
| `*.ts`         | TypeScript     |                    |
|                | TypoScript     |                    |
| `*.v`          | Coq            | :heavy_check_mark: |
|                | verilog        | :heavy_check_mark: |
| `*.xslt`       | HTML           | :heavy_check_mark: |
|                | XML            |                    |

### Long tail languages

| file extension(s)                                    | lexer            | lexer exists | note                                | done               |
| ---                                                  | ---              | ---          | ---                                 | ---                |
| `*.asax`,`*.ascx`,`*.ashx`,`*.asmx`,`*.aspx`,`*.axd` | aspx-cs          | :x:          |                                     | :heavy_check_mark: |
|                                                      | aspx-vb          | :x:          |                                     | :heavy_check_mark: |
| `*.ASM`                                              | NASM             | :x:          |                                     | :heavy_check_mark: |
|                                                      | TASM             | :x:          |                                     | :heavy_check_mark: |
| `*.S`                                                | GAS              | :x:          |                                     | :heavy_check_mark: |
|                                                      | S                | :x:          | It's implemented altogether in r.go | :heavy_check_mark: |
| `*.b`                                                | Brainfuck        | :x:          |                                     | :heavy_check_mark: |
|                                                      | Limbo            | :x:          |                                     | :heavy_check_mark: |
| `*.bas`                                              | CBM BASIC V2     | :x:          |                                     | :heavy_check_mark: |
|                                                      | QBasic           | :x:          |                                     | :heavy_check_mark: |
|                                                      | VB.net           | :x:          |                                     | :heavy_check_mark: |
| `*.bug`                                              | BUGS             | :x:          |                                     | :heavy_check_mark: |
|                                                      | JAGS             | :x:          |                                     | :heavy_check_mark: |
| `*.def`                                              | Modula-2         | :x:          |                                     | :heavy_check_mark: |
|                                                      | Singularity      | :x:          |                                     | :heavy_check_mark: |
| `*.ecl`                                              | ECL              | :x:          |                                     | :heavy_check_mark: |
|                                                      | Prolog           | :x:          |                                     | :heavy_check_mark: |
| `*.gd`                                               | GAP              | :x:          |                                     | :heavy_check_mark: |
|                                                      | GDScript         | :x:          |                                     | :heavy_check_mark: |
| `*.hy`                                               | Hy               | :x:          |                                     | :heavy_check_mark: |
|                                                      | Hybris           | :x:          |                                     | :heavy_check_mark: |
| `*.inc`                                              | Pawn             | :x:          |                                     | :heavy_check_mark: |
|                                                      | PHP              | :x:          |                                     | :heavy_check_mark: |
|                                                      | POVRay           | :x:          |                                     | :heavy_check_mark: |
| `*.inf`                                              | Inform 6         | :x:          |                                     | :heavy_check_mark: |
|                                                      | INI              | :x:          |                                     | :heavy_check_mark: |
| `*.j`                                                | Jasmin           | :x:          |                                     | :heavy_check_mark: |
|                                                      | Objective-J      | :x:          |                                     | :heavy_check_mark: |
| `*.n`                                                | Ezhil            | :x:          |                                     | :heavy_check_mark: |
|                                                      | Nemerle          | :x:          |                                     | :heavy_check_mark: |
| `*.p`                                                | OpenEdge ABL     | :x:          |                                     | :heavy_check_mark: |
|                                                      | Pawn             | :x:          |                                     | :heavy_check_mark: |
| `*.pl`                                               | Perl6            | :x:          |                                     |                    |
|                                                      | Perl             | :x:          |                                     | :heavy_check_mark: |
|                                                      | Prolog           | :x:          |                                     | :heavy_check_mark: |
| `*.pm`                                               | Perl6            | :x:          |                                     |                    |
|                                                      | Perl             | :x:          |                                     | :heavy_check_mark: |
| `*.pro`                                              | IDL              | :x:          |                                     | :heavy_check_mark: |
|                                                      | Prolog           | :x:          |                                     | :heavy_check_mark: |
| `*.s`                                                | ca65 assembler   | :x:          |                                     | :heavy_check_mark: |
|                                                      | GAS              | :x:          |                                     | :heavy_check_mark: |
| `*.sc`                                               | Python           |              |                                     |                    |
|                                                      | SuperCollider    | :x:          |                                     | :heavy_check_mark: |
| `*.scd`                                              | scdoc            | :x:          |                                     | :heavy_check_mark: |
|                                                      | SuperCollider    | :x:          |                                     | :heavy_check_mark: |
| `*.sl`                                               | Slash            | :x:          | No text analysis exists in pygments |                    |
|                                                      | Slurm            | :x:          | No text analysis exists in pygments |                    |
| `*.sql`                                              | SQL              | :x:          |                                     | :heavy_check_mark: |
|                                                      | Transact-SQL     | :x:          |                                     | :heavy_check_mark: |
| `*.t`                                                | Perl6            |              |                                     |                    |
|                                                      | Perl             |              |                                     |                    |
|                                                      | TADS 3           | :x:          |                                     | :heavy_check_mark: |
| `*.ttl`                                              | Tera Term macro  | :x:          |                                     | :heavy_check_mark: |
|                                                      | Turtle           | :x:          |                                     | :heavy_check_mark: |
| `*.u`                                                | ucode            | :x:          |                                     | :heavy_check_mark: |
|                                                      | UrbiScript       | :x:          |                                     |                    |
| `*.v`                                                | Coq              | :x:          |                                     | :heavy_check_mark: |
|                                                      | verilog          | :x:          |                                     | :heavy_check_mark: |
| `*.xsl`                                              | XML              |              |                                     |                    |
|                                                      | XSLT             | :x:          |                                     |                    |
| `*.xslt`                                             | HTML             | :x:          |                                     | :heavy_check_mark: |
|                                                      | XML              |              |                                     |                    |
|                                                      | XSLT             | :x:          |                                     |                    |

## Missing file extension support

For the following lexers, file extension support has to be added, to match the behaviour of wakatime cli. This mostly matches pygments.

### Top languages

| lexer       | filename pattern                                             | done                                  |
| ---         | ---                                                          | ---                                   |
| Elixir      | `.eex`                                                       | :heavy_check_mark:                    |
| JavaScript  | `*.mjs`                                                      | :heavy_check_mark:                    |
| JSX         | `*.jsx`                                                      | not needed, as matched by react lexer |
| Python      | `*.jy`, `*.bzl`, `BUCK`, `BUILD`, `BUILD.bazel`, `WORKSPACE` | :heavy_check_mark:                    |
| TOML        | `Pipfile`, `poetry.lock`                                     | :heavy_check_mark:                    |
| Twig        | `*.twig`                                                     | :heavy_check_mark:                    |

### Long tail

TBD

## Missing lexers

The following lexers exist in pygments, but not in chroma. They have to be added to emulate the behaviour of pygments, which is the baseline for the wakatime cli.

Only a very stripped down version is necessary here, though. No token related be functionality is needed and we really only care about 2 things here:

1. The lexer name, should match the pygments lexer name.
2. Associated file name pattern have to be identical to pygments.

### Top languages

| lexer               | filename pattern                            | not in pygments | done               |
| ---                 | ---                                         | ---             | ---                |
| Crontab             | `crontab`                                   | :x:             | :heavy_check_mark: |
| Coldfusion HTML     | `.cfm`, `*.cfml`                            |                 | :heavy_check_mark: |
| Delphi              | `.pas`, `*.dpr`                             |                 | :heavy_check_mark: |
| Gosu                | `*.gs`, `*.gsp`, `*.gst`, `*.gsx`, `*.vark` |                 | :heavy_check_mark: |
| Lasso               | `*.lasso`, `*.lasso[89]`                    |                 | :heavy_check_mark: |
| LessCss             | `*.less`                                    |                 | :heavy_check_mark: |
| liquid              | `*.liquid`                                  |                 | :heavy_check_mark: |
| Marko               | `*.marko`                                   | :x:             | :heavy_check_mark: |
| Modelica            | `*.mo`                                      |                 | :heavy_check_mark: |
| Mustache            | `*.mustache`                                | :x:             | :heavy_check_mark: |
| NewLisp             | `*.lsp`, `*.nl`, `*.kif`                    |                 | :heavy_check_mark: |
| Objective-J         | `*.j`                                       |                 | :heavy_check_mark: |
| Pawn                | `*.p`, `*.pwn`, `*.inc`                     |                 | :heavy_check_mark: |
| Pug                 | `*.pug`, `*.jade`                           |                 | :heavy_check_mark: |
| QML                 | `*.qml`, `*.qbs`                            |                 | :heavy_check_mark: |
| RPMSpec             | `*.spec`                                    |                 | :heavy_check_mark: |
| Sketch Drawing      | `*.sketch`                                  | :x:             | :heavy_check_mark: |
| Slim                | `*.slim`                                    |                 | :heavy_check_mark: |
| Smali               | `*.smali`                                   |                 | :heavy_check_mark: |
| SourcePawn          | `*.sp`                                      |                 | :heavy_check_mark: |
| Sublime Text Config | `*.sublime-settings`                        | :x:             | :heavy_check_mark: |
| Svelte              | `*.svelte`                                  | :x:             | :heavy_check_mark: |
| SWIG                | `*.swg`, `*.i`                              |                 | :heavy_check_mark: |
| VCL                 | `*.vcl`                                     |                 | :heavy_check_mark: |
| Velocity            | `*.vm`, `*.fhtml`                           |                 | :heavy_check_mark: |
| XAML                | `*.xaml`                                    | :x:             | :heavy_check_mark: |
| XSLT                | `*.xsl`, `*.xslt`, `*.xpl` # xpl is XProc   |                 | :heavy_check_mark: |

### Long tail

TBD
