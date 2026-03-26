package hackerapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	
)

func (c *Client) QueryBasedResp() (SearchBasedResp, error){
	endpoint := "/search?query=comment"
	fullURL := baseURL + endpoint
	
	req, err := http.NewRequest("GET", fullURL, nil)
	if err !=nil {
		return SearchBasedResp{}, err
	}
	
	resp,err := c.httpClient.Do(req)
	
	if err != nil {
		return  SearchBasedResp{}, err
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode > 399 {
		return SearchBasedResp{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchBasedResp{},err
	}
	
	SearchBasedResp := SearchBasedResp{}
	json.Unmarshal(dat, &SearchBasedResp)
	return SearchBasedResp, nil
}