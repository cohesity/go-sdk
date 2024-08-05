# TargetMailboxParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the id of the target mailbox. Atleast one of id or primarySMTPAddress need to be defined. In case both id and primarySMTPAddress are defined then id takes precedence. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the domain for alternate domain recovery. | [optional] 
**PrimarySmtpAddress** | Pointer to **NullableString** | Specifies the primary SMTP address of the target mailbox. Atleast one of id or primarySMTPAddress needs to be defined. In case both id and primarySMTPAddress are defined then id takes precedence. | [optional] 
**TargetFolderPath** | Pointer to **NullableString** | Specifies the path to the target folder. | [optional] 

## Methods

### NewTargetMailboxParam

`func NewTargetMailboxParam() *TargetMailboxParam`

NewTargetMailboxParam instantiates a new TargetMailboxParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetMailboxParamWithDefaults

`func NewTargetMailboxParamWithDefaults() *TargetMailboxParam`

NewTargetMailboxParamWithDefaults instantiates a new TargetMailboxParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TargetMailboxParam) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetMailboxParam) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetMailboxParam) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TargetMailboxParam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TargetMailboxParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TargetMailboxParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TargetMailboxParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetMailboxParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetMailboxParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetMailboxParam) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TargetMailboxParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TargetMailboxParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentSourceId

`func (o *TargetMailboxParam) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *TargetMailboxParam) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *TargetMailboxParam) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *TargetMailboxParam) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *TargetMailboxParam) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *TargetMailboxParam) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetPrimarySmtpAddress

`func (o *TargetMailboxParam) GetPrimarySmtpAddress() string`

GetPrimarySmtpAddress returns the PrimarySmtpAddress field if non-nil, zero value otherwise.

### GetPrimarySmtpAddressOk

`func (o *TargetMailboxParam) GetPrimarySmtpAddressOk() (*string, bool)`

GetPrimarySmtpAddressOk returns a tuple with the PrimarySmtpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySmtpAddress

`func (o *TargetMailboxParam) SetPrimarySmtpAddress(v string)`

SetPrimarySmtpAddress sets PrimarySmtpAddress field to given value.

### HasPrimarySmtpAddress

`func (o *TargetMailboxParam) HasPrimarySmtpAddress() bool`

HasPrimarySmtpAddress returns a boolean if a field has been set.

### SetPrimarySmtpAddressNil

`func (o *TargetMailboxParam) SetPrimarySmtpAddressNil(b bool)`

 SetPrimarySmtpAddressNil sets the value for PrimarySmtpAddress to be an explicit nil

### UnsetPrimarySmtpAddress
`func (o *TargetMailboxParam) UnsetPrimarySmtpAddress()`

UnsetPrimarySmtpAddress ensures that no value is present for PrimarySmtpAddress, not even an explicit nil
### GetTargetFolderPath

`func (o *TargetMailboxParam) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *TargetMailboxParam) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *TargetMailboxParam) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *TargetMailboxParam) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *TargetMailboxParam) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *TargetMailboxParam) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


