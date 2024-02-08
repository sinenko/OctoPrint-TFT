package ui_exec

import (
	"os"
	"bufio"
	"os/exec"
	"fmt"
	s "strings"
	"path/filepath"
	"encoding/json"
	"net"
	"log"
)

type Configuration struct {
    Calibrating   []string
    HomeAll   []string
    Lang    string
    WIFI_SSID    string
    WIFI_pass    string
		IPAddr string
		RouteAddr string
		DNSAddr string
		IPAddrEth0 string
		RouteAddrEth0 string
		DNSAddrEth0 string
}

type Variables struct {
    IsPrinting    			bool
    IsUpdating    			bool
    IsAllowLoadUnload   bool
}

const (
	_main_config_path     	= "/interprint/configuration"
	_output_path     				= "/interprint/output"
	_mount_flash_path     	= "/home/pi/.octoprint/uploads/"
	_wpa_suplicant_file     = "/etc/wpa_supplicant/wpa_supplicant.conf"
	_dhcpcd_file     				= "/etc/dhcpcd.conf"
	_main_config_file 			= "conf.json"
)

var (
	output_path = ""
	bash_script = ""
	main_config_path = ""
	main_config_file = ""
	Conf = Configuration{}
	Vars = Variables{}
)

func checkError( e error){
    if e != nil {
				// Logger.Warningf("Config create error")
        panic(e)
    }
}


func RebootUSB() {
	output_path = filepath.Join(_output_path)
	bash_script = filepath.Join( "_hub-ctrl.sh" )
	commands := []string{
		" -h 0 -P 2 -p 0",
		"sleep 1",
		"hub-ctrl -h 0 -P 2 -p 1",
	}
	exe_cmd(commands)
}


func Reboot() {
	output_path = filepath.Join(_output_path)
	bash_script = filepath.Join( "_reboot.sh" )
	commands := []string{
		"reboot",
	}
	exe_cmd(commands)
}


func ApplyDHCPParams(settings []string) (bool){

	if len(settings) < 8 {
		return false
	}

	
	
	if _, err := os.Stat(filepath.Join(_dhcpcd_file)); os.IsNotExist(err) {
		return false
	}
	
	
	file, err := os.Create( filepath.Join(_dhcpcd_file))
	checkError(err)
	defer file.Close()	

	file.WriteString("hostname\n")
	file.WriteString("clientid\n")
	file.WriteString("persistent\n")
	file.WriteString("option rapid_commit\n")
	file.WriteString("option domain_name_servers, domain_name, domain_search, host_name \n")
	file.WriteString("option classless_static_routes\n")
	file.WriteString("option ntp_servers \n")
	file.WriteString("option interface_mtu \n")
	file.WriteString("require dhcp_server_identifier \n")
	file.WriteString("slaac private\n")
	//WIFI
	file.WriteString("interface wlan0\n")
	file.WriteString("metric 300\n")
	if(settings[2] != "none" && settings[2] != "dhcp"){
		// file.WriteString("interface wlan0\n")
		file.WriteString(fmt.Sprintf("static ip_address=%s/24\n", settings[2]))
		if(settings[3] != "none"){
			file.WriteString(fmt.Sprintf("static routers=%s\n", settings[3]))
		}
		if(settings[4] != "none"){
			file.WriteString(fmt.Sprintf("static domain_name_servers=%s %s 8.8.8.8\n", settings[4], settings[3]))
		}
	} else if(settings[2] != "dhcp"){
		if(settings[2] == "none" && Conf.IPAddr != settings[2]){
			settings[2] = Conf.IPAddr
			file.WriteString(fmt.Sprintf("static ip_address=%s/24\n", settings[2]))
		}
		if(settings[3] == "none" && Conf.RouteAddr != settings[3]){
			settings[3] = Conf.RouteAddr
			file.WriteString(fmt.Sprintf("static routers=%s\n", settings[3]))
		}
		if(settings[4] == "none" && Conf.DNSAddr != settings[4]){
			settings[4] = Conf.DNSAddr
			file.WriteString(fmt.Sprintf("static domain_name_servers=%s %s 8.8.8.8\n", settings[4], settings[3]))
		}
	}
	//ETHERNET
	file.WriteString("interface eth0\n")
	file.WriteString("metric 200\n")
	if(settings[5] != "none" && settings[5] != "dhcp"){
		// file.WriteString("interface wlan0\n")
		file.WriteString(fmt.Sprintf("static ip_address=%s/24\n", settings[5]))
		if(settings[6] != "none"){
			file.WriteString(fmt.Sprintf("static routers=%s\n", settings[6]))
		}
		if(settings[7] != "none"){
			file.WriteString(fmt.Sprintf("static domain_name_servers=%s %s 8.8.8.8\n", settings[7], settings[6]))
		}
	} else if(settings[5] != "dhcp"){
		if(settings[5] == "none" && Conf.IPAddrEth0 != settings[5]){
			settings[5] = Conf.IPAddrEth0
			file.WriteString(fmt.Sprintf("static ip_address=%s/24\n", settings[5]))
		}
		if(settings[6] == "none" && Conf.RouteAddrEth0 != settings[6]){
			settings[6] = Conf.RouteAddrEth0
			file.WriteString(fmt.Sprintf("static routers=%s\n", settings[6]))
		}
		if(settings[7] == "none" && Conf.DNSAddrEth0 != settings[7]){
			settings[7] = Conf.DNSAddrEth0
			file.WriteString(fmt.Sprintf("static domain_name_servers=%s %s 8.8.8.8\n", settings[7], settings[6]))
		}
	}
	
	
	output_path = filepath.Join(_output_path)
	bash_script = filepath.Join( "_apply_wifi.sh" )
	
	commands := []string{
		"wpa_cli -i wlan0 reconfigure",
	}
	
	exe_cmd(commands)
	
	Conf.IPAddr = settings[2]
	Conf.RouteAddr = settings[3]
	Conf.DNSAddr = settings[4]
	Conf.IPAddrEth0 = settings[5]
	Conf.RouteAddrEth0 = settings[6]
	Conf.DNSAddrEth0 = settings[7]
	SaveConfig()
	
	return true
}


