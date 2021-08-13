# PolicyTemplateResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies a unique Policy id assigned by the Cohesity Cluster. | [optional] 
**NumLinkedPolicies** | Pointer to **NullableInt64** | Specifies the number of policies linked to this policy template. Only applicable in case of policy template. | [optional] 
**IsUsable** | Pointer to **NullableBool** | This field is set to true if this policy template qualifies to create more policies. If the template is partially filled and can not create a working policy then this field will be set to false. | [optional] 

## Methods

### NewPolicyTemplateResponseAllOf

`func NewPolicyTemplateResponseAllOf() *PolicyTemplateResponseAllOf`

NewPolicyTemplateResponseAllOf instantiates a new PolicyTemplateResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateResponseAllOfWithDefaults

`func NewPolicyTemplateResponseAllOfWithDefaults() *PolicyTemplateResponseAllOf`

NewPolicyTemplateResponseAllOfWithDefaults instantiates a new PolicyTemplateResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyTemplateResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyTemplateResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyTemplateResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyTemplateResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PolicyTemplateResponseAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PolicyTemplateResponseAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNumLinkedPolicies

`func (o *PolicyTemplateResponseAllOf) GetNumLinkedPolicies() int64`

GetNumLinkedPolicies returns the NumLinkedPolicies field if non-nil, zero value otherwise.

### GetNumLinkedPoliciesOk

`func (o *PolicyTemplateResponseAllOf) GetNumLinkedPoliciesOk() (*int64, bool)`

GetNumLinkedPoliciesOk returns a tuple with the NumLinkedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLinkedPolicies

`func (o *PolicyTemplateResponseAllOf) SetNumLinkedPolicies(v int64)`

SetNumLinkedPolicies sets NumLinkedPolicies field to given value.

### HasNumLinkedPolicies

`func (o *PolicyTemplateResponseAllOf) HasNumLinkedPolicies() bool`

HasNumLinkedPolicies returns a boolean if a field has been set.

### SetNumLinkedPoliciesNil

`func (o *PolicyTemplateResponseAllOf) SetNumLinkedPoliciesNil(b bool)`

 SetNumLinkedPoliciesNil sets the value for NumLinkedPolicies to be an explicit nil

### UnsetNumLinkedPolicies
`func (o *PolicyTemplateResponseAllOf) UnsetNumLinkedPolicies()`

UnsetNumLinkedPolicies ensures that no value is present for NumLinkedPolicies, not even an explicit nil
### GetIsUsable

`func (o *PolicyTemplateResponseAllOf) GetIsUsable() bool`

GetIsUsable returns the IsUsable field if non-nil, zero value otherwise.

### GetIsUsableOk

`func (o *PolicyTemplateResponseAllOf) GetIsUsableOk() (*bool, bool)`

GetIsUsableOk returns a tuple with the IsUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsable

`func (o *PolicyTemplateResponseAllOf) SetIsUsable(v bool)`

SetIsUsable sets IsUsable field to given value.

### HasIsUsable

`func (o *PolicyTemplateResponseAllOf) HasIsUsable() bool`

HasIsUsable returns a boolean if a field has been set.

### SetIsUsableNil

`func (o *PolicyTemplateResponseAllOf) SetIsUsableNil(b bool)`

 SetIsUsableNil sets the value for IsUsable to be an explicit nil

### UnsetIsUsable
`func (o *PolicyTemplateResponseAllOf) UnsetIsUsable()`

UnsetIsUsable ensures that no value is present for IsUsable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


