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
- [X] Better cli parsing (support `-Vwr`)
- [ ] Change command parsing to regexp (with space after "e")
- [ ] Rethink security policy/mechanism. What the heck is it right now?

**Future Server:**
- [ ] Safe PUT requests? (safe folder)
- [ ] POST requests for tracking something (mifilw)
- [ ] Derive index pages from config files

----
# Config
### How to make a config file

Example:
```
Default: [def] # todo
Class: [class] # todo
# Comment
* [name] @ [path]  => file1, file2
           scripts => script.js
           styles  => style.css

* [name] @
      [path] => file1, file2
                file3, file4
      styles => style.css ../style.css
* [path] => file2, file7
```

Default refers to when there is no name for the constructed page (like the last entry in the example). Class is appended to page title, with a delimiter in between: `[name] | [class]`.

Name it ``.httpconfig``, and place it in the folder you want it.

Files with the extension `.md` (for Github Flavored Markdown) will be converted to html when rendering the page. To support this, the command `pandoc` must be installed. At the moment, all other file types are just sent without any processing.

### Config TODO

- [ ] **Title names**
- [ ] **Default names**
- [ ] Regular Expressions
	- [ ] `?`
	- [ ] `*`
- [ ] Recursive config

---
# Access Control
### Title characters

- `@never` - These files will never be allowed and always return `403`, regardless of superuser status
- `@suRead` - Reading these files is only possible in superuser mode
- `@suWrite` - Writing these files is only possible in superuser mode
- `@readOnly` - Puts cannot happen for this path

### Comments

Like .httpconfig, these use `#` for comments

### Naming

Name it `.httpacl`

At the current moment, I'm thinking only have one acl?

### Access Control TODO

- [ ] Regular expressions
- [ ] Multiple children acls?

---
# Sources

**Markdown Support:**
- [Pandoc](https://pandoc.org) for conversions
- [MultiMarkdown User guide](https://fletcher.github.io/MultiMarkdown-6/), and the [liscense](https://github.com/fletcher/MultiMarkdown-6#license)

**HTML:**
- [Wikipedia: Status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
- [Mozilla Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference)
- [w3schools: Viewport Width](https://www.w3schools.com/html/html_responsive.asp)

**CSS:**
- [MDN: Variables](https://developer.mozilla.org/en-US/docs/Web/CSS/Using_CSS_custom_properties)
- [MDN: Media Queries](https://developer.mozilla.org/en-US/docs/Web/CSS/Media_Queries/Using_media_queries)

**JavaScript:**
- [w3resource: Embed JS in HTML](https://www.w3resource.com/javascript/introduction/html-documents.php) *note this is probably not the best source (typos)*
- [w3schools: Toggle Element in JS](https://www.w3schools.com/howto/howto_js_toggle_hide_show.asp)
- [w3schools: String Function Arguments](https://www.w3schools.com/jsref/event_onclick.asp)
- [Javatpoint: External JS files](https://www.javatpoint.com/how-to-add-javascript-to-html)

**Golang:**
- The official [website](https://golang.org): [net/http](https://golang.org/pkg/net/http/), [strings](https://golang.org/pkg/strings/), [errors](https://golang.org/doc/tutorial/handle-errors), [regexp](https://pkg.go.dev/regexp)
- Regular expression [syntax](https://github.com/google/re2/wiki/Syntax)
- How to do [enums](https://yourbasic.org/golang/iota/) (but they're called iotas)
- Golang supports goto >:) [go by example](https://golangbyexample.com/goto-statement-go/)

**Python:** guess who doesn't know python that well
- [Official documentation](https://docs.python.org/3.9/): [os](https://docs.python.org/3.9/library/os.html), [sys](https://docs.python.org/3.9/library/sys.html)
- [Classes](https://docs.python.org/3/tutorial/classes.html), [inheritance](https://stackoverflow.com/questions/576169/understanding-python-super-with-init-methods)
- [pythonbasics: Exceptions](https://pythonbasics.org/try-except/)
- [readthedocs: request](https://requests.readthedocs.io/en/master/)
