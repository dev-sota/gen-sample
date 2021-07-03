# How to use gen in the singular

### Install
```
go get -u github.com/smallnest/gen
```

### Copy my template
```
cp -r templates YOUR_WORKSPACE
```

### Option: Create new template
```
gen --save=./mytemplates
```

### Generate entity file
```
gen -c "root:@tcp(127.0.0.1:3306)/gen-sample" \
-d gen-sample \
--out ./infra \
--model=entity \
--templateDir=templates \
--mapping=templates/mapping.json \
--exec=templates/custom.gen \
--gorm \
--overwrite \
&& rm infra/entity/ar_internal_metadatum.go \
&& rm infra/entity/schema_migration.go \
&& sed -i "s/.\/.*//g" infra/entity/* \
&& sed -i '/^$/d' infra/entity/* \
&& gofmt -s -w infra/entity/.
```
