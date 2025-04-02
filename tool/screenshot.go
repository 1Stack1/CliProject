package tool

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// TakeScreenshot 截取网页截图并返回文件路径
func TakeScreenshot(url string) (string, error) {
	// 定义自定义错误
	var (
		ErrScreenshotTimeout = errors.New("截图操作超时（10秒内未完成）")
		ErrNavigationFailed  = errors.New("页面导航失败")
	)
	//  创建带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 创建浏览器上下文（合并日志配置）
	ctx, cancel = chromedp.NewContext(
		ctx, // 注意：使用带超时的父context
		chromedp.WithLogf(log.Printf),
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
		chromedp.Sleep(2*time.Second), // 等待页面加载
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 检查是否已超时
			select {
			case <-ctx.Done():
				return ctx.Err() // 返回context的原始错误
			default:
				// 正常执行截图
				var err error
				buf, err = page.CaptureScreenshot().
					WithQuality(90).
					Do(ctx)
				return err
			}
		}), // 直接截取全屏（替代手动设置视口）
	)
	// 处理错误（特别识别超时情况）
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return "", ErrScreenshotTimeout
		}
		if strings.Contains(err.Error(), "navigation failed") {
			return "", fmt.Errorf("%w: %v", ErrNavigationFailed, err)
		}
		return "", fmt.Errorf("截图失败: %w", err)
	}

	// 保存文件
	if err := os.WriteFile(filePath, buf, 0644); err != nil {
		return "", fmt.Errorf("文件保存失败: %w", err)
	}

	return filePath, nil
}
