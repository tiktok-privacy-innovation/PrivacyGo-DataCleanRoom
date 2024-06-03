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

import jupyter_server
from jupyter_server.utils import url_path_join
from ._version import __version__
from .handlers import *

def _jupyter_server_extension_points():
    return [{
        'module': 'jupyterlab_manatee'
    }]

def _jupyter_labextension_paths():
    return [{
        "src": "labextension",
        "dest": "jupyterlab_manatee"
    }]


def _load_jupyter_server_extension(serverapp: jupyter_server.serverapp.ServerApp):
    """
    Called when the extension is loaded.
    """

    web_app = serverapp.web_app
    base_url = web_app.settings['base_url']
    handlers = [
        (url_path_join(base_url, 'manatee', 'jobs'), DataCleanRoomJobHandler), 
        (url_path_join(base_url, 'manatee', 'output'), DataCleanRoomOutputHandler), (url_path_join(base_url, 'manatee', 'attestation'), DataCleanRoomAttestationHandler),
    ]
    web_app.add_handlers('.*$', handlers)