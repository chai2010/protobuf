// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate protoc --go_out=. any.proto
//go:generate protoc --go_out=. duration.proto
//go:generate protoc --go_out=. empty.proto
//go:generate protoc --go_out=. struct.proto
//go:generate protoc --go_out=. timestamp.proto
//go:generate protoc --go_out=. wrappers.proto

package google_protobuf // import "google/protobuf"
