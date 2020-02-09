package zeit

import "github.com/go-resty/resty/v2"

func (api ZeitAPI) RawProjectEnsure(name string) (*resty.Response, error) {
	return api.post("/v1/projects/ensure-project", map[string]interface{}{
		"name": name,
	})
}

func (api ZeitAPI) RawProjectCreateAlias(projectId string, domain string) (*resty.Response, error) {
	return api.post("/v1/projects/"+projectId+"/alias", map[string]interface{}{
		"domain": domain,
	})
}
