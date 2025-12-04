#!/usr/bin/env python3

import re
import sys
from collections import defaultdict


def parse_lexers(lines: list[str]) -> list[str]:
	"""Parse the output of chroma --list and return a list of lexer names"""

	lexer_name_re: re.Pattern[str] = re.compile(r"^  ([^:\s].*?)\s*$")
	lexers: list[str] = []
	in_lexers = False

	for line in lines:
		line = line.rstrip()
		if line.startswith("lexers:"):
			in_lexers = True
			continue
		if not in_lexers:
			continue

		# stop when we hit styles/formatters/etc
		if line.startswith("styles:") or line.startswith("formatters:"):
			break

		match: re.Match[str] | None = lexer_name_re.match(line)
		if match:
			name: str | None = match.group(1)
			if name:
				lexers.append(name)
	return lexers


def group_by_prefix(lexers: list[str]) -> dict[str, list[str]]:
	"""Given a list of lexer names, return a dictionary mapping prefixes
	to lists of lexers that begin with that prefix"""
	groups: defaultdict[str, list[str]] = defaultdict(list[str])
	for name in lexers:
		prefix: str = name[0].upper()
		groups[prefix].append(name)
	# sort alphabetically
	for k in groups:
		groups[k] = sorted(groups[k], key=lambda s: s.lower())
	return dict(sorted(groups.items()))


def emit_markdown(groups: dict[str, list[str]]) -> str:
	lines: list[str] = []
	longest = 0
	for prefix, lexers in groups.items():
		joined: str = ", ".join(lexers)
		l: int = len(joined)
		if l > longest:
			longest: int = l
		lines.append(f"|   {prefix}    | {joined}")
	splitter = f"| :----: | {longest * '-'}"
	markdown: list[str] = ["| Prefix | Language", splitter]
	markdown.extend(lines)
	return "\n".join(markdown)


if __name__ == "__main__":
	if sys.stdin.isatty():
		print(
			"This script parses chroma --list piped from stdin and emits a markdown table for the README"
		)
		print("Recommended usage (from repo root):")
		print(
			"env -C cmd/chroma go run . --list | uv run _tools/format_supported_langs.py"
		)
		exit(1)

	lines: list[str] | None = sys.stdin.readlines()
	if lines:
		lexers: list[str] = parse_lexers(lines)
		groups: dict[str, list[str]] = group_by_prefix(lexers)
		print(emit_markdown(groups))
	else:
		exit(1)