func ApplyWIFIParams(settings []string) (bool){

	if len(settings) < 2 {
		return false
	}

	if _, err := os.Stat(filepath.Join(_wpa_suplicant_file)); os.IsNotExist(err) {
		return false
	}
	
	
	file, err := os.Create( filepath.Join(_wpa_suplicant_file))
	checkError(err)
	// file, _ := os.Open(filepath.Join(_wpa_suplicant_file))
	defer file.Close()
	
	
	
	file.WriteString("## WPA/WPA2 secured\n")
	file.WriteString("network={\n")
	file.WriteString(fmt.Sprintf("ssid=\"%s\"\n", settings[0]))
	if(settings[1] == "none"){
		file.WriteString("key_mgmt=NONE")
	} else {
		file.WriteString(fmt.Sprintf("psk=\"%s\"\n", settings[1]))
	}
	file.WriteString("}\n")
	file.WriteString("### You should not have to change the lines below #####################\n\n")
	file.WriteString("ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev\n")
	
	output_path = filepath.Join(_output_path)
	bash_script = filepath.Join( "_apply_wifi.sh" )
	
	commands := []string{
		"wpa_cli -i wlan0 reconfigure",
	}
	
	exe_cmd(commands)
	
	Conf.WIFI_SSID = settings[0]
	Conf.WIFI_pass = settings[1]
	SaveConfig()
	
	return true
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filepath.Join(filename)); os.IsNotExist(err) {
		return false
	}
	return true
}

func LoadWIFIConfig() ([]string){
	numUSB := 1
	configFile := fmt.Sprintf("%s%s%d%s", _mount_flash_path, "usb", numUSB, "/wifi.inf")
	for (!fileExists(configFile)) {
		numUSB += 1
		configFile = fmt.Sprintf("%s%s%d%s", _mount_flash_path, "usb", numUSB, "/wifi.inf")
		if numUSB > 4 {
			return []string{ "none", "none" }
		}
	}
	
	if _, err := os.Stat(filepath.Join(configFile)); os.IsNotExist(err) {
		return []string{ "none", "none" }
	}
	
	file, _ := os.Open(filepath.Join(configFile))
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		return []string{ "none", "none" }
	}
	
	if len(lines) < 1{
		return []string{ "none", "none" }
	}
	
	if len(lines) < 2{
		if(len(lines[0]) > 128 ){
			return []string{ "none", "none" }
		}
		return []string{ lines[0], "none" }
	}
	
	if(len(lines[0]) > 128 || len(lines[1]) > 128){
		return []string{ "none", "none" }
	}
	
	 return []string{ lines[0], lines[1] }
}

