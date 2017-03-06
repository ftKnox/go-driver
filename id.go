//
// DISCLAIMER
//
// Copyright 2017 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package driver

import (
	"fmt"
	"strings"
)

// DocumentID references a document in a collection.
// Format: collection/_key
type DocumentID string

// Validate validates the given id.
func (id DocumentID) Validate() error {
	if id == "" {
		return WithStack(fmt.Errorf("DocumentID is empty"))
	}
	parts := strings.Split(string(id), "/")
	if len(parts) != 2 {
		return WithStack(fmt.Errorf("Expected 'collection/key', got '%s'", string(id)))
	}
	if parts[0] == "" {
		return WithStack(fmt.Errorf("Collection part of '%s' is empty", string(id)))
	}
	if parts[1] == "" {
		return WithStack(fmt.Errorf("Key part of '%s' is empty", string(id)))
	}
	return nil
}

// Collection returns the collection part of the ID.
func (id DocumentID) Collection() string {
	parts := strings.Split(string(id), "/")
	return pathUnescape(parts[0])
}

// Key returns the key part of the ID.
func (id DocumentID) Key() string {
	parts := strings.Split(string(id), "/")
	if len(parts) == 2 {
		return pathUnescape(parts[1])
	}
	return ""
}

// NewDocumentID creates a new document ID from the given collection, key pair.
func NewDocumentID(collection, key string) DocumentID {
	return DocumentID(pathEscape(collection) + "/" + pathEscape(key))
}
