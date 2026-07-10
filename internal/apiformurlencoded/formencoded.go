package apiformurlencoded

import "github.com/HubSpot/hubspot-sdk-go/internal/apiquery"

type Marshaler interface {
	MarshalFormEncoded() ([]byte, string, error)
}

func Marshal(val any) ([]byte, string, error) {
	formValues, err := apiquery.Marshal(val)
	if err != nil {
		return nil, "", err
	}
	return []byte(formValues.Encode()), "application/x-www-form-urlencoded", nil
}
