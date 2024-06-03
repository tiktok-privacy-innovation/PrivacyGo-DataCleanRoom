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
deploy_service() {
    app=$1
    pushd $app
    ./deploy.sh
    popd
}

if [[ $# == 0 ]]; then
    deploy_service data-clean-room
    deploy_service jupyterhub
elif [[ $# == 1 ]]; then 
    deploy_service $1
else 
    echo "ERROR: unkown parameters"
fi
