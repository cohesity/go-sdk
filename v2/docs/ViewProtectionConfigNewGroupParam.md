# ViewProtectionConfigNewGroupParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the new protection group. | 
**PolicyId** | **NullableString** | Specifies the policy id of the new protection group. | 
**StorageDomainId** | **NullableInt64** | Specifies the storage domain id of the new protection group. | 

## Methods

### NewViewProtectionConfigNewGroupParam

`func NewViewProtectionConfigNewGroupParam(name NullableString, policyId NullableString, storageDomainId NullableInt64, ) *ViewProtectionConfigNewGroupParam`

NewViewProtectionConfigNewGroupParam instantiates a new ViewProtectionConfigNewGroupParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProtectionConfigNewGroupParamWithDefaults

`func NewViewProtectionConfigNewGroupParamWithDefaults() *ViewProtectionConfigNewGroupParam`

NewViewProtectionConfigNewGroupParamWithDefaults instantiates a new ViewProtectionConfigNewGroupParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ViewProtectionConfigNewGroupParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewProtectionConfigNewGroupParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewProtectionConfigNewGroupParam) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ViewProtectionConfigNewGroupParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ViewProtectionConfigNewGroupParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPolicyId

`func (o *ViewProtectionConfigNewGroupParam) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ViewProtectionConfigNewGroupParam) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ViewProtectionConfigNewGroupParam) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### SetPolicyIdNil

`func (o *ViewProtectionConfigNewGroupParam) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ViewProtectionConfigNewGroupParam) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetStorageDomainId

`func (o *ViewProtectionConfigNewGroupParam) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ViewProtectionConfigNewGroupParam) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ViewProtectionConfigNewGroupParam) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.


### SetStorageDomainIdNil

`func (o *ViewProtectionConfigNewGroupParam) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ViewProtectionConfigNewGroupParam) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


