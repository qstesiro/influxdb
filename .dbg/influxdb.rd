# 编译与调试
{
    # 代码生成依赖c库,所以不要使用CGO_ENABLED=0
    alias gob='go build -v -tags "sqlite_foreign_keys,sqlite_json,assets" -gcflags "all=-N -l" -o influxd ./cmd/influxd/'

    alias dlv='gob && dlv exec ./influxd --init .dbg/influxd.dlv -- --reporting-disabled'
}
