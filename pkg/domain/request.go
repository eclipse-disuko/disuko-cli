// SPDX-FileCopyrightText: 2025 Mercedes-Benz Group AG and Mercedes-Benz AG
//
// SPDX-License-Identifier: Apache-2.0

package domain

type RequestCreateVersion struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RequestCssAdd struct {
	URL     string `json:"url"`
	Comment string `json:"comment"`
}

type RequestCreateTag struct {
	Tag string `json:"tag"`
}

type RequestSbomSearch struct {
	Tag string `json:"tag"`
}

type RequestCommentRR struct {
	Content string `json:"content"`
}
