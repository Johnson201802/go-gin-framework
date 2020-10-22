package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"os"
)

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

//创建OSS实例
func CreateOss(){
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := "<LTAI4FuzqYkEXt8c4EYZ7TRJ>"
	accessKeySecret := "<ZaXpqhJf13PFRG8AePWnB0OpB7LG2k>"
	bucketName := "<johnson001>"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		HandleError(err)
	}
	// 创建存储空间。
	err = client.CreateBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
}

//列举文件
func ListFile (c *gin.Context) {
	// 创建OSSClient实例。
	client, err := oss.New("https://oss-accelerate.aliyuncs.com", "LTAI4FuzqYkEXt8c4EYZ7TRJ", "ZaXpqhJf13PFRG8AePWnB0OpB7LG2k")
	if err != nil {
		HandleError(err)
	}
	// 获取存储空间。
	bucketName := "johnson001"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	// 列举文件。
	marker := ""
	files := make(map[int]interface{}, 0)
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			HandleError(err)
		}
		// 打印列举文件，默认情况下一次返回100条记录。
		for key , object := range lsRes.Objects {
			files[key] = object.Key
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	c.JSON(200,gin.H{
		"code" : 200,
		"data" : files,
	})
}

//上传文件
func UploadFile(c *gin.Context){

}



