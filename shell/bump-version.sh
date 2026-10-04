#!/usr/bin/env bash
#
# Bumps the Maroto version (SemVer) across the repo, commits, tags and
# publishes a GitHub release.
#
# Usage:
#   scripts/bump-version.sh <major|minor|patch> [-y|--yes] [-n|--dry-run]
#
# Requires: git, gh (authenticated), sed.

set -euo pipefail

usage() {
  echo "Usage: $0 <major|minor|patch> [-y|--yes] [-n|--dry-run]" >&2
  exit 1
}

BUMP_TYPE=""
ASSUME_YES=false
DRY_RUN=false

for arg in "$@"; do
  case "$arg" in
    major|minor|patch)
      BUMP_TYPE="$arg"
      ;;
    -y|--yes)
      ASSUME_YES=true
      ;;
    -n|--dry-run)
      DRY_RUN=true
      ;;
    *)
      usage
      ;;
  esac
done

[[ -z "$BUMP_TYPE" ]] && usage

command -v git >/dev/null || { echo "git is required" >&2; exit 1; }
command -v gh  >/dev/null || { echo "gh (GitHub CLI) is required" >&2; exit 1; }

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

if [[ -n "$(git status --porcelain)" ]]; then
  echo "Working tree is not clean. Commit or stash your changes first." >&2
  exit 1
fi

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
DEFAULT_BRANCH="$(git remote show origin | awk '/HEAD branch/ {print $NF}')"
if [[ "$CURRENT_BRANCH" != "$DEFAULT_BRANCH" ]]; then
  echo "You are on '$CURRENT_BRANCH', not the default branch '$DEFAULT_BRANCH'." >&2
  if ! $ASSUME_YES; then
    read -rp "Continue anyway? [y/N] " reply
    [[ "$reply" =~ ^[Yy]$ ]] || exit 1
  fi
fi

git fetch origin "$DEFAULT_BRANCH" --quiet
LOCAL_SHA="$(git rev-parse "$CURRENT_BRANCH")"
REMOTE_SHA="$(git rev-parse "origin/$DEFAULT_BRANCH")"
if [[ "$CURRENT_BRANCH" == "$DEFAULT_BRANCH" && "$LOCAL_SHA" != "$REMOTE_SHA" ]]; then
  echo "Local '$DEFAULT_BRANCH' is not up to date with origin. Pull first." >&2
  exit 1
fi

# Current version = latest SemVer git tag (source of truth for releases).
CURRENT_VERSION="$(git tag --list 'v[0-9]*.[0-9]*.[0-9]*' --sort=-v:refname | head -n1)"
if [[ -z "$CURRENT_VERSION" ]]; then
  echo "No existing vMAJOR.MINOR.PATCH tag found." >&2
  exit 1
fi

if [[ ! "$CURRENT_VERSION" =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
  echo "Latest tag '$CURRENT_VERSION' is not a valid SemVer tag." >&2
  exit 1
fi
MAJOR="${BASH_REMATCH[1]}"
MINOR="${BASH_REMATCH[2]}"
PATCH="${BASH_REMATCH[3]}"

case "$BUMP_TYPE" in
  major) MAJOR=$((MAJOR + 1)); MINOR=0; PATCH=0 ;;
  minor) MINOR=$((MINOR + 1)); PATCH=0 ;;
  patch) PATCH=$((PATCH + 1)) ;;
esac
NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"

if git rev-parse "$NEW_VERSION" >/dev/null 2>&1; then
  echo "Tag $NEW_VERSION already exists." >&2
  exit 1
fi

echo "Current version: $CURRENT_VERSION"
echo "New version:     $NEW_VERSION ($BUMP_TYPE bump)"

if [[ "$BUMP_TYPE" == "major" ]]; then
  cat >&2 <<EOF

WARNING: Go modules encode the major version in the import path.
This script only replaces the version STRING in docs; it does NOT
rename the module path (currently "github.com/johnfercher/maroto/v2")
or rewrite imports. Per Go modules convention, a real v3 release would
also require migrating the module path to ".../v3" across the codebase.
EOF
fi

# Find every tracked file that mentions the current version (ignores
# untracked/gitignored files, e.g. local IDE state).
FILES=()
while IFS= read -r line; do
  FILES+=("$line")
done < <(git grep -Il -- "$CURRENT_VERSION" || true)

if [[ ${#FILES[@]} -eq 0 ]]; then
  echo "No files reference $CURRENT_VERSION." >&2
  exit 1
fi

echo
echo "Files to update:"
printf '  %s\n' "${FILES[@]}"

if $DRY_RUN; then
  echo
  echo "Dry run: no changes made."
  exit 0
fi

if ! $ASSUME_YES; then
  echo
  read -rp "Replace $CURRENT_VERSION -> $NEW_VERSION in the files above, commit, tag, push and create a GitHub release? [y/N] " reply
  [[ "$reply" =~ ^[Yy]$ ]] || { echo "Aborted."; exit 1; }
fi

for f in "${FILES[@]}"; do
  sed -i.bak "s/${CURRENT_VERSION//./\\.}/${NEW_VERSION}/g" "$f"
  rm -f "$f.bak"
done

git add "${FILES[@]}"
git commit -m "chore: bump version to ${NEW_VERSION}"
git tag -a "$NEW_VERSION" -m "$NEW_VERSION"
git push origin "$CURRENT_BRANCH"
git push origin "$NEW_VERSION"

gh release create "$NEW_VERSION" \
  --title "$NEW_VERSION" \
  --generate-notes

echo
echo "Released $NEW_VERSION: $(gh release view "$NEW_VERSION" --json url -q .url)"
