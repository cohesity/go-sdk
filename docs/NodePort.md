# NodePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUiPort** | Pointer to **NullableBool** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Tag** | Pointer to **NullableString** | Specifies use of the nodeport kDefault - No specific service. kHttp - HTTP server. kHttps -  Secure HTTP server. kSsh - Secure shell server. | [optional] 

## Methods

### NewNodePort

`func NewNodePort() *NodePort`

NewNodePort instantiates a new NodePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePortWithDefaults

`func NewNodePortWithDefaults() *NodePort`

NewNodePortWithDefaults instantiates a new NodePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUiPort

`func (o *NodePort) GetIsUiPort() bool`

GetIsUiPort returns the IsUiPort field if non-nil, zero value otherwise.

### GetIsUiPortOk

`func (o *NodePort) GetIsUiPortOk() (*bool, bool)`

GetIsUiPortOk returns a tuple with the IsUiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUiPort

`func (o *NodePort) SetIsUiPort(v bool)`

SetIsUiPort sets IsUiPort field to given value.

### HasIsUiPort

`func (o *NodePort) HasIsUiPort() bool`

HasIsUiPort returns a boolean if a field has been set.

### SetIsUiPortNil

`func (o *NodePort) SetIsUiPortNil(b bool)`

 SetIsUiPortNil sets the value for IsUiPort to be an explicit nil

### UnsetIsUiPort
`func (o *NodePort) UnsetIsUiPort()`

UnsetIsUiPort ensures that no value is present for IsUiPort, not even an explicit nil
### GetPort

`func (o *NodePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NodePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NodePort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NodePort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *NodePort) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *NodePort) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetTag

`func (o *NodePort) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NodePort) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NodePort) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *NodePort) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *NodePort) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *NodePort) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


