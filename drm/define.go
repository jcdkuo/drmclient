package drm

const Discovery_Req = 0x01
const Discovery_Ack = 0x02

const AttType_FirmwareVersion = 0x01
const AttType_Mac = 0x02
const AttType_IP = 0x03
const AttType_Ext = 0x04
const AttType_RegTimeout = 0x05

const Ext_HTTP = 0x06
const Ext_FTP = 0x07
const Ext_Lang = 0x08
const Ext_MODEL_NAME = 0x09
const Ext_EZ_VER = 0x10
const Ext_HOSTNAME = 0x11
const Ext_CMS_PORT = 0x12
const Ext_CLOUD_DEVICE = 0x13
const Ext_P2P_PROXY = 0x14
const Ext_IOT_DEVICE = 0x15
const Ext_HTTPS_PORT = 0x16
const Ext_CLOUD_VADP = 0x17
const Ext_MODE = 0x0a

const DRM_PORT = 10000

var WaitChan chan bool = make(chan bool)
var Records = make(map[string]string)

type Record struct {
	Mac       string
	IP        string
	Model     string
	HttpPort  string
	HttpsPort string
	Hostname  string
	CloudVADP string
}
