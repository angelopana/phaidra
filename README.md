## To deploy the service


Required:
- docker
- golang

---

### Deploy

`make build-docker`

`make deploy-docker`

`make test-it`

---
### TEST
run own test cases 
Set env variables for _URL_, site endpoint and _METHOD_ as method type i.e GET/POST...

`make own-test`