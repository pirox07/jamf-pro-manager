# Utilities

## get-objects-params.go

- Mac のローカル環境で実行します。
- Go の実行環境が必要です。
  - [Go install](https://golang.org/doc/install)
- Jamf Pro の各オブジェクトのパラメータを取得、YAML 形式でファイル出力します。
- 対象オブジェクトは以下。
  - `Policies`
  - `Scripts`
  - `Computer Groups (Smart/Static)`
  - `Categories`

### Usage
```shell
export JAMF_BASE_URL=https://<your-tenant-name>.jamfcloud.com
export JAMF_USER=<User name of Jamf Pro account>
export JAMF_USER_PASSWORD=<Password of Jamf Pro account>

go get
go run get-objects-params.go

## Output
##  get-objects-params.go
##  out-conf
##   |- conf-policies.yml
##   |- conf-scripts.yml
##   |- scripts-contents
##   |   |- sample01.sh
##   |   |- sample02.sh
##   |    ~
##   |- conf-computer-groups.yml
##   └- conf-categories.yml
```
