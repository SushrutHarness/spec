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

import "encoding/json"

type Output struct {
	Alias string `json:"alias,omitempty"`
	Mask  bool   `json:"mask,omitempty"`
	Name  string `json:"name,omitempty"`
	Scope string `json:"scope,omitempty"`
}

// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Output) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Alias string `json:"alias,omitempty"`
		Mask  bool   `json:"mask,omitempty"`
		Name  string `json:"name,omitempty"`
		Scope string `json:"scope,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err != nil {
		v.Name = out1
		return nil
	}

	if err := json.Unmarshal(data, &out2); err != nil {
		v.Alias = out2.Alias
		v.Mask = out2.Mask
		v.Name = out2.Name
		v.Scope = out2.Scope
		return nil
	} else {
		return err
	}
}
