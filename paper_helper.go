package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

var paperAPI string
var paperVersion string

func main() {
	flag.StringVar(&paperVersion, "d", "", "Specify a version to download paper server.")
	flag.StringVar(&paperAPI, "a", "https://api.papermc.io", "Specify paper API server.")
	flag.Parse()
	var c = http.Client{}

	if len(paperVersion) != 0 {
		versions := getAviliableVersion(c)
		if !checkIsAviliable(versions, paperVersion) {
			fmt.Print("null")
			return
		}

		builds := getAviliableBuilds(c, paperVersion)
		if len(builds) == 0 {
			fmt.Print("null")
			return
		}

		latestBuild := builds[len(builds)-1]

		fmt.Print(getDownloadLinks(c, paperVersion, latestBuild))
		return
	}

	flag.Usage()
}

func getAviliableVersion(client http.Client) []string {
	var fetchURL = paperAPI + "/v2/projects/paper"
	response, err := client.Get(fetchURL)

	if err != nil || response.StatusCode != 200 {
		return nil
	}

	bodyData, _ := io.ReadAll(response.Body)
	unmarshalData := ProjectInfomation{}
	json.Unmarshal(bodyData, &unmarshalData)

	return unmarshalData.Versions
}

func checkIsAviliable(versions []string, version string) bool {

	if versions == nil {
		return false
	}

	for _, v := range versions {
		if v == version {
			return true
		}
	}

	return false
}

func getAviliableBuilds(client http.Client, version string) []int {
	var fetchURL = paperAPI + "/v2/projects/paper/versions/" + version

	response, err := client.Get(fetchURL)
	if err != nil || response.StatusCode != 200 {
		return nil
	}

	bodyData, _ := io.ReadAll(response.Body)

	unmarshalData := ProjectVersion{}
	json.Unmarshal(bodyData, &unmarshalData)

	return unmarshalData.Builds
}

func getDownloadLinks(client http.Client, version string, build int) string {
	var fetchURL = paperAPI + "/v2/projects/paper/versions/" + version + "/builds/" + string(fmt.Sprint(build))

	response, err := client.Get(fetchURL)
	if err != nil || response.StatusCode != 200 {
		return ""
	}

	bodyData, _ := io.ReadAll(response.Body)

	unmarshalData := VersionBuildController{}
	json.Unmarshal(bodyData, &unmarshalData)
	var name = unmarshalData.Downloads.Application.Name

	if len(name) == 0 {
		return ""
	}

	return fetchURL + "/downloads/" + name
}
