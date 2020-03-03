package client

import (
	"gcloud/gcorecloud-go"
	"gcloud/gcorecloud-go/gcore"
	"gcloud/gcorecloud-go/testhelper"
)

// Fake token to use.
const TokenID = "cbc36478b0bd8e67e89469c7749d4127"
const AccessToken = "cbc36478b0bd8e67e89469c7749d4127"
const RefreshToken = "tbc36478b0bd8e67e89469c7749d4127"
const Username = "username"
const Password = "password"
const RegionID = 1
const ProjectID = 11

// ServiceClient returns a generic service client for use in tests.
func ServiceClient() *gcorecloud.ServiceClient {
	return &gcorecloud.ServiceClient{
		ProviderClient: &gcorecloud.ProviderClient{
			AccessTokenID:  AccessToken,
			RefreshTokenID: RefreshToken,
		},
		Endpoint: testhelper.Endpoint(),
	}
}

func ServiceTokenClient(name string, version string) *gcorecloud.ServiceClient {
	options := gcorecloud.TokenOptions{
		IdentityEndpoint: testhelper.GCloudRefreshTokenIdentifyEndpoint(),
		AccessToken:      AccessToken,
		RefreshToken:     RefreshToken,
		AllowReauth:      true,
	}
	endpointOpts := gcorecloud.EndpointOpts{
		Name:    name,
		Region:  RegionID,
		Project: ProjectID,
		Version: version,
	}
	client, err := gcore.TokenClientService(options, endpointOpts)
	if err != nil {
		panic(err)
	}
	return client
}

func TaskTokenClient() *gcorecloud.ServiceClient {
	options := gcorecloud.TokenOptions{
		IdentityEndpoint: testhelper.GCloudRefreshTokenIdentifyEndpoint(),
		AccessToken:      AccessToken,
		RefreshToken:     RefreshToken,
		AllowReauth:      true,
	}
	client, err := gcore.TaskTokenClient(options)
	if err != nil {
		panic(err)
	}
	return client
}

func ServiceAuthClient(name string, version string) *gcorecloud.ServiceClient {
	options := gcorecloud.AuthOptions{
		IdentityEndpoint:     testhelper.GCoreIdentifyEndpoint(),
		RefreshTokenEndpoint: testhelper.GCoreRefreshTokenIdentifyEndpoint(),
		Username:             Username,
		Password:             Password,
		AllowReauth:          true,
	}
	endpointOpts := gcorecloud.EndpointOpts{
		Name:    name,
		Region:  RegionID,
		Project: ProjectID,
		Version: version,
	}
	client, err := gcore.AuthClientService(options, endpointOpts)
	if err != nil {
		panic(err)
	}
	return client
}

type AuthResultTest struct {
	accessToken  string
	refreshToken string
}

func (ar AuthResultTest) ExtractAccessToken() (string, error) {
	return ar.accessToken, nil
}

func (ar AuthResultTest) ExtractRefreshToken() (string, error) {
	return ar.accessToken, nil
}

func (ar AuthResultTest) ExtractTokensPair() (string, string, error) {
	return ar.accessToken, ar.refreshToken, nil
}

func NewAuthResultTest(accessToken string, refreshToken string) AuthResultTest {
	return AuthResultTest{
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}
