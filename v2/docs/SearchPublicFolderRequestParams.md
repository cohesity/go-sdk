# SearchPublicFolderRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BccRecipientAddresses** | Pointer to **[]string** | Filters the public folder items which are sent to specified email addresses in BCC. | [optional] 
**CcRecipientAddresses** | Pointer to **[]string** | Filters the public folder items which are sent to specified email addresses in CC. | [optional] 
**HasAttachment** | Pointer to **NullableBool** | Filters the public folder items which have attachment. | [optional] 
**ReceivedEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds where the received time of the public folder items is less than specified value. | [optional] 
**ReceivedStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds where the received time of the public folder item is more than specified value. | [optional] 
**RecipientAddresses** | Pointer to **[]string** | Filters the public folder items which are sent to specified email addresses. | [optional] 
**SearchString** | Pointer to **NullableString** | Specifies the search string to filter the items. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all item names are matched with the prefix string. | [optional] 
**SenderAddress** | Pointer to **NullableString** | Filters the public folder items which are received from specified user&#39;s email address. | [optional] 
**Types** | Pointer to **[]string** | Specifies a list of public folder item types. Only items within the given types will be returned. | [optional] 

## Methods

### NewSearchPublicFolderRequestParams

`func NewSearchPublicFolderRequestParams() *SearchPublicFolderRequestParams`

NewSearchPublicFolderRequestParams instantiates a new SearchPublicFolderRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPublicFolderRequestParamsWithDefaults

`func NewSearchPublicFolderRequestParamsWithDefaults() *SearchPublicFolderRequestParams`

NewSearchPublicFolderRequestParamsWithDefaults instantiates a new SearchPublicFolderRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBccRecipientAddresses

`func (o *SearchPublicFolderRequestParams) GetBccRecipientAddresses() []string`

GetBccRecipientAddresses returns the BccRecipientAddresses field if non-nil, zero value otherwise.

### GetBccRecipientAddressesOk

`func (o *SearchPublicFolderRequestParams) GetBccRecipientAddressesOk() (*[]string, bool)`

GetBccRecipientAddressesOk returns a tuple with the BccRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipientAddresses

`func (o *SearchPublicFolderRequestParams) SetBccRecipientAddresses(v []string)`

SetBccRecipientAddresses sets BccRecipientAddresses field to given value.

### HasBccRecipientAddresses

`func (o *SearchPublicFolderRequestParams) HasBccRecipientAddresses() bool`

HasBccRecipientAddresses returns a boolean if a field has been set.

### SetBccRecipientAddressesNil

`func (o *SearchPublicFolderRequestParams) SetBccRecipientAddressesNil(b bool)`

 SetBccRecipientAddressesNil sets the value for BccRecipientAddresses to be an explicit nil

### UnsetBccRecipientAddresses
`func (o *SearchPublicFolderRequestParams) UnsetBccRecipientAddresses()`

UnsetBccRecipientAddresses ensures that no value is present for BccRecipientAddresses, not even an explicit nil
### GetCcRecipientAddresses

`func (o *SearchPublicFolderRequestParams) GetCcRecipientAddresses() []string`

GetCcRecipientAddresses returns the CcRecipientAddresses field if non-nil, zero value otherwise.

### GetCcRecipientAddressesOk

`func (o *SearchPublicFolderRequestParams) GetCcRecipientAddressesOk() (*[]string, bool)`

GetCcRecipientAddressesOk returns a tuple with the CcRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipientAddresses

`func (o *SearchPublicFolderRequestParams) SetCcRecipientAddresses(v []string)`

SetCcRecipientAddresses sets CcRecipientAddresses field to given value.

### HasCcRecipientAddresses

`func (o *SearchPublicFolderRequestParams) HasCcRecipientAddresses() bool`

HasCcRecipientAddresses returns a boolean if a field has been set.

### SetCcRecipientAddressesNil

`func (o *SearchPublicFolderRequestParams) SetCcRecipientAddressesNil(b bool)`

 SetCcRecipientAddressesNil sets the value for CcRecipientAddresses to be an explicit nil

### UnsetCcRecipientAddresses
`func (o *SearchPublicFolderRequestParams) UnsetCcRecipientAddresses()`

UnsetCcRecipientAddresses ensures that no value is present for CcRecipientAddresses, not even an explicit nil
### GetHasAttachment

`func (o *SearchPublicFolderRequestParams) GetHasAttachment() bool`

GetHasAttachment returns the HasAttachment field if non-nil, zero value otherwise.

### GetHasAttachmentOk

`func (o *SearchPublicFolderRequestParams) GetHasAttachmentOk() (*bool, bool)`

GetHasAttachmentOk returns a tuple with the HasAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachment

`func (o *SearchPublicFolderRequestParams) SetHasAttachment(v bool)`

SetHasAttachment sets HasAttachment field to given value.

### HasHasAttachment

`func (o *SearchPublicFolderRequestParams) HasHasAttachment() bool`

HasHasAttachment returns a boolean if a field has been set.

### SetHasAttachmentNil

`func (o *SearchPublicFolderRequestParams) SetHasAttachmentNil(b bool)`

 SetHasAttachmentNil sets the value for HasAttachment to be an explicit nil

### UnsetHasAttachment
`func (o *SearchPublicFolderRequestParams) UnsetHasAttachment()`

