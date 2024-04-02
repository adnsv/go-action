package main

import (
	"encoding/json"
	"os"

	"github.com/adnsv/go-utils/git"
	"github.com/alecthomas/kong"
)

type gitstat struct {
	Output string `short:"o" type:"path" help:"Output filename"`
}

func (cmd *gitstat) Run(ctx *kong.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	vstats, err := git.Stat(wd)
	if err != nil {
		return err
	}

	vinfo, err := git.ParseVersion(vstats.Description)
	if err != nil {
		return err
	}

	type jret struct {
		Branch            string `json:"branch"`
		Hash              string `json:"hash"`
		ShortHash         string `json:"short-hash"`
		AuthorDate        string `json:"author-date"`
		LastTag           string `json:"last-tag"`
		AdditionalCommits int    `json:"additional-commits"`
		Semantic          string `json:"ver-semantic"`
		Quad              string `json:"ver-quad"`
		NNNN              string `json:"ver-nnnn"`
		Pre               string `json:"ver-pre,omitempty"`
		Build             string `json:"ver-build,omitempty"`
	}

	j := &jret{
		Branch:            vstats.Branch,
		Hash:              vstats.Hash,
		ShortHash:         vstats.ShortHash,
		AuthorDate:        vstats.AuthorDate,
		LastTag:           vstats.Description.Tag,
		AdditionalCommits: vstats.Description.AdditionalCommits,
		Semantic:          vinfo.Semantic.String(),
		Quad:              vinfo.Quad.String(),
		NNNN:              vinfo.Quad.CommaSeparatedString(),
		Pre:               vinfo.Pre,
		Build:             vinfo.Build,
	}
	buf, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return err
	}

	if cmd.Output != "" {
		err = os.WriteFile(cmd.Output, buf, 0666)
		if err != nil {
			return err
		}
	} else {
		_, err = os.Stdout.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}
