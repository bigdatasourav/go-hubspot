/*
 * Associations
 *
 * Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package associations

type CollectionResponsePublicAssociationDefiniton struct {
	Results []PublicAssociationDefiniton `json:"results"`
	Paging  *Paging                      `json:"paging,omitempty"`
}