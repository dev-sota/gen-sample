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
gen --connstr="root:@tcp(127.0.0.1:3306)/gen-sample" \
--database=gen-sample \
--exclude=ar_internal_metadata,schema_migrations \
--out ./infra \
--model=entity \
--templateDir=templates \
--mapping=templates/mapping.json \
--exec=templates/entity.gen \
--gorm \
--overwrite \
&& sed -i "s/.\/.*//g" infra/entity/* \
&& sed -i '/^$/d' infra/entity/* \
&& gofmt -s -w infra/entity/.
```
