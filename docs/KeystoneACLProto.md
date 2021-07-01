# KeystoneACLProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadGrantees** | Pointer to [**KeystoneACLProtoGrantees**](KeystoneACLProtoGrantees.md) |  | [optional] 
**WriteGrantees** | Pointer to [**KeystoneACLProtoGrantees**](KeystoneACLProtoGrantees.md) |  | [optional] 

## Methods

### NewKeystoneACLProto

`func NewKeystoneACLProto() *KeystoneACLProto`

NewKeystoneACLProto instantiates a new KeystoneACLProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneACLProtoWithDefaults

`func NewKeystoneACLProtoWithDefaults() *KeystoneACLProto`

NewKeystoneACLProtoWithDefaults instantiates a new KeystoneACLProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadGrantees

`func (o *KeystoneACLProto) GetReadGrantees() KeystoneACLProtoGrantees`

GetReadGrantees returns the ReadGrantees field if non-nil, zero value otherwise.

### GetReadGranteesOk

`func (o *KeystoneACLProto) GetReadGranteesOk() (*KeystoneACLProtoGrantees, bool)`

GetReadGranteesOk returns a tuple with the ReadGrantees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadGrantees

`func (o *KeystoneACLProto) SetReadGrantees(v KeystoneACLProtoGrantees)`

SetReadGrantees sets ReadGrantees field to given value.

### HasReadGrantees

`func (o *KeystoneACLProto) HasReadGrantees() bool`

HasReadGrantees returns a boolean if a field has been set.

### GetWriteGrantees

`func (o *KeystoneACLProto) GetWriteGrantees() KeystoneACLProtoGrantees`

GetWriteGrantees returns the WriteGrantees field if non-nil, zero value otherwise.

### GetWriteGranteesOk

`func (o *KeystoneACLProto) GetWriteGranteesOk() (*KeystoneACLProtoGrantees, bool)`

GetWriteGranteesOk returns a tuple with the WriteGrantees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteGrantees

`func (o *KeystoneACLProto) SetWriteGrantees(v KeystoneACLProtoGrantees)`

SetWriteGrantees sets WriteGrantees field to given value.

### HasWriteGrantees

`func (o *KeystoneACLProto) HasWriteGrantees() bool`

HasWriteGrantees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


