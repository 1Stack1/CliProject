package tool

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
	"time"
)

// TakeScreenshot 截取网页截图并返回文件路径
func TakeScreenshot(url string) (string, error) {
	// 创建上下文（合并冗余的上下文创建）
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf), // 可选：记录浏览器日志
	)
	defer cancel()

	// 生成唯一文件名
	fileName := fmt.Sprintf("screenshot_%s_%d.png",
		uuid.New().String()[:8], // 取UUID前8位
		time.Now().UnixNano(),   // 添加时间戳防止冲突
	)
	filePath := filepath.Join("./screenshots", fileName)

	// 创建存储目录
	if err := os.MkdirAll("./screenshots", 0755); err != nil {
		return "", fmt.Errorf("创建目录失败: %w", err)
	}

	// 执行截图操作
	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),     // 等待页面加载
		chromedp.FullScreenshot(&buf, 90), // 直接截取全屏（替代手动设置视口）
	)
	if err != nil {
		return "", fmt.Errorf("截图失败: %w", err)
	}

	// 保存文件
	if err := os.WriteFile(filePath, buf, 0644); err != nil {
		return "", fmt.Errorf("文件保存失败: %w", err)
	}

	return filePath, nil
}
