// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}