package Tsi


const TSI_START= "MSTART\r\n";
const TSI_STOP = "MSTOP\r\n";
const TSI_RECEIVE= "RMMEAS\r\n";

const TSI_SERVER_START  = "tsi_server_start"
const TSI_SERVER_STOP  = "tsi_server_stop"
const TSI_SERVER_EXIT  = "tsi_server_exit"

const TSI_SERVER_RECEIVE_DATA_START  = "tsi_server_receive_data"
const TSI_SERVER_RECEIVE_DATA_STOP  = "tsi_server_stop_data_receive"

const TSI_HOST = "172.16.15.214";
const TSI_PORT = 3602;

const TSI_START_POINT = 500;
const TSI_STOP_POINT = 14;
const TSI_FLAG_TIMES = 3; //多少次出现时，才认为有效

const FILE_TSI_TEST = "data_TSI_test.txt";
const FILE_TSI = "data_TSI.txt";

const TSI_RUN_TYPE_CHECK  = "tsi_run_type_check"
const TSI_RUN_TYPE_TEST  = "tsi_run_type_test"
const TSI_RUN_TYPE_TEST_PRE  = "tsi_run_type_test_pre"

