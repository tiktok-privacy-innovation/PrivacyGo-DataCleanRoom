FROM quay.io/jupyter/scipy-notebook as builder
WORKDIR /app
USER root
RUN python3 -m pip install --upgrade pip
RUN pip install build twine hatch jupyterhub==4.0.2 jupyterlab
COPY ./ ./

RUN jlpm install
RUN jlpm clean:all
RUN jlpm build
RUN python -m build

FROM quay.io/jupyter/scipy-notebook
USER root
COPY --from=builder /app/dist/jupyterlab_manatee-0.1.0-py3-none-any.whl ./
RUN pip install jupyterlab_manatee-0.1.0-py3-none-any.whl
RUN rm jupyterlab_manatee-0.1.0-py3-none-any.whl
RUN jupyter labextension disable @jupyterlab/docmanager-extension:download
RUN jupyter labextension disable @jupyterlab/filebrowser-extension:download