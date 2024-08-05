# TdmObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedByUser** | Pointer to [**User**](User.md) | Specifies the user, who created the object. | [optional] 
**Environment** | **NullableString** | Specifies the environment of the object. | 
**Id** | **NullableString** | Specifies the unique ID of the object. | 
**LastAction** | Pointer to [**TdmObjectTimelineEvent**](TdmObjectTimelineEvent.md) | Specifies the details of the last action performed on this object. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**Parent** | Pointer to [**ObjectSummary**](ObjectSummary.md) | Specifies the parent of the object. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the size (in bytes) of the object. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the current status of the object. | [optional] 
**OracleParams** | Pointer to [**OracleCloneObject**](OracleCloneObject.md) |  | [optional] 

## Methods

### NewTdmObject

`func NewTdmObject(environment NullableString, id NullableString, ) *TdmObject`

NewTdmObject instantiates a new TdmObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmObjectWithDefaults

`func NewTdmObjectWithDefaults() *TdmObject`

NewTdmObjectWithDefaults instantiates a new TdmObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedByUser

`func (o *TdmObject) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *TdmObject) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *TdmObject) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *TdmObject) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetEnvironment

`func (o *TdmObject) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TdmObject) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TdmObject) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *TdmObject) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *TdmObject) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *TdmObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TdmObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TdmObject) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TdmObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TdmObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastAction

`func (o *TdmObject) GetLastAction() TdmObjectTimelineEvent`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *TdmObject) GetLastActionOk() (*TdmObjectTimelineEvent, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *TdmObject) SetLastAction(v TdmObjectTimelineEvent)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *TdmObject) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetName

`func (o *TdmObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TdmObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TdmObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TdmObject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TdmObject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TdmObject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParent

`func (o *TdmObject) GetParent() ObjectSummary`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TdmObject) GetParentOk() (*ObjectSummary, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TdmObject) SetParent(v ObjectSummary)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TdmObject) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSizeBytes

`func (o *TdmObject) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *TdmObject) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *TdmObject) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *TdmObject) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *TdmObject) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *TdmObject) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
### GetStatus

`func (o *TdmObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TdmObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TdmObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TdmObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TdmObject) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TdmObject) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetOracleParams

`func (o *TdmObject) GetOracleParams() OracleCloneObject`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *TdmObject) GetOracleParamsOk() (*OracleCloneObject, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *TdmObject) SetOracleParams(v OracleCloneObject)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *TdmObject) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


