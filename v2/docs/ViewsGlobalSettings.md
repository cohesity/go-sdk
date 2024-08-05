# ViewsGlobalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableRemoteViewsGuiVisibility** | Pointer to **NullableBool** | Specifies the visibility of Remote Cohesity Views on Cohesity GUI. | [optional] 
**EnableRemoteViewsVisibility** | Pointer to **NullableBool** | Specifies the visibility of Remote Cohesity Views for external clients. | [optional] 
**EnableSmbAuth** | Pointer to **NullableBool** | Specifies if SMB Authentication should be enabled. | [optional] 
**EnableSmbMultiChannel** | Pointer to **NullableBool** | Specifies if SMB Multi-Channel should be enabled. | [optional] 
**S3VirtualHostedDomainNames** | Pointer to **[]string** | Specifies the list of domain names for S3 Virtual Hosted Style Paths. If set, all the Cohesity S3 Views in the cluster can be accessed using any of the specified domain names. | [optional] 

## Methods

### NewViewsGlobalSettings

`func NewViewsGlobalSettings() *ViewsGlobalSettings`

NewViewsGlobalSettings instantiates a new ViewsGlobalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsGlobalSettingsWithDefaults

`func NewViewsGlobalSettingsWithDefaults() *ViewsGlobalSettings`

NewViewsGlobalSettingsWithDefaults instantiates a new ViewsGlobalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableRemoteViewsGuiVisibility

`func (o *ViewsGlobalSettings) GetEnableRemoteViewsGuiVisibility() bool`

GetEnableRemoteViewsGuiVisibility returns the EnableRemoteViewsGuiVisibility field if non-nil, zero value otherwise.

### GetEnableRemoteViewsGuiVisibilityOk

`func (o *ViewsGlobalSettings) GetEnableRemoteViewsGuiVisibilityOk() (*bool, bool)`

GetEnableRemoteViewsGuiVisibilityOk returns a tuple with the EnableRemoteViewsGuiVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteViewsGuiVisibility

`func (o *ViewsGlobalSettings) SetEnableRemoteViewsGuiVisibility(v bool)`

SetEnableRemoteViewsGuiVisibility sets EnableRemoteViewsGuiVisibility field to given value.

### HasEnableRemoteViewsGuiVisibility

`func (o *ViewsGlobalSettings) HasEnableRemoteViewsGuiVisibility() bool`

HasEnableRemoteViewsGuiVisibility returns a boolean if a field has been set.

### SetEnableRemoteViewsGuiVisibilityNil

`func (o *ViewsGlobalSettings) SetEnableRemoteViewsGuiVisibilityNil(b bool)`

 SetEnableRemoteViewsGuiVisibilityNil sets the value for EnableRemoteViewsGuiVisibility to be an explicit nil

### UnsetEnableRemoteViewsGuiVisibility
`func (o *ViewsGlobalSettings) UnsetEnableRemoteViewsGuiVisibility()`

UnsetEnableRemoteViewsGuiVisibility ensures that no value is present for EnableRemoteViewsGuiVisibility, not even an explicit nil
### GetEnableRemoteViewsVisibility

`func (o *ViewsGlobalSettings) GetEnableRemoteViewsVisibility() bool`

GetEnableRemoteViewsVisibility returns the EnableRemoteViewsVisibility field if non-nil, zero value otherwise.

### GetEnableRemoteViewsVisibilityOk

`func (o *ViewsGlobalSettings) GetEnableRemoteViewsVisibilityOk() (*bool, bool)`

GetEnableRemoteViewsVisibilityOk returns a tuple with the EnableRemoteViewsVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteViewsVisibility

`func (o *ViewsGlobalSettings) SetEnableRemoteViewsVisibility(v bool)`

SetEnableRemoteViewsVisibility sets EnableRemoteViewsVisibility field to given value.

### HasEnableRemoteViewsVisibility

`func (o *ViewsGlobalSettings) HasEnableRemoteViewsVisibility() bool`

HasEnableRemoteViewsVisibility returns a boolean if a field has been set.

### SetEnableRemoteViewsVisibilityNil

