## Steps to use this provider artifactory

In one of you terminals checkout this repository https://github.com/suvaanshkumar/provider-jfrogartifactory amd run `make run`

- Apply all the manifest files in package/crds
- Generate an identity token on artifactory to be used here.
- create a file similar to creds.json present in examples/manifests/template folder and fill in the url and the key
- Base64 encode this file and put it in the secret.yaml  file present in examples/manifests/template in data field and apply the secret 
- Apply the examples/manifests/providerconfigartifactory.yaml
- apply the examples/manifests/genericrepository.yaml