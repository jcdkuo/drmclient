package drm

const discoveryREQ = 0x01
const discoveryACK = 0x02

const attrFirmwareVersion = 0x01
const attrMacAddress = 0x02
const attrIPAddress = 0x03
const attrExtension = 0x04
const attrRegTimeout = 0x05

const extHTTP = 0x06
const extFTP = 0x07
const extLang = 0x08
const extModelName = 0x09
const extEzVersion = 0x10
const extHostname = 0x11
const extCmsPort = 0x12
const extCloudDevice = 0x13
const extP2pProxy = 0x14
const extIotDevice = 0x15
const extHTTPSPort = 0x16
const extCloudVADP = 0x17
const extMode = 0x0a

// Record is used for show the scan results
type Record struct {
	MacAddress      string
	IPAddress       string
	ModelName       string
	HTTPPort        string
	HTTPSPort       string
	Hostname        string
	FirmwareVersion string
	CloudVADP       string
}

// Records is used for dropping redundant records
var Records = make(map[string]string)
