package main

import (
        "crypto/sha256"
        "encoding/hex"
)

var guestmode bool = true
var dblogin string = "USER:PASSWORD@/blog_cannon"
var personalpwdTMP string = "cannon"

var templatefolder string = "/home/her/CODE/blog_cannon/template"

/*############# END OF CONFIG ################*/

var foo2355523 = sha256.Sum256([]byte(personalpwdTMP))
var personalpwd = hex.EncodeToString(foo2355523[:])
