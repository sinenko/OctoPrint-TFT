package ui

import (
	l "github.com/mcuadros/OctoPrint-TFT/ui_lang"
	"github.com/gotk3/gotk3/gtk"
	"fmt"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
	"strings"
	"unicode/utf8"
)


var snetworkPanelInstance *snetworkPanel

type snetworkPanel struct {
	CommonPanel
	ssid, pass, ipaddr, iproute, ipdns, ipaddrEth0, iprouteEth0, ipdnsEth0, ipeth0, ipwlan     *LabelWithImage
	labelNEW, ssidNEW, passNEW, ipaddrNEW, iprouteNEW, ipdnsNEW, ipaddrEth0NEW, iprouteEth0NEW, ipdnsEth0NEW  *LabelWithImage
	wifi_settings  []string
	dhcp_settings  []string
}

func SNetworkPanel(ui *UI, parent Panel) Panel {
	if snetworkPanelInstance == nil {
		m := &snetworkPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		snetworkPanelInstance = m
		
	}
	snetworkPanelInstance.updateInfoBox()
	return snetworkPanelInstance
}


func (m *snetworkPanel) initialize() {
	defer m.Initialize()
	
	m.Grid().Attach(m.createInfoBox(), 0, 0, 3, 2)
	m.Grid().Attach(m.createNewWIFIInfoBox(), 3, 0, 3, 2)
	m.Grid().Attach(MustButtonImage(l.Translate("Save"), "save.svg", m.applyConfigWiFi), 0, 3, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Load..."), "usb.svg", m.loadConfigWiFi), 1, 3, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Refresh"), "refresh.svg", m.updateInfoBox), 2, 3, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Reboot"), "reboot.svg", m.rebootSystem), 3, 3, 2, 1)
}

func (m *snetworkPanel) loadConfigWiFi() {
	exe.MountFlash()
	m.wifi_settings = exe.LoadWIFIConfig()
	m.dhcp_settings = exe.LoadDHCPConfig()
	
	if (m.dhcp_settings[2] != "none" || m.wifi_settings[0] != "none"){
		m.labelNEW.Label.SetLabel(l.Translate("<b>LOADED PARAMS:</b>"))
	}
	
	if m.wifi_settings[0] == "none" {
		Logger.Error(l.Translate("File not found or wrong"))
		m.ssidNEW.Label.SetLabel("")
		m.passNEW.Label.SetLabel("")
	} else {
		if(m.wifi_settings[1] == "none"){
			Logger.Warningf(l.Translate("WIFI no pass"))
		}
		m.ssidNEW.Label.SetLabel(fmt.Sprintf(l.Translate("WIFI SSID: %s"), m.wifi_settings[0]))
		m.passNEW.Label.SetLabel(fmt.Sprintf(l.Translate("WIFI password: %s"), m.wifi_settings[1]))
	}
	
	if m.dhcp_settings[2] == "none" {
		m.ipaddrNEW.Label.SetLabel("")
		m.iprouteNEW.Label.SetLabel("")
		m.ipdnsNEW.Label.SetLabel("")
	} else {
		m.ipaddrNEW.Label.SetLabel(fmt.Sprintf(l.Translate("IP address: %s"), m.dhcp_settings[2]))
		m.iprouteNEW.Label.SetLabel(fmt.Sprintf(l.Translate("Gateway address: %s"), m.dhcp_settings[3]))
		m.ipdnsNEW.Label.SetLabel(fmt.Sprintf(l.Translate("DNS address: %s"), m.dhcp_settings[4]))
		m.ipaddrEth0NEW.Label.SetLabel(fmt.Sprintf(l.Translate("IP address: %s"), m.dhcp_settings[5]))
		m.iprouteEth0NEW.Label.SetLabel(fmt.Sprintf(l.Translate("Gateway address: %s"), m.dhcp_settings[6]))
		m.ipdnsEth0NEW.Label.SetLabel(fmt.Sprintf(l.Translate("DNS address: %s"), m.dhcp_settings[7]))
	}
}

