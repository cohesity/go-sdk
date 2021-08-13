# NewGroupParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the new protection group. | 
**PolicyId** | **NullableString** | Specifies the policy id of the new protection group. | 
**StorageDomainId** | **NullableInt64** | Specifies the storage domain id of the new protection group. | 

## Methods

### NewNewGroupParam

`func NewNewGroupParam(name NullableString, policyId NullableString, storageDomainId NullableInt64, ) *NewGroupParam`

NewNewGroupParam instantiates a new NewGroupParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewGroupParamWithDefaults

`func NewNewGroupParamWithDefaults() *NewGroupParam`

NewNewGroupParamWithDefaults instantiates a new NewGroupParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewGroupParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewGroupParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewGroupParam) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *NewGroupParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NewGroupParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPolicyId

`func (o *NewGroupParam) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *NewGroupParam) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *NewGroupParam) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### SetPolicyIdNil

`func (o *NewGroupParam) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *NewGroupParam) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetStorageDomainId

`func (o *NewGroupParam) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *NewGroupParam) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *NewGroupParam) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.


### SetStorageDomainIdNil

`func (o *NewGroupParam) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *NewGroupParam) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


