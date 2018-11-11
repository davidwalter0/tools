```
find ${PWD} \( ! -wholename '*vendor/*' -a ! -wholename '*/.git/*' -a ! -wholename '*test/*' -a -iname "*.go" \) -print | etags -
```
