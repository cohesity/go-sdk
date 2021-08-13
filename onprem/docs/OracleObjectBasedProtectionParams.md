# OracleObjectBasedProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]OracleObjectProtectionInfo**](OracleObjectProtectionInfo.md) | Specifies the list of object ids to be protected. | 
**PersistMountpoints** | Pointer to **NullableBool** | Specifies whether the mountpoints created while backing up Oracle DBs should be persisted. | [optional] 
**VlanParams** | Pointer to [**VlanParams**](VlanParams.md) |  | [optional] 

## Methods

### NewOracleObjectBasedProtectionParams

`func NewOracleObjectBasedProtectionParams(objects []OracleObjectProtectionInfo, ) *OracleObjectBasedProtectionParams`

NewOracleObjectBasedProtectionParams instantiates a new OracleObjectBasedProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleObjectBasedProtectionParamsWithDefaults

`func NewOracleObjectBasedProtectionParamsWithDefaults() *OracleObjectBasedProtectionParams`

NewOracleObjectBasedProtectionParamsWithDefaults instantiates a new OracleObjectBasedProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *OracleObjectBasedProtectionParams) GetObjects() []OracleObjectProtectionInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *OracleObjectBasedProtectionParams) GetObjectsOk() (*[]OracleObjectProtectionInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *OracleObjectBasedProtectionParams) SetObjects(v []OracleObjectProtectionInfo)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *OracleObjectBasedProtectionParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *OracleObjectBasedProtectionParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPersistMountpoints

`func (o *OracleObjectBasedProtectionParams) GetPersistMountpoints() bool`

GetPersistMountpoints returns the PersistMountpoints field if non-nil, zero value otherwise.

### GetPersistMountpointsOk

`func (o *OracleObjectBasedProtectionParams) GetPersistMountpointsOk() (*bool, bool)`

GetPersistMountpointsOk returns a tuple with the PersistMountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMountpoints

`func (o *OracleObjectBasedProtectionParams) SetPersistMountpoints(v bool)`

SetPersistMountpoints sets PersistMountpoints field to given value.

### HasPersistMountpoints

`func (o *OracleObjectBasedProtectionParams) HasPersistMountpoints() bool`

HasPersistMountpoints returns a boolean if a field has been set.

### SetPersistMountpointsNil

`func (o *OracleObjectBasedProtectionParams) SetPersistMountpointsNil(b bool)`

 SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil

### UnsetPersistMountpoints
`func (o *OracleObjectBasedProtectionParams) UnsetPersistMountpoints()`

UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil
### GetVlanParams

`func (o *OracleObjectBasedProtectionParams) GetVlanParams() VlanParams`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *OracleObjectBasedProtectionParams) GetVlanParamsOk() (*VlanParams, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *OracleObjectBasedProtectionParams) SetVlanParams(v VlanParams)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *OracleObjectBasedProtectionParams) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


