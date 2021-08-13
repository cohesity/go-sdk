# OracleSbtHostParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SbtLibraryPath** | Pointer to **NullableString** | Specifies the path of sbt library. | [optional] 
**ViewFsPath** | Pointer to **NullableString** | Specifies the Cohesity view path. | [optional] 
**VipList** | Pointer to **[]string** | Specifies the list of Cohesity primary VIPs. | [optional] 
**VlanInfoList** | Pointer to [**[]OracleVlanInfo**](OracleVlanInfo.md) | Specifies the Vlan information for Cohesity cluster. | [optional] 

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
### GetVipList

`func (o *OracleSbtHostParams) GetVipList() []string`

GetVipList returns the VipList field if non-nil, zero value otherwise.

### GetVipListOk

`func (o *OracleSbtHostParams) GetVipListOk() (*[]string, bool)`

GetVipListOk returns a tuple with the VipList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipList

`func (o *OracleSbtHostParams) SetVipList(v []string)`

SetVipList sets VipList field to given value.

### HasVipList

`func (o *OracleSbtHostParams) HasVipList() bool`

HasVipList returns a boolean if a field has been set.

### SetVipListNil

`func (o *OracleSbtHostParams) SetVipListNil(b bool)`

 SetVipListNil sets the value for VipList to be an explicit nil

### UnsetVipList
`func (o *OracleSbtHostParams) UnsetVipList()`

UnsetVipList ensures that no value is present for VipList, not even an explicit nil
### GetVlanInfoList

`func (o *OracleSbtHostParams) GetVlanInfoList() []OracleVlanInfo`

GetVlanInfoList returns the VlanInfoList field if non-nil, zero value otherwise.

### GetVlanInfoListOk

`func (o *OracleSbtHostParams) GetVlanInfoListOk() (*[]OracleVlanInfo, bool)`

GetVlanInfoListOk returns a tuple with the VlanInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfoList

`func (o *OracleSbtHostParams) SetVlanInfoList(v []OracleVlanInfo)`

SetVlanInfoList sets VlanInfoList field to given value.

### HasVlanInfoList

`func (o *OracleSbtHostParams) HasVlanInfoList() bool`

HasVlanInfoList returns a boolean if a field has been set.

### SetVlanInfoListNil

`func (o *OracleSbtHostParams) SetVlanInfoListNil(b bool)`

 SetVlanInfoListNil sets the value for VlanInfoList to be an explicit nil

### UnsetVlanInfoList
`func (o *OracleSbtHostParams) UnsetVlanInfoList()`

UnsetVlanInfoList ensures that no value is present for VlanInfoList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


