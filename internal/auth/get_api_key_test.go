package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		Name           string
		Authorization  string
		ShouldError    bool
		ExpectedResult string
	}{
		{
			Name:           "Empty Authorization Header (Missing)",
			Authorization:  "",
			ShouldError:    true,
			ExpectedResult: "",
		},
		{
			Name:           "Wrong Format, Use Bearer instead of ApiKey",
			Authorization:  "Bearer blabla",
			ShouldError:    true,
			ExpectedResult: "",
		},
		{
			Name:           "Wrong Format, No spaces",
			Authorization:  "NotNice",
			ShouldError:    true,
			ExpectedResult: "",
		},
		{
			Name:           "Wrong Format, 3 positions, no ApiKey string",
			Authorization:  "Also Not Nice",
			ShouldError:    true,
			ExpectedResult: "",
		},
		{
			Name:           "3 positions, returns position[1]",
			Authorization:  "ApiKey Surprisingly ThisIsOk",
			ShouldError:    false,
			ExpectedResult: "Surprisingly",
		},
		{
			Name:           "Correct Format, 2 position",
			Authorization:  "ApiKey ThisIsOk",
			ShouldError:    false,
			ExpectedResult: "ThisIsOk",
		},
	}
	for _, testV := range tests {
		t.Run(testV.Name, func(t *testing.T) {
			header := make(http.Header)
			header.Add("Authorization", testV.Authorization)
			v, err := GetAPIKey(header)
			if (err != nil) != testV.ShouldError {
				t.Errorf("Expected Error: %v, Error value: %v", testV.ShouldError, err)
				return
			}
			if v != testV.ExpectedResult {
				t.Errorf("ExpectedResult Differs: Received: \"%v\", Expected: \"%v\"", v, testV.ExpectedResult)
				return
			}
		})
	}
}
