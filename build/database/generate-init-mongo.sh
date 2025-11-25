#!/usr/bin/env bash
set -euo pipefail


TEMPLATE_DIR="/docker-entrypoint-initdb.d"
TEMPLATE_FILE="$TEMPLATE_DIR/init-mongo.js.template"
OUT_FILE="$TEMPLATE_DIR/init-mongo.js"

echo "[mongo-init] running generator to produce $OUT_FILE from template"

if [ ! -f "$TEMPLATE_FILE" ]; then
  echo "[mongo-init] template not found: $TEMPLATE_FILE" >&2
  exit 0
fi

if [ -z "${DB_NAME:-}" ] || [ -z "${DB_USER:-}" ] || [ -z "${DB_PASSWORD:-}" ]; then
  echo "[mongo-init] WARNING: one of DB_NAME, APP_DB_USER or APP_DB_PASSWORD is empty or unset"
  echo "[mongo-init] The template will not be rendered and no app user will be created unless these are set."
  exit 0
fi

escape_replacement() {
  printf '%s' "$1" | sed -e 's/[\/&]/\\&/g'
}

DB_NAME_R=$(escape_replacement "${DB_NAME}")
APP_DB_USER_R=$(escape_replacement "${APP_DB_USER}")
APP_DB_PASSWORD_R=$(escape_replacement "${APP_DB_PASSWORD}")

sed \
  -e "s|\$\{DB_NAME\}|$DB_NAME_R|g" \
  -e "s|\$\{APP_DB_USER\}|$APP_DB_USER_R|g" \
  -e "s|\$\{APP_DB_PASSWORD\}|$APP_DB_PASSWORD_R|g" \
  "$TEMPLATE_FILE" > "$OUT_FILE"

chmod 644 "$OUT_FILE"

echo "[mongo-init] generated $OUT_FILE"

exit 0
