# M365SelfServiceSecurityGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableDownload** | Pointer to **NullableBool** | Specifies whether the Security Group members are allowed to download content of their Mailbox or OneDrive. | [optional] 
**GlobalId** | **string** | Specifies the Global ID of this Security Group | 
**Name** | Pointer to **string** | Specifies the name of the Security Group | [optional] 

## Methods

### NewM365SelfServiceSecurityGroupInfo

`func NewM365SelfServiceSecurityGroupInfo(globalId string, ) *M365SelfServiceSecurityGroupInfo`

NewM365SelfServiceSecurityGroupInfo instantiates a new M365SelfServiceSecurityGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM365SelfServiceSecurityGroupInfoWithDefaults

`func NewM365SelfServiceSecurityGroupInfoWithDefaults() *M365SelfServiceSecurityGroupInfo`

NewM365SelfServiceSecurityGroupInfoWithDefaults instantiates a new M365SelfServiceSecurityGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableDownload

`func (o *M365SelfServiceSecurityGroupInfo) GetEnableDownload() bool`

GetEnableDownload returns the EnableDownload field if non-nil, zero value otherwise.

### GetEnableDownloadOk

`func (o *M365SelfServiceSecurityGroupInfo) GetEnableDownloadOk() (*bool, bool)`

GetEnableDownloadOk returns a tuple with the EnableDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDownload

`func (o *M365SelfServiceSecurityGroupInfo) SetEnableDownload(v bool)`

SetEnableDownload sets EnableDownload field to given value.

### HasEnableDownload

`func (o *M365SelfServiceSecurityGroupInfo) HasEnableDownload() bool`

HasEnableDownload returns a boolean if a field has been set.

### SetEnableDownloadNil

`func (o *M365SelfServiceSecurityGroupInfo) SetEnableDownloadNil(b bool)`

 SetEnableDownloadNil sets the value for EnableDownload to be an explicit nil

### UnsetEnableDownload
`func (o *M365SelfServiceSecurityGroupInfo) UnsetEnableDownload()`

UnsetEnableDownload ensures that no value is present for EnableDownload, not even an explicit nil
### GetGlobalId

`func (o *M365SelfServiceSecurityGroupInfo) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *M365SelfServiceSecurityGroupInfo) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *M365SelfServiceSecurityGroupInfo) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetName

`func (o *M365SelfServiceSecurityGroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *M365SelfServiceSecurityGroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *M365SelfServiceSecurityGroupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *M365SelfServiceSecurityGroupInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


