// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build aix

package osext

import "os"

func executable() (string, error) {
	return os.Executable()
}
