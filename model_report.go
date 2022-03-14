/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Report struct {
	Type_ string `json:"type"`
	// The UUID that can be used to identify the report.
	Uuid string `json:"uuid,omitempty"`
	// The title of the report.
	Title string `json:"title,omitempty"`
	// A string to describe the purpose of the report.
	Details string `json:"details,omitempty"`
	// ID of the report provided by the report creator. It can be used to identify the report as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the report creator for updating or deleting this specific report. Needs to be unique.
	ExternalId string `json:"external_id,omitempty"`
	// A string to describe the tool or company who created the report.
	Reporter string `json:"reporter,omitempty"`
	// A URL linking to the results of the report in an external tool.
	Link string `json:"link,omitempty"`
	// If enabled, a remote link is created in Jira for the issue associated with the commit the report belongs to.
	RemoteLinkEnabled bool `json:"remote_link_enabled,omitempty"`
	// A URL to the report logo. If none is provided, the default insights logo will be used.
	LogoUrl string `json:"logo_url,omitempty"`
	// The type of the report.
	ReportType string `json:"report_type,omitempty"`
	// The state of the report. May be set to PENDING and later updated.
	Result string `json:"result,omitempty"`
	// An array of data fields to display information on the report. Maximum 10.
	Data []ReportData `json:"data,omitempty"`
	// The timestamp when the report was created.
	CreatedOn time.Time `json:"created_on,omitempty"`
	// The timestamp when the report was updated.
	UpdatedOn time.Time `json:"updated_on,omitempty"`
}