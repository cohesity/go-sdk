# OracleSbtHostParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SbtLibraryPath** | Pointer to **NullableString** | The path of sbt library, This is set only when backup type is kSbt. | [optional] 
**ViewFsPath** | Pointer to **NullableString** | Cohesity view path. | [optional] 
**VipVec** | Pointer to **[]string** | Vector of Cohesity primary VIPs. | [optional] 
**VlanInfoVec** | Pointer to [**[]OracleVlanInfo**](OracleVlanInfo.md) | Vlan information for Cohesity cluster. | [optional] 

## Methods

### NewOracleSbtHostParams

`func NewOracleSbtHostParams() *OracleSbtHostParams`

NewOracleSbtHostParams instantiates a new OracleSbtHostParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleSbtHostParamsWithDefaults

`func NewOracleSbtHostParamsWithDefaults() *OracleSbtHostParams`

NewOracleSbtHostParamsWithDefaults instantiates a new OracleSbtHostParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSbtLibraryPath

`func (o *OracleSbtHostParams) GetSbtLibraryPath() string`

GetSbtLibraryPath returns the SbtLibraryPath field if non-nil, zero value otherwise.

### GetSbtLibraryPathOk

`func (o *OracleSbtHostParams) GetSbtLibraryPathOk() (*string, bool)`

GetSbtLibraryPathOk returns a tuple with the SbtLibraryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtLibraryPath

`func (o *OracleSbtHostParams) SetSbtLibraryPath(v string)`

SetSbtLibraryPath sets SbtLibraryPath field to given value.

### HasSbtLibraryPath

`func (o *OracleSbtHostParams) HasSbtLibraryPath() bool`

HasSbtLibraryPath returns a boolean if a field has been set.

### SetSbtLibraryPathNil

`func (o *OracleSbtHostParams) SetSbtLibraryPathNil(b bool)`

 SetSbtLibraryPathNil sets the value for SbtLibraryPath to be an explicit nil

### UnsetSbtLibraryPath
`func (o *OracleSbtHostParams) UnsetSbtLibraryPath()`

UnsetSbtLibraryPath ensures that no value is present for SbtLibraryPath, not even an explicit nil
### GetViewFsPath

`func (o *OracleSbtHostParams) GetViewFsPath() string`

GetViewFsPath returns the ViewFsPath field if non-nil, zero value otherwise.

### GetViewFsPathOk

`func (o *OracleSbtHostParams) GetViewFsPathOk() (*string, bool)`

GetViewFsPathOk returns a tuple with the ViewFsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewFsPath

`func (o *OracleSbtHostParams) SetViewFsPath(v string)`

SetViewFsPath sets ViewFsPath field to given value.

### HasViewFsPath

`func (o *OracleSbtHostParams) HasViewFsPath() bool`

HasViewFsPath returns a boolean if a field has been set.

### SetViewFsPathNil

`func (o *OracleSbtHostParams) SetViewFsPathNil(b bool)`

 SetViewFsPathNil sets the value for ViewFsPath to be an explicit nil

### UnsetViewFsPath
`func (o *OracleSbtHostParams) UnsetViewFsPath()`

UnsetViewFsPath ensures that no value is present for ViewFsPath, not even an explicit nil
### GetVipVec

`func (o *OracleSbtHostParams) GetVipVec() []string`

GetVipVec returns the VipVec field if non-nil, zero value otherwise.

### GetVipVecOk

`func (o *OracleSbtHostParams) GetVipVecOk() (*[]string, bool)`

GetVipVecOk returns a tuple with the VipVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipVec

`func (o *OracleSbtHostParams) SetVipVec(v []string)`

SetVipVec sets VipVec field to given value.

### HasVipVec

`func (o *OracleSbtHostParams) HasVipVec() bool`

HasVipVec returns a boolean if a field has been set.

### SetVipVecNil

`func (o *OracleSbtHostParams) SetVipVecNil(b bool)`

 SetVipVecNil sets the value for VipVec to be an explicit nil

### UnsetVipVec
`func (o *OracleSbtHostParams) UnsetVipVec()`

UnsetVipVec ensures that no value is present for VipVec, not even an explicit nil
### GetVlanInfoVec

`func (o *OracleSbtHostParams) GetVlanInfoVec() []OracleVlanInfo`

GetVlanInfoVec returns the VlanInfoVec field if non-nil, zero value otherwise.

### GetVlanInfoVecOk

`func (o *OracleSbtHostParams) GetVlanInfoVecOk() (*[]OracleVlanInfo, bool)`

GetVlanInfoVecOk returns a tuple with the VlanInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfoVec

`func (o *OracleSbtHostParams) SetVlanInfoVec(v []OracleVlanInfo)`

SetVlanInfoVec sets VlanInfoVec field to given value.

### HasVlanInfoVec

`func (o *OracleSbtHostParams) HasVlanInfoVec() bool`

HasVlanInfoVec returns a boolean if a field has been set.

### SetVlanInfoVecNil

`func (o *OracleSbtHostParams) SetVlanInfoVecNil(b bool)`

 SetVlanInfoVecNil sets the value for VlanInfoVec to be an explicit nil

### UnsetVlanInfoVec
`func (o *OracleSbtHostParams) UnsetVlanInfoVec()`

UnsetVlanInfoVec ensures that no value is present for VlanInfoVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


