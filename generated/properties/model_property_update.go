/*
 * Properties
 *
 * All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package properties

type PropertyUpdate struct {
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label,omitempty"`
	// The data type of the property.
	Type_ string `json:"type,omitempty"`
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType,omitempty"`
	// The name of the property group the property belongs to.
	GroupName string `json:"groupName,omitempty"`
	// A description of the property that will be shown as help text in HubSpot.
	Description string `json:"description,omitempty"`
	// A list of valid options for the property.
	Options []OptionInput `json:"options,omitempty"`
	// Properties are displayed in order starting with the lowest positive integer value. Values of -1 will cause the Property to be displayed after any positive values.
	DisplayOrder int32 `json:"displayOrder,omitempty"`
	// If true, the property won't be visible and can't be used in HubSpot.
	Hidden bool `json:"hidden,omitempty"`
	// Whether or not the property can be used in a HubSpot form.
	FormField bool `json:"formField,omitempty"`
}