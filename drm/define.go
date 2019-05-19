package drm

const DISCOVERY_REQ = 0x01
const DISCOVERY_ACK = 0x02

const ATTR_FIRMEWARE_VERSION = 0x01
const ATTR_MAC = 0x02
const ATTR_IP = 0x03
const ATTR_EXT = 0x04
const ATTR_REG_TIMEOUT = 0x05

const EXT_HTTP = 0x06
const EXT_FTP = 0x07
const EXT_LANG = 0x08
const EXT_MODEL_NAME = 0x09
const EXT_EZ_VER = 0x10
const EXT_HOSTNAME = 0x11
const EXT_CMS_PORT = 0x12
const EXT_CLOUD_DEVICE = 0x13
const EXT_P2P_PROXY = 0x14
const EXT_IOT_DEVICE = 0x15
const EXT_HTTPS_PORT = 0x16
const EXT_CLOUD_VADP = 0x17
const EXT_MODE = 0x0a

type Argument struct {
	IP   string
	Port int
}

type Record struct {
	Mac              string
	IP               string
	Model            string
	HttpPort         string
	HttpsPort        string
	Hostname         string
	Firmware_version string
	CloudVADP        string
}

var WaitChan chan bool = make(chan bool)
var Records = make(map[string]string)
var Args = Argument{}
