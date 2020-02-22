# Install

```
go get .
```

# Run

```
PORT=7896 VMESS_T_URL=https://edmo.com/ VMESS_T_SELECT="#btn"  VMESS_TATTR="data-clipboard-text" go run main.go
curl http://localhost:7896/vmess
```

# Deploy

```
go build
PORT=7896 VMESS_T_URL=https://edmo.com/ VMESS_T_SELECT="#btn"  VMESS_TATTR="data-clipboard-text" vmess-sub > all.log 2>&1 &
```
