# note

```
go get -u github.com/smallnest/gen
gen --save=./templates
gen -c "root:@tcp(127.0.0.1:3306)/gen-sample" -d gen-sample --out ./infra --model=entity --templateDir=templates --mapping=templates/mapping.json  --gorm --overwrite
```
