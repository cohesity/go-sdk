# MssqlObjectEntityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostInfo** | Pointer to [**HostInformation**](HostInformation.md) |  | [optional] 
**AagInfo** | Pointer to [**AAGInfo**](AAGInfo.md) |  | [optional] 

## Methods

### NewMssqlObjectEntityParams

`func NewMssqlObjectEntityParams() *MssqlObjectEntityParams`

NewMssqlObjectEntityParams instantiates a new MssqlObjectEntityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlObjectEntityParamsWithDefaults

`func NewMssqlObjectEntityParamsWithDefaults() *MssqlObjectEntityParams`

NewMssqlObjectEntityParamsWithDefaults instantiates a new MssqlObjectEntityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostInfo

`func (o *MssqlObjectEntityParams) GetHostInfo() HostInformation`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *MssqlObjectEntityParams) GetHostInfoOk() (*HostInformation, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *MssqlObjectEntityParams) SetHostInfo(v HostInformation)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *MssqlObjectEntityParams) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.

### GetAagInfo

`func (o *MssqlObjectEntityParams) GetAagInfo() AAGInfo`

GetAagInfo returns the AagInfo field if non-nil, zero value otherwise.

### GetAagInfoOk

`func (o *MssqlObjectEntityParams) GetAagInfoOk() (*AAGInfo, bool)`

GetAagInfoOk returns a tuple with the AagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagInfo

`func (o *MssqlObjectEntityParams) SetAagInfo(v AAGInfo)`

SetAagInfo sets AagInfo field to given value.

### HasAagInfo

`func (o *MssqlObjectEntityParams) HasAagInfo() bool`

HasAagInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


