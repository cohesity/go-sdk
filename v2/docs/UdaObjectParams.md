# UdaObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasEntitySupport** | Pointer to **NullableBool** | Specifies whether this Object belongs to a source having entity support. | [optional] 
**SourceType** | Pointer to **NullableString** | Specifies the source type for Universal Data Adapter object. | [optional] 

## Methods

### NewUdaObjectParams

`func NewUdaObjectParams() *UdaObjectParams`

NewUdaObjectParams instantiates a new UdaObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaObjectParamsWithDefaults

`func NewUdaObjectParamsWithDefaults() *UdaObjectParams`

NewUdaObjectParamsWithDefaults instantiates a new UdaObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasEntitySupport

`func (o *UdaObjectParams) GetHasEntitySupport() bool`

GetHasEntitySupport returns the HasEntitySupport field if non-nil, zero value otherwise.

### GetHasEntitySupportOk

`func (o *UdaObjectParams) GetHasEntitySupportOk() (*bool, bool)`

GetHasEntitySupportOk returns a tuple with the HasEntitySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitySupport

`func (o *UdaObjectParams) SetHasEntitySupport(v bool)`

SetHasEntitySupport sets HasEntitySupport field to given value.

### HasHasEntitySupport

`func (o *UdaObjectParams) HasHasEntitySupport() bool`

HasHasEntitySupport returns a boolean if a field has been set.

### SetHasEntitySupportNil

`func (o *UdaObjectParams) SetHasEntitySupportNil(b bool)`

 SetHasEntitySupportNil sets the value for HasEntitySupport to be an explicit nil

### UnsetHasEntitySupport
`func (o *UdaObjectParams) UnsetHasEntitySupport()`

UnsetHasEntitySupport ensures that no value is present for HasEntitySupport, not even an explicit nil
### GetSourceType

`func (o *UdaObjectParams) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UdaObjectParams) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UdaObjectParams) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *UdaObjectParams) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *UdaObjectParams) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *UdaObjectParams) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


