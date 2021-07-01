# HBaseRecoverJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdfsConnectParams** | Pointer to [**HdfsConnectParams**](HdfsConnectParams.md) |  | [optional] 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered entities | [optional] 

## Methods

### NewHBaseRecoverJobParams

`func NewHBaseRecoverJobParams() *HBaseRecoverJobParams`

NewHBaseRecoverJobParams instantiates a new HBaseRecoverJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHBaseRecoverJobParamsWithDefaults

`func NewHBaseRecoverJobParamsWithDefaults() *HBaseRecoverJobParams`

NewHBaseRecoverJobParamsWithDefaults instantiates a new HBaseRecoverJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfsConnectParams

`func (o *HBaseRecoverJobParams) GetHdfsConnectParams() HdfsConnectParams`

GetHdfsConnectParams returns the HdfsConnectParams field if non-nil, zero value otherwise.

### GetHdfsConnectParamsOk

`func (o *HBaseRecoverJobParams) GetHdfsConnectParamsOk() (*HdfsConnectParams, bool)`

GetHdfsConnectParamsOk returns a tuple with the HdfsConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnectParams

`func (o *HBaseRecoverJobParams) SetHdfsConnectParams(v HdfsConnectParams)`

SetHdfsConnectParams sets HdfsConnectParams field to given value.

### HasHdfsConnectParams

`func (o *HBaseRecoverJobParams) HasHdfsConnectParams() bool`

HasHdfsConnectParams returns a boolean if a field has been set.

### GetSuffix

`func (o *HBaseRecoverJobParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *HBaseRecoverJobParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *HBaseRecoverJobParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *HBaseRecoverJobParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *HBaseRecoverJobParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *HBaseRecoverJobParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


