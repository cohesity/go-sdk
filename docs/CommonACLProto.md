# CommonACLProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadGrantees** | Pointer to [**CommonACLProtoGrantees**](CommonACLProtoGrantees.md) |  | [optional] 
**WriteRlistings** | Pointer to **NullableBool** | Write permissions granted to grantees. | [optional] 

## Methods

### NewCommonACLProto

`func NewCommonACLProto() *CommonACLProto`

NewCommonACLProto instantiates a new CommonACLProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonACLProtoWithDefaults

`func NewCommonACLProtoWithDefaults() *CommonACLProto`

NewCommonACLProtoWithDefaults instantiates a new CommonACLProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadGrantees

`func (o *CommonACLProto) GetReadGrantees() CommonACLProtoGrantees`

GetReadGrantees returns the ReadGrantees field if non-nil, zero value otherwise.

### GetReadGranteesOk

`func (o *CommonACLProto) GetReadGranteesOk() (*CommonACLProtoGrantees, bool)`

GetReadGranteesOk returns a tuple with the ReadGrantees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadGrantees

`func (o *CommonACLProto) SetReadGrantees(v CommonACLProtoGrantees)`

SetReadGrantees sets ReadGrantees field to given value.

### HasReadGrantees

`func (o *CommonACLProto) HasReadGrantees() bool`

HasReadGrantees returns a boolean if a field has been set.

### GetWriteRlistings

`func (o *CommonACLProto) GetWriteRlistings() bool`

GetWriteRlistings returns the WriteRlistings field if non-nil, zero value otherwise.

### GetWriteRlistingsOk

`func (o *CommonACLProto) GetWriteRlistingsOk() (*bool, bool)`

GetWriteRlistingsOk returns a tuple with the WriteRlistings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteRlistings

`func (o *CommonACLProto) SetWriteRlistings(v bool)`

SetWriteRlistings sets WriteRlistings field to given value.

### HasWriteRlistings

`func (o *CommonACLProto) HasWriteRlistings() bool`

HasWriteRlistings returns a boolean if a field has been set.

### SetWriteRlistingsNil

`func (o *CommonACLProto) SetWriteRlistingsNil(b bool)`

 SetWriteRlistingsNil sets the value for WriteRlistings to be an explicit nil

### UnsetWriteRlistings
`func (o *CommonACLProto) UnsetWriteRlistings()`

UnsetWriteRlistings ensures that no value is present for WriteRlistings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


