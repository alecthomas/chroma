#!/usr/bin/env python3
import re
from collections import defaultdict
from subprocess import check_output

README_FILE: str = "README.md"

"""

Example output:

lexers:
  ABAP
    aliases: abap
    filenames: *.abap *.ABAP
    mimetypes: text/x-abap
  ABNF
    aliases: abnf
    filenames: *.abnf
    mimetypes: text/x-abnf

"""

if __name__ == "__main__":
    lines: list[str] = check_output(["go", "run", "-C", "./cmd/chroma", ".", "--list"]).decode("utf-8").splitlines()

    # Find only the names of the lexers, which are intended by *only* two spaces.
    lines = [line.strip() for line in lines if line.startswith("  ") and not line.startswith("   ")]
    lines = sorted(lines, key=lambda line: line.lower())

    table: defaultdict[str, list[str]] = defaultdict(list)

    for line in lines:
        table[line[0].upper()].append(line)

    header: str = "| Prefix | Language"
    separator: str = "| :----: | {language_seperator}"

    rows: list[str] = []

    language_seperator_length: int = 0

    for key, value in table.items():
        # Spacing here based on the length of "Prefix"
        lexers = ", ".join(value)
        language_seperator_length = max(language_seperator_length, len(lexers))
        rows.append(f"|   {key}    | {lexers}")

    tbody = "\n".join(rows)

    lexer_table: str = (
        separator.format(language_seperator="-" * language_seperator_length) + "\n" + tbody
    )

    with open(README_FILE, "r") as f:
        content = f.read()

    with open(README_FILE, "w") as f:
        # Want to replace content between (and including) "| Prefix | Language" and the first two blank lines after it.
        marker = re.compile(r"(?P<start>" + re.escape(header) + r"\n).*?(?P<end>\n\n)", re.DOTALL)
        replacement = r"\g<start>%s\g<end>" % lexer_table
        updated_content = marker.sub(replacement, content)
        f.write(updated_content)

    print(tbody)