func (m *snetworkPanel) createInfoBox() *gtk.Box {
	m.ssid = MustLabelWithImage("", "")
	m.pass = MustLabelWithImage("", "")
	m.ipaddr = MustLabelWithImage("", "")
	m.iproute = MustLabelWithImage("", "")
	m.ipdns = MustLabelWithImage("", "")
	m.ipaddrEth0 = MustLabelWithImage("", "")
	m.iprouteEth0 = MustLabelWithImage("", "")
	m.ipdnsEth0 = MustLabelWithImage("", "")
	m.ipeth0 = MustLabelWithImage("", "")
	m.ipwlan = MustLabelWithImage("", "")

	m.ssid.Label.SetLabel(fmt.Sprintf(l.Translate("WIFI SSID: %s"), exe.Conf.WIFI_SSID))
	m.pass.Label.SetLabel(fmt.Sprintf(l.Translate("WIFI password: %s"), replacePass(exe.Conf.WIFI_pass)))
	m.ipaddr.Label.SetLabel(fmt.Sprintf(l.Translate("IP address: %s"), exe.Conf.IPAddr))
	m.iproute.Label.SetLabel(fmt.Sprintf(l.Translate("Gateway address: %s"), exe.Conf.RouteAddr))
	m.ipdns.Label.SetLabel(fmt.Sprintf(l.Translate("DNS address: %s"), exe.Conf.DNSAddr))
	m.ipaddrEth0.Label.SetLabel(fmt.Sprintf(l.Translate("LAN IP address: %s"), exe.Conf.IPAddrEth0))
	m.iprouteEth0.Label.SetLabel(fmt.Sprintf(l.Translate("LAN Gateway address: %s"), exe.Conf.RouteAddrEth0))
	m.ipdnsEth0.Label.SetLabel(fmt.Sprintf(l.Translate("LAN DNS address: %s"), exe.Conf.DNSAddrEth0))
	m.ipeth0.Label.SetLabel(fmt.Sprintf(l.Translate("<b>IP LAN: %s</b>"), exe.GetIpAddrByIface("eth0")))
	m.ipwlan.Label.SetLabel(fmt.Sprintf(l.Translate("<b>IP wifi: %s</b>"), exe.GetIpAddrByIface("wlan0")))

	info := MustBox(gtk.ORIENTATION_VERTICAL, 0)
	info.SetHExpand(true)
	info.SetVExpand(true)
	
	info.Add(MustLabelWithImage("", l.Translate("<b>CURRENT PARAMS:</b>")))
	info.Add(m.ssid)
	info.Add(m.pass)
	info.Add(m.ipaddr)
	info.Add(m.iproute)
	info.Add(m.ipdns)
	info.Add(m.ipaddrEth0)
	info.Add(m.iprouteEth0)
	info.Add(m.ipdnsEth0)
	info.Add(m.ipeth0)
	info.Add(m.ipwlan)
	info.SetMarginStart(10)

	return info
}

func (m *snetworkPanel) updateInfoBox() {
	m.ssid.Label.SetLabel(fmt.Sprintf(l.Translate("WIFI SSID: %s"), exe.Conf.WIFI_SSID))
	m.pass.Label.SetLabel(fmt.Sprintf(l.Translate("WIFI password: %s"), replacePass(exe.Conf.WIFI_pass)))
	m.ipaddr.Label.SetLabel(fmt.Sprintf(l.Translate("IP address: %s"), exe.Conf.IPAddr))
	m.iproute.Label.SetLabel(fmt.Sprintf(l.Translate("Gateway address: %s"), exe.Conf.RouteAddr))
	m.ipdns.Label.SetLabel(fmt.Sprintf(l.Translate("DNS address: %s"), exe.Conf.DNSAddr))
	m.ipaddrEth0.Label.SetLabel(fmt.Sprintf(l.Translate("LAN IP address: %s"), exe.Conf.IPAddrEth0))
	m.iprouteEth0.Label.SetLabel(fmt.Sprintf(l.Translate("LAN Gateway address: %s"), exe.Conf.RouteAddrEth0))
	m.ipdnsEth0.Label.SetLabel(fmt.Sprintf(l.Translate("LAN DNS address: %s"), exe.Conf.DNSAddrEth0))
	m.ipeth0.Label.SetLabel(fmt.Sprintf(l.Translate("<b>IP LAN: %s</b>"), exe.GetIpAddrByIface("eth0")))
	m.ipwlan.Label.SetLabel(fmt.Sprintf(l.Translate("<b>IP wifi: %s</b>"), exe.GetIpAddrByIface("wlan0")))
}

