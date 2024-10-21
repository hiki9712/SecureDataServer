/*
* @desc:工具
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/3/4 22:16
 */

package libUtils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/tiger1103/gfast/v3/internal/app/common/consts"
)

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

// Upload 将文件上传到指定服务器
func Upload(ctx context.Context, params g.Map, jsonFileList []g.Map) (err error) {
	var (
		username string
		password string
		addr     string
		path     string
	)
	username = params["username"].(string)
	password = params["password"].(string)
	addr = params["addr"].(string)
	path = params["path"].(string)
	//ssh 连接
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		g.Log().Error(ctx, "connect ssh error", g.Map{"err": err})
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)
	// 创建SFTP客户端
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		g.Log().Error(ctx, "sftp new client error", g.Map{"err": err})
		return
	}
	defer func(sftpClient *sftp.Client) {
		err := sftpClient.Close()
		if err != nil {
			return
		}
	}(sftpClient)
	for i, jsonFile := range jsonFileList {
		jsonBytes, _ := json.Marshal(jsonFile)
		remoteFile, _ := sftpClient.Create(path + gconv.String(i) + ".json")
		//if err != nil {
		//	g.Log().Error(ctx, "create json file error", g.Map{"err": err})
		//	return
		//}
		remoteFile.Write(jsonBytes)
		defer func(remoteFile *sftp.File) {
			err := remoteFile.Close()
			if err != nil {

			}
		}(remoteFile)
	}
	return
}

// JsonFileSplit 将json文件按不同大小的要求切分成小json
func JsonFileSplit(ctx context.Context, originJson []map[string]interface{}, chunkSize int) (splitJson [][]map[string]interface{}, err error) {
	var (
		currentChunk []map[string]interface{}
		currentSize  int
	)
	for _, item := range originJson {
		jsonData, err := json.Marshal(item)
		if err != nil {
			g.Log().Error(ctx, "json marshal failed!: ", err)
			return nil, err
		}
		newSize := currentSize + len(jsonData)

		if newSize >= chunkSize {
			splitJson = append(splitJson, currentChunk)
			currentChunk = []map[string]interface{}{item}
			currentSize = len(jsonData)
		} else {
			currentChunk = append(currentChunk, item)
			currentSize += len(jsonData)
		}
	}

	if len(currentChunk) > 0 {
		splitJson = append(splitJson, currentChunk)
	}

	return splitJson, nil
}

// ResolveReq 接口解析为g.Map
func ResolveReq(ctx context.Context, req interface{}) (data g.Map, err error) {
	reqJson, err := json.Marshal(req)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = json.Unmarshal(reqJson, &data)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

// GenUniqId 生成一个int64类型的id
func GenUniqId(ctx context.Context) (res int64) {
	node, _ := snowflake.NewNode(1)
	res = node.Generate().Int64()
	return
}

// GetDomain 获取当前请求接口域名
func GetDomain(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	pathInfo, err := gurl.ParseURL(r.GetUrl(), -1)
	if err != nil {
		g.Log().Error(ctx, err)
		return ""
	}
	return fmt.Sprintf("%s://%s:%s/", pathInfo["scheme"], pathInfo["host"], pathInfo["port"])
}

// GetClientIp 获取客户端IP
func GetClientIp(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetClientIp()
}

// GetUserAgent 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

// GetLocalIP 服务端ip
func GetLocalIP() (ip string, err error) {
	var addrs []net.Addr
	addrs, err = net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// GetCityByIp 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "::1" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(context.TODO(), url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		return city
	} else {
		return ""
	}
}

// 写入文件
func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	n, _ := f.Seek(0, io.SeekEnd)
	_, err = f.WriteAt([]byte(content), n)
	defer f.Close()
	return err
}

// 文件或文件夹是否存在
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// 解析路径获取文件名称及后缀
func ParseFilePath(pathStr string) (fileName string, fileType string) {
	fileNameWithSuffix := path.Base(pathStr)
	fileType = path.Ext(fileNameWithSuffix)
	fileName = strings.TrimSuffix(fileNameWithSuffix, fileType)
	return
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !FileIsExisted(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// GetType 获取文件类型
func GetType(p string) (result string, err error) {
	file, err := os.Open(p)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	defer file.Close()
	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	filetype := http.DetectContentType(buff)
	return filetype, nil
}

// GetFilesPath 获取附件相对路径
func GetFilesPath(ctx context.Context, fileUrl string) (path string, err error) {
	upType := g.Cfg().MustGet(ctx, "upload.default").Int()
	if upType != 0 || (upType == 0 && !gstr.ContainsI(fileUrl, consts.UploadPath)) {
		path = fileUrl
		return
	}
	pathInfo, err := gurl.ParseURL(fileUrl, 32)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("解析附件路径失败")
		return
	}
	pos := gstr.PosI(pathInfo["path"], consts.UploadPath)
	if pos >= 0 {
		path = gstr.SubStr(pathInfo["path"], pos)
	}
	return
}
