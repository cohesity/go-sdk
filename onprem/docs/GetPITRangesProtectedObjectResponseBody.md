# GetPITRangesProtectedObjectResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment for which restore ranges are returned. | [optional] 
**OracleRestoreRangeInfo** | Pointer to [**OracleRestoreRangeInfo**](OracleRestoreRangeInfo.md) |  | [optional] 

## Methods

### NewGetPITRangesProtectedObjectResponseBody

`func NewGetPITRangesProtectedObjectResponseBody() *GetPITRangesProtectedObjectResponseBody`

NewGetPITRangesProtectedObjectResponseBody instantiates a new GetPITRangesProtectedObjectResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPITRangesProtectedObjectResponseBodyWithDefaults

`func NewGetPITRangesProtectedObjectResponseBodyWithDefaults() *GetPITRangesProtectedObjectResponseBody`

NewGetPITRangesProtectedObjectResponseBodyWithDefaults instantiates a new GetPITRangesProtectedObjectResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *GetPITRangesProtectedObjectResponseBody) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetPITRangesProtectedObjectResponseBody) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetPITRangesProtectedObjectResponseBody) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetPITRangesProtectedObjectResponseBody) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *GetPITRangesProtectedObjectResponseBody) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *GetPITRangesProtectedObjectResponseBody) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetOracleRestoreRangeInfo

`func (o *GetPITRangesProtectedObjectResponseBody) GetOracleRestoreRangeInfo() OracleRestoreRangeInfo`

GetOracleRestoreRangeInfo returns the OracleRestoreRangeInfo field if non-nil, zero value otherwise.

### GetOracleRestoreRangeInfoOk

`func (o *GetPITRangesProtectedObjectResponseBody) GetOracleRestoreRangeInfoOk() (*OracleRestoreRangeInfo, bool)`

GetOracleRestoreRangeInfoOk returns a tuple with the OracleRestoreRangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRestoreRangeInfo

`func (o *GetPITRangesProtectedObjectResponseBody) SetOracleRestoreRangeInfo(v OracleRestoreRangeInfo)`

SetOracleRestoreRangeInfo sets OracleRestoreRangeInfo field to given value.

### HasOracleRestoreRangeInfo

`func (o *GetPITRangesProtectedObjectResponseBody) HasOracleRestoreRangeInfo() bool`

HasOracleRestoreRangeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


