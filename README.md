# rehttp
Golang HTTP client wrapping httpcontrol.Transport, http.Client &amp; redial.Dialer

[![Build Status](https://travis-ci.org/linkosmos/rehttp.svg?branch=master)](https://travis-ci.org/linkosmos/rehttp)
[![GoDoc](http://godoc.org/github.com/linkosmos/rehttp?status.svg)](http://godoc.org/github.com/linkosmos/rehttp)
[![BSD License](http://img.shields.io/badge/license-BSD-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)

### Features

 - Connection pool
 - Layered timeout's
 - Request stats
 - Lower/Higher level API
 - Cached DNS records with round-robin or random access
 - Request retry
 - Keep alive
 - Compression
 - HTTPS

### Wrapper

This package wraps underlying layers of http request procedure into readable & configurable options.
Package exposes 2 major components http Client & Transport.
HTTP Client is swappable where Transport but satisfies standard library interface.

### License

Copyright (c) 2015, linkosmos
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of rehttp nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
