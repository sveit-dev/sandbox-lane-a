#!/usr/bin/env bash
# Self-test for deny-list.sh: plants a random canary word in a temporary
# repository and checks the contract. The canary is generated at run time, so
# it never appears in this repository.

set -u

script="$(cd "$(dirname "$0")" && pwd)/deny-list.sh"
work=$(mktemp -d "${TMPDIR:-/tmp}/deny-list-test.XXXXXX") || exit 1
trap 'rm -rf "$work"' EXIT

canary="canary$(od -An -N6 -tx1 /dev/urandom | tr -d ' \n')"
upper=$(printf '%s' "$canary" | tr '[:lower:]' '[:upper:]')
failures=0

check() {
  if [ "$1" = ok ]; then
    echo "ok   $2"
  else
    echo "FAIL $2"
    failures=$((failures + 1))
  fi
}

run() {
  # Runs the script in the temporary repository; sets out and rc.
  out=$(cd "$work/repo" && DENY_WORDS="$1" bash "$script" 2>&1)
  rc=$?
}

mkdir -p "$work/repo/docs"
(
  cd "$work/repo" || exit 1
  git init -q
  printf 'first line\nsecond line mentions %s here\n' "$upper" > docs/note.md
  printf 'clean file\n' > clean.txt
  printf 'untracked %s\n' "$canary" > untracked.txt
  git add docs/note.md clean.txt
)

run "other,$canary"
[ $rc -eq 1 ] && check ok "exit 1 on a hit" || check fail "exit 1 on a hit (got $rc)"
printf '%s\n' "$out" | grep -qx 'docs/note.md:2' && check ok "prints path:line" || check fail "prints path:line"
printf '%s\n' "$out" | grep -qi "$canary" && check fail "output carries the canary" || check ok "output carries no canary"
printf '%s\n' "$out" | grep -q 'untracked.txt' && check fail "scans untracked files" || check ok "skips untracked files"

run " $canary , "
[ $rc -eq 1 ] && check ok "trims blanks and empty entries" || check fail "trims blanks and empty entries (got $rc)"

run "absentword"
[ $rc -eq 0 ] && check ok "exit 0 when clean" || check fail "exit 0 when clean (got $rc)"

run ""
[ $rc -eq 2 ] && check ok "exit 2 when empty" || check fail "exit 2 when empty (got $rc)"
out=$(cd "$work/repo" && env -u DENY_WORDS bash "$script" 2>&1)
rc=$?
[ $rc -eq 2 ] && check ok "exit 2 when unset" || check fail "exit 2 when unset (got $rc)"

(cd "$work/repo" && printf 'x\n' > "$canary-file.txt" && git add "$canary-file.txt")
run "$canary"
printf '%s\n' "$out" | grep -qx '(path withheld):0' && check ok "withholds a path containing the word" || check fail "withholds a path containing the word"
printf '%s\n' "$out" | grep -qi "$canary" && check fail "output carries the canary in a path" || check ok "output carries no canary in a path"

if [ $failures -ne 0 ]; then
  echo "deny-list-test: $failures failed"
  exit 1
fi
echo "deny-list-test: passed"