func replacePass(pass string) (string) {
	return strings.Repeat("*", utf8.RuneCountInString(pass))
}

func (m *snetworkPanel) createNewWIFIInfoBox() *gtk.Box {
	m.labelNEW = MustLabelWithImage("", "")
	m.ssidNEW = MustLabelWithImage("", "")
	m.passNEW = MustLabelWithImage("", "")
	m.ipaddrNEW = MustLabelWithImage("", "")
	m.iprouteNEW = MustLabelWithImage("", "")
	m.ipdnsNEW = MustLabelWithImage("", "")
	m.ipaddrEth0NEW = MustLabelWithImage("", "")
	m.iprouteEth0NEW = MustLabelWithImage("", "")
	m.ipdnsEth0NEW = MustLabelWithImage("", "")
	
	info := MustBox(gtk.ORIENTATION_VERTICAL, 0)
	info.SetHExpand(true)
	info.SetVExpand(true)
	info.Add(m.labelNEW)
	info.Add(m.ssidNEW)
	info.Add(m.passNEW)
	info.Add(m.ipaddrNEW)
	info.Add(m.iprouteNEW)
	info.Add(m.ipdnsNEW)
	info.Add(m.ipaddrEth0NEW)
	info.Add(m.iprouteEth0NEW)
	info.Add(m.ipdnsEth0NEW)
	info.SetMarginStart(10)

	return info
}

func (m *snetworkPanel) applyConfigWiFi() {
	if(len(m.wifi_settings) < 2){
		exe.MountFlash()
		m.wifi_settings = exe.LoadWIFIConfig()
		
		if m.wifi_settings[0] == "none" {
			Logger.Error(l.Translate("File not found or wrong"))
			m.labelNEW.Label.SetLabel("")
			m.ssidNEW.Label.SetLabel("")
			m.passNEW.Label.SetLabel("")
		} else {
			if(m.wifi_settings[1] == "none"){
				Logger.Warningf(l.Translate("WIFI no pass"))
			}
		}
	}
	
	if(exe.ApplyWIFIParams(m.wifi_settings) == true){
		Logger.Warningf(l.Translate("New settings saved"))
	} else {
		Logger.Error(l.Translate("Settings not saved"))
	}
	
	
	if(len(m.wifi_settings) < 8){
		exe.MountFlash()
		m.wifi_settings = exe.LoadDHCPConfig()
		
		if m.wifi_settings[2] == "none" {
			//Настройки Ip адреса не найдены
			// Logger.Error(l.Translate("IP address LAN not found"))
		} else {
			if(m.wifi_settings[3] == "none"){
				Logger.Warningf(l.Translate("Gateway IP  not found"))
			}
			if(m.wifi_settings[4] == "none"){
				Logger.Warningf(l.Translate("DNS IP not found"))
			}
		}
		if m.wifi_settings[5] == "none" {
			//Настройки Ip адреса не найдены
			// Logger.Error(l.Translate("IP address LAN settings not found"))
		} else {
			if(m.wifi_settings[6] == "none"){
				Logger.Warningf(l.Translate("Gateway IP LAN not found"))
			}
			if(m.wifi_settings[7] == "none"){
				Logger.Warningf(l.Translate("DNS IP LAN not found"))
			}
		}
	}
	
	if(exe.ApplyDHCPParams(m.wifi_settings) == true){
		Logger.Warningf(l.Translate("IP address settings saved"))
	} else {
		Logger.Error(l.Translate("IP address settings not saved"))
		//IP адрес не сохранен
	}
	
}

func (m *snetworkPanel) rebootSystem() {
	if(exe.Vars.IsPrinting) { 
		Logger.Error(l.Translate("Error: printer is busy"))
		return 
	}
	exe.Reboot()
}