func LoadDHCPConfig() ([]string){
	numUSB := 1
	configFile := fmt.Sprintf("%s%s%d%s", _mount_flash_path, "usb", numUSB, "/wifi.inf")
	for (!fileExists(configFile)) {
		numUSB += 1
		configFile = fmt.Sprintf("%s%s%d%s", _mount_flash_path, "usb", numUSB, "/wifi.inf")
		if numUSB > 4 {
			return []string{ "none", "none", "none", "none", "none", "none", "none", "none" }
		}
	}
	
	if _, err := os.Stat(filepath.Join(configFile)); os.IsNotExist(err) {
		return []string{ "none", "none" }
	}
	
	file, _ := os.Open(filepath.Join(configFile))
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		return []string{ "none", "none", "none", "none", "none", "none", "none", "none" }
	}
	
	if len(lines) < 3{
		return []string{ "none", "none", "none", "none", "none", "none", "none", "none" }
	}
	
	_ip_ := "none"
	_ipgateway_ := "none"
	_ipdns_ := "none"
	_ip_eth0_ := "none"
	_ipgateway_eth0_ := "none"
	_ipdns_eth0_ := "none"
	
	
	////WIFI
	
	if len(lines) >= 3{
		if(lines[2] == "dhcp"){
			_ip_ = "dhcp"
		} else {
			_ipaddr := net.ParseIP(lines[2])
			if _ipaddr != nil {
				if(_ipaddr.To4() != nil && !_ipaddr.IsLoopback()) {
					_ip_ = _ipaddr.String()
				}
			}
		}
	}
	if len(lines) < 4{
		return []string{ "none", "none", _ip_, "none", "none", "none", "none", "none" }
	}
	
	
	if len(lines) >= 4{
		_ipgateway := net.ParseIP(lines[3])
		
		if _ipgateway != nil {
			if(_ipgateway.To4() != nil && !_ipgateway.IsLoopback()) {
				_ipgateway_ = _ipgateway.String()
			}
		}
	}
	if len(lines) < 5{
		return []string{ "none", "none", _ip_, _ipgateway_, "none", "none", "none", "none" }
	}
	
	
	if len(lines) >= 5{
		_ipdns := net.ParseIP(lines[4])
		if _ipdns != nil {
			if(_ipdns.To4() != nil && !_ipdns.IsLoopback()) {
				_ipdns_ = _ipdns.String()
			}
		}
	}
	if len(lines) < 6{
		return []string{ "none", "none", _ip_, _ipgateway_, _ipdns_, "none", "none", "none" }
	}
	
	/////WIFI END

	/////ETHERNET
	
		if len(lines) >= 6{
			if(lines[5] == "dhcp"){
				_ip_eth0_ = "dhcp"
			} else {
				_ipaddr_eth0 := net.ParseIP(lines[5])
				if _ipaddr_eth0 != nil {
					if(_ipaddr_eth0.To4() != nil && !_ipaddr_eth0.IsLoopback()) {
						_ip_eth0_ = _ipaddr_eth0.String()
					}
				}
			}
		}
		if len(lines) < 7{	
			return []string{ "none", "none", _ip_, _ipgateway_, _ipdns_, _ip_eth0_, "none", "none" }
		}
		
		if len(lines) >= 7{	
			_ipgateway_eth0 := net.ParseIP(lines[6])
			
			if _ipgateway_eth0 != nil {
				if(_ipgateway_eth0.To4() != nil && !_ipgateway_eth0.IsLoopback()) {
					_ipgateway_eth0_ = _ipgateway_eth0.String()
				}
			}
		}
		if len(lines) < 8{	
			return []string{ "none", "none", _ip_, _ipgateway_, _ipdns_, _ip_eth0_, _ipgateway_eth0_, "none" }
		}
		
		if len(lines) >= 8{	
			_ipdns_eth0 := net.ParseIP(lines[7])
			if _ipdns_eth0 != nil {
				if(_ipdns_eth0.To4() != nil && !_ipdns_eth0.IsLoopback()) {
					_ipdns_eth0_ = _ipdns_eth0.String()
				}
			}
		}
		if len(lines) < 9{	
			return []string{ "none", "none", _ip_, _ipgateway_, _ipdns_, _ip_eth0_, _ipgateway_eth0_, _ipdns_eth0_ }
		}
	
	/////ETHERNET END
	
	return lines
}

