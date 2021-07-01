# HiveRecoverJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdfsConnectParams** | Pointer to [**HdfsConnectParams**](HdfsConnectParams.md) |  | [optional] 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered entities | [optional] 

## Methods

### NewHiveRecoverJobParams

`func NewHiveRecoverJobParams() *HiveRecoverJobParams`

NewHiveRecoverJobParams instantiates a new HiveRecoverJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveRecoverJobParamsWithDefaults

`func NewHiveRecoverJobParamsWithDefaults() *HiveRecoverJobParams`

NewHiveRecoverJobParamsWithDefaults instantiates a new HiveRecoverJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfsConnectParams

`func (o *HiveRecoverJobParams) GetHdfsConnectParams() HdfsConnectParams`

GetHdfsConnectParams returns the HdfsConnectParams field if non-nil, zero value otherwise.

### GetHdfsConnectParamsOk

`func (o *HiveRecoverJobParams) GetHdfsConnectParamsOk() (*HdfsConnectParams, bool)`

GetHdfsConnectParamsOk returns a tuple with the HdfsConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnectParams

`func (o *HiveRecoverJobParams) SetHdfsConnectParams(v HdfsConnectParams)`

SetHdfsConnectParams sets HdfsConnectParams field to given value.

### HasHdfsConnectParams

`func (o *HiveRecoverJobParams) HasHdfsConnectParams() bool`

HasHdfsConnectParams returns a boolean if a field has been set.

### GetSuffix

`func (o *HiveRecoverJobParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *HiveRecoverJobParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *HiveRecoverJobParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *HiveRecoverJobParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *HiveRecoverJobParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *HiveRecoverJobParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