UnsetHasAttachment ensures that no value is present for HasAttachment, not even an explicit nil
### GetReceivedEndTimeSecs

`func (o *SearchPublicFolderRequestParams) GetReceivedEndTimeSecs() int64`

GetReceivedEndTimeSecs returns the ReceivedEndTimeSecs field if non-nil, zero value otherwise.

### GetReceivedEndTimeSecsOk

`func (o *SearchPublicFolderRequestParams) GetReceivedEndTimeSecsOk() (*int64, bool)`

GetReceivedEndTimeSecsOk returns a tuple with the ReceivedEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEndTimeSecs

`func (o *SearchPublicFolderRequestParams) SetReceivedEndTimeSecs(v int64)`

SetReceivedEndTimeSecs sets ReceivedEndTimeSecs field to given value.

### HasReceivedEndTimeSecs

`func (o *SearchPublicFolderRequestParams) HasReceivedEndTimeSecs() bool`

HasReceivedEndTimeSecs returns a boolean if a field has been set.

### SetReceivedEndTimeSecsNil

`func (o *SearchPublicFolderRequestParams) SetReceivedEndTimeSecsNil(b bool)`

 SetReceivedEndTimeSecsNil sets the value for ReceivedEndTimeSecs to be an explicit nil

### UnsetReceivedEndTimeSecs
`func (o *SearchPublicFolderRequestParams) UnsetReceivedEndTimeSecs()`

UnsetReceivedEndTimeSecs ensures that no value is present for ReceivedEndTimeSecs, not even an explicit nil
### GetReceivedStartTimeSecs

`func (o *SearchPublicFolderRequestParams) GetReceivedStartTimeSecs() int64`

GetReceivedStartTimeSecs returns the ReceivedStartTimeSecs field if non-nil, zero value otherwise.

### GetReceivedStartTimeSecsOk

`func (o *SearchPublicFolderRequestParams) GetReceivedStartTimeSecsOk() (*int64, bool)`

GetReceivedStartTimeSecsOk returns a tuple with the ReceivedStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedStartTimeSecs

`func (o *SearchPublicFolderRequestParams) SetReceivedStartTimeSecs(v int64)`

SetReceivedStartTimeSecs sets ReceivedStartTimeSecs field to given value.

### HasReceivedStartTimeSecs

`func (o *SearchPublicFolderRequestParams) HasReceivedStartTimeSecs() bool`

HasReceivedStartTimeSecs returns a boolean if a field has been set.

### SetReceivedStartTimeSecsNil

`func (o *SearchPublicFolderRequestParams) SetReceivedStartTimeSecsNil(b bool)`

 SetReceivedStartTimeSecsNil sets the value for ReceivedStartTimeSecs to be an explicit nil

### UnsetReceivedStartTimeSecs
`func (o *SearchPublicFolderRequestParams) UnsetReceivedStartTimeSecs()`

UnsetReceivedStartTimeSecs ensures that no value is present for ReceivedStartTimeSecs, not even an explicit nil
### GetRecipientAddresses

`func (o *SearchPublicFolderRequestParams) GetRecipientAddresses() []string`

GetRecipientAddresses returns the RecipientAddresses field if non-nil, zero value otherwise.

### GetRecipientAddressesOk

`func (o *SearchPublicFolderRequestParams) GetRecipientAddressesOk() (*[]string, bool)`

GetRecipientAddressesOk returns a tuple with the RecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddresses

`func (o *SearchPublicFolderRequestParams) SetRecipientAddresses(v []string)`

SetRecipientAddresses sets RecipientAddresses field to given value.

### HasRecipientAddresses

`func (o *SearchPublicFolderRequestParams) HasRecipientAddresses() bool`

HasRecipientAddresses returns a boolean if a field has been set.

### SetRecipientAddressesNil

`func (o *SearchPublicFolderRequestParams) SetRecipientAddressesNil(b bool)`

 SetRecipientAddressesNil sets the value for RecipientAddresses to be an explicit nil

### UnsetRecipientAddresses
`func (o *SearchPublicFolderRequestParams) UnsetRecipientAddresses()`

UnsetRecipientAddresses ensures that no value is present for RecipientAddresses, not even an explicit nil
### GetSearchString

`func (o *SearchPublicFolderRequestParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *SearchPublicFolderRequestParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *SearchPublicFolderRequestParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.

### HasSearchString

`func (o *SearchPublicFolderRequestParams) HasSearchString() bool`

HasSearchString returns a boolean if a field has been set.

### SetSearchStringNil

`func (o *SearchPublicFolderRequestParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *SearchPublicFolderRequestParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetSenderAddress

`func (o *SearchPublicFolderRequestParams) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *SearchPublicFolderRequestParams) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *SearchPublicFolderRequestParams) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *SearchPublicFolderRequestParams) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### SetSenderAddressNil

`func (o *SearchPublicFolderRequestParams) SetSenderAddressNil(b bool)`

 SetSenderAddressNil sets the value for SenderAddress to be an explicit nil

### UnsetSenderAddress
`func (o *SearchPublicFolderRequestParams) UnsetSenderAddress()`

UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
### GetTypes

`func (o *SearchPublicFolderRequestParams) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchPublicFolderRequestParams) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchPublicFolderRequestParams) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchPublicFolderRequestParams) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *SearchPublicFolderRequestParams) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *SearchPublicFolderRequestParams) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


