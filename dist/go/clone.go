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

type Clone struct {
	Depth      int64     `json:"depth,omitempty"`
	Disabled   bool      `json:"disabled,omitempty"`
	Insecure   bool      `json:"insecure,omitempty"`
	Lfs        bool      `json:"lfs,omitempty"`
	Ref        *CloneRef `json:"ref,omitempty"`
	Strategy   string    `json:"strategy,omitempty"`
	Submodules bool      `json:"submodules,omitempty"`
	Tags       bool      `json:"tags,omitempty"`
	Trace      bool      `json:"trace,omitempty"`
}

// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Clone) UnmarshalJSON(data []byte) error {
	var out1 bool
	var out2 = struct {
		Depth      int64     `json:"depth,omitempty"`
		Disabled   bool      `json:"disabled,omitempty"`
		Insecure   bool      `json:"insecure,omitempty"`
		Lfs        bool      `json:"lfs,omitempty"`
		Ref        *CloneRef `json:"ref,omitempty"`
		Strategy   string    `json:"strategy,omitempty"`
		Submodules bool      `json:"submodules,omitempty"`
		Tags       bool      `json:"tags,omitempty"`
		Trace      bool      `json:"trace,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err != nil {
		v.Disabled = !out1
		return nil
	}

	if err := json.Unmarshal(data, &out2); err != nil {
		v.Depth = out2.Depth
		v.Disabled = out2.Disabled
		v.Insecure = out2.Insecure
		v.Lfs = out2.Lfs
		v.Ref = out2.Ref
		v.Strategy = out2.Strategy
		v.Submodules = out2.Submodules
		v.Tags = out2.Tags
		v.Trace = out2.Trace
		return nil
	} else {
		return err
	}
}
