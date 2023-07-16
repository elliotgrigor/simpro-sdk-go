// Package simpro improves the simPRO Software API experience with a simpler,
// developer-friendly interface.
package simpro

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SimPROSDK
type SimPROSDK struct {
	simPRODomain   string       // "<ORGANISATION>.simprocloud.com"
	apiBase        string       // "/api/v1.0/companies/"
	apiCompanyID   uint         // 0
	apiAccessToken string       // "_______FortyCharacterRandomString_______"
	apiHttpClient  *http.Client
}

// NewSimPROSDK
func NewSimPROSDK(domain, token string, company ...uint) (*SimPROSDK, error) {
	if token == "" {
		return nil, ErrorEmptyAccessToken
	}
	if domain == "" {
		return nil, ErrorEmptyDomain
	}

	sdk := &SimPROSDK{
		simPRODomain:   domain,
		apiBase:        "/api/v1.0/companies/",
		apiAccessToken: token,
		apiHttpClient:  &http.Client{
			Timeout: time.Second * 60,
		},
	}

	if len(company) > 0 {
		sdk.apiCompanyID = company[0]
	}

	return sdk, nil
}

// SetCompany
func (sdk *SimPROSDK) SetCompany(id uint) {
	sdk.apiCompanyID = id
}

// GetCompanies retrieves a slice of company IDs and names. This can be used in
// conjunction with an omitted company ID in NewSimPROSDK to check the available
// companies before committing to one with SetCompany.
func (sdk *SimPROSDK) GetCompanies() ([]*CompanyListResponse, error) {
	var (
		emptyCpList = []*CompanyListResponse{}

		url = fmt.Sprintf("https://%s%s", sdk.simPRODomain, sdk.apiBase)
	)

	resp, err := sdk.makeHTTPRequest("GET", url, nil)
	if err != nil {
		return emptyCpList, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyCpList, ErrorFailedReadingBody(err.Error())
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return emptyCpList, nil
	}

	var cpList []*CompanyListResponse

	err = json.Unmarshal(body, &cpList)
	if err != nil {
		return emptyCpList, ErrorFailedJSONUnmarshal(err.Error())
	}

	return cpList, nil
}

// GetCompanyInfo retrieves details about the currently set company.
func (sdk *SimPROSDK) GetCompanyInfo() (*CompanyResponse, error) {
	var (
		emptyCp = &CompanyResponse{}

		url = fmt.Sprintf("https://%s%s%d",
			sdk.simPRODomain,
			sdk.apiBase,
			sdk.apiCompanyID,
		)
	)

	resp, err := sdk.makeHTTPRequest("GET", url, nil)
	if err != nil {
		return emptyCp, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyCp, ErrorFailedReadingBody(err.Error())
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return emptyCp, nil
	}

	var cp *CompanyResponse

	err = json.Unmarshal(body, &cp)
	if err != nil {
		return emptyCp, ErrorFailedJSONUnmarshal(err.Error())
	}

	return cp, nil
}

// GetSecurityGroups
func (sdk *SimPROSDK) GetSecurityGroups() (
	[]*SecurityGroupListResponse, error,
) {
	var (
		emptySgList = []*SecurityGroupListResponse{}

		url = fmt.Sprintf("https://%s%s%d/setup/securityGroups/",
			sdk.simPRODomain,
			sdk.apiBase,
			sdk.apiCompanyID,
		)
	)

	resp, err := sdk.makeHTTPRequest("GET", url, nil)
	if err != nil {
		return emptySgList, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptySgList, ErrorFailedReadingBody(err.Error())
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return emptySgList, nil
	}

	var sgList []*SecurityGroupListResponse

	err = json.Unmarshal(body, &sgList)
	if err != nil {
		return emptySgList, ErrorFailedJSONUnmarshal(err.Error())
	}

	return sgList, nil
}

// GetSecurityGroup
func (sdk *SimPROSDK) GetSecurityGroup(id uint) (
	*SecurityGroupResponse, error,
) {
	var (
		emptySg = &SecurityGroupResponse{}

		url = fmt.Sprintf("https://%s%s%d/setup/securityGroups/%d",
			sdk.simPRODomain,
			sdk.apiBase,
			sdk.apiCompanyID,
			id,
		)
	)

	resp, err := sdk.makeHTTPRequest("GET", url, nil)
	if err != nil {
		return emptySg, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptySg, ErrorFailedReadingBody(err.Error())
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return emptySg, nil
	}

	var sg *SecurityGroupResponse

	// TODO:
	// Find a way to properly ignore missing fields in JSON response
	// i.e. "BusinessGroup":{} => BusinessGroup:{ID:nil Name:nil}
	//
	// Currently unmarshals to zero values. No bueno :(
	// i.e. BusinessGroup:{ID:0 Name:}

	err = json.Unmarshal(body, &sg)
	if err != nil {
		return emptySg, ErrorFailedJSONUnmarshal(err.Error())
	}

	return sg, nil
}

// ------------ Private ------------

// makeHTTPRequest
func (sdk *SimPROSDK) makeHTTPRequest(
	method, url string,
	body        io.Reader,
) (
	*http.Response, error,
) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, ErrorFailedCreatingRequest(err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sdk.apiAccessToken))
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := sdk.apiHttpClient.Do(req)
	if err != nil {
		return nil, ErrorFailedMakingRequest(err.Error())
	}

	var (
		statusNot200 = resp.StatusCode != http.StatusOK
		statusNot404 = resp.StatusCode != http.StatusNotFound
	)

	if statusNot200 && statusNot404 {
		return nil, ErrorUnexpectedResponse(resp.StatusCode)
	}

	return resp, nil
}
