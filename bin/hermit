#!/bin/bash

set -eo pipefail

if [ -z "${HERMIT_STATE_DIR}" ]; then
  case "$(uname -s)" in
  Darwin)
    export HERMIT_STATE_DIR="${HOME}/Library/Caches/hermit"
    ;;
  Linux)
    export HERMIT_STATE_DIR="${XDG_CACHE_HOME:-${HOME}/.cache}/hermit"
    ;;
  esac
fi

export HERMIT_DIST_URL="${HERMIT_DIST_URL:-https://github.com/cashapp/hermit/releases/download/stable}"
HERMIT_CHANNEL="$(basename "${HERMIT_DIST_URL}")"
export HERMIT_CHANNEL
export HERMIT_EXE=${HERMIT_EXE:-${HERMIT_STATE_DIR}/pkg/hermit@${HERMIT_CHANNEL}/hermit}

if [ ! -x "${HERMIT_EXE}" ]; then
  echo "Bootstrapping ${HERMIT_EXE} from ${HERMIT_DIST_URL}" 1>&2
  curl -fsSL "${HERMIT_DIST_URL}/install.sh" | /bin/bash 1>&2
fi

exec "${HERMIT_EXE}" --level=fatal exec "$0" -- "$@"
