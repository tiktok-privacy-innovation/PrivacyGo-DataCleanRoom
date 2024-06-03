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

locals {
    dcr_k8s_namespace = var.username != "" ? "data-clean-room-${var.username}" : "data-clean-room-${var.env}"
    jupyter_k8s_namespace = var.username != "" ? "jupyterhub-${var.username}" : "jupyterhub-${var.env}"
}

resource "kubernetes_namespace" "data_clean_room_k8s_namespace" {
  metadata {
    name = local.dcr_k8s_namespace
  }

  depends_on = [ kubernetes_cluster_role_binding.cluster_admin_binding ]
}

resource "kubernetes_namespace" "jupyterhub_k8s_namespace" {
  metadata {
    name = local.jupyter_k8s_namespace
  }
}