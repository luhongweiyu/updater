package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var 永无BUG string = `
 * _ooOoo_
 * o8888888o
 * 88" . "88
 * (| -_- |)
 *  O\ = /O
 * ___/'---'\____
 * .   ' \\| |// '.
 * / \\||| : |||// \
 * / _||||| -:- |||||- \
 * | | \\\ - /// | |
 * | \_| ''\---/'' | |
 * \ .-\__ '-' ___/-. /
 * ___'. .' /--.--\ '. . __
 * ."" '< '.___\_<|>_/___.' >'"".
 * | | : '- \'.;'\ _ /';.'/ - ' : | |
 * \ \ '-. \_ __\ /__ _/ .-' / /
 * ======'-.____'-.___\_____/___.-'____.-'======
 * '=---='
 *          .............................................
 *           佛曰：bug泛滥，我已瘫痪！
 *
 *                    _ooOoo_
 *                   o8888888o
 *                   88" . "88
 *                   (| -_- |)
 *                    O\ = /O
 *                ____/'---'\____
 *              .   ' \\| |// '.
 *               / \\||| : |||// \
 *             / _||||| -:- |||||- \
 *               | | \\\ - /// | |
 *             | \_| ''\---/'' | |
 *              \ .-\__ '-' ___/-. /
 *           ___'. .' /--.--\ '. . __
 *        ."" '< '.___\_<|>_/___.' >'"".
 *       | | : '- \'.;'\ _ /';.'/ - ' : | |
 *         \ \ '-. \_ __\ /__ _/ .-' / /
 * ======'-.____'-.___\_____/___.-'____.-'======
 *                    '=---='
 *
 * .............................................
 *          佛祖保佑             永无BUG
 *
 *                  ___====-_  _-====___
 *            _--^^^#####//      \\#####^^^--_
 *         _-^##########// (    ) \\##########^-_
 *        -############//  |\^^/|  \\############-
 *      _/############//   (@::@)   \\############\_
 *     /#############((     \\//     ))#############\
 *    -###############\\    (oo)    //###############-
 *   -#################\\  / VV \  //#################-
 *  -###################\\/      \//###################-
 * _#/|##########/\######(   /\   )######/\##########|\#_
 * |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
 * '  |/  V  V  '   V  \#\| |  | |/#/  V   '  V  V  \|  '
 *    '   '  '      '   / | |  | | \   '      '  '   '
 *                     (  | |  | |  )
 *                    __\ | |  | | /__
 *                   (vvv(VVV)(VVV)vvv)                
 *                        神兽保佑
 *                       代码无BUG!
 *
 *
 *
 *                                                    __----~~~~~~~~~~~------___
 *                                   .  .   ~~//====......          __--~ ~~
 *                   -.            \_|//     |||\\  ~~~~~~::::... /~
 *                ___-==_       _-~o~  \/    |||  \\            _/~~-
 *        __---~~~.==~||\=_    -_--~/_-~|-   |\\   \\        _/~
 *    _-~~     .=~    |  \\-_    '-~7  /-   /  ||    \      /
 *  .~       .~       |   \\ -_    /  /-   /   ||      \   /
 * /  ____  /         |     \\ ~-_/  /|- _/   .||       \ /
 * |~~    ~~|--~~~~--_ \     ~==-/   | \~--===~~        .\
 *          '         ~-|      /|    |-~\~~       __--~~
 *                      |-~~-_/ |    |   ~\_   _-~            /\
 *                           /  \     \__   \/~                \__
 *                       _--~ _/ | .-~~____--~-/                  ~~==.
 *                      ((->/~   '.|||' -_|    ~~-/ ,              . _||
 *                                 -_     ~\      ~~---l__i__i__i--~~_/
 *                                 _-~-__   ~)  \--______________--~~
 *                               //.-~~~-~_--~- |-------~~~~~~~~
 *                                      //.-~~~--\
 *                               神兽保佑
 *                              代码无BUG!
 *
`
var 全局_密码 string
var 全局_端口 string
var 全局_地址 string
var 全局_项目 string

