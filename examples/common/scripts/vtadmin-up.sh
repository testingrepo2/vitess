#!/bin/bash

<<<<<<< HEAD
=======
# Copyright 2023 The Vitess Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

function output() {
  echo -e "$@"
}

>>>>>>> 7603abc534 (Update VTAdmin build script (#15839))
script_dir="$(dirname "${BASH_SOURCE[0]:-$0}")"
source "${script_dir}/../env.sh"

cluster_name="local"
log_dir="${VTDATAROOT}/tmp"
web_dir="${script_dir}/../../../web/vtadmin"

vtadmin_api_port=14200
vtadmin_web_port=14201

<<<<<<< HEAD
=======
case_insensitive_hostname=$(echo "$hostname" | tr '[:upper:]' '[:lower:]')

output "\n\033[1;32mvtadmin-api expects vtadmin-web at, and set http-origin to \"http://${case_insensitive_hostname}:${vtadmin_web_port}\"\033[0m"

>>>>>>> 7603abc534 (Update VTAdmin build script (#15839))
vtadmin \
  --addr "${case_insensitive_hostname}:${vtadmin_api_port}" \
  --http-origin "http://${case_insensitive_hostname}:${vtadmin_web_port}" \
  --http-tablet-url-tmpl "http://{{ .Tablet.Hostname }}:15{{ .Tablet.Alias.Uid }}" \
  --tracer "opentracing-jaeger" \
  --grpc-tracing \
  --http-tracing \
  --logtostderr \
  --alsologtostderr \
  --rbac \
  --rbac-config="${script_dir}/../vtadmin/rbac.yaml" \
  --cluster "id=${cluster_name},name=${cluster_name},discovery=staticfile,discovery-staticfile-path=${script_dir}/../vtadmin/discovery.json,tablet-fqdn-tmpl=http://{{ .Tablet.Hostname }}:15{{ .Tablet.Alias.Uid }}" \
  > "${log_dir}/vtadmin-api.out" 2>&1 &

vtadmin_api_pid=$!
echo ${vtadmin_api_pid} > "${log_dir}/vtadmin-api.pid"

echo "\
vtadmin-api is running!
  - API: http://${case_insensitive_hostname}:${vtadmin_api_port}
  - Logs: ${log_dir}/vtadmin-api.out
  - PID: ${vtadmin_api_pid}
"

# Wait for vtadmin to successfully discover the cluster
expected_cluster_result="{\"result\":{\"clusters\":[{\"id\":\"${cluster_name}\",\"name\":\"${cluster_name}\"}]},\"ok\":true}"
<<<<<<< HEAD
for _ in {0..300}; do
  result=$(curl -s "http://${hostname}:${vtadmin_api_port}/api/clusters")
=======
for _ in {0..100}; do
  result=$(curl -s "http://${case_insensitive_hostname}:${vtadmin_api_port}/api/clusters")
>>>>>>> 7603abc534 (Update VTAdmin build script (#15839))
  if [[ ${result} == "${expected_cluster_result}" ]]; then
    break
  fi
  sleep 0.1
done

# Check one last time
[[ $(curl -s "http://${case_insensitive_hostname}:${vtadmin_api_port}/api/clusters") == "${expected_cluster_result}" ]] || fail "vtadmin failed to discover the running example Vitess cluster."

# Download nvm and node
if [[ -z ${NVM_DIR} ]]; then
    export NVM_DIR="$HOME/.nvm"
fi

if [[ -z ${NODE_VERSION} ]]; then
    export NODE_VERSION="16"
fi

output "\nInstalling nvm...\n"

if [ -d "$NVM_DIR" ]; then
  output "\033[1;32mnvm is already installed!\033[0m"
else
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash && output "\033[1;32mnvm is installed!\033[0m" || fail "\033[1;32mnvm failed to install!\033[0m" 
fi

source "$NVM_DIR/nvm.sh"

output "\nConfiguring Node.js $NODE_VERSION\n"
nvm install "$NODE_VERSION" || fail "Could not install and use nvm $NODE_VERSION."

# As a TODO, it'd be nice to make the assumption that vtadmin-web is already
# installed and built (since we assume that `make` has already been run for
# other Vitess components.)
npm --prefix "$web_dir" --silent install

VITE_VTADMIN_API_ADDRESS="http://${hostname}:${vtadmin_api_port}" \
  VITE_ENABLE_EXPERIMENTAL_TABLET_DEBUG_VARS="true" \
  npm run --prefix "$web_dir" build

"${web_dir}/node_modules/.bin/serve" --no-clipboard -l $vtadmin_web_port -s "${web_dir}/build" \
  > "${log_dir}/vtadmin-web.out" 2>&1 &

vtadmin_web_pid=$!
echo ${vtadmin_web_pid} > "${log_dir}/vtadmin-web.pid"

echo "\
vtadmin-web is running!
  - Browser: http://${case_insensitive_hostname}:${vtadmin_web_port}
  - Logs: ${log_dir}/vtadmin-web.out
  - PID: ${vtadmin_web_pid}
"
