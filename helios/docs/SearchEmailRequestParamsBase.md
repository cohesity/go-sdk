# SearchEmailRequestParamsBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to **[]string** | Specifies a list of mailbox item types. Only items within the given types will be returned. | [optional] 
**HasAttachment** | Pointer to **NullableBool** | Filters the emails which have attachment. | [optional] 
**SenderAddress** | Pointer to **NullableString** | Filters the emails which are received from specified User&#39;s email address. | [optional] 
**RecipientAddresses** | Pointer to **[]string** | Filters the emails which are sent to specified email addresses. | [optional] 
**CcRecipientAddresses** | Pointer to **[]string** | Filters the emails which are sent to specified email addresses in CC. | [optional] 
**BccRecipientAddresses** | Pointer to **[]string** | Filters the emails which are sent to specified email addresses in BCC. | [optional] 
**ReceivedStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds where the received time of the email is more than specified value. | [optional] 
**ReceivedEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds where the received time of the email is less than specified value. | [optional] 
**EmailSubject** | Pointer to **NullableString** | Filters the emails which have the specified text in its subject. | [optional] 
**FolderNames** | Pointer to **[]string** | Filters the emails which are categorized to specified folders. | [optional] 
**SourceEnvironment** | Pointer to **NullableString** | Specifies the source environment. | [optional] 

## Methods

### NewSearchEmailRequestParamsBase

`func NewSearchEmailRequestParamsBase() *SearchEmailRequestParamsBase`

NewSearchEmailRequestParamsBase instantiates a new SearchEmailRequestParamsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEmailRequestParamsBaseWithDefaults

`func NewSearchEmailRequestParamsBaseWithDefaults() *SearchEmailRequestParamsBase`

NewSearchEmailRequestParamsBaseWithDefaults instantiates a new SearchEmailRequestParamsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *SearchEmailRequestParamsBase) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchEmailRequestParamsBase) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchEmailRequestParamsBase) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchEmailRequestParamsBase) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *SearchEmailRequestParamsBase) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *SearchEmailRequestParamsBase) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetHasAttachment

`func (o *SearchEmailRequestParamsBase) GetHasAttachment() bool`

GetHasAttachment returns the HasAttachment field if non-nil, zero value otherwise.

### GetHasAttachmentOk

`func (o *SearchEmailRequestParamsBase) GetHasAttachmentOk() (*bool, bool)`

GetHasAttachmentOk returns a tuple with the HasAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachment

`func (o *SearchEmailRequestParamsBase) SetHasAttachment(v bool)`

SetHasAttachment sets HasAttachment field to given value.

### HasHasAttachment

`func (o *SearchEmailRequestParamsBase) HasHasAttachment() bool`

HasHasAttachment returns a boolean if a field has been set.

### SetHasAttachmentNil

`func (o *SearchEmailRequestParamsBase) SetHasAttachmentNil(b bool)`

 SetHasAttachmentNil sets the value for HasAttachment to be an explicit nil

### UnsetHasAttachment
`func (o *SearchEmailRequestParamsBase) UnsetHasAttachment()`

UnsetHasAttachment ensures that no value is present for HasAttachment, not even an explicit nil
### GetSenderAddress

