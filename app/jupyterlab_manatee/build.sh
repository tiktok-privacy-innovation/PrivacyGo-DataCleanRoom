#!/bin/bash
# Copyright 2024 TikTok Pte. Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
set -e
VAR_FILE="../../.env"
if [ ! -f "$VAR_FILE" ]; then
    echo "Error: Variables file does not exist."
    exit 1
fi

VAR_FILE=$(realpath $VAR_FILE)
source $VAR_FILE

docker_repo="dcr-$env-images"
image_url="us-docker.pkg.dev/${project_id}/${docker_repo}/scipy-notebook-with-dcr"
if [ -z "$username" ]; then
    tag=$env
else
    tag=$username
fi

docker build --platform linux/amd64 --tag "$image_url:$tag" .
docker push "$image_url:$tag"