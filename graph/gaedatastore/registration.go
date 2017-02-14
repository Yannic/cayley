// Copyright 2014 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gaedatastore

import (
	"github.com/cayleygraph/cayley/graph"
)

const QuadStoreType = "gaedatastore"

func init() {
	graph.RegisterQuadStore(QuadStoreType, &QuadStoreRegistration{})
}

type QuadStoreRegistration struct{}

func (q *QuadStoreRegistration) New(_ string, options graph.Options) (graph.QuadStore, error) {
	return newQuadStore("", options)
}

func (q *QuadStoreRegistration) NewForRequest(qs graph.QuadStore, options graph.Options) (graph.QuadStore, error) {
	return newQuadStoreForRequest(qs, options)
}

func (q *QuadStoreRegistration) Init(_ string, opts graph.Options) error {
	return initQuadStore("", opts)
}

func (q *QuadStoreRegistration) Upgrade(string, graph.Options) error {
	return graph.ErrOperationNotSupported
}

func (q *QuadStoreRegistration) IsPersistent() bool {
	return true
}
