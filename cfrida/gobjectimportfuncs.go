package cfrida

import (
	"fmt"
	"runtime"
	"strings"
)

var(
	g_free = libfrida.NewProc(fixImportName("g_free"))
	g_object_ref = libfrida.NewProc(fixImportName("g_object_ref"))
	g_object_unref = libfrida.NewProc(fixImportName("g_object_unref"))
	g_signal_connect_data = libfrida.NewProc(fixImportName("g_signal_connect_data"))
	g_signal_handler_disconnect = libfrida.NewProc(fixImportName("g_signal_handler_disconnect"))
	g_strlen = libfrida.NewProc(fixImportName("g_strlen"))
	g_ref_string_length = libfrida.NewProc(fixImportName("g_ref_string_length"))
	g_error_get_message = libfrida.NewProc(fixImportName("g_error_get_message"))
	g_error_get_code = libfrida.NewProc(fixImportName("g_error_get_code"))
	g_error_free = libfrida.NewProc(fixImportName("g_error_free"))
	g_io_stream_close = libfrida.NewProc(fixImportName("g_io_stream_close"))
	g_output_stream_close = libfrida.NewProc(fixImportName("g_output_stream_close"))
	g_input_stream_close = libfrida.NewProc(fixImportName("g_input_stream_close"))
	g_output_stream_close_async = libfrida.NewProc(fixImportName("g_output_stream_close_async"))
	g_input_stream_close_async = libfrida.NewProc(fixImportName("g_input_stream_close_async"))
	g_input_stream_read_bytes = libfrida.NewProc(fixImportName("g_input_stream_read_bytes"))
	g_output_stream_write_bytes = libfrida.NewProc(fixImportName("g_output_stream_write_bytes"))
	g_output_stream_write_all = libfrida.NewProc(fixImportName("g_output_stream_write_all"))
	g_input_stream_read_all = libfrida.NewProc(fixImportName("g_input_stream_read_all"))
	g_io_stream_is_closed = libfrida.NewProc(fixImportName("g_io_stream_is_closed"))
	g_io_stream_get_input_stream = libfrida.NewProc(fixImportName("g_io_stream_get_input_stream"))
	g_io_stream_get_output_stream = libfrida.NewProc(fixImportName("g_io_stream_get_output_stream"))
	g_io_stream_has_pending = libfrida.NewProc(fixImportName("g_io_stream_has_pending"))
	g_io_stream_set_pending = libfrida.NewProc(fixImportName("g_io_stream_set_pending"))
	g_io_stream_clear_pending = libfrida.NewProc(fixImportName("g_io_stream_clear_pending"))
	g_bytes_get_data = libfrida.NewProc(fixImportName("g_bytes_get_data"))
	g_bytes_get_size = libfrida.NewProc(fixImportName("g_bytes_get_size"))
	g_bytes_ref = libfrida.NewProc(fixImportName("g_bytes_ref"))
	g_bytes_unref = libfrida.NewProc(fixImportName("g_bytes_unref"))
	g_bytes_new = libfrida.NewProc(fixImportName("g_bytes_new"))
	g_hash_table_unref = libfrida.NewProc(fixImportName("g_hash_table_unref"))
	g_hash_table_ref = libfrida.NewProc(fixImportName("g_hash_table_ref"))
	g_hash_table_iter_init = libfrida.NewProc(fixImportName("g_hash_table_iter_init"))
	g_hash_table_iter_new = libfrida.NewProc(fixImportName("g_hash_table_iter_new"))
	g_hash_table_iter_free = libfrida.NewProc(fixImportName("g_hash_table_iter_free"))
	g_hash_table_iter_next = libfrida.NewProc(fixImportName("g_hash_table_iter_next"))
	g_variant_is_of_type = libfrida.NewProc(fixImportName("g_variant_is_of_type"))
	g_variant_get_type_string = libfrida.NewProc(fixImportName("g_variant_get_type_string"))
	g_variant_get_string = libfrida.NewProc(fixImportName("g_variant_get_string"))
	g_variant_get_int64 = libfrida.NewProc(fixImportName("g_variant_get_int64"))
	g_variant_get_boolean = libfrida.NewProc(fixImportName("g_variant_get_boolean"))
	g_variant_get_variant = libfrida.NewProc(fixImportName("g_variant_get_variant"))
	g_variant_get_fixed_array = libfrida.NewProc(fixImportName("g_variant_get_fixed_array"))
	g_variant_iter_next_value = libfrida.NewProc(fixImportName("g_variant_iter_next_value"))
	g_variant_get_child_value = libfrida.NewProc(fixImportName("g_variant_get_child_value"))
	g_variant_unref = libfrida.NewProc(fixImportName("g_variant_unref"))
	g_variant_iter_init = libfrida.NewProc(fixImportName("g_variant_iter_init"))
	g_variant_type_new = libfrida.NewProc(fixImportName("g_variant_type_new"))
	g_variant_type_free = libfrida.NewProc(fixImportName("g_variant_type_free"))
	g_variant_iter_new = libfrida.NewProc(fixImportName("g_variant_iter_new"))
	g_variant_iter_free = libfrida.NewProc(fixImportName("g_variant_iter_free"))
	g_strv_length = libfrida.NewProc(fixImportName("g_strv_length"))
)

func fixImportName(name string)string{
	if runtime.GOOS=="windows"{
		return name
	}
	if strings.HasPrefix(name,"frida_")==false{
		return fmt.Sprintf("_frida_%s",name)
	}
	return name
}