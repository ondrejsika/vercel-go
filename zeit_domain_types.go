package zeit

type ListDomainsResponse struct {
	Domains []struct {
		Id                  string   `json:"id"`
		Name                string   `json:"name"`
		ServiceType         string   `json:"serviceType"`
		NsVerifiedAt        int      `json:"nsVerifiedAt"`
		TxtVerifiedAt       int      `json:"txtVerifiedAt"`
		CdnEnabled          bool     `json:"cdnEnabled"`
		CreatedAt           int      `json:"createdAt"`
		BoughtAt            int      `json:"boughtAt"`
		VerificationRecord  string   `json:"verificationRecord"`
		Verified            bool     `json:"verified"`
		Nameservers         []string `json:"nameservers"`
		IntendedNameservers []string `json:"intendedNameservers"`
		Creator             struct {
			Id       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		} `json:"creator"`
	} `json:"domains"`
}