`func (o *ViewsGlobalSettings) SetEnableRemoteViewsVisibilityNil(b bool)`

 SetEnableRemoteViewsVisibilityNil sets the value for EnableRemoteViewsVisibility to be an explicit nil

### UnsetEnableRemoteViewsVisibility
`func (o *ViewsGlobalSettings) UnsetEnableRemoteViewsVisibility()`

UnsetEnableRemoteViewsVisibility ensures that no value is present for EnableRemoteViewsVisibility, not even an explicit nil
### GetEnableSmbAuth

`func (o *ViewsGlobalSettings) GetEnableSmbAuth() bool`

GetEnableSmbAuth returns the EnableSmbAuth field if non-nil, zero value otherwise.

### GetEnableSmbAuthOk

`func (o *ViewsGlobalSettings) GetEnableSmbAuthOk() (*bool, bool)`

GetEnableSmbAuthOk returns a tuple with the EnableSmbAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbAuth

`func (o *ViewsGlobalSettings) SetEnableSmbAuth(v bool)`

SetEnableSmbAuth sets EnableSmbAuth field to given value.

### HasEnableSmbAuth

`func (o *ViewsGlobalSettings) HasEnableSmbAuth() bool`

HasEnableSmbAuth returns a boolean if a field has been set.

### SetEnableSmbAuthNil

`func (o *ViewsGlobalSettings) SetEnableSmbAuthNil(b bool)`

 SetEnableSmbAuthNil sets the value for EnableSmbAuth to be an explicit nil

### UnsetEnableSmbAuth
`func (o *ViewsGlobalSettings) UnsetEnableSmbAuth()`

UnsetEnableSmbAuth ensures that no value is present for EnableSmbAuth, not even an explicit nil
### GetEnableSmbMultiChannel

`func (o *ViewsGlobalSettings) GetEnableSmbMultiChannel() bool`

GetEnableSmbMultiChannel returns the EnableSmbMultiChannel field if non-nil, zero value otherwise.

### GetEnableSmbMultiChannelOk

`func (o *ViewsGlobalSettings) GetEnableSmbMultiChannelOk() (*bool, bool)`

GetEnableSmbMultiChannelOk returns a tuple with the EnableSmbMultiChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbMultiChannel

`func (o *ViewsGlobalSettings) SetEnableSmbMultiChannel(v bool)`

SetEnableSmbMultiChannel sets EnableSmbMultiChannel field to given value.

### HasEnableSmbMultiChannel

`func (o *ViewsGlobalSettings) HasEnableSmbMultiChannel() bool`

HasEnableSmbMultiChannel returns a boolean if a field has been set.

### SetEnableSmbMultiChannelNil

`func (o *ViewsGlobalSettings) SetEnableSmbMultiChannelNil(b bool)`

 SetEnableSmbMultiChannelNil sets the value for EnableSmbMultiChannel to be an explicit nil

### UnsetEnableSmbMultiChannel
`func (o *ViewsGlobalSettings) UnsetEnableSmbMultiChannel()`

UnsetEnableSmbMultiChannel ensures that no value is present for EnableSmbMultiChannel, not even an explicit nil
### GetS3VirtualHostedDomainNames

`func (o *ViewsGlobalSettings) GetS3VirtualHostedDomainNames() []string`

GetS3VirtualHostedDomainNames returns the S3VirtualHostedDomainNames field if non-nil, zero value otherwise.

### GetS3VirtualHostedDomainNamesOk

`func (o *ViewsGlobalSettings) GetS3VirtualHostedDomainNamesOk() (*[]string, bool)`

GetS3VirtualHostedDomainNamesOk returns a tuple with the S3VirtualHostedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3VirtualHostedDomainNames

`func (o *ViewsGlobalSettings) SetS3VirtualHostedDomainNames(v []string)`

SetS3VirtualHostedDomainNames sets S3VirtualHostedDomainNames field to given value.

### HasS3VirtualHostedDomainNames

`func (o *ViewsGlobalSettings) HasS3VirtualHostedDomainNames() bool`

HasS3VirtualHostedDomainNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


