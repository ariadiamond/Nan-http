## ![favicon.png](Root/favicon.png) Notes Server

# High level goal

aria likes to have really pretty notes and things, and that is what keeps her sane when doing computer science things. She also has grandiose ideas sometimes, so there is a lot of cool ideas that might never happen.

# Table Of Contents

- [To do](#TODO)
- [Config](#Config)
- [Access Control](#Access-Control)
- [Sources](#Sources)

---
# TODO

**Immediate Server**:
- [ ] Test scripts
	- [ ] unit tests
		- [ ] Config
        - [ ] Access Control
        - [ ] PUT
	- [ ] functional
        - [ ] Config
        - [ ] Access Control
        - [ ] PUT
- [ ] Support for PUT requests
    - [ ] Override for shorter/equal length files

**Future Server:**
- [ ] Safe PUT requests? (safe folder)
- [ ] POST requests for tracking something
- [ ] Server config and folder specific config
- [ ] Javascripts for things?
- [ ] Access control
- [ ] Derive index pages from config files

----
# Config
### How to make a config file

Comments can be made using `#`

``name @ path => file1, file2, ...`` with each not including the path to the folder

Name her ``.httpconfig``

### Config TODO

- [X] Title names
- [ ] Default names
- [ ] Regular Expressions
	- [ ] `?`
	- [ ] `*`
- [ ] Recursive config

---
# Access Control
### Title characters

- `@never` - These files will never be allowed and always return a 403, regardless of superuser status
- `@suRead` - Only reading is only possible in superuser mode
- `@suWrite` - This will not happen for a while
- `@readOnly` - This requires PUT/POST requests, so will not happen for a while.

### Comments

Like .httpconfig, these use `#` for comments

### Naming

Name it `.httpacl`

At the current moment, I'm thinking only have one acl?

### Access Control TODO

- [ ] Regular expressions
- [ ] Multiple children acls?

---
### Sources

**HTTP:**
- [Status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)

**HTML:**
- [Mozilla Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference)
- [Viewport Width](https://www.w3schools.com/html/html_responsive.asp)

**JavaScript:**
- [Embed JS in HTML](https://www.w3resource.com/javascript/introduction/html-documents.php) *note this is probably not the best source (typos)*
- [Toggle Element in JS](https://www.w3schools.com/howto/howto_js_toggle_hide_show.asp)
- [String Function Arguments](https://www.w3schools.com/jsref/event_onclick.asp)
- [External JS files](https://www.javatpoint.com/how-to-add-javascript-to-html)

**Golang:**
- The official [website](https://golang.org): [net/http](https://golang.org/pkg/net/http/), [strings](https://golang.org/pkg/strings/), [errors](https://golang.org/doc/tutorial/handle-errors)
- How to do [enums](https://yourbasic.org/golang/iota/) (but they're called iotas)
- Golang supports goto >:) [go by example](https://golangbyexample.com/goto-statement-go/)

**Python:** guess who doesn't know python that well
- [Official documentation](https://docs.python.org/3.9/): [os](https://docs.python.org/3.9/library/os.html), [sys](https://docs.python.org/3.9/library/sys.html)
- [Classes](https://docs.python.org/3/tutorial/classes.html), [inheritance](https://stackoverflow.com/questions/576169/understanding-python-super-with-init-methods)
- [Exceptions](https://pythonbasics.org/try-except/)
- [request](https://requests.readthedocs.io/en/master/)
