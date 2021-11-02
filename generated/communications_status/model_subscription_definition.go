/*
 * Subscriptions
 *
 * Subscriptions allow contacts to control what forms of communications they receive. Contacts can decide whether they want to receive communication pertaining to a specific topic, brand, or an entire HubSpot account.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package communications_status

import (
	"time"
)

type SubscriptionDefinition struct {
	// The ID of the definition.
	Id string `json:"id"`
	// The name of the subscription.
	Name string `json:"name"`
	// A description of the subscription.
	Description string `json:"description"`
	// The purpose of this subscription or the department in your organization that uses it.
	Purpose string `json:"purpose,omitempty"`
	// The method or technology used to contact.
	CommunicationMethod string `json:"communicationMethod,omitempty"`
	// Whether the definition is active or archived.
	IsActive bool `json:"isActive"`
	// A subscription definition created by HubSpot.
	IsDefault bool `json:"isDefault"`
	// A default description that is used by some HubSpot tools and cannot be edited.
	IsInternal bool `json:"isInternal"`
	// Time at which the definition was created.
	CreatedAt time.Time `json:"createdAt"`
	// Time at which the definition was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}