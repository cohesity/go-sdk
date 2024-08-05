# GcpProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NativeProtectionTypeParams** | Pointer to [**GcpNativeProtectionGroupParams**](GcpNativeProtectionGroupParams.md) |  | [optional] 
**ProtectionType** | **string** | Specifies the GCP Protection Group type. | 

## Methods

### NewGcpProtectionGroupParams

`func NewGcpProtectionGroupParams(protectionType string, ) *GcpProtectionGroupParams`

NewGcpProtectionGroupParams instantiates a new GcpProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpProtectionGroupParamsWithDefaults

`func NewGcpProtectionGroupParamsWithDefaults() *GcpProtectionGroupParams`

NewGcpProtectionGroupParamsWithDefaults instantiates a new GcpProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNativeProtectionTypeParams

`func (o *GcpProtectionGroupParams) GetNativeProtectionTypeParams() GcpNativeProtectionGroupParams`

GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field if non-nil, zero value otherwise.

### GetNativeProtectionTypeParamsOk

`func (o *GcpProtectionGroupParams) GetNativeProtectionTypeParamsOk() (*GcpNativeProtectionGroupParams, bool)`

GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeProtectionTypeParams

`func (o *GcpProtectionGroupParams) SetNativeProtectionTypeParams(v GcpNativeProtectionGroupParams)`

SetNativeProtectionTypeParams sets NativeProtectionTypeParams field to given value.

### HasNativeProtectionTypeParams

`func (o *GcpProtectionGroupParams) HasNativeProtectionTypeParams() bool`

HasNativeProtectionTypeParams returns a boolean if a field has been set.

### GetProtectionType

`func (o *GcpProtectionGroupParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *GcpProtectionGroupParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *GcpProtectionGroupParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


