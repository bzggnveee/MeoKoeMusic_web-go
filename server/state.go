package server

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type AppState struct {
	mu       sync.RWMutex
	path     string
	Settings map[string]string `json:"settings"`
	Cookies  map[string]string `json:"cookies"`
	Origin   string            `json:"origin"`    // "real" or "synced"
	DeviceID string            `json:"device_id"`
	CreateAt string            `json:"created_at"`
}

var appState *AppState

func InitState(filePath string) *AppState {
	s := &AppState{
		path:     filePath,
		Settings: map[string]string{},
		Cookies:  map[string]string{},
		Origin:   "synced",
	}
	if data, err := os.ReadFile(filePath); err == nil {
		json.Unmarshal(data, s)
	}
	if s.Origin == "" {
		s.Origin = "synced"
	}
	appState = s
	log.Printf("[state] 加载: %s (origin:%s cookies:%d)", filePath, s.Origin, len(s.Cookies))
	return s
}

func GetState() *AppState { return appState }

func (s *AppState) Save() {
	s.mu.RLock()
	data, err := json.MarshalIndent(s, "", "  ")
	s.mu.RUnlock()
	if err != nil {
		return
	}
	os.WriteFile(s.path, data, 0644)
}

// MergeCookies 真实登录时全量覆盖并标记 origin=real
func (s *AppState) MergeCookies(cookies map[string]string) {
	s.mu.Lock()
	if _, has := cookies["token"]; has {
		s.Cookies = make(map[string]string)
		for k, v := range cookies {
			if v != "" { s.Cookies[k] = v }
		}
		s.Origin = "real"
		s.CreateAt = time.Now().Format(time.RFC3339)
	}
	s.mu.Unlock()
	s.Save()
}

// MarkReal 标记当前登录为真实登录
func (s *AppState) MarkReal(deviceID string) {
	s.mu.Lock()
	s.Origin = "real"
	s.DeviceID = deviceID
	s.CreateAt = time.Now().Format(time.RFC3339)
	s.mu.Unlock()
	s.Save()
}

// IsRealLogin 判断当前是否为真实登录
func (s *AppState) IsRealLogin() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Origin == "real"
}

// GetOrigin 获取当前登录来源
func (s *AppState) GetOrigin() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Origin
}

func (s *AppState) GetSetting(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Settings[key]
}

func (s *AppState) SetSetting(key, value string) {
	s.mu.Lock()
	s.Settings[key] = value
	s.mu.Unlock()
	s.Save()
}

func (s *AppState) GetCookies() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r := make(map[string]string)
	for k, v := range s.Cookies {
		r[k] = v
	}
	return r
}