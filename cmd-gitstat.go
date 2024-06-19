package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adnsv/go-utils/git"
	"github.com/alecthomas/kong"
)

type gitstat struct {
	Verbose bool   `help:"Show detailed output"`
	Output  string `short:"o" type:"path" help:"Output filename"`
}

func (cmd *gitstat) Run(ctx *kong.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to obtain current wd: %w", err)
	}

	if cmd.Verbose {
		fmt.Printf("running in directory: %s\n", wd)
	}

	vstats, err := git.Stat(wd)
	if err != nil {
		return fmt.Errorf("failed to obtain stats dir %q: %w", wd, err)
	}

	if cmd.Verbose {
		fmt.Printf("parsing tag: %s\n", vstats.Description.Tag)
	}
	vinfo, err := git.ParseVersion(vstats.Description)
	if err != nil {
		return fmt.Errorf("failed to parse version info: %w", err)
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
		return fmt.Errorf("failed to marshal results into JSON: %w", err)
	}

	if cmd.Output != "" {
		if cmd.Verbose {
			fmt.Printf("writing output to %q\n", cmd.Output)
		}
		err = os.WriteFile(cmd.Output, buf, 0666)
		if err != nil {
			return fmt.Errorf("failed to write results into a file: %w", err)
		}
	} else {
		_, err = os.Stdout.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}
