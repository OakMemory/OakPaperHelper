package main

import "time"

type GetProject struct {
	Projects []string `json:"projects"`
}

type ProjectInfomation struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

type ProjectVersion struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
}

type GetAllBuildOfProject struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []struct {
		Build    int       `json:"build"`
		Time     time.Time `json:"time"`
		Channel  string    `json:"channel"`
		Promoted bool      `json:"promoted"`
		Changes  []struct {
			Commit  string `json:"commit"`
			Summary string `json:"summary"`
			Message string `json:"message"`
		} `json:"changes"`
		Downloads struct {
			Application struct {
				Name   string `json:"name"`
				Sha256 string `json:"sha256"`
			} `json:"application"`
			MojangMappings struct {
				Name   string `json:"name"`
				Sha256 string `json:"sha256"`
			} `json:"mojang-mappings"`
		} `json:"downloads"`
	} `json:"builds"`
}

type VersionBuildController struct {
	ProjectID   string    `json:"project_id"`
	ProjectName string    `json:"project_name"`
	Version     string    `json:"version"`
	Build       int       `json:"build"`
	Time        time.Time `json:"time"`
	Channel     string    `json:"channel"`
	Promoted    bool      `json:"promoted"`
	Changes     []struct {
		Commit  string `json:"commit"`
		Summary string `json:"summary"`
		Message string `json:"message"`
	} `json:"changes"`
	Downloads struct {
		Application struct {
			Name   string `json:"name"`
			Sha256 string `json:"sha256"`
		} `json:"application"`
		MojangMappings struct {
			Name   string `json:"name"`
			Sha256 string `json:"sha256"`
		} `json:"mojang-mappings"`
	} `json:"downloads"`
}
