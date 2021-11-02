/*
 * Calling Extensions API
 *
 * Provides a way for apps to add custom calling options to a contact record. This works in conjunction with the [Calling SDK](#), which is used to build your phone/calling UI. The endpoints here allow your service to appear as an option to HubSpot users when they access the *Call* action on a contact record. Once accessed, your custom phone/calling UI will be displayed in an iframe at the specified URL with the specified dimensions on that record.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package calling

import (
	"time"
)

// Current settings state
type SettingsResponse struct {
	// The name of your calling service to display to users.
	Name string `json:"name"`
	// The URL to your phone/calling UI, built with the [Calling SDK](#).
	Url string `json:"url"`
	// The target height of the iframe that will contain your phone/calling UI.
	Height int32 `json:"height"`
	// The target width of the iframe that will contain your phone/calling UI.
	Width int32 `json:"width"`
	// When true, your service will appear as an option under the *Call* action in contact records of connected accounts.
	IsReady bool `json:"isReady"`
	// When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects.
	SupportsCustomObjects bool `json:"supportsCustomObjects"`
	// When this calling extension was created.
	CreatedAt time.Time `json:"createdAt"`
	// The last time the settings for this calling extension were modified.
	UpdatedAt time.Time `json:"updatedAt"`
}