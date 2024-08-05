# ViewProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Mode of protocol access.   &#39;ReadOnly&#39;   &#39;ReadWrite&#39; | [optional] 
**Type** | Pointer to **string** | Type of protocol. Specifies the supported Protocols for the View.   &#39;NFS&#39; enables protocol access to NFS v3.   &#39;NFS4&#39; enables protocol access to NFS v4.1.   &#39;SMB&#39; enables protocol access to SMB.   &#39;S3&#39; enables protocol access to S3.   &#39;Swift&#39; enables protocol access to Swift. | [optional] 

## Methods

### NewViewProtocol

`func NewViewProtocol() *ViewProtocol`

NewViewProtocol instantiates a new ViewProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProtocolWithDefaults

`func NewViewProtocolWithDefaults() *ViewProtocol`

NewViewProtocolWithDefaults instantiates a new ViewProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ViewProtocol) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ViewProtocol) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ViewProtocol) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ViewProtocol) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetType

`func (o *ViewProtocol) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewProtocol) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewProtocol) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewProtocol) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


