// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (amd64 || arm64 || ppc64le)

#include <errno.h>
#include <stddef.h>
#include <stdint.h>
#include <string.h>
#include <signal.h>

#include "libcgo.h"

// go_sigaction_t is a C version of the sigactiont struct from
// defs_linux_amd64.go.  This definition — and its conversion to and from struct
// sigaction — are specific to linux/amd64.
typedef struct {
	uintptr_t handler;
	uint64_t flags;
	uintptr_t restorer;
	uint64_t mask;
} go_sigaction_t;

// SA_RESTORER is part of the kernel interface.
// This is Linux i386/amd64 specific.
#ifndef SA_RESTORER
#define SA_RESTORER 0x4000000
#endif

int32_t
x_cgo_sigaction(intptr_t signum, const go_sigaction_t *goact, go_sigaction_t *oldgoact) {
	return ENOTSUP;
}
