package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PinterestResponse struct {
	ResourceResponse struct {
		Data struct {
			Results []struct {
				Images struct {
					Orig struct {
						URL string `json:"url"`
					} `json:"orig"`
				} `json:"images"`
			} `json:"results"`
		} `json:"data"`
		Next string `json:"next"`
	} `json:"resource_response"`
	Resource struct {
		Next string `json:"next"`
	} `json:"resource"`
}

type Pinterest struct{}

func NewPinterest() *Pinterest {
	pin := Pinterest{}
	return &pin
}

func (p *Pinterest) Scrap(query string, limit int) ([]string, error) {
	images := make([]string, 0)

	apiURL := fmt.Sprintf("https://in.pinterest.com/resource/BaseSearchResource/get/?source_url=%%2Fsearch%%2Fpins%%2F%%3Fq%%3D%s%%26rs%%3Dtyped&data=%%7B%%22options%%22%%3A%%7B%%22applied_filters%%22%%3Anull%%2C%%22appliedProductFilters%%22%%3A%%22---%%22%%2C%%22article%%22%%3Anull%%2C%%22auto_correction_disabled%%22%%3Afalse%%2C%%22corpus%%22%%3Anull%%2C%%22customized_rerank_type%%22%%3Anull%%2C%%22domains%%22%%3Anull%%2C%%22filters%%22%%3Anull%%2C%%22first_page_size%%22%%3Anull%%2C%%22page_size%%22%%3Anull%%2C%%22price_max%%22%%3Anull%%2C%%22price_min%%22%%3Anull%%2C%%22query_pin_sigs%%22%%3Anull%%2C%%22query%%22%%3A%%22%s%%22%%2C%%22redux_normalize_feed%%22%%3Atrue%%2C%%22rs%%22%%3A%%22typed%%22%%2C%%22scope%%22%%3A%%22pins%%22%%2C%%22source_id%%22%%3Anull%%2C%%22top_pin_id%%22%%3Anull%%7D%%2C%%22context%%22%%3A%%7B%%7D%%7D", url.QueryEscape(query), url.QueryEscape(query))

	for len(images) < limit {
		resp, err := http.Get(apiURL)
		if err != nil {
			return images, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return images, err
		}

		var pinterestResp PinterestResponse
		err = json.Unmarshal(body, &pinterestResp)
		if err != nil {
			return images, err
		}

		for _, result := range pinterestResp.ResourceResponse.Data.Results {
			images = append(images, result.Images.Orig.URL)
		}

		if pinterestResp.Resource.Next == "" {
			break
		}

		apiURL = pinterestResp.Resource.Next
	}

	if len(images) > limit {
		images = images[:limit]
	}

	return images, nil
}
