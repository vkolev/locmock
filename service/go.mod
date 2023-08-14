module github.com/vkolev/locmock/service

go 1.20

replace github.com/vkolev/locmock/action => ../action

require github.com/vkolev/locmock/action v0.0.0-20230814175333-079dce7814fd

require gopkg.in/yaml.v3 v3.0.1 // indirect
