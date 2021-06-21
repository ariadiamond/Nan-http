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

### TODO

- [ ] Everything
- [ ] Regular expressions
- [ ] Multiple children acls?
