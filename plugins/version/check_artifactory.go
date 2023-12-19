package version

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Masterminds/semver"
	"github.com/sirupsen/logrus"
)

const baseURL = "https://github.com/josh-silvas/nbot/releases/nbot/"

type (
	// Index : Represents the base index of files in Artifactory storage.
	Index struct {
		Repo         string    `json:"repo"`
		Path         string    `json:"path"`
		Created      time.Time `json:"created"`
		CreatedBy    string    `json:"createdBy"`
		LastModified time.Time `json:"lastModified"`
		ModifiedBy   string    `json:"modifiedBy"`
		LastUpdated  time.Time `json:"lastUpdated"`
		Children     []struct {
			URI    string `json:"uri"`
			Folder bool   `json:"folder"`
		} `json:"children"`
		URI string `json:"uri"`
	}
)

// FromArtifactory : Function to fetch the latest version from Artifactory
func FromArtifactory(proxyURL string) (*semver.Version, error) {
	var res Index
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	// Build the request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("version:Fetch:NewRequest:%w", err)
	}
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // nolint: gosec
			},
		},
		Timeout: 10 * time.Second,
	}

	if proxyURL != "" {
		pURL, err := url.Parse(proxyURL)
		if err != nil {
			logrus.Fatal(err)
		}
		client.Transport = &http.Transport{
			MaxIdleConnsPerHost: 20,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // nolint: gosec
			},
			Proxy: http.ProxyURL(pURL),
		}
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("VersionFromArtifactory:Do:%w", err)
	}
	if r != nil {
		defer func() {
			if defErr := r.Body.Close(); defErr != nil {
				closeErr := defErr.Error()
				err = fmt.Errorf("%w/%s", err, closeErr)
			}
		}()
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("VersionFromArtifactory:Fetch:%s", r.Status)
	}

	if err = json.NewDecoder(r.Body).Decode(&res); err != nil {
		return nil, err
	}
	latestVer, err := semver.NewVersion("0.0.0")
	if err != nil {
		return nil, err
	}
	for i := range res.Children {
		if !res.Children[i].Folder {
			continue
		}
		if v, err := semver.NewVersion(strings.TrimPrefix(res.Children[i].URI, "/")); err == nil {
			if v.GreaterThan(latestVer) {
				latestVer = v
			}
		}

	}

	return latestVer, nil
}
