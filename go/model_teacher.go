/*
 * ITU API
 *
 * Nice api yep
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Teacher struct {

	Name string `json:"name,omitempty"`

	Popularity int32 `json:"popularity,omitempty"`

	Id int32 `json:"id,omitempty"`

	Course interface{} `json:"course,omitempty"`
}
