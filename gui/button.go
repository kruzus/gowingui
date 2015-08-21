// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"image"
	"syscall"

	"github.com/alexbrainman/gowingui/uit"
	"github.com/alexbrainman/gowingui/winapi"
)

// Button represents button.
type Button struct {
	*WinControl
}

func (c *WinControl) AddButton(id int, text string, r image.Rectangle) (*Button, error) {
	h, err := uit.M.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr("BUTTON"),
		syscall.StringToUTF16Ptr(text),
		winapi.WS_CHILD|winapi.WS_VISIBLE|winapi.BS_DEFPUSHBUTTON,
		int32(r.Min.X), int32(r.Min.Y), int32(r.Dx()), int32(r.Dy()),
		c.Handle, syscall.Handle(id), moduleHandle, 0)
	if err != nil {
		return nil, err
	}
	return &Button{
		WinControl: &WinControl{Handle: h},
	}, nil
}
