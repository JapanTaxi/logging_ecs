package main

import (
	"errors"
	"fmt"
	"logging_ecs"
	"time"
)

func main() {

	fmt.Println("Test Start.........")

	detail_int := 0
	detail_string := "detail_string"
	detail_time := time.Now()

	detail_struct := &struct {
		Param1 int    `json:"param_int"`
		Param2 string `json:"param_string"`
		Param3 bool   `json:"param_bool"`
	}{
		Param1: 123,
		Param2: "test_param",
		Param3: true,
	}

	err := errors.New("Test Error")

	// Debug
	logger.Debug("test_debug_1")
	logger.DebugWithDetail("test_debug_2", detail_int)
	logger.DebugWithDetail("test_debug_3", detail_string)
	logger.DebugWithDetail("test_debug_4", detail_time)
	logger.DebugWithDetail("test_debug_5", detail_struct)

	// Info
	logger.Info("test_info_1")
	logger.InfoWithDetail("test_info_2", detail_int)
	logger.InfoWithDetail("test_info_3", detail_string)
	logger.InfoWithDetail("test_info_4", detail_time)
	logger.InfoWithDetail("test_info_5", detail_struct)

	// Warn
	logger.Warn("test_warn_1")
	logger.WarnWithDetail("test_warn_2", detail_int)
	logger.WarnWithDetail("test_warn_3", detail_string)
	logger.WarnWithDetail("test_warn_4", detail_time)
	logger.WarnWithDetail("test_warn_5", detail_struct)

	// Error
	logger.Error("test_error_1", err)
	logger.ErrorWithDetail("test_error_2", err, detail_int)
	logger.ErrorWithDetail("test_error_3", err, detail_string)
	logger.ErrorWithDetail("test_error_4", err, detail_time)
	logger.ErrorWithDetail("test_error_5", err, detail_struct)

	fmt.Println(".........Test End 出力結果を確認すること。5 * 4 レベルの計20行が出力されて入ればOK。")

}
