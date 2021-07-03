# note

### install
```
go get -u github.com/smallnest/gen
```

### generate file
```
gen -c "root:@tcp(127.0.0.1:3306)/gen-sample" -d gen-sample \
--out ./infra \
--model=entity --templateDir=templates \
--mapping=templates/mapping.json \
--exec=templates/custom.gen \
--gorm \
--overwrite \
&& rm infra/entity/ar_internal_metadatum.go \
&& rm infra/entity/schema_migration.go
```

### option: create new template
```
gen --save=./mytemplates
```