func MountFlash() {
	output_path = filepath.Join(_output_path)
	bash_script = filepath.Join( "_mount_flash.sh" )
		
	commands := []string{
		fmt.Sprintf("%s%s%s%s%s","oldfolder=`sudo find " ,_mount_flash_path, " -name usb* | grep -oE '" , _mount_flash_path, "usb[0-9]{1}'`"),
		"oldfolders=$(echo $oldfolder | tr \"\\n\" \"\\n\")",
		"for olddir in $oldfolders",
		"do",
		"sudo umount $olddir",
		"sudo rmdir $olddir",
		"done",
		"flashcard=`sudo blkid | grep \\/sd | grep -oE '/dev/sd[a-z]{1,3}[0-9]{1,3}:[^\\\"]+\\\"[^\\\"]+\\\"[^\\\"]+\"[^\\\"]+\\\"'`",
		"flashcards=$(echo \"${flashcard/\": \"/\":\"}\")",
		"flashcards=$(echo \"${flashcards/\"\\\" T\"/\"\\\"T\"}\")",
		"flashcards=$(echo $flashcards | tr \" \" \"\\n\")",
		"i=$((0))",
		"for usb in $flashcards",
		"do",
		"i=$((i + 1))",
		"usbp=`echo $usb | grep \\/sd | grep -oE '/dev/sd[a-z]{1,3}[0-9]{1,3}'`",
		"typef=`echo $usb | grep \\/sd | grep -oE 'TYPE=\\\"[a-z]{1,5}[0-9]{0,3}\\\"'`",
		fmt.Sprintf("%s%s%s", "sudo mkdir ", _mount_flash_path, "usb$i"),
		fmt.Sprintf("%s%s%s", "sudo chmod 775 ", _mount_flash_path, "usb$i"),
		"if [ \"$typef\" = \"TYPE=\\\"ntfs\\\"\" ]; then",
		fmt.Sprintf("%s%s%s", "sudo mount -t ntfs -o uid=pi,gid=pi $usbp ", _mount_flash_path, "usb$i"),
		"else",
		fmt.Sprintf("%s%s%s", "sudo mount -t vfat -o uid=pi,gid=pi $usbp ", _mount_flash_path, "usb$i"),
		"fi",
		// fmt.Sprintf("%s%s%s", "isconfig=`sudo find ", _mount_flash_path, "usb$i -name config.txt | grep -oE '/config.txt'`"),
		// fmt.Sprintf("%s%s%s", "isgcode=`sudo find ", _mount_flash_path, "usb$i -name *.gcode`"),
		// "if [ \"$isconfig\"==\"/config.txt\" ] && [ -z \"$isgcode\" ]; then ",
		// fmt.Sprintf("%s%s%s", "sudo umount ", _mount_flash_path, "usb$i"),
		// fmt.Sprintf("%s%s%s", "sudo rm -r -f ", _mount_flash_path, "usb$i"),
		// "fi",
		"done",
		
	}
	exe_cmd(commands)
}

func UpdateSoftware(version string) int {
	output_path = filepath.Join(_output_path)
	bash_script = filepath.Join( "_check_update_software.sh" )
	os.RemoveAll(output_path)
			
	commands := []string{
		"rm update.deb",
		fmt.Sprintf("%s%s%s","wget -q \"http://interprint.kz/updates/?model=h2&get&v=", version, "\" -O update.deb -o /dev/null"),
		"sizefile=$(stat update.deb -c %s)",
		"if [ \"$sizefile\" -eq 0 ]; then",
		"rm upd_status",
		"echo \"NOUPDATES\" > upd_status",
		"exit",
		"else",
		"wget -q \"http://interprint.kz/updates/?model=h2&get_md5\" -O sum.md5 -o /dev/null",
		"checkstatus=$(md5sum -c sum.md5 | grep \"OK\" -oE)",
		"if [ \"$checkstatus\" = \"OK\" ]; then",
		"rm upd_status",
		"echo \"UPD OK\" > upd_status",
		"exit",
		"else",
		"rm upd_status",
		"echo \"NOEQ\" > upd_status",
		"exit",
		"fi",
		"fi",
	}
	exe_cmd(commands)
	
	file, _ := os.Open(filepath.Join(output_path, "upd_status"))
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		return 0
	}
	
	if len(lines) < 1{
		return 0
	}
	
	if(lines[0] == "NOUPDATES"){
		return -1
	}
	
	if(lines[0] == "NOEQ"){
		return -2
	}
	
	if(lines[0] == "UPD OK"){
		Vars.IsUpdating = true
		output_path = filepath.Join(_output_path)
		bash_script = filepath.Join( "_update_software.sh" )
		commandsUpd := []string{
			"dpkg -i update.deb",
			fmt.Sprintf("%s%s%s","rm ", output_path, "/*"),
		}
		exe_cmd_at(commandsUpd)
		return 1
	}
	
	return 0
}



