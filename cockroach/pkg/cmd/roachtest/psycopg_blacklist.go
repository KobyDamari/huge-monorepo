// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

package main

var psycopgBlacklists = blacklistsForVersion{
	{"v2.2", "psycopgBlackList19_1", psycopgBlackList19_1, "psycopgIgnoreList19_1", psycopgIgnoreList19_1},
	{"v19.1", "psycopgBlackList19_1", psycopgBlackList19_1, "psycopgIgnoreList19_1", psycopgIgnoreList19_1},
	{"v19.2", "psycopgBlackList19_2", psycopgBlackList19_2, "psycopgIgnoreList19_2", psycopgIgnoreList19_2},
}

// These are lists of known psycopg test errors and failures.
// When the psycopg test suite is run, the results are compared to this list.
// Any passed test that is not on this list is reported as PASS - expected
// Any passed test that is on this list is reported as PASS - unexpected
// Any failed test that is on this list is reported as FAIL - expected
// Any failed test that is not on this list is reported as FAIL - unexpected
// Any test on this list that is not run is reported as FAIL - not run
//
// Please keep these lists alphabetized for easy diffing.
// After a failed run, an updated version of this blacklist should be available
// in the test log.
var psycopgBlackList19_2 = psycopgBlackList19_1

var psycopgBlackList19_1 = blacklist{
	"tests.test_async.AsyncTests.test_async_after_async":                                                     "5807",
	"tests.test_async.AsyncTests.test_async_callproc":                                                        "5807",
	"tests.test_async.AsyncTests.test_async_connection_error_message":                                        "5807",
	"tests.test_async.AsyncTests.test_async_cursor_gone":                                                     "5807",
	"tests.test_async.AsyncTests.test_async_dont_read_all":                                                   "5807",
	"tests.test_async.AsyncTests.test_async_executemany":                                                     "5807",
	"tests.test_async.AsyncTests.test_async_fetch_wrong_cursor":                                              "5807",
	"tests.test_async.AsyncTests.test_async_iter":                                                            "5807",
	"tests.test_async.AsyncTests.test_async_named_cursor":                                                    "5807",
	"tests.test_async.AsyncTests.test_async_scroll":                                                          "5807",
	"tests.test_async.AsyncTests.test_async_select":                                                          "5807",
	"tests.test_async.AsyncTests.test_async_subclass":                                                        "5807",
	"tests.test_async.AsyncTests.test_commit_while_async":                                                    "5807",
	"tests.test_async.AsyncTests.test_connection_setup":                                                      "5807",
	"tests.test_async.AsyncTests.test_copy_no_hang":                                                          "5807",
	"tests.test_async.AsyncTests.test_copy_while_async":                                                      "5807",
	"tests.test_async.AsyncTests.test_error":                                                                 "5807",
	"tests.test_async.AsyncTests.test_error_two_cursors":                                                     "5807",
	"tests.test_async.AsyncTests.test_fetch_after_async":                                                     "5807",
	"tests.test_async.AsyncTests.test_flush_on_write":                                                        "5807",
	"tests.test_async.AsyncTests.test_lobject_while_async":                                                   "5807",
	"tests.test_async.AsyncTests.test_non_block_after_notification":                                          "5807",
	"tests.test_async.AsyncTests.test_notices":                                                               "5807",
	"tests.test_async.AsyncTests.test_notify":                                                                "5807",
	"tests.test_async.AsyncTests.test_poll_conn_for_notification":                                            "5807",
	"tests.test_async.AsyncTests.test_poll_noop":                                                             "5807",
	"tests.test_async.AsyncTests.test_reset_while_async":                                                     "5807",
	"tests.test_async.AsyncTests.test_rollback_while_async":                                                  "5807",
	"tests.test_async.AsyncTests.test_scroll":                                                                "5807",
	"tests.test_async.AsyncTests.test_set_parameters_while_async":                                            "5807",
	"tests.test_async.AsyncTests.test_stop_on_first_error":                                                   "5807",
	"tests.test_async.AsyncTests.test_sync_poll":                                                             "5807",
	"tests.test_async_keyword.AsyncTests.test_async_connection_error_message":                                "5807",
	"tests.test_async_keyword.AsyncTests.test_async_subclass":                                                "5807",
	"tests.test_async_keyword.AsyncTests.test_connection_setup":                                              "5807",
	"tests.test_async_keyword.CancelTests.test_async_cancel":                                                 "5807",
	"tests.test_async_keyword.CancelTests.test_async_connection_cancel":                                      "5807",
	"tests.test_cancel.CancelTests.test_async_cancel":                                                        "5807",
	"tests.test_cancel.CancelTests.test_async_connection_cancel":                                             "5807",
	"tests.test_cancel.CancelTests.test_cancel":                                                              "5807",
	"tests.test_cancel.CancelTests.test_empty_cancel":                                                        "5807",
	"tests.test_connection.AutocommitTests.test_set_session_autocommit":                                      "35879",
	"tests.test_connection.ConnectionTests.test_cleanup_on_badconn_close":                                    "35897",
	"tests.test_connection.ConnectionTests.test_encoding_name":                                               "35882",
	"tests.test_connection.ConnectionTests.test_notices":                                                     "5807",
	"tests.test_connection.ConnectionTests.test_notices_consistent_order":                                    "5807",
	"tests.test_connection.ConnectionTests.test_notices_deque":                                               "5807",
	"tests.test_connection.ConnectionTests.test_notices_limited":                                             "5807",
	"tests.test_connection.ConnectionTests.test_notices_noappend":                                            "5807",
	"tests.test_connection.ConnectionTests.test_reset":                                                       "35879",
	"tests.test_connection.ConnectionTwoPhaseTests.test_cancel_fails_prepared":                               "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_recovered_xids":                                      "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_status_after_recover":                                "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_tpc_commit":                                          "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_tpc_commit_one_phase":                                "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_tpc_commit_recovered":                                "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_tpc_recover_non_dbapi_connection":                    "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_tpc_rollback":                                        "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_tpc_rollback_one_phase":                              "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_tpc_rollback_recovered":                              "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_unparsed_roundtrip":                                  "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_xid_construction":                                    "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_xid_encoding":                                        "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_xid_from_string":                                     "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_xid_roundtrip":                                       "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_xid_to_string":                                       "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_xid_unicode":                                         "22329",
	"tests.test_connection.ConnectionTwoPhaseTests.test_xid_unicode_unparsed":                                "22329",
	"tests.test_connection.IsolationLevelsTestCase.test_encoding":                                            "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_isolation_level":                                     "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_isolation_level_autocommit":                          "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_isolation_level_closed":                              "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_isolation_level_read_committed":                      "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_isolation_level_serializable":                        "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_set_isolation_level":                                 "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_set_isolation_level_abort":                           "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_set_isolation_level_autocommit":                      "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_set_isolation_level_default":                         "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_setattr_isolation_level_int":                         "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_setattr_isolation_level_invalid":                     "12123",
	"tests.test_connection.IsolationLevelsTestCase.test_setattr_isolation_level_str":                         "12123",
	"tests.test_connection.TestConnectionInfo.test_backend_pid":                                              "unknown",
	"tests.test_connection.TransactionControlTests.test_idempotence_check":                                   "35879",
	"tests.test_connection.TransactionControlTests.test_mixing_session_attribs":                              "35879",
	"tests.test_connection.TransactionControlTests.test_set_deferrable":                                      "31632",
	"tests.test_connection.TransactionControlTests.test_set_isolation_level":                                 "12123",
	"tests.test_connection.TransactionControlTests.test_set_isolation_level_str":                             "12123",
	"tests.test_connection.TransactionControlTests.test_setattr_deferrable":                                  "12123",
	"tests.test_copy.CopyTests.test_copy_bytes":                                                              "5807",
	"tests.test_copy.CopyTests.test_copy_expert_file_refcount":                                               "5807",
	"tests.test_copy.CopyTests.test_copy_expert_textiobase":                                                  "5807",
	"tests.test_copy.CopyTests.test_copy_from":                                                               "5807",
	"tests.test_copy.CopyTests.test_copy_from_cols":                                                          "5807",
	"tests.test_copy.CopyTests.test_copy_from_cols_err":                                                      "5807",
	"tests.test_copy.CopyTests.test_copy_from_insane_size":                                                   "5807",
	"tests.test_copy.CopyTests.test_copy_from_propagate_error":                                               "5807",
	"tests.test_copy.CopyTests.test_copy_from_segfault":                                                      "5807",
	"tests.test_copy.CopyTests.test_copy_no_column_limit":                                                    "5807",
	"tests.test_copy.CopyTests.test_copy_rowcount":                                                           "5807",
	"tests.test_copy.CopyTests.test_copy_rowcount_error":                                                     "5807",
	"tests.test_copy.CopyTests.test_copy_text":                                                               "5807",
	"tests.test_copy.CopyTests.test_copy_to":                                                                 "5807",
	"tests.test_copy.CopyTests.test_copy_to_propagate_error":                                                 "5807",
	"tests.test_copy.CopyTests.test_copy_to_segfault":                                                        "5807",
	"tests.test_cursor.CursorTests.test_callproc_dict":                                                       "17511",
	"tests.test_cursor.CursorTests.test_description_attribs":                                                 "unknown",
	"tests.test_cursor.CursorTests.test_description_extra_attribs":                                           "unknown",
	"tests.test_cursor.CursorTests.test_executemany_propagate_exceptions":                                    "5807",
	"tests.test_cursor.CursorTests.test_external_close_async":                                                "35897",
	"tests.test_cursor.CursorTests.test_external_close_sync":                                                 "35897",
	"tests.test_cursor.CursorTests.test_invalid_name":                                                        "5807",
	"tests.test_cursor.CursorTests.test_iter_named_cursor_default_itersize":                                  "30352",
	"tests.test_cursor.CursorTests.test_iter_named_cursor_efficient":                                         "30352",
	"tests.test_cursor.CursorTests.test_iter_named_cursor_itersize":                                          "30352",
	"tests.test_cursor.CursorTests.test_iter_named_cursor_rownumber":                                         "30352",
	"tests.test_cursor.CursorTests.test_named_cursor_stealing":                                               "30352",
	"tests.test_cursor.CursorTests.test_named_noop_close":                                                    "30352",
	"tests.test_cursor.CursorTests.test_not_scrollable":                                                      "30352",
	"tests.test_cursor.CursorTests.test_scroll_named":                                                        "30352",
	"tests.test_cursor.CursorTests.test_scrollable":                                                          "30352",
	"tests.test_cursor.CursorTests.test_stolen_named_cursor_close":                                           "30352",
	"tests.test_cursor.CursorTests.test_withhold":                                                            "30352",
	"tests.test_cursor.CursorTests.test_withhold_autocommit":                                                 "30352",
	"tests.test_cursor.CursorTests.test_withhold_no_begin":                                                   "30352",
	"tests.test_dates.DatetimeTests.test_adapt_datetime":                                                     "36115",
	"tests.test_dates.DatetimeTests.test_adapt_infinity_tz":                                                  "36116",
	"tests.test_dates.DatetimeTests.test_adapt_negative_timedelta":                                           "35807",
	"tests.test_dates.DatetimeTests.test_adapt_timedelta":                                                    "35807",
	"tests.test_dates.DatetimeTests.test_interval_iso_8601_not_supported":                                    "32562",
	"tests.test_dates.DatetimeTests.test_time_24":                                                            "36118",
	"tests.test_dates.DatetimeTests.test_type_roundtrip_timetz":                                              "26097",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorRealWithNamedCursorFetchAll":       "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorRealWithNamedCursorFetchMany":      "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorRealWithNamedCursorFetchManyNoarg": "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorRealWithNamedCursorFetchOne":       "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorRealWithNamedCursorIter":           "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorRealWithNamedCursorIterRowNumber":  "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorRealWithNamedCursorNotGreedy":      "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorWithPlainCursorRealFetchAll":       "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorWithPlainCursorRealFetchMany":      "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorWithPlainCursorRealFetchManyNoarg": "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorWithPlainCursorRealFetchOne":       "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorWithPlainCursorRealIter":           "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testDictCursorWithPlainCursorRealIterRowNumber":  "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testPickleRealDictRow":                           "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.testRealMeansReal":                               "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.test_iter_methods_2":                             "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.test_mod":                                        "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.test_order":                                      "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.test_order_iter":                                 "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorRealTests.test_pop":                                        "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictConnCursorArgs":                              "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithNamedCursorFetchAll":               "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithNamedCursorFetchMany":              "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithNamedCursorFetchManyNoarg":         "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithNamedCursorFetchOne":               "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithNamedCursorIter":                   "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithNamedCursorIterRowNumber":          "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithNamedCursorNotGreedy":              "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithPlainCursorFetchAll":               "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithPlainCursorFetchMany":              "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithPlainCursorFetchManyNoarg":         "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithPlainCursorFetchOne":               "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithPlainCursorIter":                   "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testDictCursorWithPlainCursorIterRowNumber":          "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testPickleDictRow":                                   "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.testUpdateRow":                                       "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.test_iter_methods_2":                                 "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.test_order":                                          "5807",
	"tests.test_extras_dictcursor.ExtrasDictCursorTests.test_order_iter":                                     "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_bad_col_names":                                   "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_cache":                                           "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_cursor_args":                                     "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_executemany":                                     "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_fetchall":                                        "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_fetchmany":                                       "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_fetchmany_noarg":                                 "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_fetchone":                                        "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_iter":                                            "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_max_cache":                                       "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_minimal_generation":                              "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_named":                                           "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_named_fetchall":                                  "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_named_fetchmany":                                 "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_named_fetchone":                                  "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_named_rownumber":                                 "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_no_result_no_surprise":                           "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_not_greedy":                                      "5807",
	"tests.test_extras_dictcursor.NamedTupleCursorTest.test_record_updated":                                  "5807",
	"tests.test_green.CallbackErrorTestCase.test_errors_named_cursor":                                        "30352",
	"tests.test_green.GreenTestCase.test_non_block_after_notification":                                       "5807",
	"tests.test_lobject.LargeObject64Tests.test_seek_tell_truncate_greater_than_2gb":                         "35902",
	"tests.test_lobject.LargeObjectTests.test_close":                                                         "35902",
	"tests.test_lobject.LargeObjectTests.test_close_after_commit":                                            "243",
	"tests.test_lobject.LargeObjectTests.test_close_connection_gone":                                         "243",
	"tests.test_lobject.LargeObjectTests.test_close_twice":                                                   "243",
	"tests.test_lobject.LargeObjectTests.test_create":                                                        "243",
	"tests.test_lobject.LargeObjectTests.test_create_with_existing_oid":                                      "243",
	"tests.test_lobject.LargeObjectTests.test_create_with_oid":                                               "243",
	"tests.test_lobject.LargeObjectTests.test_export":                                                        "243",
	"tests.test_lobject.LargeObjectTests.test_export_after_close":                                            "243",
	"tests.test_lobject.LargeObjectTests.test_export_after_commit":                                           "243",
	"tests.test_lobject.LargeObjectTests.test_factory":                                                       "243",
	"tests.test_lobject.LargeObjectTests.test_import":                                                        "243",
	"tests.test_lobject.LargeObjectTests.test_mode_defaults":                                                 "243",
	"tests.test_lobject.LargeObjectTests.test_open_existing":                                                 "243",
	"tests.test_lobject.LargeObjectTests.test_open_for_write":                                                "243",
	"tests.test_lobject.LargeObjectTests.test_open_mode_n":                                                   "243",
	"tests.test_lobject.LargeObjectTests.test_open_non_existent":                                             "243",
	"tests.test_lobject.LargeObjectTests.test_read":                                                          "243",
	"tests.test_lobject.LargeObjectTests.test_read_after_close":                                              "243",
	"tests.test_lobject.LargeObjectTests.test_read_after_commit":                                             "243",
	"tests.test_lobject.LargeObjectTests.test_read_binary":                                                   "243",
	"tests.test_lobject.LargeObjectTests.test_read_large":                                                    "243",
	"tests.test_lobject.LargeObjectTests.test_read_text":                                                     "243",
	"tests.test_lobject.LargeObjectTests.test_seek_after_close":                                              "243",
	"tests.test_lobject.LargeObjectTests.test_seek_after_commit":                                             "243",
	"tests.test_lobject.LargeObjectTests.test_seek_tell":                                                     "243",
	"tests.test_lobject.LargeObjectTests.test_tell_after_close":                                              "243",
	"tests.test_lobject.LargeObjectTests.test_tell_after_commit":                                             "243",
	"tests.test_lobject.LargeObjectTests.test_unlink":                                                        "243",
	"tests.test_lobject.LargeObjectTests.test_unlink_after_close":                                            "243",
	"tests.test_lobject.LargeObjectTests.test_unlink_after_commit":                                           "243",
	"tests.test_lobject.LargeObjectTests.test_write":                                                         "243",
	"tests.test_lobject.LargeObjectTests.test_write_after_close":                                             "243",
	"tests.test_lobject.LargeObjectTests.test_write_after_commit":                                            "243",
	"tests.test_lobject.LargeObjectTests.test_write_large":                                                   "243",
	"tests.test_lobject.LargeObjectTruncateTests.test_truncate":                                              "243",
	"tests.test_lobject.LargeObjectTruncateTests.test_truncate_after_close":                                  "243",
	"tests.test_lobject.LargeObjectTruncateTests.test_truncate_after_commit":                                 "243",
	"tests.test_module.ExceptionsTestCase.test_9_3_diagnostics":                                              "5807",
	"tests.test_module.ExceptionsTestCase.test_diagnostics_copy":                                             "5807",
	"tests.test_module.ExceptionsTestCase.test_diagnostics_from_commit":                                      "5807",
	"tests.test_notify.NotifiesTests.test_many_notifies":                                                     "6130",
	"tests.test_notify.NotifiesTests.test_notifies_received_on_execute":                                      "6130",
	"tests.test_notify.NotifiesTests.test_notifies_received_on_poll":                                         "6130",
	"tests.test_notify.NotifiesTests.test_notify_attributes":                                                 "6130",
	"tests.test_notify.NotifiesTests.test_notify_deque":                                                      "6130",
	"tests.test_notify.NotifiesTests.test_notify_noappend":                                                   "6130",
	"tests.test_notify.NotifiesTests.test_notify_object":                                                     "6130",
	"tests.test_notify.NotifiesTests.test_notify_payload":                                                    "6130",
	"tests.test_quote.QuotingTestCase.test_koi8":                                                             "35882",
	"tests.test_quote.QuotingTestCase.test_latin1":                                                           "35882",
	"tests.test_sql.SqlFormatTests.test_copy":                                                                "5807",
	"tests.test_transaction.DeadlockSerializationTests.test_deadlock":                                        "6583",
	"tests.test_transaction.TransactionTests.test_commit":                                                    "5807",
	"tests.test_transaction.TransactionTests.test_failed_commit":                                             "5807",
	"tests.test_transaction.TransactionTests.test_rollback":                                                  "5807",
	"tests.test_types_basic.TypesBasicTests.testArray":                                                       "32552",
	"tests.test_types_basic.TypesBasicTests.testArrayOfNulls":                                                "32552",
	"tests.test_types_basic.TypesBasicTests.testEmptyArray":                                                  "23299",
	"tests.test_types_basic.TypesBasicTests.testEmptyArrayRegression":                                        "36179",
	"tests.test_types_basic.TypesBasicTests.testGenericArrayNull":                                            "25123",
	"tests.test_types_basic.TypesBasicTests.testNestedArrays":                                                "32552",
	"tests.test_types_basic.TypesBasicTests.testNestedEmptyArray":                                            "32552",
	"tests.test_types_basic.TypesBasicTests.testNetworkArray":                                                "18846",
	"tests.test_types_extras.AdaptTypeTestCase.test_cast_composite":                                          "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_cast_nested":                                             "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_composite_array":                                         "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_composite_namespace":                                     "26443",
	"tests.test_types_extras.AdaptTypeTestCase.test_empty_string":                                            "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_from_tables":                                             "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_non_dbapi_connection":                                    "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_register_globally":                                       "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_register_on_connection":                                  "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_register_on_cursor":                                      "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_subclass":                                                "27793",
	"tests.test_types_extras.AdaptTypeTestCase.test_wrong_schema":                                            "27793",
	"tests.test_types_extras.JsonTestCase.test_default_cast":                                                 "23468",
	"tests.test_types_extras.JsonTestCase.test_register_default":                                             "unknown",
	"tests.test_types_extras.JsonTestCase.test_scs":                                                          "36215",
	"tests.test_types_extras.JsonbTestCase.test_default_cast":                                                "23468",
	"tests.test_types_extras.JsonbTestCase.test_loads":                                                       "23468",
	"tests.test_types_extras.JsonbTestCase.test_null":                                                        "23468",
	"tests.test_types_extras.JsonbTestCase.test_register_default":                                            "23468",
	"tests.test_types_extras.RangeCasterTestCase.test_adapt_date_range":                                      "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_adapt_number_range":                                    "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_adapt_numeric_range":                                   "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_cast_date":                                             "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_cast_empty":                                            "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_cast_inf":                                              "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_cast_null":                                             "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_cast_numbers":                                          "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_cast_timestamp":                                        "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_cast_timestamptz":                                      "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_range_escaping":                                        "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_register_range_adapter":                                "27791",
	"tests.test_types_extras.RangeCasterTestCase.test_schema_range":                                          "26443",
	"tests.test_with.WithCursorTestCase.test_exception_swallow":                                              "30352",
	"tests.test_with.WithCursorTestCase.test_named_with_noop":                                                "30352",
}

var psycopgIgnoreList19_2 = psycopgIgnoreList19_1

var psycopgIgnoreList19_1 = blacklist{
	"tests.test_green.GreenTestCase.test_flush_on_write": "flakey",
}
