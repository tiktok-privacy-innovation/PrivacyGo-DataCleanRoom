/**
 * Copyright 2024 TikTok Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Contents } from '@jupyterlab/services';
import { IDocumentManager } from '@jupyterlab/docmanager';
import { PanelWithToolbar, ToolbarButton, fileUploadIcon, } from '@jupyterlab/ui-components';
import { filter } from '@lumino/algorithm';
import { ITranslator, nullTranslator } from '@jupyterlab/translation';
import { FileBrowser, FilterFileBrowserModel } from '@jupyterlab/filebrowser';
import { showDialog, Dialog } from '@jupyterlab/apputils';
import { ServerConnection } from '@jupyterlab/services';

/*
This class overrides items() to make the filebrowser list only ipynb files.
We're doing this soley for the demo purpose, and the actual product may have additional files
(e.g., local python modules)
*/
class NotebookOnlyFilterFileBrowserModel extends FilterFileBrowserModel {
    override items(): IterableIterator<Contents.IModel> {
        return filter(super.items(), value => {
            if (value.type === 'notebook') {
                return true;
            } else {
                return false;
            }
        });
    }
}

export class DataCleanRoomSources extends PanelWithToolbar {
    constructor(options: DataCleanRoomSources.IOptions) {
        super();
        const { manager } = options;
        this._manager = manager;
        const trans = (options.translator ?? nullTranslator).load('jupyterlab');
        this.title.label = trans.__('Sources');

        const fbModel = new NotebookOnlyFilterFileBrowserModel({
            manager: manager,
        });
        this._browser = new FileBrowser({
            id: 'jupyterlab_manatee:plugin:sources',
            model: fbModel
        });
        this.toolbar.addItem(
            'submit',
            new ToolbarButton({
                icon: fileUploadIcon,
                onClick: () => this.sendSelectedFilesToAPI(),
                tooltip: trans.__('Submit Job to Data Clean Room')
            })
        );

        this.addWidget(this._browser);
    };

    async sendSelectedFilesToAPI() {
        for (const item of this._browser.selectedItems()) {
            const result = await showDialog({
                title: "Submitting a Job to Data Clean Room?",
                body: 'Path: ' + item.path,
                buttons: [Dialog.okButton(), Dialog.cancelButton()]
            });
            
            if (result.button.accept) {
                const file = await this._manager.services.contents.get(item.path);
                // Prepare data
                const data = JSON.stringify({
                    path: item.path,
                    filename: file.name
                });

                const settings = ServerConnection.makeSettings();

                console.log("Sending... %s", settings.baseUrl);
                ServerConnection.makeRequest(settings.baseUrl + "manatee/jobs", {
                    body: data, method: "POST"
                }, settings).then(response => {
                    if (response.status !== 200) {
                        console.log("Error has occured!");
                    }
                    response.body?.getReader().read().then(({done, value}) => {
                        if (done) {
                            console.log("stream is closed");
                            return;
                        }
                        let decoder = new TextDecoder('utf-8');
                        console.log("value:", decoder.decode(value));
                    });
                });
            }
        }
    }

    protected _manager : IDocumentManager;
    protected _browser : FileBrowser;
}

export namespace DataCleanRoomSources {
    export interface IOptions {
        manager: IDocumentManager;
        translator?: ITranslator;
    }
}
