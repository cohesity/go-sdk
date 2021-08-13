# DataLockConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **NullableString** | Specifies the type of WORM retention type. &lt;br&gt;&#39;Compliance&#39; implies WORM retention is set for compliance reason. &lt;br&gt;&#39;Administrative&#39; implies WORM retention is set for administrative purposes. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewDataLockConstraints

`func NewDataLockConstraints() *DataLockConstraints`

NewDataLockConstraints instantiates a new DataLockConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLockConstraintsWithDefaults

`func NewDataLockConstraintsWithDefaults() *DataLockConstraints`

NewDataLockConstraintsWithDefaults instantiates a new DataLockConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *DataLockConstraints) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DataLockConstraints) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DataLockConstraints) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DataLockConstraints) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *DataLockConstraints) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *DataLockConstraints) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *DataLockConstraints) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *DataLockConstraints) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *DataLockConstraints) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *DataLockConstraints) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *DataLockConstraints) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *DataLockConstraints) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


