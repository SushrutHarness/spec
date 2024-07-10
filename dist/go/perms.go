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

type Permissions struct {
	Actions            string `json:"actions,omitempty"`
	Checks             string `json:"checks,omitempty"`
	Contents           string `json:"contents,omitempty"`
	Deployments        string `json:"deployments,omitempty"`
	Discussions        string `json:"discussions,omitempty"`
	IdToken            string `json:"id-token,omitempty"`
	Issues             string `json:"issues,omitempty"`
	Packages           string `json:"packages,omitempty"`
	Pages              string `json:"pages,omitempty"`
	PullRequests       string `json:"pull-requests,omitempty"`
	RepositoryProjects string `json:"repository-projects,omitempty"`
	SecurityEvents     string `json:"security-events,omitempty"`
	Statuses           string `json:"statuses,omitempty"`
}

// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Permissions) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Actions            string `json:"actions,omitempty"`
		Checks             string `json:"checks,omitempty"`
		Contents           string `json:"contents,omitempty"`
		Deployments        string `json:"deployments,omitempty"`
		Discussions        string `json:"discussions,omitempty"`
		IdToken            string `json:"id-token,omitempty"`
		Issues             string `json:"issues,omitempty"`
		Packages           string `json:"packages,omitempty"`
		Pages              string `json:"pages,omitempty"`
		PullRequests       string `json:"pull-requests,omitempty"`
		RepositoryProjects string `json:"repository-projects,omitempty"`
		SecurityEvents     string `json:"security-events,omitempty"`
		Statuses           string `json:"statuses,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err != nil {
		switch out1 {
		case "read-all":
			out1 = "read"
		case "write-all":
			out1 = "write"
		default:
			return fmt.Errorf("unknown permissions value: %s", out1)
		}
		v.Actions = out1
		v.Checks = out1
		v.Contents = out1
		v.Deployments = out1
		v.Discussions = out1
		v.IdToken = out1
		v.Issues = out1
		v.Packages = out1
		v.Pages = out1
		v.PullRequests = out1
		v.RepositoryProjects = out1
		v.SecurityEvents = out1
		v.Statuses = out1
		return nil
	}

	if err := json.Unmarshal(data, &out2); err != nil {
		v.Actions = out2.Actions
		v.Checks = out2.Checks
		v.Contents = out2.Contents
		v.Deployments = out2.Deployments
		v.Discussions = out2.Discussions
		v.IdToken = out2.IdToken
		v.Issues = out2.Issues
		v.Packages = out2.Packages
		v.Pages = out2.Pages
		v.PullRequests = out2.PullRequests
		v.RepositoryProjects = out2.RepositoryProjects
		v.SecurityEvents = out2.SecurityEvents
		v.Statuses = out2.Statuses
		return nil
	} else {
		return err
	}
}

// "read" | "write" | "none";
// "write-all" | "read-all"