func exe_cmd(cmds []string) {
    os.Remove(filepath.Join(output_path, bash_script))
    err := os.MkdirAll( output_path, os.ModePerm|os.ModeDir )
    checkError(err)
    file, err := os.Create( filepath.Join(output_path, bash_script))
    checkError(err)
    defer file.Close()
    file.WriteString("#!/bin/sh\n")
    file.WriteString( s.Join(cmds, "\n"))
    err = os.Chdir(output_path)
    checkError(err)
    // out, err := 
		exec.Command("bash", bash_script).Output()
    // checkError(err)
    // fmt.Println(string(out))
}

func exe_cmd_at(cmds []string) {
    os.Remove(filepath.Join(output_path, bash_script))
    err := os.MkdirAll( output_path, os.ModePerm|os.ModeDir )
    checkError(err)
    file, err := os.Create( filepath.Join(output_path, bash_script))
    checkError(err)
    defer file.Close()
    file.WriteString("#!/bin/sh\n")
    file.WriteString( s.Join(cmds, "\n"))
    err = os.Chdir(output_path)
    checkError(err)
		
		commands := []string{
			fmt.Sprintf("%s%s%s%s%s","echo \"cd ", output_path, "; sh ", filepath.Join(output_path, bash_script), "\"  | at -m now +1 minute"),
		}
		
		output_path = filepath.Join(_output_path)
		bash_script = filepath.Join( "_create_at_task.sh" )
		
		exe_cmd(commands)
}

func exe_cmd_nohup(cmds []string) {
		os.Remove(filepath.Join(output_path, bash_script))
    err := os.MkdirAll( output_path, os.ModePerm|os.ModeDir )
    checkError(err)
    file, err := os.Create( filepath.Join(output_path, bash_script))
    checkError(err)
    defer file.Close()
    file.WriteString("#!/bin/sh\n")
    file.WriteString( s.Join(cmds, "\n"))
    err = os.Chdir(output_path)
    checkError(err)
		err = os.Chmod(bash_script, 0755)
		checkError(err)
    err = os.Chdir(output_path)
    checkError(err)
		cmdres := exec.Command("/usr/bin/nohup", "sh", filepath.Join(output_path, bash_script), "&")
		
		
		err2 := cmdres.Run()
		if err2 != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err2)
		}
}

func SaveDefaultConfig(){
	main_config_path = filepath.Join(_main_config_path)
	main_config_file = filepath.Join(_main_config_file)
	os.RemoveAll(main_config_path)
  err := os.MkdirAll( main_config_path, os.ModePerm|os.ModeDir )
  checkError(err)
	file, err := os.Create( filepath.Join(main_config_path, main_config_file))
	checkError(err)
	defer file.Close()
	encoder := json.NewEncoder(file)
	configuration := Configuration{
    HomeAll: []string{"G28"},
    Calibrating: []string{"G28", "G29", "M500"},
    Lang: "en",
    WIFI_SSID: "HostPointName",
    WIFI_pass: "12345678",
		IPAddr: "none",
		RouteAddr: "none",
		DNSAddr: "none",
		IPAddrEth0: "none",
		RouteAddrEth0: "none",
		DNSAddrEth0: "none",
	}
		
	err = encoder.Encode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func SaveConfig(){
	main_config_path = filepath.Join(_main_config_path)
	main_config_file = filepath.Join(_main_config_file)
	os.RemoveAll(main_config_path)
  err := os.MkdirAll( main_config_path, os.ModePerm|os.ModeDir )
  checkError(err)
	file, err := os.Create( filepath.Join(main_config_path, main_config_file))
	checkError(err)
	defer file.Close()
	encoder := json.NewEncoder(file)
	configuration := Conf
		
	err = encoder.Encode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func LoadConf(){
	main_config_path = filepath.Join(_main_config_path)
	main_config_file = filepath.Join(_main_config_file)
	if _, err := os.Stat(filepath.Join(main_config_path, main_config_file)); os.IsNotExist(err) {
		SaveDefaultConfig()
	}
	
	file, _ := os.Open(filepath.Join(main_config_path, main_config_file))
	defer file.Close()
	decoder := json.NewDecoder(file)
	// Conf := Configuration{}
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func GetIpAddrByIface(intname string) (string) {
	i, _ := net.InterfaceByName(intname) //here your interface
	addrs, err := i.Addrs()
	// handle err
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			if v.IP.To4() != nil && !v.IP.IsLoopback() {
				ip = v.IP
				return ip.String()
			}
		case *net.IPAddr:
			if v.IP.To4() != nil && !v.IP.IsLoopback() {
				ip = v.IP
				return ip.String()
			}
		}
	}
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
	}
	return ""
}
