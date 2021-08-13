# OracleObjectEntityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseEntityInfo** | Pointer to [**DatabaseEntityInfo**](DatabaseEntityInfo.md) |  | [optional] 
**HostInfo** | Pointer to [**HostInformation**](HostInformation.md) |  | [optional] 

## Methods

### NewOracleObjectEntityParams

`func NewOracleObjectEntityParams() *OracleObjectEntityParams`

NewOracleObjectEntityParams instantiates a new OracleObjectEntityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleObjectEntityParamsWithDefaults

`func NewOracleObjectEntityParamsWithDefaults() *OracleObjectEntityParams`

NewOracleObjectEntityParamsWithDefaults instantiates a new OracleObjectEntityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseEntityInfo

`func (o *OracleObjectEntityParams) GetDatabaseEntityInfo() DatabaseEntityInfo`

GetDatabaseEntityInfo returns the DatabaseEntityInfo field if non-nil, zero value otherwise.

### GetDatabaseEntityInfoOk

`func (o *OracleObjectEntityParams) GetDatabaseEntityInfoOk() (*DatabaseEntityInfo, bool)`

GetDatabaseEntityInfoOk returns a tuple with the DatabaseEntityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEntityInfo

`func (o *OracleObjectEntityParams) SetDatabaseEntityInfo(v DatabaseEntityInfo)`

SetDatabaseEntityInfo sets DatabaseEntityInfo field to given value.

### HasDatabaseEntityInfo

`func (o *OracleObjectEntityParams) HasDatabaseEntityInfo() bool`

HasDatabaseEntityInfo returns a boolean if a field has been set.

### GetHostInfo

`func (o *OracleObjectEntityParams) GetHostInfo() HostInformation`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *OracleObjectEntityParams) GetHostInfoOk() (*HostInformation, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *OracleObjectEntityParams) SetHostInfo(v HostInformation)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *OracleObjectEntityParams) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