`func (o *SearchEmailRequestParamsBase) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *SearchEmailRequestParamsBase) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *SearchEmailRequestParamsBase) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *SearchEmailRequestParamsBase) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### SetSenderAddressNil

`func (o *SearchEmailRequestParamsBase) SetSenderAddressNil(b bool)`

 SetSenderAddressNil sets the value for SenderAddress to be an explicit nil

### UnsetSenderAddress
`func (o *SearchEmailRequestParamsBase) UnsetSenderAddress()`

UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
### GetRecipientAddresses

`func (o *SearchEmailRequestParamsBase) GetRecipientAddresses() []string`

GetRecipientAddresses returns the RecipientAddresses field if non-nil, zero value otherwise.

### GetRecipientAddressesOk

`func (o *SearchEmailRequestParamsBase) GetRecipientAddressesOk() (*[]string, bool)`

GetRecipientAddressesOk returns a tuple with the RecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddresses

`func (o *SearchEmailRequestParamsBase) SetRecipientAddresses(v []string)`

SetRecipientAddresses sets RecipientAddresses field to given value.

### HasRecipientAddresses

`func (o *SearchEmailRequestParamsBase) HasRecipientAddresses() bool`

HasRecipientAddresses returns a boolean if a field has been set.

### SetRecipientAddressesNil

`func (o *SearchEmailRequestParamsBase) SetRecipientAddressesNil(b bool)`

 SetRecipientAddressesNil sets the value for RecipientAddresses to be an explicit nil

### UnsetRecipientAddresses
`func (o *SearchEmailRequestParamsBase) UnsetRecipientAddresses()`

UnsetRecipientAddresses ensures that no value is present for RecipientAddresses, not even an explicit nil
### GetCcRecipientAddresses

`func (o *SearchEmailRequestParamsBase) GetCcRecipientAddresses() []string`

GetCcRecipientAddresses returns the CcRecipientAddresses field if non-nil, zero value otherwise.

### GetCcRecipientAddressesOk

`func (o *SearchEmailRequestParamsBase) GetCcRecipientAddressesOk() (*[]string, bool)`

GetCcRecipientAddressesOk returns a tuple with the CcRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipientAddresses

`func (o *SearchEmailRequestParamsBase) SetCcRecipientAddresses(v []string)`

SetCcRecipientAddresses sets CcRecipientAddresses field to given value.

### HasCcRecipientAddresses

`func (o *SearchEmailRequestParamsBase) HasCcRecipientAddresses() bool`

HasCcRecipientAddresses returns a boolean if a field has been set.

### SetCcRecipientAddressesNil

`func (o *SearchEmailRequestParamsBase) SetCcRecipientAddressesNil(b bool)`

 SetCcRecipientAddressesNil sets the value for CcRecipientAddresses to be an explicit nil

### UnsetCcRecipientAddresses
`func (o *SearchEmailRequestParamsBase) UnsetCcRecipientAddresses()`

UnsetCcRecipientAddresses ensures that no value is present for CcRecipientAddresses, not even an explicit nil
### GetBccRecipientAddresses

`func (o *SearchEmailRequestParamsBase) GetBccRecipientAddresses() []string`

GetBccRecipientAddresses returns the BccRecipientAddresses field if non-nil, zero value otherwise.

### GetBccRecipientAddressesOk

`func (o *SearchEmailRequestParamsBase) GetBccRecipientAddressesOk() (*[]string, bool)`

GetBccRecipientAddressesOk returns a tuple with the BccRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipientAddresses

`func (o *SearchEmailRequestParamsBase) SetBccRecipientAddresses(v []string)`

SetBccRecipientAddresses sets BccRecipientAddresses field to given value.

### HasBccRecipientAddresses

`func (o *SearchEmailRequestParamsBase) HasBccRecipientAddresses() bool`

HasBccRecipientAddresses returns a boolean if a field has been set.

### SetBccRecipientAddressesNil

`func (o *SearchEmailRequestParamsBase) SetBccRecipientAddressesNil(b bool)`

 SetBccRecipientAddressesNil sets the value for BccRecipientAddresses to be an explicit nil

### UnsetBccRecipientAddresses
`func (o *SearchEmailRequestParamsBase) UnsetBccRecipientAddresses()`

UnsetBccRecipientAddresses ensures that no value is present for BccRecipientAddresses, not even an explicit nil
### GetReceivedStartTimeSecs

`func (o *SearchEmailRequestParamsBase) GetReceivedStartTimeSecs() int64`

GetReceivedStartTimeSecs returns the ReceivedStartTimeSecs field if non-nil, zero value otherwise.

### GetReceivedStartTimeSecsOk

`func (o *SearchEmailRequestParamsBase) GetReceivedStartTimeSecsOk() (*int64, bool)`

GetReceivedStartTimeSecsOk returns a tuple with the ReceivedStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedStartTimeSecs

`func (o *SearchEmailRequestParamsBase) SetReceivedStartTimeSecs(v int64)`

SetReceivedStartTimeSecs sets ReceivedStartTimeSecs field to given value.

### HasReceivedStartTimeSecs

`func (o *SearchEmailRequestParamsBase) HasReceivedStartTimeSecs() bool`

HasReceivedStartTimeSecs returns a boolean if a field has been set.

### SetReceivedStartTimeSecsNil

`func (o *SearchEmailRequestParamsBase) SetReceivedStartTimeSecsNil(b bool)`

 SetReceivedStartTimeSecsNil sets the value for ReceivedStartTimeSecs to be an explicit nil

### UnsetReceivedStartTimeSecs
`func (o *SearchEmailRequestParamsBase) UnsetReceivedStartTimeSecs()`

UnsetReceivedStartTimeSecs ensures that no value is present for ReceivedStartTimeSecs, not even an explicit nil
### GetReceivedEndTimeSecs

`func (o *SearchEmailRequestParamsBase) GetReceivedEndTimeSecs() int64`

GetReceivedEndTimeSecs returns the ReceivedEndTimeSecs field if non-nil, zero value otherwise.

### GetReceivedEndTimeSecsOk

`func (o *SearchEmailRequestParamsBase) GetReceivedEndTimeSecsOk() (*int64, bool)`

GetReceivedEndTimeSecsOk returns a tuple with the ReceivedEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEndTimeSecs

`func (o *SearchEmailRequestParamsBase) SetReceivedEndTimeSecs(v int64)`

SetReceivedEndTimeSecs sets ReceivedEndTimeSecs field to given value.

### HasReceivedEndTimeSecs

`func (o *SearchEmailRequestParamsBase) HasReceivedEndTimeSecs() bool`

HasReceivedEndTimeSecs returns a boolean if a field has been set.

### SetReceivedEndTimeSecsNil

`func (o *SearchEmailRequestParamsBase) SetReceivedEndTimeSecsNil(b bool)`

 SetReceivedEndTimeSecsNil sets the value for ReceivedEndTimeSecs to be an explicit nil

### UnsetReceivedEndTimeSecs
`func (o *SearchEmailRequestParamsBase) UnsetReceivedEndTimeSecs()`

UnsetReceivedEndTimeSecs ensures that no value is present for ReceivedEndTimeSecs, not even an explicit nil
### GetEmailSubject

`func (o *SearchEmailRequestParamsBase) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *SearchEmailRequestParamsBase) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *SearchEmailRequestParamsBase) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *SearchEmailRequestParamsBase) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *SearchEmailRequestParamsBase) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *SearchEmailRequestParamsBase) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetFolderNames

