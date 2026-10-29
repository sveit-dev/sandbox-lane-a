#!/usr/bin/env bash
# Deny-list check, a contract shared with sveit-protocols.
#
# Reads DENY_WORDS (comma-separated, lower-case) from the environment and
# scans every file tracked in the current git repository for any of them,
# case-insensitively, as fixed strings. Prints only path:line per hit, never
# the word or the line. A tracked path that itself contains a word is printed
# as "(path withheld):0".
#
# Exit codes: 0 clean, 1 at least one hit, 2 DENY_WORDS unset or empty, or a
# file that could not be read.
#
# Runs on macOS bash 3.2 and Ubuntu. Never enable tracing here: the list must
# not reach the log.

set +x
set -u

if [ -z "${DENY_WORDS:-}" ]; then
  echo "deny-list: DENY_WORDS is unset or empty" >&2
  exit 2
fi

root=$(git rev-parse --show-toplevel 2>/dev/null) || {
  echo "deny-list: not inside a git repository" >&2
  exit 2
}
cd "$root" || exit 2

patterns=$(mktemp "${TMPDIR:-/tmp}/deny-list.XXXXXX") || exit 2
trap 'rm -f "$patterns"' EXIT
chmod 600 "$patterns"

# One word per line, surrounding blanks trimmed, empty entries dropped. The
# words go through a file, never through argv.
printf '%s\n' "$DENY_WORDS" | tr ',' '\n' \
  | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//' \
  | grep -v '^$' > "$patterns"
if [ ! -s "$patterns" ]; then
  echo "deny-list: DENY_WORDS is unset or empty" >&2
  exit 2
fi

status=0
while IFS= read -r -d '' path; do
  if printf '%s\n' "$path" | grep -q -i -F -f "$patterns"; then
    echo "(path withheld):0"
    status=1
    continue
  fi
  # Symlinks and submodules are listed but have no content of their own here.
  if [ -L "$path" ] || [ ! -f "$path" ]; then
    continue
  fi
  lines=$(grep -a -n -i -F -f "$patterns" -- "$path" 2>/dev/null)
  rc=$?
  if [ $rc -eq 0 ]; then
    printf '%s\n' "$lines" | cut -d: -f1 | while IFS= read -r line; do
      echo "$path:$line"
    done
    status=1
  elif [ $rc -gt 1 ]; then
    echo "deny-list: $path: could not be read" >&2
    [ $status -eq 0 ] && status=2
  fi
done < <(git ls-files -z)

if [ $status -eq 0 ]; then
  echo "deny-list: clean" >&2
fi
exit $status
