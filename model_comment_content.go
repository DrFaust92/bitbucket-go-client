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

type CommentContent struct {
	// The text as it was typed by a user.
	Raw string `json:"raw,omitempty"`
	// The type of markup language the raw content is to be interpreted in.
	Markup string `json:"markup,omitempty"`
	// The user's content rendered as HTML.
	Html string `json:"html,omitempty"`
}
