package hackerapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	
)

func (c *Client) FirstPageReq() (FrontPageRes, error){
	endpoint := "/search?tags=front_page"
	fullURL := baseURL + endpoint
	
	req, err := http.NewRequest("GET", fullURL, nil)
	if err !=nil {
		return FrontPageRes{}, err
	}
	
	resp,err := c.httpClient.Do(req)
	
	if err != nil {
		return  FrontPageRes{}, err
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode > 399 {
		return FrontPageRes{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return FrontPageRes{},err
	}
	
	frontPageResp := FrontPageRes{}
	json.Unmarshal(dat, &frontPageResp)
	return frontPageResp, nil
}