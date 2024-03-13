#!/usr/bin/env bash

set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

TAGGER="/tmp/buildtagger"
REPO_ROOT="."
EXTRA_BUILDTAGGER_ARGS="${EXTRA_BUILDTAGGER_ARGS:-}"
RESTORE_DEEPCOPY_TAGS="${RESTORE_DEEPCOPY_TAGS:-false}"

if [[ ! -f "${TAGGER}" ]]; then
  curl -fsSL  "${BUILDTAGGER_DOWNLOAD_URL}" -o "${TAGGER}"
  chmod +x "${TAGGER}"
fi

"${TAGGER}" --parent-dir "${REPO_ROOT}"/apis --regex "(.+)/.+/.+\.go" --tag-format "(%s || all) && !ignore_autogenerated" --mode dir ${EXTRA_BUILDTAGGER_ARGS}
"${TAGGER}" --parent-dir "${REPO_ROOT}"/apis --regex "(.+)/.+/zz_groupversion_info\.go" --tag-format "(%s || all) && !ignore_autogenerated" --mode dir ${EXTRA_BUILDTAGGER_ARGS}
"${TAGGER}" --parent-dir "${REPO_ROOT}"/internal/controller --regex "(.+)/.+/zz_controller\.go" --tag-format "(%s || all) && !ignore_autogenerated" --mode dir ${EXTRA_BUILDTAGGER_ARGS}
"${TAGGER}" --parent-dir "${REPO_ROOT}"/internal/controller --regex "zz_(.+)_setup\.go" --tag-format "(%s || all) && !ignore_autogenerated" --mode file ${EXTRA_BUILDTAGGER_ARGS}
"${TAGGER}" --parent-dir "${REPO_ROOT}"/cmd/provider --regex "(.+)/zz_main\.go" --tag-format "(%s || all) && !ignore_autogenerated" --mode dir ${EXTRA_BUILDTAGGER_ARGS}
"${TAGGER}" --parent-dir "${REPO_ROOT}"/config --regex "(.+)/config\.go" --tag-format "(%s || all) && !ignore_autogenerated" --mode dir ${EXTRA_BUILDTAGGER_ARGS}

# constant tags
# apis/zz_register.go -> (apiregister || register || all) && !ignore_autogenerated
"${TAGGER}" --parent-dir "${REPO_ROOT}"/apis/zz_register.go --tag-format "all && !ignore_autogenerated" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# cmd/generator/main.go -> config || generate || all
"${TAGGER}" --parent-dir "${REPO_ROOT}"/cmd/generator/main.go --tag-format "all" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# config/autoscaling/config_test.go -> (autoscaling || config || all) && !ignore_autogenerated
"${TAGGER}" --parent-dir "${REPO_ROOT}"/config/autoscaling/config_test.go --tag-format "(autoscaling || all) && !ignore_autogenerated" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# config/common/apis/lambda/extractor.go -> config || lambda || all
"${TAGGER}" --parent-dir "${REPO_ROOT}"/config/common/apis/lambda/extractor.go --tag-format "lambda || all" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# config/config_registry.go -> configregistry || register || config || all
"${TAGGER}" --parent-dir "${REPO_ROOT}"/config/registry.go --tag-format "configregistry || all" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# config/provider.go -> configprovider || register || config || all
"${TAGGER}" --parent-dir "${REPO_ROOT}"/config/provider.go --tag-format "(configprovider || all) && !linter_run" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# config/overrides.go -> configprovider || register || config || all
"${TAGGER}" --parent-dir "${REPO_ROOT}"/config/overrides.go --tag-format "configprovider || all" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# internal/controller/eks/clusterauth/controller.go -> eks || all
"${TAGGER}" --parent-dir "${REPO_ROOT}"/internal/controller/eks/clusterauth/controller.go --tag-format "eks || all" --mode file ${EXTRA_BUILDTAGGER_ARGS}
# internal/controller/eks/clusterauth/eks.go -> eks || all
"${TAGGER}" --parent-dir "${REPO_ROOT}"/internal/controller/eks/clusterauth/eks.go --tag-format "eks || all" --mode file ${EXTRA_BUILDTAGGER_ARGS}

if [[ "${RESTORE_DEEPCOPY_TAGS}" == "true" ]]; then
  "${TAGGER}" --parent-dir "${REPO_ROOT}"/apis --regex "zz_generated.deepcopy.go" --tag-format "!ignore_autogenerated" --mode file
fi