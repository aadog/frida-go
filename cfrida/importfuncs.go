package cfrida

var (
	frida_init                                       = libfrida.NewProc("frida_init")
	frida_shutdown                                   = libfrida.NewProc("frida_shutdown")
	frida_deinit                                     = libfrida.NewProc("frida_deinit")
	frida_get_main_context                           = libfrida.NewProc("frida_get_main_context")
	frida_unref                                      = libfrida.NewProc("frida_unref")
	frida_version                                    = libfrida.NewProc("frida_version")
	frida_version_string                             = libfrida.NewProc("frida_version_string")
	frida_device_manager_new                         = libfrida.NewProc("frida_device_manager_new")
	frida_device_manager_close                       = libfrida.NewProc("frida_device_manager_close")
	frida_device_manager_close_finish                = libfrida.NewProc("frida_device_manager_close_finish")
	frida_device_manager_close_sync                  = libfrida.NewProc("frida_device_manager_close_sync")
	frida_device_manager_get_device_by_id            = libfrida.NewProc("frida_device_manager_get_device_by_id")
	frida_device_manager_get_device_by_id_finish     = libfrida.NewProc("frida_device_manager_get_device_by_id_finish")
	frida_device_manager_get_device_by_id_sync       = libfrida.NewProc("frida_device_manager_get_device_by_id_sync")
	frida_device_manager_get_device_by_type          = libfrida.NewProc("frida_device_manager_get_device_by_type")
	frida_device_manager_get_device_by_type_finish   = libfrida.NewProc("frida_device_manager_get_device_by_type_finish")
	frida_device_manager_get_device_by_type_sync     = libfrida.NewProc("frida_device_manager_get_device_by_type_sync")
	frida_device_manager_get_device                  = libfrida.NewProc("frida_device_manager_get_device")
	frida_device_manager_get_device_finish           = libfrida.NewProc("frida_device_manager_get_device_finish")
	frida_device_manager_get_device_sync             = libfrida.NewProc("frida_device_manager_get_device_sync")
	frida_device_manager_find_device_by_id           = libfrida.NewProc("frida_device_manager_find_device_by_id")
	frida_device_manager_find_device_by_id_finish    = libfrida.NewProc("frida_device_manager_find_device_by_id_finish")
	frida_device_manager_find_device_by_id_sync      = libfrida.NewProc("frida_device_manager_find_device_by_id_sync")
	frida_device_manager_find_device_by_type         = libfrida.NewProc("frida_device_manager_find_device_by_type")
	frida_device_manager_find_device_by_type_finish  = libfrida.NewProc("frida_device_manager_find_device_by_type_finish")
	frida_device_manager_find_device_by_type_sync    = libfrida.NewProc("frida_device_manager_find_device_by_type_sync")
	frida_device_manager_find_device                 = libfrida.NewProc("frida_device_manager_find_device")
	frida_device_manager_find_device_finish          = libfrida.NewProc("frida_device_manager_find_device_finish")
	frida_device_manager_find_device_sync            = libfrida.NewProc("frida_device_manager_find_device_sync")
	frida_device_manager_enumerate_devices           = libfrida.NewProc("frida_device_manager_enumerate_devices")
	frida_device_manager_enumerate_devices_finish    = libfrida.NewProc("frida_device_manager_enumerate_devices_finish")
	frida_device_manager_enumerate_devices_sync      = libfrida.NewProc("frida_device_manager_enumerate_devices_sync")
	frida_device_manager_add_remote_device           = libfrida.NewProc("frida_device_manager_add_remote_device")
	frida_device_manager_add_remote_device_finish    = libfrida.NewProc("frida_device_manager_add_remote_device_finish")
	frida_device_manager_add_remote_device_sync      = libfrida.NewProc("frida_device_manager_add_remote_device_sync")
	frida_device_manager_remove_remote_device        = libfrida.NewProc("frida_device_manager_remove_remote_device")
	frida_device_manager_remove_remote_device_finish = libfrida.NewProc("frida_device_manager_remove_remote_device_finish")
	frida_device_manager_remove_remote_device_sync   = libfrida.NewProc("frida_device_manager_remove_remote_device_sync")
	frida_session_options_new   = libfrida.NewProc("frida_session_options_new")
	frida_process_query_options_new   = libfrida.NewProc("frida_process_query_options_new")
	frida_remote_device_options_new   = libfrida.NewProc("frida_remote_device_options_new")
	frida_application_query_options_new   = libfrida.NewProc("frida_application_query_options_new")
	frida_frontmost_query_options_new   = libfrida.NewProc("frida_frontmost_query_options_new")
	frida_remote_device_options_set_certificate   = libfrida.NewProc("frida_remote_device_options_set_certificate")
	frida_remote_device_options_set_origin   = libfrida.NewProc("frida_remote_device_options_set_origin")
	frida_remote_device_options_set_token   = libfrida.NewProc("frida_remote_device_options_set_token")
	frida_remote_device_options_set_keepalive_interval   = libfrida.NewProc("frida_remote_device_options_set_keepalive_interval")
	frida_device_list_size                           = libfrida.NewProc("frida_device_list_size")
	frida_device_list_get                            = libfrida.NewProc("frida_device_list_get")
	frida_device_get_id                              = libfrida.NewProc("frida_device_get_id")
	frida_device_get_name                            = libfrida.NewProc("frida_device_get_name")
	frida_device_get_icon                            = libfrida.NewProc("frida_device_get_icon")
	frida_device_get_dtype                           = libfrida.NewProc("frida_device_get_dtype")
	frida_device_get_manager                         = libfrida.NewProc("frida_device_get_manager")
	frida_device_is_lost                             = libfrida.NewProc("frida_device_is_lost")
	frida_device_get_frontmost_application           = libfrida.NewProc("frida_device_get_frontmost_application")
	frida_device_get_frontmost_application_finish    = libfrida.NewProc("frida_device_get_frontmost_application_finish")
	frida_device_get_frontmost_application_sync      = libfrida.NewProc("frida_device_get_frontmost_application_sync")
	frida_device_enumerate_applications              = libfrida.NewProc("frida_device_enumerate_applications")
	frida_device_enumerate_applications_finish       = libfrida.NewProc("frida_device_enumerate_applications_finish")
	frida_device_enumerate_applications_sync         = libfrida.NewProc("frida_device_enumerate_applications_sync")
	frida_application_query_options_select_identifier         = libfrida.NewProc("frida_application_query_options_select_identifier")
	frida_application_query_options_set_scope         = libfrida.NewProc("frida_application_query_options_set_scope")
	frida_frontmost_query_options_set_scope         = libfrida.NewProc("frida_frontmost_query_options_set_scope")
	frida_process_query_options_select_pid         = libfrida.NewProc("frida_process_query_options_select_pid")
	frida_process_query_options_set_scope         = libfrida.NewProc("frida_process_query_options_set_scope")
	frida_device_get_process_by_pid                  = libfrida.NewProc("frida_device_get_process_by_pid")
	frida_device_get_process_by_pid_finish           = libfrida.NewProc("frida_device_get_process_by_pid_finish")
	frida_device_get_process_by_pid_sync             = libfrida.NewProc("frida_device_get_process_by_pid_sync")
	frida_device_get_process_by_name                 = libfrida.NewProc("frida_device_get_process_by_name")
	frida_device_get_process_by_name_finish          = libfrida.NewProc("frida_device_get_process_by_name_finish")
	frida_device_get_process_by_name_sync            = libfrida.NewProc("frida_device_get_process_by_name_sync")
	frida_device_get_process                         = libfrida.NewProc("frida_device_get_process")
	frida_device_get_process_finish                  = libfrida.NewProc("frida_device_get_process_finish")
	frida_device_get_process_sync                    = libfrida.NewProc("frida_device_get_process_sync")
	frida_device_find_process_by_pid                 = libfrida.NewProc("frida_device_find_process_by_pid")
	frida_device_find_process_by_pid_finish          = libfrida.NewProc("frida_device_find_process_by_pid_finish")
	frida_device_find_process_by_pid_sync            = libfrida.NewProc("frida_device_find_process_by_pid_sync")
	frida_device_find_process_by_name                = libfrida.NewProc("frida_device_find_process_by_name")
	frida_device_find_process_by_name_finish         = libfrida.NewProc("frida_device_find_process_by_name_finish")
	frida_device_find_process_by_name_sync           = libfrida.NewProc("frida_device_find_process_by_name_sync")
	frida_device_find_process                        = libfrida.NewProc("frida_device_find_process")
	frida_device_find_process_finish                 = libfrida.NewProc("frida_device_find_process_finish")
	frida_device_find_process_sync                   = libfrida.NewProc("frida_device_find_process_sync")
	frida_device_enumerate_processes                 = libfrida.NewProc("frida_device_enumerate_processes")
	frida_device_enumerate_processes_finish          = libfrida.NewProc("frida_device_enumerate_processes_finish")
	frida_device_enumerate_processes_sync            = libfrida.NewProc("frida_device_enumerate_processes_sync")
	frida_device_enable_spawn_gating                 = libfrida.NewProc("frida_device_enable_spawn_gating")
	frida_device_enable_spawn_gating_finish          = libfrida.NewProc("frida_device_enable_spawn_gating_finish")
	frida_device_enable_spawn_gating_sync            = libfrida.NewProc("frida_device_enable_spawn_gating_sync")
	frida_device_disable_spawn_gating                = libfrida.NewProc("frida_device_disable_spawn_gating")
	frida_device_disable_spawn_gating_finish         = libfrida.NewProc("frida_device_disable_spawn_gating_finish")
	frida_device_disable_spawn_gating_sync           = libfrida.NewProc("frida_device_disable_spawn_gating_sync")
	frida_device_enumerate_pending_spawn             = libfrida.NewProc("frida_device_enumerate_pending_spawn")
	frida_device_enumerate_pending_spawn_finish      = libfrida.NewProc("frida_device_enumerate_pending_spawn_finish")
	frida_device_enumerate_pending_spawn_sync        = libfrida.NewProc("frida_device_enumerate_pending_spawn_sync")
	frida_device_enumerate_pending_children          = libfrida.NewProc("frida_device_enumerate_pending_children")
	frida_device_enumerate_pending_children_finish   = libfrida.NewProc("frida_device_enumerate_pending_children_finish")
	frida_device_enumerate_pending_children_sync     = libfrida.NewProc("frida_device_enumerate_pending_children_sync")
	frida_device_spawn                               = libfrida.NewProc("frida_device_spawn")
	frida_device_spawn_finish                        = libfrida.NewProc("frida_device_spawn_finish")
	frida_device_spawn_sync                          = libfrida.NewProc("frida_device_spawn_sync")
	frida_device_input                               = libfrida.NewProc("frida_device_input")
	frida_device_input_finish                        = libfrida.NewProc("frida_device_input_finish")
	frida_device_input_sync                          = libfrida.NewProc("frida_device_input_sync")
	frida_device_resume                              = libfrida.NewProc("frida_device_resume")
	frida_device_resume_finish                       = libfrida.NewProc("frida_device_resume_finish")
	frida_device_resume_sync                         = libfrida.NewProc("frida_device_resume_sync")
	frida_device_kill                                = libfrida.NewProc("frida_device_kill")
	frida_device_kill_finish                         = libfrida.NewProc("frida_device_kill_finish")
	frida_device_kill_sync                           = libfrida.NewProc("frida_device_kill_sync")
	frida_device_open_channel_sync                           = libfrida.NewProc("frida_device_open_channel_sync")
	frida_device_get_host_session_sync                           = libfrida.NewProc("frida_device_get_host_session_sync")
	frida_device_attach                              = libfrida.NewProc("frida_device_attach")
	frida_device_attach_finish                       = libfrida.NewProc("frida_device_attach_finish")
	frida_device_attach_sync                         = libfrida.NewProc("frida_device_attach_sync")
	frida_device_inject_library_file                 = libfrida.NewProc("frida_device_inject_library_file")
	frida_device_inject_library_file_finish          = libfrida.NewProc("frida_device_inject_library_file_finish")
	frida_device_inject_library_file_sync            = libfrida.NewProc("frida_device_inject_library_file_sync")
	frida_device_inject_library_blob                 = libfrida.NewProc("frida_device_inject_library_blob")
	frida_device_inject_library_blob_finish          = libfrida.NewProc("frida_device_inject_library_blob_finish")
	frida_device_inject_library_blob_sync            = libfrida.NewProc("frida_device_inject_library_blob_sync")
	frida_application_list_size                      = libfrida.NewProc("frida_application_list_size")
	frida_application_list_get                       = libfrida.NewProc("frida_application_list_get")
	frida_application_get_identifier                 = libfrida.NewProc("frida_application_get_identifier")
	frida_application_get_name                       = libfrida.NewProc("frida_application_get_name")
	frida_application_get_pid                        = libfrida.NewProc("frida_application_get_pid")
	frida_application_get_parameters                        = libfrida.NewProc("frida_application_get_parameters")
	frida_application_get_small_icon                 = libfrida.NewProc("frida_application_get_small_icon")
	frida_application_get_large_icon                 = libfrida.NewProc("frida_application_get_large_icon")
	frida_process_list_size                          = libfrida.NewProc("frida_process_list_size")
	frida_process_list_get                           = libfrida.NewProc("frida_process_list_get")
	frida_process_get_pid                            = libfrida.NewProc("frida_process_get_pid")
	frida_process_get_name                           = libfrida.NewProc("frida_process_get_name")
	frida_process_get_small_icon                     = libfrida.NewProc("frida_process_get_small_icon")
	frida_process_get_large_icon                     = libfrida.NewProc("frida_process_get_large_icon")
	frida_spawn_options_new                          = libfrida.NewProc("frida_spawn_options_new")
	frida_spawn_options_get_argv                     = libfrida.NewProc("frida_spawn_options_get_argv")
	frida_spawn_options_get_envp                     = libfrida.NewProc("frida_spawn_options_get_envp")
	frida_spawn_options_get_env                      = libfrida.NewProc("frida_spawn_options_get_env")
	frida_spawn_options_get_cwd                      = libfrida.NewProc("frida_spawn_options_get_cwd")
	frida_spawn_options_get_stdio                    = libfrida.NewProc("frida_spawn_options_get_stdio")
	frida_spawn_options_get_aux                      = libfrida.NewProc("frida_spawn_options_get_aux")
	frida_spawn_options_set_argv                     = libfrida.NewProc("frida_spawn_options_set_argv")
	frida_session_options_set_realm                     = libfrida.NewProc("frida_session_options_set_realm")
	frida_session_options_set_persist_timeout                     = libfrida.NewProc("frida_session_options_set_persist_timeout")
	frida_spawn_options_set_envp                     = libfrida.NewProc("frida_spawn_options_set_envp")
	frida_spawn_options_set_env                      = libfrida.NewProc("frida_spawn_options_set_env")
	frida_spawn_options_set_cwd                      = libfrida.NewProc("frida_spawn_options_set_cwd")
	frida_spawn_options_set_stdio                    = libfrida.NewProc("frida_spawn_options_set_stdio")
	frida_spawn_list_size                            = libfrida.NewProc("frida_spawn_list_size")
	frida_spawn_list_get                             = libfrida.NewProc("frida_spawn_list_get")
	frida_spawn_get_pid                              = libfrida.NewProc("frida_spawn_get_pid")
	frida_spawn_get_identifier                       = libfrida.NewProc("frida_spawn_get_identifier")
	frida_child_list_size                            = libfrida.NewProc("frida_child_list_size")
	frida_child_list_get                             = libfrida.NewProc("frida_child_list_get")
	frida_child_get_pid                              = libfrida.NewProc("frida_child_get_pid")
	frida_child_get_parent_pid                       = libfrida.NewProc("frida_child_get_parent_pid")
	frida_child_get_origin                           = libfrida.NewProc("frida_child_get_origin")
	frida_child_get_identifier                       = libfrida.NewProc("frida_child_get_identifier")
	frida_child_get_path                             = libfrida.NewProc("frida_child_get_path")
	frida_child_get_argv                             = libfrida.NewProc("frida_child_get_argv")
	frida_child_get_envp                             = libfrida.NewProc("frida_child_get_envp")
	frida_crash_get_pid                              = libfrida.NewProc("frida_crash_get_pid")
	frida_crash_get_process_name                     = libfrida.NewProc("frida_crash_get_process_name")
	frida_crash_get_summary                          = libfrida.NewProc("frida_crash_get_summary")
	frida_crash_get_report                           = libfrida.NewProc("frida_crash_get_report")
	frida_crash_load_parameters                      = libfrida.NewProc("frida_crash_load_parameters")
	frida_icon_get_width                             = libfrida.NewProc("frida_icon_get_width")
	frida_icon_get_height                            = libfrida.NewProc("frida_icon_get_height")
	frida_icon_get_rowstride                         = libfrida.NewProc("frida_icon_get_rowstride")
	frida_icon_get_pixels                            = libfrida.NewProc("frida_icon_get_pixels")
	frida_session_get_pid                            = libfrida.NewProc("frida_session_get_pid")
	frida_session_get_persist_timeout                            = libfrida.NewProc("frida_session_get_persist_timeout")
	frida_session_get_device                         = libfrida.NewProc("frida_session_get_device")
	frida_session_is_detached                        = libfrida.NewProc("frida_session_is_detached")
	frida_session_detach                             = libfrida.NewProc("frida_session_detach")
	frida_session_detach_finish                      = libfrida.NewProc("frida_session_detach_finish")
	frida_session_detach_sync                        = libfrida.NewProc("frida_session_detach_sync")
	frida_session_resume_sync                        = libfrida.NewProc("frida_session_resume_sync")
	frida_session_enable_child_gating                = libfrida.NewProc("frida_session_enable_child_gating")
	frida_session_enable_child_gating_finish         = libfrida.NewProc("frida_session_enable_child_gating_finish")
	frida_session_enable_child_gating_sync           = libfrida.NewProc("frida_session_enable_child_gating_sync")
	frida_session_disable_child_gating               = libfrida.NewProc("frida_session_disable_child_gating")
	frida_session_disable_child_gating_finish        = libfrida.NewProc("frida_session_disable_child_gating_finish")
	frida_session_disable_child_gating_sync          = libfrida.NewProc("frida_session_disable_child_gating_sync")
	frida_session_create_script                      = libfrida.NewProc("frida_session_create_script")
	frida_session_create_script_finish               = libfrida.NewProc("frida_session_create_script_finish")
	frida_session_create_script_sync                 = libfrida.NewProc("frida_session_create_script_sync")
	frida_session_create_script_from_bytes           = libfrida.NewProc("frida_session_create_script_from_bytes")
	frida_session_create_script_from_bytes_finish    = libfrida.NewProc("frida_session_create_script_from_bytes_finish")
	frida_session_create_script_from_bytes_sync      = libfrida.NewProc("frida_session_create_script_from_bytes_sync")
	frida_session_compile_script                     = libfrida.NewProc("frida_session_compile_script")
	frida_session_compile_script_finish              = libfrida.NewProc("frida_session_compile_script_finish")
	frida_session_compile_script_sync                = libfrida.NewProc("frida_session_compile_script_sync")
	frida_session_enable_debugger                    = libfrida.NewProc("frida_session_enable_debugger")
	frida_session_enable_debugger_finish             = libfrida.NewProc("frida_session_enable_debugger_finish")
	frida_session_enable_debugger_sync               = libfrida.NewProc("frida_session_enable_debugger_sync")
	frida_session_disable_debugger                   = libfrida.NewProc("frida_session_disable_debugger")
	frida_session_disable_debugger_finish            = libfrida.NewProc("frida_session_disable_debugger_finish")
	frida_session_disable_debugger_sync              = libfrida.NewProc("frida_session_disable_debugger_sync")
	frida_session_setup_peer_connection_sync              = libfrida.NewProc("frida_session_setup_peer_connection_sync")
	frida_relay_get_address              = libfrida.NewProc("frida_relay_get_address")
	frida_relay_new              = libfrida.NewProc("frida_relay_new")
	frida_relay_get_username              = libfrida.NewProc("frida_relay_get_username")
	frida_relay_get_password              = libfrida.NewProc("frida_relay_get_password")
	frida_relay_get_kind              = libfrida.NewProc("frida_relay_get_kind")
	frida_session_enable_jit                         = libfrida.NewProc("frida_session_enable_jit")
	frida_session_enable_jit_finish                  = libfrida.NewProc("frida_session_enable_jit_finish")
	frida_session_enable_jit_sync                    = libfrida.NewProc("frida_session_enable_jit_sync")
	frida_script_get_id                              = libfrida.NewProc("frida_script_get_id")
	frida_script_is_destroyed                        = libfrida.NewProc("frida_script_is_destroyed")
	frida_script_load                                = libfrida.NewProc("frida_script_load")
	frida_script_load_finish                         = libfrida.NewProc("frida_script_load_finish")
	frida_script_load_sync                           = libfrida.NewProc("frida_script_load_sync")
	frida_script_unload                              = libfrida.NewProc("frida_script_unload")
	frida_script_unload_finish                       = libfrida.NewProc("frida_script_unload_finish")
	frida_script_unload_sync                         = libfrida.NewProc("frida_script_unload_sync")
	frida_script_eternalize                          = libfrida.NewProc("frida_script_eternalize")
	frida_script_eternalize_finish                   = libfrida.NewProc("frida_script_eternalize_finish")
	frida_script_eternalize_sync                     = libfrida.NewProc("frida_script_eternalize_sync")
	frida_script_post                                = libfrida.NewProc("frida_script_post")
	frida_script_post_finish                         = libfrida.NewProc("frida_script_post_finish")
	frida_script_post_sync                           = libfrida.NewProc("frida_script_post_sync")
	frida_script_options_new                         = libfrida.NewProc("frida_script_options_new")
	frida_peer_options_new                         = libfrida.NewProc("frida_peer_options_new")
	frida_script_options_get_name                    = libfrida.NewProc("frida_script_options_get_name")
	frida_script_options_get_runtime                 = libfrida.NewProc("frida_script_options_get_runtime")
	frida_script_options_set_name                    = libfrida.NewProc("frida_script_options_set_name")
	frida_script_options_set_runtime                 = libfrida.NewProc("frida_script_options_set_runtime")
	frida_peer_options_clear_relays                 = libfrida.NewProc("frida_peer_options_clear_relays")
	frida_peer_options_add_relay                 = libfrida.NewProc("frida_peer_options_add_relay")
	frida_peer_options_enumerate_relays                 = libfrida.NewProc("frida_peer_options_enumerate_relays")
	frida_peer_options_set_stun_server                 = libfrida.NewProc("frida_peer_options_set_stun_server")
	frida_injector_new                               = libfrida.NewProc("frida_injector_new")
	frida_injector_new_inprocess                     = libfrida.NewProc("frida_injector_new_inprocess")
	frida_injector_close                             = libfrida.NewProc("frida_injector_close")
	frida_injector_close_finish                      = libfrida.NewProc("frida_injector_close_finish")
	frida_injector_close_sync                        = libfrida.NewProc("frida_injector_close_sync")
	frida_injector_inject_library_file               = libfrida.NewProc("frida_injector_inject_library_file")
	frida_injector_inject_library_file_finish        = libfrida.NewProc("frida_injector_inject_library_file_finish")
	frida_injector_inject_library_file_sync          = libfrida.NewProc("frida_injector_inject_library_file_sync")
	frida_injector_inject_library_blob               = libfrida.NewProc("frida_injector_inject_library_blob")
	frida_injector_inject_library_blob_finish        = libfrida.NewProc("frida_injector_inject_library_blob_finish")
	frida_injector_inject_library_blob_sync          = libfrida.NewProc("frida_injector_inject_library_blob_sync")
	frida_injector_demonitor_and_clone_state         = libfrida.NewProc("frida_injector_demonitor_and_clone_state")
	frida_injector_demonitor_and_clone_state_finish  = libfrida.NewProc("frida_injector_demonitor_and_clone_state_finish")
	frida_injector_demonitor_and_clone_state_sync    = libfrida.NewProc("frida_injector_demonitor_and_clone_state_sync")
	frida_injector_recreate_thread                   = libfrida.NewProc("frida_injector_recreate_thread")
	frida_injector_recreate_thread_finish            = libfrida.NewProc("frida_injector_recreate_thread_finish")
	frida_injector_recreate_thread_sync              = libfrida.NewProc("frida_injector_recreate_thread_sync")
	frida_file_monitor_new                           = libfrida.NewProc("frida_file_monitor_new")
	frida_file_monitor_get_path                      = libfrida.NewProc("frida_file_monitor_get_path")
	frida_file_monitor_enable                        = libfrida.NewProc("frida_file_monitor_enable")
	frida_file_monitor_enable_finish                 = libfrida.NewProc("frida_file_monitor_enable_finish")
	frida_file_monitor_enable_sync                   = libfrida.NewProc("frida_file_monitor_enable_sync")
	frida_file_monitor_disable                       = libfrida.NewProc("frida_file_monitor_disable")
	frida_file_monitor_disable_finish                = libfrida.NewProc("frida_file_monitor_disable_finish")
	frida_file_monitor_disable_sync                  = libfrida.NewProc("frida_file_monitor_disable_sync")
	frida_error_quark                                = libfrida.NewProc("frida_error_quark")
	frida_runtime_get_type                           = libfrida.NewProc("frida_runtime_get_type")
	frida_device_type_get_type                       = libfrida.NewProc("frida_device_type_get_type")
	frida_child_origin_get_type                      = libfrida.NewProc("frida_child_origin_get_type")
	frida_script_runtime_get_type                    = libfrida.NewProc("frida_script_runtime_get_type")
	frida_session_detach_reason_get_type             = libfrida.NewProc("frida_session_detach_reason_get_type")
	frida_stdio_get_type                             = libfrida.NewProc("frida_stdio_get_type")
	frida_unload_policy_get_type                     = libfrida.NewProc("frida_unload_policy_get_type")
	frida_device_manager_get_type                    = libfrida.NewProc("frida_device_manager_get_type")
	frida_device_list_get_type                       = libfrida.NewProc("frida_device_list_get_type")
	frida_device_get_type                            = libfrida.NewProc("frida_device_get_type")
	frida_application_list_get_type                  = libfrida.NewProc("frida_application_list_get_type")
	frida_application_get_type                       = libfrida.NewProc("frida_application_get_type")
	frida_process_list_get_type                      = libfrida.NewProc("frida_process_list_get_type")
	frida_process_get_type                           = libfrida.NewProc("frida_process_get_type")
	frida_spawn_options_get_type                     = libfrida.NewProc("frida_spawn_options_get_type")
	frida_spawn_list_get_type                        = libfrida.NewProc("frida_spawn_list_get_type")
	frida_spawn_get_type                             = libfrida.NewProc("frida_spawn_get_type")
	frida_child_list_get_type                        = libfrida.NewProc("frida_child_list_get_type")
	frida_child_get_type                             = libfrida.NewProc("frida_child_get_type")
	frida_crash_get_type                             = libfrida.NewProc("frida_crash_get_type")
	frida_icon_get_type                              = libfrida.NewProc("frida_icon_get_type")
	frida_session_get_type                           = libfrida.NewProc("frida_session_get_type")
	frida_script_get_type                            = libfrida.NewProc("frida_script_get_type")
	frida_script_options_get_type                    = libfrida.NewProc("frida_script_options_get_type")
	frida_injector_get_type                          = libfrida.NewProc("frida_injector_get_type")
	frida_file_monitor_get_type                      = libfrida.NewProc("frida_file_monitor_get_type")
)