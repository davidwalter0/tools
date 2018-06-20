
At their very core, cookies are just key/value pairs stored on an end
user’s computer. As a result, the only thing you really need to do to
create one is to set the Name and Value fields of the http.Cookie
type, then call the http.SetCookie function to tell the end user’s
browser to set that cookie.

```
    // https://www.calhoun.io/securing-cookies-in-go/

func someHandler(w http.ResponseWriter, r *http.Request) {
  c := http.Cookie{
    Name: "theme",
    Value: "dark",
  }
  http.SetCookie(w, &c)
}
```

SetCookie doesn't return an error

http.SetCookie doesn’t return an error, but may instead silently drop
an invalid cookie. This isn’t an awesome experience, but it is what it
is, so definitely keep this in mind when using the function.


1. Cookie  theft - Attackers can  attempt to steal cookies  in various
   ways. We will discuss how to prevent/mitigate most of these, but at
   the  end of  the day  we can’t  completely prevent  physical device
   access.

2. Cookie tampering - Whether intentional  or not, data in cookies can
   be altered.  We will  want to  discuss how we  can verify  the data
   stored in a cookie is indeed valid data we wrote.

3. Data  leaks -  Cookies are  stored on end  user’s computers,  so we
   should  be conscious  of what  data we  store there  in case  it is
   leaked.

4. Cross-site scripting (XSS) - While not directly related to cookies,
   XSS attacks are more powerful when  they have access to cookies. We
   should  consider  limiting our  cookies  from  being accessible  to
   scripts where it isn’t needed.

5. Cross-site  Request Forgery  (CSRF) - These  attacks often  rely on
   users being logged in with a session stored in a cookie, so we will
   discuss how to prevent them even  when we are using cookies in this
   manner.


#### Cookies are often stolen in one of two ways:

- Man in the middle attacks, or something similar where an attacker
intercepts your web request and steals the cookie data from it.
- Gaining access to the hardware.

Preventing man in the middle attacks basically boils down to always using SSL when your website uses cookies. By using SSL you make it essentially impossible for others to sit in the middle of a request, because they won’t be able to decrypt the data.

For those of you thinking, “ahh a man-in-the-middle attack isn’t very likely…” I would strongly encourage you to check out firesheep, a simple tool designed to illustrate just how easy it is to steal unencrypted cookies sent over public wifi.

If you want to ensure that this doesn’t happen to your users, SET UP SSL!. Caddy Server makes this a breeze with Let’s Encrypt. Just use it. Its really freaking simple to setup for prod environments. For example, you could easily proxy your Go application with 4 lines of code:

```
calhoun.io {
  gzip
  proxy / localhost:3000
}
```

#### Cookie tampering (aka users faking data)

1. Digitally sign the data

Digitally signing data is the act of adding a “signature” to the data
so that you can verify it’s authenticity. Data doesn’t need to be
encrypted or masked from the end user, but we do need to add enough
data to our cookie so that if a user alters the data we can detect it.

The way this works with cookies is via a hash - we hash the data, then
store both the data along with the hash of the data in the
cookie. Then later when the user sends the cookie to us, we hash the
data again and verify that it matches the original hash we created.

We don’t want users also creating new hashes, so you will often see
hashing algorithms like HMAC being used so that the data can be hashed
using a secret key. This prevents end users from editing both the data
and the digital signature (the hash).

Digitally signing data is built into JSON Web Tokens (JWT) by default,
so you might already be familiar with this approach.

This can be done in Go using a package like Gorilla’s securecookie,
where you provide it with a hash key when creating a SecureCookie and
then use that object to secure your cookies.

```
// Examples adapted from examples at
// http://www.gorillatoolkit.org/pkg/securecookie.

// The data here IS NOT encrypted, but is only encoded. We discuss how
// to encrypt the data in the “data leaks” section.

// It is recommended to use a key with 32 or 64 bytes, but
// this key is less for simplicity.

var hashKey = []byte("very-secret")
var s = securecookie.New(hashKey, nil)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
  encoded, err := s.Encode("cookie-name", "cookie-value")
  if err == nil {
    cookie := &http.Cookie{
      Name:  "cookie-name",
      Value: encoded,
      Path:  "/",
    }
    http.SetCookie(w, cookie)
    fmt.Fprintln(w, encoded)
  }
}

``` 

You could then read this cookie by using the same SecureCookie object
in another handler.

```
func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
  if cookie, err := r.Cookie("cookie-name"); err == nil {
    var value string
    if err = s.Decode("cookie-name", cookie.Value, &value); err == nil {
      fmt.Fprintln(w, value)
    }
  }
}
```

#### Obfuscate data

Another solution is to mask your data in a way that makes it
impossible for users to fake it. Eg, rather than storing a cookie
like:

```
// Don't do this
http.Cookie{
  Name: "user_id",
  Value: "123",
}
```

We could instead store data that maps to the real data in our
database. This is often done with session IDs or remember tokens,
where we have a table named remember_tokens and then store data in it
like so:

```
remember_token: LAKJFD098afj0jasdf08jad08AJFs9aj2ASfd1
user_id: 123
```