`func (o *SearchEmailRequestParamsBase) GetFolderNames() []string`

GetFolderNames returns the FolderNames field if non-nil, zero value otherwise.

### GetFolderNamesOk

`func (o *SearchEmailRequestParamsBase) GetFolderNamesOk() (*[]string, bool)`

GetFolderNamesOk returns a tuple with the FolderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderNames

`func (o *SearchEmailRequestParamsBase) SetFolderNames(v []string)`

SetFolderNames sets FolderNames field to given value.

### HasFolderNames

`func (o *SearchEmailRequestParamsBase) HasFolderNames() bool`

HasFolderNames returns a boolean if a field has been set.

### SetFolderNamesNil

`func (o *SearchEmailRequestParamsBase) SetFolderNamesNil(b bool)`

 SetFolderNamesNil sets the value for FolderNames to be an explicit nil

### UnsetFolderNames
`func (o *SearchEmailRequestParamsBase) UnsetFolderNames()`

UnsetFolderNames ensures that no value is present for FolderNames, not even an explicit nil
### GetSourceEnvironment

`func (o *SearchEmailRequestParamsBase) GetSourceEnvironment() string`

GetSourceEnvironment returns the SourceEnvironment field if non-nil, zero value otherwise.

### GetSourceEnvironmentOk

`func (o *SearchEmailRequestParamsBase) GetSourceEnvironmentOk() (*string, bool)`

GetSourceEnvironmentOk returns a tuple with the SourceEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironment

`func (o *SearchEmailRequestParamsBase) SetSourceEnvironment(v string)`

SetSourceEnvironment sets SourceEnvironment field to given value.

### HasSourceEnvironment

`func (o *SearchEmailRequestParamsBase) HasSourceEnvironment() bool`

HasSourceEnvironment returns a boolean if a field has been set.

### SetSourceEnvironmentNil

`func (o *SearchEmailRequestParamsBase) SetSourceEnvironmentNil(b bool)`

 SetSourceEnvironmentNil sets the value for SourceEnvironment to be an explicit nil

### UnsetSourceEnvironment
`func (o *SearchEmailRequestParamsBase) UnsetSourceEnvironment()`

UnsetSourceEnvironment ensures that no value is present for SourceEnvironment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


