// Do not edit. Bootstrap copy of /usr/local/google/buildbot/src/android/build-tools/out/obj/go/src/cmd/internal/obj/funcdata.go

//line /usr/local/google/buildbot/src/android/build-tools/out/obj/go/src/cmd/internal/obj/funcdata.go:1
// Inferno utils/5c/list.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/list.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package obj

// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file defines the IDs for PCDATA and FUNCDATA instructions
// in Go binaries. It is included by assembly sources, so it must
// be written using #defines.
//
// The Go compiler also #includes this file, for now.
//
// symtab.go also contains a copy of these constants.

// Pseudo-assembly statements.

// GO_ARGS, GO_RESULTS_INITIALIZED, and NO_LOCAL_POINTERS are macros
// that communicate to the runtime information about the location and liveness
// of pointers in an assembly function's arguments, results, and stack frame.
// This communication is only required in assembly functions that make calls
// to other functions that might be preempted or grow the stack.
// NOSPLIT functions that make no calls do not need to use these macros.

// GO_ARGS indicates that the Go prototype for this assembly function
// defines the pointer map for the function's arguments.
// GO_ARGS should be the first instruction in a function that uses it.
// It can be omitted if there are no arguments at all.
// GO_ARGS is inserted implicitly by the linker for any function
// that also has a Go prototype and therefore is usually not necessary
// to write explicitly.

// GO_RESULTS_INITIALIZED indicates that the assembly function
// has initialized the stack space for its results and that those results
// should be considered live for the remainder of the function.

// NO_LOCAL_POINTERS indicates that the assembly function stores
// no pointers to heap objects in its local stack variables.

// ArgsSizeUnknown is set in Func.argsize to mark all functions
// whose argument size is unknown (C vararg functions, and
// assembly code without an explicit specification).
// This value is generated by the compiler, assembler, or linker.
const (
	PCDATA_StackMapIndex       = 0
	FUNCDATA_ArgsPointerMaps   = 0
	FUNCDATA_LocalsPointerMaps = 1
	FUNCDATA_DeadValueMaps     = 2
	ArgsSizeUnknown            = -0x80000000
)
