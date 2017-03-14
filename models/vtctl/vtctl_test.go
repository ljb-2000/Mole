package vtctl

import (
	"fmt"
	"testing"
	"time"
)

const (
	server = "http://192.168.80.172:15000"
)

// TestKeyspacesList
func TestKeyspacesList(t *testing.T) {
	list, err := KeyspacesList(server)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("%v", list)
}

// TestCreateKeyspace
func TestCreateKeyspace(t *testing.T) {
	param := []string{
		"CreateKeyspace",
		fmt.Sprintf("db%v", time.Now().Unix()),
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestDeleteKeyspace
func TestDeleteKeyspace(t *testing.T) {
	param := []string{
		"DeleteKeyspace",
		"-recursive",
		"gwgggg",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestValidateKeyspace
func TestValidateKeyspace(t *testing.T) {
	param := []string{
		"ValidateKeyspace",
		"-ping-tablets",
		"gw_keyspace",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestValidateSchemaKeyspace
func TestValidateSchemaKeyspace(t *testing.T) {
	param := []string{
		"ValidateSchemaKeyspace",
		"gw_keyspace",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestValidateVersionKeyspace
func TestValidateVersionKeyspace(t *testing.T) {
	param := []string{
		"ValidateVersionKeyspace",
		"test",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestRebuildKeyspaceGraph
func TestRebuildKeyspaceGraph(t *testing.T) {
	param := []string{
		"RebuildKeyspaceGraph",
		"gw_keyspace",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestCreateShard
func TestCreateShard(t *testing.T) {
	param := []string{
		"CreateShard",
		"test/-10",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestDeleteShard
func TestDeleteShard(t *testing.T) {
	param := []string{
		"DeleteShard",
		"-recursive",
		"-even_if_serving",
		"test/-10",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestGetSchema
func TestGetSchema(t *testing.T) {
	param := []string{
		"GetSchema",
		"test-1201",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestInitShardMaster
func TestInitShardMaster(t *testing.T) {
	param := []string{
		"InitShardMaster",
		"-force",
		"gw_keyspace/10-",
		"test-001",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestReparentTablet
func TestReparentTablet(t *testing.T) {
	param := []string{
		"ReparentTablet",
		"test-1301",
	}

	resp, err := Vtctl(server, param)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}

// TestShards
func TestShards(t *testing.T) {
	keyspace := "keyspace"

	list, err := Shards(server, keyspace)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("list:%v", list)
}

// TestTablets
func TestTablets(t *testing.T) {
	shard := "gw_keyspace/-10"

	list, err := Tablets(server, shard)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("list:%v", list)
}

// TestSchemaApply
func TestSchemaApply(t *testing.T) {
	keyspace := "gw_keyspace"
	sql := "create table tt1(id int(11))"

	resp, err := SchemaApply(server, keyspace, sql)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("error:%v output:%v", resp.Error, resp.Output)
}
