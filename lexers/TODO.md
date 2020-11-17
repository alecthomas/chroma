## Wakatime lexer support

For the following lexers, text analysis capabilities of pygments have to be ported to chroma for the wakatime golang client to work. In general we focus on lexers, which are associated with the same file extension. If the lexer doesn't even exist in chroma yet, it has to be added.

## Top languages

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
| `*.sql`        | MySQL          |                    |
|                | SQL            | :heavy_check_mark: |
| `*.ts`         | TypeScript     |                    |
|                | TypoScript     |                    |
| `*.v`          | Coq            | :heavy_check_mark: |
|                | verilog        | :heavy_check_mark: |
| `*.xslt`       | HTML           |                    |
|                | XML            |                    |

## Long tail languages

| file extension(s)                                    | lexer            | lexer exists | note                                | done               |
| ---                                                  | ---              | ---          | ---                                 | ---                |
| `*.asax`,`*.ascx`,`*.ashx`,`*.asmx`,`*.aspx`,`*.axd` | aspx-cs          | :x:          |                                     | :heavy_check_mark: |
|                                                      | aspx-vb          | :x:          |                                     | :heavy_check_mark: |
| `*.ASM`                                              | NASM             |              |                                     |                    |
|                                                      | TASM             |              |                                     |                    |
| `*.S`                                                | GAS              |              |                                     |                    |
|                                                      | S                | :x:          |                                     |                    |
| `*.b`                                                | Brainfuck        |              |                                     |                    |
|                                                      | Limbo            | :x:          |                                     | :heavy_check_mark: |
| `*.bas`                                              | CBM BASIC V2     | :x:          |                                     |                    |
|                                                      | QBasic           |              |                                     |                    |
|                                                      | VB.net           |              |                                     |                    |
| `*.bug`                                              | BUGS             | :x:          |                                     |                    |
|                                                      | JAGS             | :x:          |                                     |                    |
| `*.def`                                              | Modula-2         |              |                                     |                    |
|                                                      | Singularity      | :x:          |                                     |                    |
| `*.ecl`                                              | ECL              | :x:          |                                     |                    |
|                                                      | Prolog           |              |                                     |                    |
| `*.gd`                                               | GAP              | :x:          |                                     |                    |
|                                                      | GDScript         |              |                                     |                    |
| `*.hy`                                               | Hy               |              |                                     |                    |
|                                                      | Hybris           | :x:          |                                     |                    |
| `*.inc`                                              | Pawn             | :x:          |                                     |                    |
|                                                      | PHP              |              |                                     |                    |
|                                                      | POVRay           |              |                                     |                    |
| `*.inf`                                              | Inform 6         | :x:          |                                     |                    |
|                                                      | INI              |              |                                     |                    |
| `*.j`                                                | Jasmin           | :x:          |                                     |                    |
|                                                      | Objective-J      |              |                                     |                    |
| `*.n`                                                | Ezhil            | :x:          |                                     |                    |
|                                                      | Nemerle          | :x:          |                                     |                    |
| `*.p`                                                | OpenEdge ABL     | :x:          |                                     |                    |
|                                                      | Pawn             | :x:          |                                     |                    |
| `*.pl`                                               | Perl6            | :x:          |                                     |                    |
|                                                      | Perl             |              |                                     |                    |
|                                                      | Prolog           |              |                                     |                    |
| `*.pm`                                               | Perl6            | :x:          |                                     |                    |
|                                                      | Perl             |              |                                     |                    |
| `*.pro`                                              | IDL              | :x:          |                                     |                    |
|                                                      | Prolog           |              |                                     |                    |
| `*.s`                                                | ca65 assembler   | :x:          |                                     |                    |
|                                                      | GAS              |              |                                     |                    |
| `*.sc`                                               | Python           |              |                                     |                    |
|                                                      | SuperCollider    | :x:          |                                     |                    |
| `*.scd`                                              | scdoc            | :x:          |                                     |                    |
|                                                      | SuperCollider    | :x:          |                                     |                    |
| `*.sl`                                               | Slash            | :x:          | No text analysis exists in pygments |                    |
|                                                      | Slurm            | :x:          | No text analysis exists in pygments |                    |
| `*.sql`                                              | SQL              |              |                                     |                    |
|                                                      | Transact-SQL     |              |                                     |                    |
| `*.t`                                                | Perl6            |              |                                     |                    |
|                                                      | Perl             |              |                                     |                    |
|                                                      | TADS 3           | :x:          |                                     |                    |
| `*.ttl`                                              | Tera Term macro  | :x:          |                                     |                    |
|                                                      | Turtle           |              |                                     |                    |
| `*.u`                                                | ucode            | :x:          |                                     |                    |
|                                                      | UrbiScript       | :x:          |                                     |                    |
| `*.v`                                                | Coq              |              |                                     |                    |
|                                                      | verilog          |              |                                     |                    |
| `*.xsl`                                              | XML              |              |                                     |                    |
|                                                      | XSLT             | :x:          |                                     |                    |
| `*.xslt`                                             | HTML             |              |                                     |                    |
|                                                      | XML              |              |                                     |                    |
|                                                      | XSLT             | :x:          |                                     |                    |
