# ChannelParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentLibraryParams** | Pointer to [**[]OneDriveParam**](OneDriveParam.md) | Specifies the list of doclibs of the Channel to recover. It is populated iff recoverEntireChannel is false. | [optional] 
**Id** | **NullableString** | Specifies the Channel id. | 
**Name** | Pointer to **NullableString** | Specifies the Channel name. | [optional] 
**RecoverEntireChannel** | **NullableBool** | Specifies whether to recover the whole Microsoft 365 Channel. | 
**Type** | Pointer to **string** | Specifies the type of channel public or private | [optional] 

## Methods

### NewChannelParam

`func NewChannelParam(id NullableString, recoverEntireChannel NullableBool, ) *ChannelParam`

NewChannelParam instantiates a new ChannelParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelParamWithDefaults

`func NewChannelParamWithDefaults() *ChannelParam`

NewChannelParamWithDefaults instantiates a new ChannelParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentLibraryParams

`func (o *ChannelParam) GetDocumentLibraryParams() []OneDriveParam`

GetDocumentLibraryParams returns the DocumentLibraryParams field if non-nil, zero value otherwise.

### GetDocumentLibraryParamsOk

`func (o *ChannelParam) GetDocumentLibraryParamsOk() (*[]OneDriveParam, bool)`

GetDocumentLibraryParamsOk returns a tuple with the DocumentLibraryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLibraryParams

`func (o *ChannelParam) SetDocumentLibraryParams(v []OneDriveParam)`

SetDocumentLibraryParams sets DocumentLibraryParams field to given value.

### HasDocumentLibraryParams

`func (o *ChannelParam) HasDocumentLibraryParams() bool`

HasDocumentLibraryParams returns a boolean if a field has been set.

### SetDocumentLibraryParamsNil

`func (o *ChannelParam) SetDocumentLibraryParamsNil(b bool)`

 SetDocumentLibraryParamsNil sets the value for DocumentLibraryParams to be an explicit nil

### UnsetDocumentLibraryParams
`func (o *ChannelParam) UnsetDocumentLibraryParams()`

UnsetDocumentLibraryParams ensures that no value is present for DocumentLibraryParams, not even an explicit nil
### GetId

`func (o *ChannelParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelParam) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ChannelParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ChannelParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ChannelParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelParam) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ChannelParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ChannelParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRecoverEntireChannel

`func (o *ChannelParam) GetRecoverEntireChannel() bool`

GetRecoverEntireChannel returns the RecoverEntireChannel field if non-nil, zero value otherwise.

### GetRecoverEntireChannelOk

`func (o *ChannelParam) GetRecoverEntireChannelOk() (*bool, bool)`

GetRecoverEntireChannelOk returns a tuple with the RecoverEntireChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireChannel

`func (o *ChannelParam) SetRecoverEntireChannel(v bool)`

SetRecoverEntireChannel sets RecoverEntireChannel field to given value.


### SetRecoverEntireChannelNil

`func (o *ChannelParam) SetRecoverEntireChannelNil(b bool)`

 SetRecoverEntireChannelNil sets the value for RecoverEntireChannel to be an explicit nil

### UnsetRecoverEntireChannel
`func (o *ChannelParam) UnsetRecoverEntireChannel()`

UnsetRecoverEntireChannel ensures that no value is present for RecoverEntireChannel, not even an explicit nil
### GetType

`func (o *ChannelParam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelParam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelParam) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChannelParam) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


