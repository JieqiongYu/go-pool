# Static file server

## static file location

```bash
mkdir $HOME/static
```

## preparation
Create some `text.txt` in `$HOME/static`

## Test

```bash
curl -X GET http://localhost:8000/api/v1/show-file/text.txt
```