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

import (
	"time"
)

type Repository struct {
	Type_       string    `json:"type"`
	CreatedOn   time.Time `json:"created_on,omitempty"`
	Description string    `json:"description,omitempty"`
	//  Controls the rules for forking this repository.  * **allow_forks**: unrestricted forking * **no_public_forks**: restrict forking to private forks (forks cannot   be made public later) * **no_forks**: deny all forking
	ForkPolicy string `json:"fork_policy,omitempty"`
	// The concatenation of the repository owner's username and the slugified name, e.g. \"evzijst/interruptingcow\". This is the same string used in Bitbucket URLs.
	FullName   string           `json:"full_name,omitempty"`
	HasIssues  bool             `json:"has_issues,omitempty"`
	HasWiki    bool             `json:"has_wiki,omitempty"`
	IsPrivate  bool             `json:"is_private,omitempty"`
	Language   string           `json:"language,omitempty"`
	Links      *RepositoryLinks `json:"links,omitempty"`
	Mainbranch *Branch          `json:"mainbranch,omitempty"`
	Name       string           `json:"name,omitempty"`
	Owner      *Account         `json:"owner,omitempty"`
	Parent     *Repository      `json:"parent,omitempty"`
	Project    *Project         `json:"project,omitempty"`
	Scm        string           `json:"scm,omitempty"`
	Size       int32            `json:"size,omitempty"`
	// Slugified name, the same string that is used in Bitbucket URLs.
	Slug      string    `json:"slug,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
	// The repository's immutable id. This can be used as a substitute for the slug segment in URLs. Doing this guarantees your URLs will survive renaming of the repository by its owner, or even transfer of the repository to a different user.
	Uuid string `json:"uuid,omitempty"`
}
