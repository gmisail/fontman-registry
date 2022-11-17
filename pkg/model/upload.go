package model

type StyleUploadRequest struct {
	Style string `json:"style"`
	Url   string `json:"url"`
}

type FamilyUploadRequest struct {
	Name    string               `json:"name"`
	License string               `json:"license"`
	Creator string               `json:"creator"`
	Styles  []StyleUploadRequest `json:"styles"`
}
