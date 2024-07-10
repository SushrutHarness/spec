// Copyright 2022 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

import (
	"encoding/json"
	"strings"
)

type Mount struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
}

// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Mount) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Source string `json:"source,omitempty"`
		Target string `json:"target,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err != nil {
		parts := strings.SplitN(out1, ":", 2)
		switch len(parts) {
		case 1:
			v.Source = parts[0]
			v.Target = parts[0]
		case 2:
			v.Source = parts[0]
			v.Target = parts[1]
		}
		return nil
	}

	if err := json.Unmarshal(data, &out2); err != nil {
		v.Source = out2.Source
		v.Target = out2.Target
		return nil
	} else {
		return err
	}
}
