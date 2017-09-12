# Copyright 2017 The Cayley Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    commit = "65df68f4f64e9c59eb571290eb86bf07766393b6",
)
load("@io_bazel_rules_docker//docker:docker.bzl", "docker_repositories")

docker_repositories()

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "572e3179a26eeabe38c9b29163f7f53fc7da0f00",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

go_rules_dependencies()
go_register_toolchains()


go_repository(
  name = "com_github_badgerodon_peg",
  importpath = "github.com/badgerodon/peg",
  commit = "9e5f7f4d07ca576562618c23e8abadda278b684f",
)

go_repository(
  name = "com_github_boltdb_bolt",
  importpath = "github.com/boltdb/bolt",
  tag = "v1.3.1",
)

go_repository(
  name = "com_github_cznic_mathutil",
  importpath = "github.com/cznic/mathutil",
  commit = "91af0ce59d1753e5eeaa6a3fd8f7aeeba01f1210",
)

go_repository(
  name = "com_github_davecgh_go_spew",
  importpath = "github.com/davecgh/go-spew",
  commit = "a476722483882dd40b8111f0eb64e1d7f43f56e4",
)

go_repository(
  name = "com_github_dennwc_graphql",
  importpath = "github.com/dennwc/graphql",
  commit = "5cf1f05cb1945ca40ca374e21ccf26f01fc67d10",
)

go_repository(
  name = "com_github_dlclark_regexp2",
  importpath = "github.com/dlclark/regexp2",
  tag = "v1.1.6",
)

go_repository(
  name = "com_github_dop251_goja",
  importpath = "github.com/dop251/goja",
  commit = "0aeff7545edfeafa4535f10f7fa4970c230915e8",
)

go_repository(
  name = "com_github_fsnotify_fsnotify",
  importpath = "github.com/fsnotify/fsnotify",
  commit = "4da3e2cfbabc9f751898f250b49f2439785783a1",
)

go_repository(
  name = "com_github_go_sql_driver_mysql",
  importpath = "github.com/go-sql-driver/mysql",
  commit = "26471af196a17ee75a22e6481b5a5897fb16b081",
)

go_repository(
  name = "com_github_gogo_protobuf",
  importpath = "github.com/gogo/protobuf",
  tag = "v0.4",
)

go_repository(
  name = "com_github_golang_snappy",
  importpath = "github.com/golang/snappy",
  commit = "553a641470496b2327abcac10b36396bd98e45c9",
)

go_repository(
  name = "com_github_hashicorp_hcl",
  importpath = "github.com/hashicorp/hcl",
  commit = "8f6b1344a92ff8877cf24a5de9177bf7d0a2a187",
)

go_repository(
  name = "com_github_julienschmidt_httprouter",
  importpath = "github.com/julienschmidt/httprouter",
  commit = "975b5c4c7c21c0e3d2764200bf2aa8e34657ae6e",
)

go_repository(
  name = "com_github_lib_pq",
  importpath = "github.com/lib/pq",
  commit = "e42267488fe361b9dc034be7a6bffef5b195bceb",
)

go_repository(
  name = "com_github_linkeddata_gojsonld",
  importpath = "github.com/linkeddata/gojsonld",
  commit = "4f5db6791326b8962ede4edbba693edcf20fd1ad",
)

go_repository(
  name = "com_github_linkeddata_gojsonld",
  importpath = "github.com/linkeddata/gojsonld",
  commit = "4f5db6791326b8962ede4edbba693edcf20fd1ad",
)

go_repository(
  name = "com_github_magiconair_properties",
  importpath = "github.com/magiconair/properties",
  tag = "v1.7.3",
)

go_repository(
  name = "com_github_mitchellh_mapstructure",
  importpath = "github.com/mitchellh/mapstructure",
  commit = "d0303fe809921458f417bcf828397a65db30a7e4",
)

go_repository(
  name = "com_github_pborman_uuid",
  importpath = "github.com/pborman/uuid",
  tag = "v1.1",
)

go_repository(
  name = "com_github_pelletier_go_buffruneio",
  importpath = "github.com/pelletier/go-buffruneio",
  tag = "v0.2.0",
)

go_repository(
  name = "com_github_pelletier_go_toml",
  importpath = "github.com/pelletier/go-toml",
  tag = "v1.0.0",
)

go_repository(
  name = "com_github_peterh_liner",
  importpath = "github.com/peterh/liner",
  commit = "a37ad39843113264dae84a5d89fcee28f50b35c6",
)

go_repository(
  name = "com_github_pmezard_go_difflib",
  importpath = "github.com/pmezard/go-difflib",
  tag = "v1.0.0",
)

go_repository(
  name = "com_github_russross_blackfriday",
  importpath = "github.com/russross/blackfriday",
  commit = "4048872b16cc0fc2c5fd9eacf0ed2c2fedaa0c8c",
)

go_repository(
  name = "com_github_shurcooL_sanitized_anchor_name",
  importpath = "github.com/shurcooL/sanitized_anchor_name",
  commit = "541ff5ee47f1dddf6a5281af78307d921524bcb5",
)

go_repository(
  name = "com_github_spf13_afero",
  importpath = "github.com/spf13/afero",
  commit = "ee1bd8ee15a1306d1f9201acc41ef39cd9f99a1b",
)

go_repository(
  name = "com_github_spf13_cast",
  importpath = "github.com/spf13/cast",
  tag = "v1.1.0",
)

go_repository(
  name = "com_github_spf13_cobra",
  importpath = "github.com/spf13/cobra",
  commit = "b78744579491c1ceeaaa3b40205e56b0591b93a3",
)

go_repository(
  name = "com_github_spf13_jwalterweatherman",
  importpath = "github.com/spf13/jwalterweatherman",
  commit = "12bd96e66386c1960ab0f74ced1362f66f552f7b",
)

go_repository(
  name = "com_github_spf13_pflag",
  importpath = "github.com/spf13/pflag",
  commit = "7aff26db30c1be810f9de5038ec5ef96ac41fd7c",
)

go_repository(
  name = "com_github_spf13_viper",
  importpath = "github.com/spf13/viper",
  tag = "v1.0.0",
)

go_repository(
  name = "com_github_stretchr_testify",
  importpath = "github.com/stretchr/testify",
  commit = "890a5c3458b43e6104ff5da8dfa139d013d77544",
)

go_repository(
  name = "com_github_syndtr_goleveldb",
  importpath = "github.com/syndtr/goleveldb",
  commit = "b89cc31ef7977104127d34c1bd31ebd1a9db2199",
)

go_repository(
  name = "com_github_tylertreat_BoomFilters",
  importpath = "github.com/tylertreat/BoomFilters",
  commit = "894fc0d07ef71050281111ca3953677cfe84e0df",
)

go_repository(
  name = "in_gopkg_mgo_v2",
  importpath = "gopkg.in/mgo.v2",
  tag = "v2",
)

go_repository(
  name = "in_gopkg_yaml_v2",
  importpath = "gopkg.in/yaml.v2",
  tag = "v2",
)

go_repository(
  name = "org_golang_google_appengine",
  importpath = "google.golang.org/appengine",
  commit = "d9a072cfa7b9736e44311ef77b3e09d804bfa599",
)

go_repository(
  name = "org_golang_x_sys",
  importpath = "golang.org/x/sys",
  commit = "7a85bfad8bb9558204b9cc9b3624680d2d443a35",
)

go_repository(
  name = "org_golang_x_text",
  importpath = "golang.org/x/text",
  commit = "bd91bbf73e9a4a801adbfb97133c992678533126",
)
