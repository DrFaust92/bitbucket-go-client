/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bitbucket

type Branchrestriction struct {
	Type_ string                       `json:"type"`
	Links *BranchingModelSettingsLinks `json:"links,omitempty"`
	// The branch restriction status' id.
	Id int32 `json:"id,omitempty"`
	// The type of restriction that is being applied.
	Kind string `json:"kind"`
	// Indicates how the restriction is matched against a branch. The default is `glob`.
	BranchMatchKind string `json:"branch_match_kind"`
	// Apply the restriction to branches of this type. Active when `branch_match_kind` is `branching_model`. The branch type will be calculated using the branching model configured for the repository.
	BranchType string `json:"branch_type,omitempty"`
	// Apply the restriction to branches that match this pattern. Active when `branch_match_kind` is `glob`. Will be empty when `branch_match_kind` is `branching_model`.
	Pattern string    `json:"pattern"`
	Users   []Account `json:"users,omitempty"`
	Groups  []Group   `json:"groups,omitempty"`
	// <staticmethod object at 0x7f5f04539d50>
	Value int32 `json:"value,omitempty"`
}
