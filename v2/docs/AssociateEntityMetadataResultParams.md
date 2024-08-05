# AssociateEntityMetadataResultParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **int64** | Specifies the entity id of the object. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies an error message (if any) corresponding to this entityId. | [optional] 
**ErrorType** | Pointer to **NullableString** | Specifies an error type(if any) corresponding to this entityId. | [optional] 

## Methods

### NewAssociateEntityMetadataResultParams

`func NewAssociateEntityMetadataResultParams() *AssociateEntityMetadataResultParams`

NewAssociateEntityMetadataResultParams instantiates a new AssociateEntityMetadataResultParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociateEntityMetadataResultParamsWithDefaults

`func NewAssociateEntityMetadataResultParamsWithDefaults() *AssociateEntityMetadataResultParams`

NewAssociateEntityMetadataResultParamsWithDefaults instantiates a new AssociateEntityMetadataResultParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *AssociateEntityMetadataResultParams) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AssociateEntityMetadataResultParams) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AssociateEntityMetadataResultParams) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AssociateEntityMetadataResultParams) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AssociateEntityMetadataResultParams) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AssociateEntityMetadataResultParams) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AssociateEntityMetadataResultParams) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AssociateEntityMetadataResultParams) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AssociateEntityMetadataResultParams) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AssociateEntityMetadataResultParams) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorType

`func (o *AssociateEntityMetadataResultParams) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *AssociateEntityMetadataResultParams) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *AssociateEntityMetadataResultParams) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *AssociateEntityMetadataResultParams) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### SetErrorTypeNil

`func (o *AssociateEntityMetadataResultParams) SetErrorTypeNil(b bool)`

 SetErrorTypeNil sets the value for ErrorType to be an explicit nil

### UnsetErrorType
`func (o *AssociateEntityMetadataResultParams) UnsetErrorType()`

UnsetErrorType ensures that no value is present for ErrorType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


