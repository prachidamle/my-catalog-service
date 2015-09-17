package main

import "github.com/rancher/go-rancher/client"

type Template struct {
	client.Resource
	Name           string `json:"name"`
	Category       string `json:"category"`
	Description    string `json:"description"`
	DefaultVersion string `json:"version"`
}

type Templates []Template