func 计算md5(filePath string) string {
	// filePath := "example.txt" // 文件路径
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Value := fmt.Sprintf("%x", hashInBytes)
	// fmt.Println("MD5:", md5Value)
	return md5Value
}
func 获取服务器文件() map[string]string {
	return json转map(`{"file.zip":"d286c323eb6bdfc03fa6c55d9380609a","go.mod":"502fc277dc5f44330e384230fe0f3051","go.sum":"7ae2d46e13b63380dc89765ae4e52784","lua5.4\\src\\Makefile":"dd81f51713f238be66fc7fc5ed762310","lua5.4\\src\\lapi.c":"ea7c6042378f4a1ecbd9b49a84a93090","lua5.4\\src\\lapi.h":"e41c82147fa08cabdfc934b33594759d","lua5.4\\src\\lauxlib.c":"00f381c1e81eb7550a40b3404cea9f8f","lua5.4\\src\\lauxlib.h":"4b4efd6105f3949ca042c476aa949e30","lua5.4\\src\\lbaselib.c":"dd028fd8e9ffe47dd91b1dc025fca653","lua5.4\\src\\lcode.c":"79e093dee98a8ba1deb3bf7efd3d3c9d","lua5.4\\src\\lcode.h":"871898f08f1a2948fba0c315de9b7b95","lua5.4\\src\\lcorolib.c":"82c60028d4ff74f59efa9bc2c8c60bc7","lua5.4\\src\\lctype.c":"73e6e7a1d62ea4776bc78033b14598a2","lua5.4\\src\\lctype.h":"4f5a7f445c2cb239ab324f8622e8894b","lua5.4\\src\\ldblib.c":"4b047e7ae84dfbf063b8e45282db5678","lua5.4\\src\\ldebug.c":"544575fd514f2c40edae1376c1960c0d","lua5.4\\src\\ldebug.h":"7accf269cd02aaa2038ab45e691c446d","lua5.4\\src\\ldo.c":"d51991e277724b3792bff5d6d4ef829d","lua5.4\\src\\ldo.h":"971751e7ecfb5b2ea91a0fbfc856ff07","lua5.4\\src\\ldump.c":"4c0e44138608254375f676c79a869829","lua5.4\\src\\lfunc.c":"f829a6c19aaf626b9d0db81026a4bb77","lua5.4\\src\\lfunc.h":"f219c9719d615926a8a507ce53332e2b","lua5.4\\src\\lgc.c":"042c0b9eec666353437279db02924b2d","lua5.4\\src\\lgc.h":"b92dda7c779b1e93346004b636f8c9d8","lua5.4\\src\\linit.c":"db04fcd55233e996bd594b6d41352af7","lua5.4\\src\\liolib.c":"85df713df8d709b6bd8c250da3dd98d8","lua5.4\\src\\ljumptab.h":"559c809d2c65e02b15d2cbe1814f760c","lua5.4\\src\\llex.c":"99433de61bb10e3ffb9ba03cce9b5274","lua5.4\\src\\llex.h":"da06f8f4d5cae5a2a65fa68f35a93cca","lua5.4\\src\\llimits.h":"10b960e5689db782cca6de2f235a6df5","lua5.4\\src\\lmathlib.c":"a5ffa48d2f18f6758adfa91b54e41edb","lua5.4\\src\\lmem.c":"3f1771d1aab29052bc768c435909c091","lua5.4\\src\\lmem.h":"313e02a40d047c422219b8cdf3b31411","lua5.4\\src\\loadlib.c":"9b49b9ce5c570e97a6bf56a21df1726f","lua5.4\\src\\lobject.c":"41a23be18b21a448c7c1c1666cba4e73","lua5.4\\src\\lobject.h":"db0c9e3dbbde604e34b0f38c15563696","lua5.4\\src\\lopcodes.c":"2bdb450e467e3d232e25bb84368b1433","lua5.4\\src\\lopcodes.h":"f08793852bd567c7700b704d5ee05d04","lua5.4\\src\\lopnames.h":"4796bd9cda2cf3b8b90fa429ff3f2d4d","lua5.4\\src\\loslib.c":"1fe9110bf7ba845bc85069cd3b1e404f","lua5.4\\src\\lparser.c":"a73744ee32952d243fcf9ddb622d4af7","lua5.4\\src\\lparser.h":"321483322c302043ba75f9b04ade80d5","lua5.4\\src\\lprefix.h":"c9b706105782437eef7616320c468ef1","lua5.4\\src\\lstate.c":"f204fa736bb619485b85e373dbf9642c","lua5.4\\src\\lstate.h":"94147a47f2a39c48e9eb9a1d486a8863","lua5.4\\src\\lstring.c":"0206cabf0540a673ac9ee291abb04c22","lua5.4\\src\\lstring.h":"5111c938a6aab9498fbd60daf46fbc78","lua5.4\\src\\lstrlib.c":"1ac9811628628344f562ce54976ff8c3","lua5.4\\src\\ltable.c":"e292ce661ca0cb0e10fc06965b97cfed","lua5.4\\src\\ltable.h":"c623c34e385fb441fc99c594fad29953","lua5.4\\src\\ltablib.c":"aafc2acc559334c08b5d0e40fb69a834","lua5.4\\src\\ltm.c":"dee47309e78b4b9ae1760c3ba5ad41b3","lua5.4\\src\\ltm.h":"550ecc448a2b932870d8240a6cb7e444","lua5.4\\src\\lua.c":"99fc7c604651d3b474c52eedbbc081a2","lua5.4\\src\\lua.h":"36ad4a649d5d5e1c03aa39e4af217cbe","lua5.4\\src\\lua.hpp":"9fac7d7bf2a86babec4b57d3904495c8","lua5.4\\src\\luac.c":"b7fe867c6c009ac032a814cd7132ceae","lua5.4\\src\\luaconf.h":"5005168517fc6cb63dfcdee40f181638","lua5.4\\src\\lualib.h":"d1f65ceeb5b9caee5b89bac768a3666a","lua5.4\\src\\lundump.c":"800676223c6a9016e7a47e6fc2fa7b26","lua5.4\\src\\lundump.h":"b099819326432a971d0768cc7d203c82","lua5.4\\src\\lutf8lib.c":"06fa418a6744d488d45201f84eac7065","lua5.4\\src\\lvm.c":"6d95c4f6f6c528dfb8adc9885ad1c292","lua5.4\\src\\lvm.h":"e360c8a71d6ab43e5777604ead88c858","lua5.4\\src\\lzio.c":"480ddbad8144af1118a8d2d376785192","lua5.4\\src\\lzio.h":"00622612a9c5cd1c5e7b63f3fab5afd8","main.go":"f48ef2a8a221370f70733144e197a703","static\\static2\\1.txt":"201cd22db3c51ce685cf6dcc12a028b2","static\\小狗.png":"5b51073028182f023c39b8a70706a57c","templates\\index.html":"9c2e23863a118c888cfb4935d2baf1b8","下载.go":"a0a59f85fdb94dd81387018d92b5bdfe","记录.txt":"1709b9903a1bf03596336c2ac9073d6e"}`)
}
func 获取本地文件(路径 string) map[string]string {
	dir := 路径 // 目标目录./static  "."
	filejson := make(map[string]string)
	files, err := GetAllFiles(dir)
	if err != nil {
		// log.Fatal(err)
		return map[string]string{}
	}
	// 打印所有文件的路径
	for _, file := range files {
		filejson[file] = 计算md5(file)
	}
	return filejson
	// jsonString, _ := (json.Marshal(filejson))
	// fmt.Println(string(jsonString))
	// fmt.Println(string(jsonString), _ := json.Marshal(filejson))
}
func json转map(files string) map[string]string {
	data := make(map[string]string)
	json.Unmarshal([]byte(files), &data)
	return data
}
func map转json(files map[string]string) string {
	jsonString, _ := (json.Marshal(files))
	return (string(jsonString))
}

