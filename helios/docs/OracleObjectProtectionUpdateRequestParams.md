# OracleObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]OracleObjectProtectionInfo**](OracleObjectProtectionInfo.md) | Specifies the list of object ids to be protected. | 
**PersistMountpoints** | Pointer to **NullableBool** | Specifies whether the mountpoints created while backing up Oracle DBs should be persisted. | [optional] 
**VlanParams** | Pointer to [**VlanParams**](VlanParams.md) |  | [optional] 

## Methods

### NewOracleObjectProtectionUpdateRequestParams

`func NewOracleObjectProtectionUpdateRequestParams(objects []OracleObjectProtectionInfo, ) *OracleObjectProtectionUpdateRequestParams`

NewOracleObjectProtectionUpdateRequestParams instantiates a new OracleObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleObjectProtectionUpdateRequestParamsWithDefaults

`func NewOracleObjectProtectionUpdateRequestParamsWithDefaults() *OracleObjectProtectionUpdateRequestParams`

NewOracleObjectProtectionUpdateRequestParamsWithDefaults instantiates a new OracleObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *OracleObjectProtectionUpdateRequestParams) GetObjects() []OracleObjectProtectionInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *OracleObjectProtectionUpdateRequestParams) GetObjectsOk() (*[]OracleObjectProtectionInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *OracleObjectProtectionUpdateRequestParams) SetObjects(v []OracleObjectProtectionInfo)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *OracleObjectProtectionUpdateRequestParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *OracleObjectProtectionUpdateRequestParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPersistMountpoints

`func (o *OracleObjectProtectionUpdateRequestParams) GetPersistMountpoints() bool`

GetPersistMountpoints returns the PersistMountpoints field if non-nil, zero value otherwise.

### GetPersistMountpointsOk

`func (o *OracleObjectProtectionUpdateRequestParams) GetPersistMountpointsOk() (*bool, bool)`

GetPersistMountpointsOk returns a tuple with the PersistMountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMountpoints

`func (o *OracleObjectProtectionUpdateRequestParams) SetPersistMountpoints(v bool)`

SetPersistMountpoints sets PersistMountpoints field to given value.

### HasPersistMountpoints

`func (o *OracleObjectProtectionUpdateRequestParams) HasPersistMountpoints() bool`

HasPersistMountpoints returns a boolean if a field has been set.

### SetPersistMountpointsNil

`func (o *OracleObjectProtectionUpdateRequestParams) SetPersistMountpointsNil(b bool)`

 SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil

### UnsetPersistMountpoints
`func (o *OracleObjectProtectionUpdateRequestParams) UnsetPersistMountpoints()`

UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil
### GetVlanParams

`func (o *OracleObjectProtectionUpdateRequestParams) GetVlanParams() VlanParams`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *OracleObjectProtectionUpdateRequestParams) GetVlanParamsOk() (*VlanParams, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *OracleObjectProtectionUpdateRequestParams) SetVlanParams(v VlanParams)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *OracleObjectProtectionUpdateRequestParams) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


