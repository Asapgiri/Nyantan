# Nihility
Random manga translation site

# Inti steps
- Install `go`
- Install [mongoDB](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-ubuntu/)
Example database is in the docs folder

## Mandatory Environmentals
```bash
export NYANTAN_URI="mongodb://localhost:27017"
export NYANTAN_DATABASE_NAME="nyantan"
export NYANTAN_SESSION_KEY="1589334fcad0ddd6b393ce5bab75ed43" # random hash
```

# Todo/Progress:
- [x] Base template
- [ ] Templates/Pages
    - [ ] Home
    - [ ] Brows Stuff/Manga page
        - [ ] Fandom (filter??)
    - [ ] Profile
    - [x] Login
    - [ ] Edit
        - [x] List
        - [ ] Create/Upload
        - [ ] Translate
        - [ ] Proofreed
        - [ ] Check
- [ ] Current page highlighting
- [ ] Page DTO-s (Should wait normal backend!?)
- [ ] Userpages
- [ ] Backend user logic
- [ ] Roles
- [ ] Apis
- [ ] Proper user auth
- [x] Lazy load /translations drop down contents trough apis
- [ ] Provide proper error messages
- [ ] Solve TODO-s
- [ ] Solve FEXME-s