// func main() {
// 	服务器文件 := 获取服务器文件()
// 	for file, 本地md5 := range 获取本地文件() {
// 		if 服务器文件[file] == 本地md5 {
// 			fmt.Println("相同的", file, 本地md5)
// 		} else {
// 			fmt.Println("不相同", file, 本地md5)
// 		}
// 	}
// }

// GetAllFiles 获取指定目录下所有文件的路径
func GetAllFiles(dir string) ([]string, error) {
	// var files []string
	// err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// regexp.MustCompile()
	// 	if !info.IsDir() {
	// 		// path = strings.Replace(path, dir, "", 1)
	// 		fmt.Println(path)
	// 		files = append(files, path)
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return files, nil

	var files []string
	err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// regexp.MustCompile()
		if !info.IsDir() {
			// path = strings.Replace(path, dir, "", 1)
			// fmt.Println(path)
			// files = append(files, path)
			files = append(files, strings.Replace(path, "\\", "/", -1))
			// fmt.Println(files)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil

}

func httpPost(url2 string, data url.Values) string {

	resp, err := http.PostForm(url2, data)
	if err != nil {
		// handle error
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}

	// fmt.Println(string(body))
	return string(body)

	// // url := "http://127.0.0.1:8080/getFileList"
	// // data := []byte(`{"name": "John", "age": 30}`)
	// data, _ := json.Marshal(datamap)
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	// if err != nil {
	// 	// 处理错误
	// 	return ""
	// }
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	// 处理错误
	// 	return ""
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	// 处理错误
	// 	return ""
	// }
	// // fmt.Println("响应体:", string(body))

	// return string(body)

}
func 初始化设置() {
	// context, _ := ioutil.ReadFile("下载设置")
	// fmt.Println(string(context))
	// L_全局_密码 := regexp.MustCompile(`password\[(\w+)\]`).FindStringSubmatch(string(context))
	// L_全局_端口 := regexp.MustCompile(`port\[(\d+)\]`).FindStringSubmatch(string(context))
	// L_全局_地址 := regexp.MustCompile(`address\[([^\]]+)\]`).FindStringSubmatch(string(context))
	// L_全局_项目 := regexp.MustCompile(`project\[(\w+)\]`).FindStringSubmatch(string(context))

	// if len(L_全局_密码) == 2 {
	// 	全局_密码 = L_全局_密码[1]
	// }
	// if len(L_全局_端口) == 2 {
	// 	全局_端口 = L_全局_端口[1]
	// }
	// if len(L_全局_地址) == 2 {
	// 	全局_地址 = L_全局_地址[1]
	// }
	// if len(L_全局_项目) == 2 {
	// 	全局_项目 = L_全局_项目[1]
	// }
	viper.SetConfigFile("./下载设置.yaml")
	viper.ReadInConfig()
	全局_密码 = viper.GetString("password")
	全局_端口 = viper.GetString("port")
	全局_地址 = viper.GetString("address")
	全局_项目 = viper.GetString("project")
	// fmt.Println("密码", 全局_密码)
	// fmt.Println("端口", 全局_端口)
	// fmt.Println("地址", 全局_地址)
	// fmt.Println("项目", 全局_项目)

}
