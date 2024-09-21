// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && (amd64 || arm64 || loong64 || ppc64le)) || (freebsd && amd64)

#include <errno.h>
#include <stdint.h>
#include <stdlib.h>
// #include <sys/mman.h>

#include "libcgo.h"

uintptr_t
x_cgo_mmap(void *addr, uintptr_t length, int32_t prot, int32_t flags, int32_t fd, uint32_t offset) {
	return (uintptr_t)ENOTSUP;
}

void
x_cgo_munmap(void *addr, uintptr_t length) {
}
