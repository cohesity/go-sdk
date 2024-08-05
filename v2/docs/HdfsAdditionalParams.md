# HdfsAdditionalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **NullableString** | Authentication type. | [optional] [readonly] 
**NamenodeAddress** | Pointer to **string** | The HDFS Namenode IP or hostname. | [optional] [readonly] 
**WebhdfsPort** | Pointer to **int32** | The HDFS WebHDFS port. | [optional] [readonly] 

## Methods

### NewHdfsAdditionalParams

`func NewHdfsAdditionalParams() *HdfsAdditionalParams`

NewHdfsAdditionalParams instantiates a new HdfsAdditionalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsAdditionalParamsWithDefaults

`func NewHdfsAdditionalParamsWithDefaults() *HdfsAdditionalParams`

NewHdfsAdditionalParamsWithDefaults instantiates a new HdfsAdditionalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *HdfsAdditionalParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *HdfsAdditionalParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *HdfsAdditionalParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *HdfsAdditionalParams) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *HdfsAdditionalParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *HdfsAdditionalParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetNamenodeAddress

`func (o *HdfsAdditionalParams) GetNamenodeAddress() string`

GetNamenodeAddress returns the NamenodeAddress field if non-nil, zero value otherwise.

### GetNamenodeAddressOk

`func (o *HdfsAdditionalParams) GetNamenodeAddressOk() (*string, bool)`

GetNamenodeAddressOk returns a tuple with the NamenodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamenodeAddress

`func (o *HdfsAdditionalParams) SetNamenodeAddress(v string)`

SetNamenodeAddress sets NamenodeAddress field to given value.

### HasNamenodeAddress

`func (o *HdfsAdditionalParams) HasNamenodeAddress() bool`

HasNamenodeAddress returns a boolean if a field has been set.

### GetWebhdfsPort

`func (o *HdfsAdditionalParams) GetWebhdfsPort() int32`

GetWebhdfsPort returns the WebhdfsPort field if non-nil, zero value otherwise.

### GetWebhdfsPortOk

`func (o *HdfsAdditionalParams) GetWebhdfsPortOk() (*int32, bool)`

GetWebhdfsPortOk returns a tuple with the WebhdfsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhdfsPort

`func (o *HdfsAdditionalParams) SetWebhdfsPort(v int32)`

SetWebhdfsPort sets WebhdfsPort field to given value.

### HasWebhdfsPort

`func (o *HdfsAdditionalParams) HasWebhdfsPort() bool`

HasWebhdfsPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


