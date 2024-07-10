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
	"fmt"
)

type Action struct {
	Abort              bool          `json:"abort,omitempty"`
	Fail               bool          `json:"fail,omitempty"`
	Ignore             bool          `json:"ignore,omitempty"`
	ManualIntervention *ActionManual `json:"manual-intervention,omitempty"`
	PipelineRollback   bool          `json:"pipeline-rollback,omitempty"`
	Retry              *ActionRetry  `json:"retry,omitempty"`
	RetryStepGroup     bool          `json:"retry-step-group,omitempty"`
	StageRollback      bool          `json:"stage-rollback,omitempty"`
	Success            bool          `json:"success,omitempty"`
}

// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Action) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Abort              bool          `json:"abort,omitempty"`
		Fail               bool          `json:"fail,omitempty"`
		Ignore             bool          `json:"ignore,omitempty"`
		ManualIntervention *ActionManual `json:"manual-intervention,omitempty"`
		PipelineRollback   bool          `json:"pipeline-rollback,omitempty"`
		Retry              *ActionRetry  `json:"retry,omitempty"`
		RetryStepGroup     bool          `json:"retry-step-group,omitempty"`
		StageRollback      bool          `json:"stage-rollback,omitempty"`
		Success            bool          `json:"success,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err != nil {
		switch out1 {
		case "abort":
			v.Abort = true
		case "fail":
			v.Fail = true
		case "ignore":
			v.Ignore = true
		case "manual-intervention":
			v.ManualIntervention = new(ActionManual)
		case "pipeline-rollback":
			v.PipelineRollback = true
		case "retry":
			v.Retry = new(ActionRetry)
		case "retry-step-group":
			v.RetryStepGroup = true
		case "stage-rollback":
			v.StageRollback = true
		case "success":
			v.Success = true
		default:
			return fmt.Errorf("unknown action type: %s", out2)
		}
		return nil
	}

	if err := json.Unmarshal(data, &out2); err != nil {
		v.Abort = out2.Abort
		v.Fail = out2.Fail
		v.Ignore = out2.Ignore
		v.ManualIntervention = out2.ManualIntervention
		v.PipelineRollback = out2.PipelineRollback
		v.Retry = out2.Retry
		v.RetryStepGroup = out2.RetryStepGroup
		v.StageRollback = out2.StageRollback
		v.Success = out2.Success
		return nil
	} else {
		return err
	}
}
