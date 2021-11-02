/*
 * CRM Imports
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package imports

// A summary detailing which list contains the imported objects.
type PublicObjectListRecord struct {
	// The ID of the list containing the imported objects.
	ListId string `json:"listId"`
	// The type of object contained in the list.
	ObjectType string `json:"objectType"`
}