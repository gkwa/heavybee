# heavybee

Purpose:
Learn CRUD with dynamodb + golang.

## example usage

```bash


```

## install heavybee


on macos/linux:
```bash

brew install gkwa/homebrew-tools/heavybee

```


on windows:

```powershell

TBD

```


## Getting started


```bash

cd heavybee
terraform init
terraform plan -out=tfplan
terraform apply tfplan

# bulk load data.json
make && ./heavybee run --json data.json

# query the table
make && ./heavybee query

# cleanup
terraform plan -destroy -out=tfplan && terraform apply tfplan



```