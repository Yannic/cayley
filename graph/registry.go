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

package graph

import (
	"fmt"
	"reflect"
)

var (
	ErrQuadStoreNotRegistred = fmt.Errorf("This QuadStore is not registered.")
	ErrOperationNotSupported = fmt.Errorf("This Operation is not supported.")
)

var storeRegistry = make(map[string]QuadStoreRegistration)

// TODO: better documentation
type QuadStoreRegistration interface {
	// Create new QuadStore
	// Must return ErrOperationNotSupported if operation is not supported
	New(string, Options) (QuadStore, error)

	// Create new QuadStore for request
	// Must return ErrOperationNotSupported if operation is not supported
	NewForRequest(QuadStore, Options) (QuadStore, error)

	// Upgrade existing QuadStore
	// Must return ErrOperationNotSupported if operation is not supported
	Upgrade(string, Options) error

	// Initialize new QuadStore
	// Must return ErrOperationNotSupported if operation is not supported
	Init(string, Options) error

	// QuadStore is persistent
	// Must return ErrOperationNotSupported if operation is not supported
	IsPersistent() bool
}

func RegisterQuadStore(name string, register QuadStoreRegistration) {
	// Register QuadStore with friendly name
	if _, found := storeRegistry[name]; found {
		panic(fmt.Sprintf("Already registered QuadStore \"%s\".", name))
	}
	storeRegistry[name] = register

	// Also Register QuadStore with pkgPath
	var pkgPath string
	_type := reflect.TypeOf(register)
	if _type.Kind() == reflect.Ptr {
		pkgPath = _type.Elem().PkgPath()
	} else {
		pkgPath = _type.PkgPath()
	}

	if _, found := storeRegistry[pkgPath]; found {
		panic(fmt.Sprintf("Already registered QuadStore \"%s\".", pkgPath))
	}
	storeRegistry[pkgPath] = register
}

func NewQuadStore(name string, dbpath string, opts Options) (QuadStore, error) {
	r, registered := storeRegistry[name]
	if !registered {
		return nil, ErrQuadStoreNotRegistred
	}

	return r.New(dbpath, opts)
}

func NewQuadStoreForRequest(qs QuadStore, opts Options) (QuadStore, error) {
	pkgPath := reflect.TypeOf(qs).PkgPath()
	r, registered := storeRegistry[pkgPath]
	if !registered {
		return nil, ErrQuadStoreNotRegistred
	}

	return r.NewForRequest(qs, opts)
}

func UpgradeQuadStore(name string, dbpath string, opts Options) error {
	r, registered := storeRegistry[name]
	if !registered {
		return ErrQuadStoreNotRegistred
	}

	return r.Upgrade(dbpath, opts)
}

func InitQuadStore(name string, dbpath string, opts Options) error {
	r, registered := storeRegistry[name]
	if !registered {
		return ErrQuadStoreNotRegistred
	}

	return r.Init(dbpath, opts)
}

func IsPersistent(name string) bool {
	r, registered := storeRegistry[name]
	if !registered {
		return false
	}

	return r.IsPersistent()
}

func QuadStores() []string {
	t := make([]string, 0, len(storeRegistry))
	for n := range storeRegistry {
		t = append(t, n)
	}
	return t
}