We would then only store the remember token in the cookie, and even if
the user wanted to fake it they wouldn’t know what to change it to. It
just looks like gibberish.

Later when a user visits our application we would look up the remember
token in our database and determine which user they are logged in as.

In order for this to work well, you need to ensure that your
obfuscated data is:

- Maps to a user (or some other resource)
- Random
- Has significant entropy
- Can be invalidated (eg delete/change the token stored in the DB)

#### Data leaks

Data leaks often require another attack vector - like cookie theft -
before they can become a real concern, but it is always good to err on
the side of caution. Just because a cookie gets stolen doesn’t mean we
want to accidentally tell the attacker the user’s password as well.


Whenever storing data in cookies, always minimize the amount of
sensitive data stored there. Don’t store things like a user’s
password, and be sure that any encoded data also doesn’t have
this. [Articles like this one](https://hackernoon.com/your-node-js-authentication-tutorial-is-wrong-f1a3bf831a46#2491)
point out a few instances where developers have unknowingly stored
sensitive data in cookies or JWTs because it was base64 encoded, but
in reality anyone can decode this data. It is encoded, NOT encrypted.

This is a pretty big blunder to make, so if you are concerned with
accidentally storing something sensitive I suggest you look into
packages like Gorilla’s securecookie.

Earlier we discussed how it can be used to digitally sign your
cookies, but securecookie can also be used to encrypt and decrypt your
cookie data so that it can’t be decoded and read easily.

To enable encryption with the package, you simply need to pass in a
block key when creating your SecureCookie instance.


#### Cross-site scripting (XSS)

Cross-site scripting, often denoted as XSS, occurs when someone
manages to inject some JavaScript into your site that you didn’t
write, but because of the way the attack works the browser doesn’t
know that and runs the JavaScript as if your server did provide the
code.

You should be doing your best to prevent XSS in general, and we won’t
go into too much detail about what it is here, but JUST IN CASE it
slips through I suggest disabling JavaScript access to cookies
whenever it isn’t needed. You can always enable it later if you need
it, so don’t let that be an excuse to leave yourself vulnerable
either.

```
cookie := http.Cookie{
  // true means no scripts, http requests only. This has
  // nothing to do with https vs http
  HttpOnly: true,
}
```


#### CSRF (Cross Site Request Forgery)

CSRF occurs when a user visits a site that isn’t yours, but that site
has a form submitting to your web application. Because the end user
submits the form and this isn’t done via a script, the browser treats
this as a user-initiated action and passes cookies along with the form
submission.

This doesn’t seem too bad at first, but what happens when that
external site starts sending over data the user didn’t intend? For
example, badsite.com might have a form that submits a request to
transfer $100 to their bank account to chase.com hoping that you are
logged into a bank account there, and this could lead to money being
transfered without the end user intending to.

Cookies aren’t directly at fault for this, but if you are using
cookies for things like authentication you need to prevent this using
a package like Gorilla’s csrf.

This package works by providing you with a CSRF token you can insert
into every web form, and whenever a form is submitted without a token
the csrf package’s middleware will reject the form, making it
impossible for external sites to trick users into submitting forms.

For more on what CSRF is, see the following:

https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
https://en.wikipedia.org/wiki/Cross-site_request_forgery

Limit cookie access where it isn’t needed
The last thing we are going to discuss isn’t related to a specific attack, but is more of a guiding principle that I suggest when working with cookies - limit access wherever you can, and only provide access when it is needed.

We touched on this briefly when discussing XSS, but the general idea is that you should limit access to your cookies wherever you can. For example, if your web application doesn’t use subdomains, you have no reason to give all subdomains access to your cookies. This is the default for cookies, so you don’t actually need to do anything to limit to a specific domain.

On the other hand, if you do need share your cookie with subdomains you can do so with something like this:

```
c := Cookie{
  // Defaults to host-only, which means exact subdomain
  // matching. Only change this to enable subdomains if you
  // need to! The below code would work on any subdomain for
  // yoursite.com
  Domain: "yoursite.com",
}
```

For more information on how the domain is resolved, see
https://tools.ietf.org/html/rfc6265#section-5.1.3. You can also see
the source code where this gets its default value at
https://golang.org/src/net/http/cookie.go#L157.


On top of limiting to specific domains, you can also limit your
cookies to specific paths.

```
c := Cookie{
  // Defaults to any path on your app, but you can use this
  // to limit to a specific subdirectory. Eg:
  Path: "/app/",
}
```

The TL;DR is you can set path prefixes with something like /blah/, but
if you would like to read more about how this field works you can
check out https://tools.ietf.org/html/rfc6265#section-5.1.4.

#### Why not just use JWTs?

This will inevitably come up, so I’m going to address it briefly.

Despite what many people may tell you, cookies can be just as secure
as JWTs. In fact, JWTs and cookies don’t really even solve the same
issue, as JWTs could be stored inside of cookies and used virtually
identical to how they are used when provided as a header.

Regardless, cookies can be used for non-authentication data, and even
in those cases I find knowing about proper security measures is
useful.



#### Summary


```
var cookie = http.Cookie{
  Name: "user_id",
  Value: "123",
  HttpOnly: true,
  Domain: "yoursite.com",
  Path: "/app/",
}
```
