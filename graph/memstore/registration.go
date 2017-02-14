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

package memstore

import (
	"github.com/cayleygraph/cayley/graph"
)

const QuadStoreType = "memstore"

func init() {
	graph.RegisterQuadStore(QuadStoreType, &QuadStoreRegistration{})
}

type QuadStoreRegistration struct{}

func (q *QuadStoreRegistration) New(path string, options graph.Options) (graph.QuadStore, error) {
	return newQuadStore(), nil
}

func (q *QuadStoreRegistration) NewForRequest(graph.QuadStore, graph.Options) (graph.QuadStore, error) {
	return nil, graph.ErrOperationNotSupported
}

func (q *QuadStoreRegistration) Init(path string, opts graph.Options) error {
	return graph.ErrOperationNotSupported
}

func (q *QuadStoreRegistration) Upgrade(path string, opts graph.Options) error {
	return graph.ErrOperationNotSupported
}

func (q *QuadStoreRegistration) IsPersistent() bool {
	return false
}
