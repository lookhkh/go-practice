package pkgquery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type pkgRegisterResult struct {
	ID string `json:"id"`
}

func RegisterPackageData(url string, data pkgData) (pkgRegisterResult, error) {
	p := pkgRegisterResult{}
	b, err := json.Marshal(data)
	if err != nil {
		return p, nil
	}

	reader := bytes.NewReader(b)
	r, err := http.Post(url, "application/json", reader)
	if err != nil {
		return p, nil
	}

	defer r.Body.Close()

	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return p, err
	}

	if r.StatusCode != http.StatusOK {
		return p, errors.New(string(respData))
	}
	err = json.Unmarshal(respData, &p)

	return p, err
}

func FetchPackageData(url string) ([]pkgData, error) {

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	// if r.Header.Get("Content-Type") != "application/json" {
	// 	return packages, nil
	// }

	data, err := io.ReadAll(r.Body)
	fmt.Println(data)

	var packages []pkgData

	if err != nil {
		return packages, nil
	}

	err = json.Unmarshal(data, &packages)
	return packages, err

}
