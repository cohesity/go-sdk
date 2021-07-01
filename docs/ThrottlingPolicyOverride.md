# ThrottlingPolicyOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | Pointer to **NullableInt64** | Specifies the Protection Source id of the Datastore. | [optional] 
**DatastoreName** | Pointer to **NullableString** | Specifies the display name of the Datastore. | [optional] 
**ThrottlingPolicy** | Pointer to [**ThrottlingPolicyParameters**](ThrottlingPolicyParameters.md) |  | [optional] 

## Methods

### NewThrottlingPolicyOverride

`func NewThrottlingPolicyOverride() *ThrottlingPolicyOverride`

NewThrottlingPolicyOverride instantiates a new ThrottlingPolicyOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThrottlingPolicyOverrideWithDefaults

`func NewThrottlingPolicyOverrideWithDefaults() *ThrottlingPolicyOverride`

NewThrottlingPolicyOverrideWithDefaults instantiates a new ThrottlingPolicyOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *ThrottlingPolicyOverride) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ThrottlingPolicyOverride) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ThrottlingPolicyOverride) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *ThrottlingPolicyOverride) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *ThrottlingPolicyOverride) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *ThrottlingPolicyOverride) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDatastoreName

`func (o *ThrottlingPolicyOverride) GetDatastoreName() string`

GetDatastoreName returns the DatastoreName field if non-nil, zero value otherwise.

### GetDatastoreNameOk

`func (o *ThrottlingPolicyOverride) GetDatastoreNameOk() (*string, bool)`

GetDatastoreNameOk returns a tuple with the DatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreName

`func (o *ThrottlingPolicyOverride) SetDatastoreName(v string)`

SetDatastoreName sets DatastoreName field to given value.

### HasDatastoreName

`func (o *ThrottlingPolicyOverride) HasDatastoreName() bool`

HasDatastoreName returns a boolean if a field has been set.

### SetDatastoreNameNil

`func (o *ThrottlingPolicyOverride) SetDatastoreNameNil(b bool)`

 SetDatastoreNameNil sets the value for DatastoreName to be an explicit nil

### UnsetDatastoreName
`func (o *ThrottlingPolicyOverride) UnsetDatastoreName()`

UnsetDatastoreName ensures that no value is present for DatastoreName, not even an explicit nil
### GetThrottlingPolicy

`func (o *ThrottlingPolicyOverride) GetThrottlingPolicy() ThrottlingPolicyParameters`

GetThrottlingPolicy returns the ThrottlingPolicy field if non-nil, zero value otherwise.

### GetThrottlingPolicyOk

`func (o *ThrottlingPolicyOverride) GetThrottlingPolicyOk() (*ThrottlingPolicyParameters, bool)`

GetThrottlingPolicyOk returns a tuple with the ThrottlingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingPolicy

`func (o *ThrottlingPolicyOverride) SetThrottlingPolicy(v ThrottlingPolicyParameters)`

SetThrottlingPolicy sets ThrottlingPolicy field to given value.

### HasThrottlingPolicy

`func (o *ThrottlingPolicyOverride) HasThrottlingPolicy() bool`

HasThrottlingPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


