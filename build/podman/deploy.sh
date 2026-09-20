#!/usr/bin/env bash
set -euo pipefail

QUADLET_DIR="/etc/containers/systemd"
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

: "${ALIYUN_REGISTRY:?ALIYUN_REGISTRY is required}"
: "${ALIYUN_NAMESPACE:?ALIYUN_NAMESPACE is required}"
: "${ALIYUN_USERNAME:?ALIYUN_USERNAME is required}"
: "${ALIYUN_PASSWORD:?ALIYUN_PASSWORD is required}"
: "${IMG_TAG:?IMG_TAG is required}"

echo "===== Deploy Info ====="
echo "Registry : ${ALIYUN_REGISTRY}"
echo "Namespace: ${ALIYUN_NAMESPACE}"
echo "Tag      : ${IMG_TAG}"
echo

mkdir -p "${QUADLET_DIR}"

echo "===== Login Registry ====="

printf '%s' "${ALIYUN_PASSWORD}" |
  podman login "${ALIYUN_REGISTRY}" \
    -u "${ALIYUN_USERNAME}" \
    --password-stdin

echo
echo "===== Deploy Network ====="

cp "${BASE_DIR}/app.network" \
   "${QUADLET_DIR}/app.network"

deploy_service() {
  local file="$1"
  local type="$2"

  local filename
  local service
  local name
  local image

  filename="$(basename "${file}")"
  service="${filename%.container}"
  name="${service%-${type}}"

  image="${ALIYUN_REGISTRY}/${ALIYUN_NAMESPACE}/${name}-${type}:${IMG_TAG}"

  echo
  echo "Deploying ${service}"
  echo "Image: ${image}"

  podman pull "${image}"

  sed \
    "s|^Image=.*|Image=${image}|" \
    "${file}" \
    > "${QUADLET_DIR}/${filename}"
}

echo
echo "===== Prepare RPC ====="

for file in "${BASE_DIR}"/*-rpc.container; do
  [ -f "${file}" ] || continue
  deploy_service "${file}" "rpc"
done

echo
echo "===== Prepare API ====="

for file in "${BASE_DIR}"/*-api.container; do
  [ -f "${file}" ] || continue
  deploy_service "${file}" "api"
done

echo
echo "===== Reload systemd ====="

systemctl daemon-reload

echo
echo "===== Restart RPC ====="

for file in "${BASE_DIR}"/*-rpc.container; do
  [ -f "${file}" ] || continue

  service="$(basename "${file}" .container)"

  echo "Restarting ${service}"
  systemctl restart "${service}.service"
done

echo
echo "===== Restart API ====="

for file in "${BASE_DIR}"/*-api.container; do
  [ -f "${file}" ] || continue

  service="$(basename "${file}" .container)"

  echo "Restarting ${service}"
  systemctl restart "${service}.service"
done

echo
echo "===== Podman Containers ====="

podman ps

echo
echo "===== Deployment Complete ====="