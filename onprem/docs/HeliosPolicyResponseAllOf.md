# HeliosPolicyResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies a unique policy id assigned by the Helios. | [optional] 
**NumLinkedPolicies** | Pointer to **NullableInt64** | In case of global policy response, specifies the number of policies linked to this global policy on the cluster. | [optional] 
**NumObjectProtections** | Pointer to **NullableInt64** | Specifies the number of object protections using the protection policy. | [optional] 

## Methods

### NewHeliosPolicyResponseAllOf

`func NewHeliosPolicyResponseAllOf() *HeliosPolicyResponseAllOf`

NewHeliosPolicyResponseAllOf instantiates a new HeliosPolicyResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosPolicyResponseAllOfWithDefaults

`func NewHeliosPolicyResponseAllOfWithDefaults() *HeliosPolicyResponseAllOf`

NewHeliosPolicyResponseAllOfWithDefaults instantiates a new HeliosPolicyResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeliosPolicyResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosPolicyResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosPolicyResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosPolicyResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosPolicyResponseAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosPolicyResponseAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNumLinkedPolicies

`func (o *HeliosPolicyResponseAllOf) GetNumLinkedPolicies() int64`

GetNumLinkedPolicies returns the NumLinkedPolicies field if non-nil, zero value otherwise.

### GetNumLinkedPoliciesOk

`func (o *HeliosPolicyResponseAllOf) GetNumLinkedPoliciesOk() (*int64, bool)`

GetNumLinkedPoliciesOk returns a tuple with the NumLinkedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLinkedPolicies

`func (o *HeliosPolicyResponseAllOf) SetNumLinkedPolicies(v int64)`

SetNumLinkedPolicies sets NumLinkedPolicies field to given value.

### HasNumLinkedPolicies

`func (o *HeliosPolicyResponseAllOf) HasNumLinkedPolicies() bool`

HasNumLinkedPolicies returns a boolean if a field has been set.

### SetNumLinkedPoliciesNil

`func (o *HeliosPolicyResponseAllOf) SetNumLinkedPoliciesNil(b bool)`

 SetNumLinkedPoliciesNil sets the value for NumLinkedPolicies to be an explicit nil

### UnsetNumLinkedPolicies
`func (o *HeliosPolicyResponseAllOf) UnsetNumLinkedPolicies()`

UnsetNumLinkedPolicies ensures that no value is present for NumLinkedPolicies, not even an explicit nil
### GetNumObjectProtections

`func (o *HeliosPolicyResponseAllOf) GetNumObjectProtections() int64`

GetNumObjectProtections returns the NumObjectProtections field if non-nil, zero value otherwise.

### GetNumObjectProtectionsOk

`func (o *HeliosPolicyResponseAllOf) GetNumObjectProtectionsOk() (*int64, bool)`

GetNumObjectProtectionsOk returns a tuple with the NumObjectProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectProtections

`func (o *HeliosPolicyResponseAllOf) SetNumObjectProtections(v int64)`

SetNumObjectProtections sets NumObjectProtections field to given value.

### HasNumObjectProtections

`func (o *HeliosPolicyResponseAllOf) HasNumObjectProtections() bool`

HasNumObjectProtections returns a boolean if a field has been set.

### SetNumObjectProtectionsNil

`func (o *HeliosPolicyResponseAllOf) SetNumObjectProtectionsNil(b bool)`

 SetNumObjectProtectionsNil sets the value for NumObjectProtections to be an explicit nil

### UnsetNumObjectProtections
`func (o *HeliosPolicyResponseAllOf) UnsetNumObjectProtections()`

UnsetNumObjectProtections ensures that no value is present for NumObjectProtections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


