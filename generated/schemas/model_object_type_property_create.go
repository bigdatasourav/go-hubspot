/*
 * Schemas
 *
 * The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package schemas

// Defines a property to create.
type ObjectTypePropertyCreate struct {
	// The internal property name, which must be used when referencing the property from the API.
	Name string `json:"name"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label"`
	// The name of the group this property belongs to.
	GroupName string `json:"groupName,omitempty"`
	// A description of the property that will be shown as help text in HubSpot.
	Description string `json:"description,omitempty"`
	// A list of available options for the property. This field is only required for enumerated properties.
	Options []OptionInput `json:"options,omitempty"`
	// The order that this property should be displayed in the HubSpot UI relative to other properties for this object type. Properties are displayed in order starting with the lowest positive integer value. A value of -1 will cause the property to be displayed **after** any positive values.
	DisplayOrder int32 `json:"displayOrder,omitempty"`
	// Whether or not the property's value must be unique. Once set, this can't be changed.
	HasUniqueValue bool `json:"hasUniqueValue,omitempty"`
	Hidden         bool `json:"hidden,omitempty"`
	// The data type of the property.
	Type_ string `json:"type"`
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType"`
}