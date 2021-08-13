package helios

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
    "time"

	. "github.com/cohesity/go-sdk/helios/utils"
	. "github.com/cohesity/go-sdk/helios/models"
)

type CohesityClientConfig struct {
    Username string
    Password string
    Domain string
    ClusterVIP string
    AccessClusterId *int64 // optional
    RegionId *string // optional
    Timeout time.Duration // optional
}

func NewCohesityClient(c CohesityClientConfig) (*APIClient, error) {
    configuration := NewConfiguration()

	// Disable certifate checking 
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    
    // set url
    configuration.Servers[0].URL = "https://" + c.ClusterVIP + configuration.Servers[0].URL
    apiClient := NewAPIClient(configuration)
    
    // set timeout
    if c.Timeout > 0 {
        apiClient.cfg.HTTPClient.Timeout = c.Timeout
    } 
    
    if c.Username != "" {
        apiClient.username = &(c.Username) 
    } else {
        return nil, GenericOpenAPIError{
            error: "Username is required.",
        }
    }
    
    if c.Password != "" {
        apiClient.password = &(c.Password) 
    } else {
        return nil, GenericOpenAPIError{
            error: "Password is required.",
        }
    }

    if c.Domain != "" {
        apiClient.domain = &(c.Domain) 
    } else {
        return nil, GenericOpenAPIError{
            error: "Domain is required.",
        }
    }
    
    if c.AccessClusterId != nil {
        apiClient.accessClusterId = c.AccessClusterId
    }

    if c.RegionId != nil {
        apiClient.regionId = c.RegionId
    }

    if _ , err := GetTokenHelper(apiClient); err != nil {
        return nil, err 
    }

	return apiClient, nil
}


func GetTokenHelper(a *APIClient) (string, error) {
    body := NewCreateAccessTokenRequestParams(*NewNullableString(a.password), *NewNullableString(a.username))
    body.SetDomain(*(a.domain))

    request := ApiCreateAccessTokenRequest {
        Body: body, 
    }
    resp, r, err := a.AccessToken.CreateAccessToken(request)

    if err != nil {
        if r == nil {
            return "", GenericOpenAPIError{
                error: "Fail to connect. " + err.Error(),
            }
        }

        if b, err := ioutil.ReadAll(r.Body); err == nil {
            newErr := GenericOpenAPIError{
                error: "Fail to authenticate. Backend response: " + string(b),
            }
            return "", newErr
        } else {
            return "", GenericOpenAPIError{
                error: "Fail to authentiate. Fail to decode backend response: " + err.Error(),
            }
        }
    }

    accessToken := resp.GetAccessToken()
    // tokenType := resp.GetTokenType()

    // fmt.Println(accessToken)
    // fmt.Println(tokenType)

    return accessToken, nil
}