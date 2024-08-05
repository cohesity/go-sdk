# DownloadAgentRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AixParams** | Pointer to [**AixAgentParams**](AixAgentParams.md) |  | [optional] 
**LinuxParams** | Pointer to [**LinuxAgentParams**](LinuxAgentParams.md) |  | [optional] 
**MySqlParams** | Pointer to [**MySqlAgentParams**](MySqlAgentParams.md) |  | [optional] 
**Platform** | **string** | Specifies the platform for which agent needs to be downloaded. | 
**SapHanaParams** | Pointer to [**SapHanaAgentParams**](SapHanaAgentParams.md) |  | [optional] 
**SapOracleParams** | Pointer to [**SapOracleAgentParams**](SapOracleAgentParams.md) |  | [optional] 
**VmwareCDPFilterParams** | Pointer to [**VMWareCDPFilterParams**](VMWareCDPFilterParams.md) |  | [optional] 

## Methods

### NewDownloadAgentRequestParams

`func NewDownloadAgentRequestParams(platform string, ) *DownloadAgentRequestParams`

NewDownloadAgentRequestParams instantiates a new DownloadAgentRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadAgentRequestParamsWithDefaults

`func NewDownloadAgentRequestParamsWithDefaults() *DownloadAgentRequestParams`

NewDownloadAgentRequestParamsWithDefaults instantiates a new DownloadAgentRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAixParams

`func (o *DownloadAgentRequestParams) GetAixParams() AixAgentParams`

GetAixParams returns the AixParams field if non-nil, zero value otherwise.

### GetAixParamsOk

`func (o *DownloadAgentRequestParams) GetAixParamsOk() (*AixAgentParams, bool)`

GetAixParamsOk returns a tuple with the AixParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAixParams

`func (o *DownloadAgentRequestParams) SetAixParams(v AixAgentParams)`

SetAixParams sets AixParams field to given value.

### HasAixParams

`func (o *DownloadAgentRequestParams) HasAixParams() bool`

HasAixParams returns a boolean if a field has been set.

### GetLinuxParams

`func (o *DownloadAgentRequestParams) GetLinuxParams() LinuxAgentParams`

GetLinuxParams returns the LinuxParams field if non-nil, zero value otherwise.

### GetLinuxParamsOk

`func (o *DownloadAgentRequestParams) GetLinuxParamsOk() (*LinuxAgentParams, bool)`

GetLinuxParamsOk returns a tuple with the LinuxParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxParams

`func (o *DownloadAgentRequestParams) SetLinuxParams(v LinuxAgentParams)`

SetLinuxParams sets LinuxParams field to given value.

### HasLinuxParams

`func (o *DownloadAgentRequestParams) HasLinuxParams() bool`

HasLinuxParams returns a boolean if a field has been set.

### GetMySqlParams

`func (o *DownloadAgentRequestParams) GetMySqlParams() MySqlAgentParams`

GetMySqlParams returns the MySqlParams field if non-nil, zero value otherwise.

### GetMySqlParamsOk

`func (o *DownloadAgentRequestParams) GetMySqlParamsOk() (*MySqlAgentParams, bool)`

GetMySqlParamsOk returns a tuple with the MySqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMySqlParams

`func (o *DownloadAgentRequestParams) SetMySqlParams(v MySqlAgentParams)`

SetMySqlParams sets MySqlParams field to given value.

### HasMySqlParams

`func (o *DownloadAgentRequestParams) HasMySqlParams() bool`

HasMySqlParams returns a boolean if a field has been set.

### GetPlatform

`func (o *DownloadAgentRequestParams) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DownloadAgentRequestParams) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DownloadAgentRequestParams) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetSapHanaParams

`func (o *DownloadAgentRequestParams) GetSapHanaParams() SapHanaAgentParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *DownloadAgentRequestParams) GetSapHanaParamsOk() (*SapHanaAgentParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *DownloadAgentRequestParams) SetSapHanaParams(v SapHanaAgentParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *DownloadAgentRequestParams) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSapOracleParams

`func (o *DownloadAgentRequestParams) GetSapOracleParams() SapOracleAgentParams`

GetSapOracleParams returns the SapOracleParams field if non-nil, zero value otherwise.

### GetSapOracleParamsOk

`func (o *DownloadAgentRequestParams) GetSapOracleParamsOk() (*SapOracleAgentParams, bool)`

GetSapOracleParamsOk returns a tuple with the SapOracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapOracleParams

`func (o *DownloadAgentRequestParams) SetSapOracleParams(v SapOracleAgentParams)`

SetSapOracleParams sets SapOracleParams field to given value.

### HasSapOracleParams

`func (o *DownloadAgentRequestParams) HasSapOracleParams() bool`

HasSapOracleParams returns a boolean if a field has been set.

### GetVmwareCDPFilterParams

`func (o *DownloadAgentRequestParams) GetVmwareCDPFilterParams() VMWareCDPFilterParams`

GetVmwareCDPFilterParams returns the VmwareCDPFilterParams field if non-nil, zero value otherwise.

### GetVmwareCDPFilterParamsOk

`func (o *DownloadAgentRequestParams) GetVmwareCDPFilterParamsOk() (*VMWareCDPFilterParams, bool)`

GetVmwareCDPFilterParamsOk returns a tuple with the VmwareCDPFilterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareCDPFilterParams

`func (o *DownloadAgentRequestParams) SetVmwareCDPFilterParams(v VMWareCDPFilterParams)`

SetVmwareCDPFilterParams sets VmwareCDPFilterParams field to given value.

### HasVmwareCDPFilterParams

`func (o *DownloadAgentRequestParams) HasVmwareCDPFilterParams() bool`

HasVmwareCDPFilterParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


