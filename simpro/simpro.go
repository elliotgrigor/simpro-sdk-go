/**
 * Package `simpro` improves the simPRO Software API experience with a simpler,
 * developer-friendly interface.
 **/
package simpro

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

// GetCompanies
func (sdk *SimPROSDK) GetCompanies() ([]CompanyListResponse, error) {
	var (
		emptyResp = []CompanyListResponse{}

		url = fmt.Sprintf("https://%s%s", sdk.simPRODomain, sdk.apiBase)
	)

	resp, err := sdk.makeHTTPRequest("GET", url, nil)
	if err != nil {
		return emptyResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyResp, ErrorFailedReadingBody(err.Error())
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return emptyResp, nil
	}

	var data []CompanyListResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return emptyResp, ErrorFailedJSONUnmarshal(err.Error())
	}

	return data, nil
}

// GetCompanyInfo
func (sdk *SimPROSDK) GetCompanyInfo() (CompanyResponse, error) {
	var (
		emptyResp = CompanyResponse{}

		url = fmt.Sprintf("https://%s%s%d",
			sdk.simPRODomain,
			sdk.apiBase,
			sdk.apiCompanyID,
		)
	)

	resp, err := sdk.makeHTTPRequest("GET", url, nil)
	if err != nil {
		return emptyResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyResp, ErrorFailedReadingBody(err.Error())
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return emptyResp, nil
	}

	var data CompanyResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return emptyResp, ErrorFailedJSONUnmarshal(err.Error())
	}

	return data, nil
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
