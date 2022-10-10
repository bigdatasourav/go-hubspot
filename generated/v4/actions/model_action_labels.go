/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
)

// ActionLabels User-facing labels for this custom workflow action.
type ActionLabels struct {
	// A map of input field names to the user-facing labels.
	InputFieldLabels *map[string]string `json:"inputFieldLabels,omitempty"`
	// A map of input field names to descriptions for the fields. These will show up as tooltips when users are editing your action.
	InputFieldDescriptions *map[string]string `json:"inputFieldDescriptions,omitempty"`
	// The name of this custom action. This is what will show up when users are selecting an action in the workflows app.
	ActionName string `json:"actionName"`
	// A description for this custom action. This will show up in the action editor along with the input fields.
	ActionDescription *string `json:"actionDescription,omitempty"`
	// The name to be displayed at the top of the action editor in the workflows app.
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// The label to be displayed in the action card of the workflow editor once this custom action has been added to a workflow.
	ActionCardContent *string `json:"actionCardContent,omitempty"`
}

// NewActionLabels instantiates a new ActionLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionLabels(actionName string) *ActionLabels {
	this := ActionLabels{}
	this.ActionName = actionName
	return &this
}

// NewActionLabelsWithDefaults instantiates a new ActionLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionLabelsWithDefaults() *ActionLabels {
	this := ActionLabels{}
	return &this
}

// GetInputFieldLabels returns the InputFieldLabels field value if set, zero value otherwise.
func (o *ActionLabels) GetInputFieldLabels() map[string]string {
	if o == nil || o.InputFieldLabels == nil {
		var ret map[string]string
		return ret
	}
	return *o.InputFieldLabels
}

// GetInputFieldLabelsOk returns a tuple with the InputFieldLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionLabels) GetInputFieldLabelsOk() (*map[string]string, bool) {
	if o == nil || o.InputFieldLabels == nil {
		return nil, false
	}
	return o.InputFieldLabels, true
}

// HasInputFieldLabels returns a boolean if a field has been set.
func (o *ActionLabels) HasInputFieldLabels() bool {
	if o != nil && o.InputFieldLabels != nil {
		return true
	}

	return false
}

// SetInputFieldLabels gets a reference to the given map[string]string and assigns it to the InputFieldLabels field.
func (o *ActionLabels) SetInputFieldLabels(v map[string]string) {
	o.InputFieldLabels = &v
}

// GetInputFieldDescriptions returns the InputFieldDescriptions field value if set, zero value otherwise.
func (o *ActionLabels) GetInputFieldDescriptions() map[string]string {
	if o == nil || o.InputFieldDescriptions == nil {
		var ret map[string]string
		return ret
	}
	return *o.InputFieldDescriptions
}

// GetInputFieldDescriptionsOk returns a tuple with the InputFieldDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionLabels) GetInputFieldDescriptionsOk() (*map[string]string, bool) {
	if o == nil || o.InputFieldDescriptions == nil {
		return nil, false
	}
	return o.InputFieldDescriptions, true
}

// HasInputFieldDescriptions returns a boolean if a field has been set.
func (o *ActionLabels) HasInputFieldDescriptions() bool {
	if o != nil && o.InputFieldDescriptions != nil {
		return true
	}

	return false
}

// SetInputFieldDescriptions gets a reference to the given map[string]string and assigns it to the InputFieldDescriptions field.
func (o *ActionLabels) SetInputFieldDescriptions(v map[string]string) {
	o.InputFieldDescriptions = &v
}

// GetActionName returns the ActionName field value
func (o *ActionLabels) GetActionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value
// and a boolean to check if the value has been set.
func (o *ActionLabels) GetActionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionName, true
}

// SetActionName sets field value
func (o *ActionLabels) SetActionName(v string) {
	o.ActionName = v
}

// GetActionDescription returns the ActionDescription field value if set, zero value otherwise.
func (o *ActionLabels) GetActionDescription() string {
	if o == nil || o.ActionDescription == nil {
		var ret string
		return ret
	}
	return *o.ActionDescription
}

// GetActionDescriptionOk returns a tuple with the ActionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionLabels) GetActionDescriptionOk() (*string, bool) {
	if o == nil || o.ActionDescription == nil {
		return nil, false
	}
	return o.ActionDescription, true
}

// HasActionDescription returns a boolean if a field has been set.
func (o *ActionLabels) HasActionDescription() bool {
	if o != nil && o.ActionDescription != nil {
		return true
	}

	return false
}

// SetActionDescription gets a reference to the given string and assigns it to the ActionDescription field.
func (o *ActionLabels) SetActionDescription(v string) {
	o.ActionDescription = &v
}

// GetAppDisplayName returns the AppDisplayName field value if set, zero value otherwise.
func (o *ActionLabels) GetAppDisplayName() string {
	if o == nil || o.AppDisplayName == nil {
		var ret string
		return ret
	}
	return *o.AppDisplayName
}

// GetAppDisplayNameOk returns a tuple with the AppDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionLabels) GetAppDisplayNameOk() (*string, bool) {
	if o == nil || o.AppDisplayName == nil {
		return nil, false
	}
	return o.AppDisplayName, true
}

// HasAppDisplayName returns a boolean if a field has been set.
func (o *ActionLabels) HasAppDisplayName() bool {
	if o != nil && o.AppDisplayName != nil {
		return true
	}

	return false
}

// SetAppDisplayName gets a reference to the given string and assigns it to the AppDisplayName field.
func (o *ActionLabels) SetAppDisplayName(v string) {
	o.AppDisplayName = &v
}

// GetActionCardContent returns the ActionCardContent field value if set, zero value otherwise.
func (o *ActionLabels) GetActionCardContent() string {
	if o == nil || o.ActionCardContent == nil {
		var ret string
		return ret
	}
	return *o.ActionCardContent
}

// GetActionCardContentOk returns a tuple with the ActionCardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionLabels) GetActionCardContentOk() (*string, bool) {
	if o == nil || o.ActionCardContent == nil {
		return nil, false
	}
	return o.ActionCardContent, true
}

// HasActionCardContent returns a boolean if a field has been set.
func (o *ActionLabels) HasActionCardContent() bool {
	if o != nil && o.ActionCardContent != nil {
		return true
	}

	return false
}

// SetActionCardContent gets a reference to the given string and assigns it to the ActionCardContent field.
func (o *ActionLabels) SetActionCardContent(v string) {
	o.ActionCardContent = &v
}

func (o ActionLabels) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InputFieldLabels != nil {
		toSerialize["inputFieldLabels"] = o.InputFieldLabels
	}
	if o.InputFieldDescriptions != nil {
		toSerialize["inputFieldDescriptions"] = o.InputFieldDescriptions
	}
	if true {
		toSerialize["actionName"] = o.ActionName
	}
	if o.ActionDescription != nil {
		toSerialize["actionDescription"] = o.ActionDescription
	}
	if o.AppDisplayName != nil {
		toSerialize["appDisplayName"] = o.AppDisplayName
	}
	if o.ActionCardContent != nil {
		toSerialize["actionCardContent"] = o.ActionCardContent
	}
	return json.Marshal(toSerialize)
}

type NullableActionLabels struct {
	value *ActionLabels
	isSet bool
}

func (v NullableActionLabels) Get() *ActionLabels {
	return v.value
}

func (v *NullableActionLabels) Set(val *ActionLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableActionLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableActionLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionLabels(val *ActionLabels) *NullableActionLabels {
	return &NullableActionLabels{value: val, isSet: true}
}

func (v NullableActionLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}