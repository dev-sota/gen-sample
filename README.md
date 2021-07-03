# note

```
go get -u github.com/smallnest/gen
gen --save=./templates
gen -c "root:@tcp(127.0.0.1:3306)/gen-sample" -d gen-sample --out ./entity --model=entity --gorm --mapping=extra.json
```
