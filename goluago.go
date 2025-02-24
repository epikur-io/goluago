package goluago

import (
	"github.com/epikur-io/go-lua"
	"github.com/epikur-io/goluago/pkg/crypto/aes"
	"github.com/epikur-io/goluago/pkg/crypto/hmac"
	"github.com/epikur-io/goluago/pkg/crypto/md5"
	"github.com/epikur-io/goluago/pkg/crypto/sha256"
	"github.com/epikur-io/goluago/pkg/encoding/base64"
	"github.com/epikur-io/goluago/pkg/encoding/hex"
	"github.com/epikur-io/goluago/pkg/encoding/json"
	"github.com/epikur-io/goluago/pkg/env"
	"github.com/epikur-io/goluago/pkg/fmt"
	"github.com/epikur-io/goluago/pkg/net/url"
	"github.com/epikur-io/goluago/pkg/regexp"
	"github.com/epikur-io/goluago/pkg/strings"
	"github.com/epikur-io/goluago/pkg/time"
	"github.com/epikur-io/goluago/pkg/uuid"
	"github.com/epikur-io/goluago/util"
)

func Open(l *lua.State) {
	regexp.Open(l)
	strings.Open(l)
	json.Open(l)
	time.Open(l)
	fmt.Open(l)
	url.Open(l)
	util.Open(l)
	hmac.Open(l)
	base64.Open(l)
	env.Open(l)
	uuid.Open(l)
	hex.Open(l)
	sha256.Open(l)
	aes.Open(l)
	md5.Open(l)
}
